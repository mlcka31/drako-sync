package main

import (
	"backend/chatgpt"
	"backend/contract"
	"backend/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"math/big"
	"os"
	"time"
)

type Main struct {
	Client              *ethclient.Client
	Auth                *bind.TransactOpts
	ChainID             *big.Int
	ContractAddress     common.Address
	Instance            *contract.Contract
	InitialMessagePrice *big.Int
	CoeffIncrease       *big.Int
	AiAgentAddress      common.Address

	logger         *zap.Logger
	accountManager *AccountManager
	server         *Server
	ChatGPTClient  chatgpt.ChatGPTClient
}

func NewMain(logger *zap.Logger, chatGPTClient chatgpt.ChatGPTClient, accountManager *AccountManager, server *Server) *Main {
	logger.Info("NewMain")
	env := utils.LoadEnvs()

	client, err := ethclient.Dial(env.RPC_URL)
	if err != nil {
		logger.Fatal("Failed to connect to Ethereum client", zap.Error(err))
	}

	if err := accountManager.WritePrivateKeyToFile(); err != nil {
		logger.Fatal("Failed to write private key to file", zap.Error(err))
	}

	// Load the private key from the file
	privateKeyBytes, err := os.ReadFile(accountManager.PrivateKeyPath)
	if err != nil {
		logger.Fatal("Failed to read private key file", zap.Error(err))
	}
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		logger.Fatal("Failed to parse private key", zap.Error(err))
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.Fatal("Failed to fetch chain ID", zap.Error(err))
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logger.Fatal("Failed to create transaction auth", zap.Error(err))
	}

	contractAddress := common.HexToAddress(env.GAME_CONTRACT_ADDRESS)

	// Create an instance of the contract
	instance, err := contract.NewContract(contractAddress, client)

	chatGPTClient.SetAIAddress(env.ADMIN_ADDRESS)

	return &Main{
		Client:              client,
		Auth:                auth,
		ChainID:             chainID,
		Instance:            instance,
		ChatGPTClient:       chatGPTClient,
		accountManager:      accountManager,
		server:              server,
		logger:              logger,
		InitialMessagePrice: big.NewInt(1e15),
		CoeffIncrease:       big.NewInt(1100000000000000000),
		AiAgentAddress:      common.HexToAddress(env.ADMIN_ADDRESS),
	}
}

func (bc *Main) MonitorMessages() {
	bc.logger.Info("MonitorMessages")

	for {
		stateInt, err := bc.Instance.GameState(&bind.CallOpts{})
		state := parseGameState(stateInt)
		if state != "AgentAction" {
			time.Sleep(3 * time.Second)
			continue
		}

		// Retrieve all messages from the contract
		messages, err := bc.Instance.GetMessages(&bind.CallOpts{})
		if err != nil {
			bc.logger.Error("Failed to retrieve messages", zap.Error(err))
			time.Sleep(3 * time.Second)
			continue
		}

		// Extract content from each message
		last10Messages := takeAtMax10(messages)

		// Send all messages to ChatGPT in a single request
		chatResponse, err := bc.ChatGPTClient.SendAllMessages(context.Background(), last10Messages)
		if err != nil {
			bc.logger.Error("Failed to send messages to ChatGPT", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}

		// Process ChatGPT's response
		bc.logger.Info("ChatGPT response", zap.Any("response", chatResponse))

		if chatResponse.IsSuccess {
			tx, err := bc.Instance.PayoutPrize(bc.Auth)
			if err != nil {
				bc.logger.Error("Failed to PayoutPrize", zap.Error(err))
				continue
			}
			bc.logger.Info("Reply transaction sent", zap.String("tx", tx.Hash().Hex()))
			break
		}

		// Optionally, call the smart contract's reply function with the ChatGPT response
		tx, err := bc.Instance.Reply(bc.Auth, chatResponse.Content)
		if err != nil {
			bc.logger.Error("Failed to call reply function", zap.Error(err))
			continue
		}
		bc.logger.Info("Reply transaction sent", zap.String("tx", tx.Hash().Hex()))

		// Wait before checking for new messages
		time.Sleep(3 * time.Second)
	}
}

// CheckWalletBalance retrieves and prints the balance of the wallet associated with the Main instance
func (bc *Main) CheckWalletBalance() {
	// Retrieve the wallet address from the Auth field
	walletAddress := bc.Auth.From

	// Fetch the balance at the latest block
	balance, err := bc.Client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve wallet balance: %v", err)
	}

	// Convert balance from Wei to Ether
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	// Print the wallet balance
	fmt.Printf("Balance of wallet %s: %f ETH\n", walletAddress.Hex(), etherValue)
}

func main() {
	env := utils.LoadEnvs()

	logger, err := setupLogger()
	defer logger.Sync()

	chatGPTClient := chatgpt.NewChatGPTClient(env.OPEN_AI_KEY, env.INIT_PROMPT, logger)

	//accountManager, err := NewAccountManager("private_key.hex", env, logger)
	accountManager, err := NewAccountManagerSecure("private_key.hex", logger)
	if err != nil {
		logger.Fatal("Error creating account", zap.Error(err))
	}
	pubKey := accountManager.GetPublicKey()

	server := NewServer("zk-dungeon", logger)
	server.SetPublicKey(pubKey)

	go server.Run(env.WEB_HOST + ":" + env.WEB_PORT)

	blockchain := NewMain(logger, *chatGPTClient, accountManager, server)
	blockchain.MonitorMessages()
}

func setupLogger() (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""   // Omit timestamp
	encoderConfig.CallerKey = "" // Omit caller information
	config := zap.NewProductionConfig()
	config.Sampling = nil
	//logger, err := config.Build()

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel)

	// Construct the logger
	logger := zap.New(core)

	return logger, nil
}

func parseGameState(state uint8) string {
	switch state {
	case 0:
		return "UserAction"
	case 1:
		return "AgentAction"
	case 2:
		return "Complete"
	case 3:
		return "NotStarted"
	default:
		return "Unknown"
	}
}

func takeAtMax10[T any](items []T) []T {
	if len(items) > 10 {
		return items[len(items)-10:]
	}
	return items
}
