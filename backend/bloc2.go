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
	"strings"
	"time"
)

type Blockchain struct {
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

func NewBlockchain(logger *zap.Logger, chatGPTClient chatgpt.ChatGPTClient, accountManager *AccountManager, server *Server) *Blockchain {
	logger.Info("NewBlockchain")
	env := utils.LoadEnvs()

	client, err := ethclient.Dial(env.RPC_URL)
	if err != nil {
		logger.Fatal("Failed to connect to Ethereum client", zap.Error(err))
	}

	privateKey, err := crypto.HexToECDSA(env.PRIVATE_KEY)
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

	return &Blockchain{
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

func (bc *Blockchain) DeployContract() {
	address, tx, instance, err := contract.DeployContract(bc.Auth, bc.Client, bc.InitialMessagePrice, bc.CoeffIncrease)
	if err != nil {
		bc.logger.Fatal("Failed to deploy contract", zap.Error(err))
	}

	//fmt.Printf("Contract deployed at: %s\n", address.Hex())
	//fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
	bc.logger.Info("Contract deployed", zap.String("address", address.Hex()))
	bc.logger.Info("Transaction hash", zap.String("tx", tx.Hash().Hex()))

	_, err = bind.WaitDeployed(context.Background(), bc.Client, tx)
	if err != nil {
		bc.logger.Fatal("Failed to confirm contract deployment", zap.Error(err))
	}

	bc.ContractAddress = address
	bc.Instance = instance
}

func (bc *Blockchain) StartGame() {
	tx, err := bc.Instance.SetAIAgentAddress(bc.Auth, bc.AiAgentAddress)
	if err != nil {
		bc.logger.Fatal("Failed to set AI agent address", zap.Error(err))
	}
	bc.logger.Info("Set AI Agent address TX", zap.String("tx", tx.Hash().Hex()))

	tx, err = bc.Instance.InitGame(bc.Auth, bc.AiAgentAddress)
	if err != nil {
		bc.logger.Fatal("Failed to initialize game", zap.Error(err))
	}
	bc.logger.Info("Game initialized TX", zap.String("tx", tx.Hash().Hex()))

	messageContent := "Hello, AI!"
	msgPrice, err := bc.Instance.GetmessagePrice(&bind.CallOpts{})
	if err != nil {
		bc.logger.Fatal("Failed to fetch message price", zap.Error(err))
	}

	bc.Auth.Value = msgPrice
	tx, err = bc.Instance.SendMessage(bc.Auth, messageContent)
	if err != nil {
		bc.logger.Fatal("Failed to send message", zap.Error(err))
	}
	bc.logger.Info("Message sent TX", zap.String("tx", tx.Hash().Hex()))

	state, err := bc.Instance.GetGameState(&bind.CallOpts{})
	if err != nil {
		bc.logger.Fatal("Failed to retrieve game state", zap.Error(err))
	}
	bc.logger.Info("Current Game State", zap.String("state", parseGameState(state)))

	prizePool, err := bc.Instance.GetPrizePool(&bind.CallOpts{})
	if err != nil {
		bc.logger.Fatal("Failed to retrieve prize pool", zap.Error(err))
	}
	bc.logger.Info("Current Prize Pool", zap.String("prizePool", prizePool.String()))
}

func (bc *Blockchain) MonitorMessages() {
	bc.logger.Info("MonitorMessages")

	for {
		stateInt, err := bc.Instance.GetGameState(&bind.CallOpts{})
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
		messageContents := make([]string, len(messages))
		for i, msg := range messages {
			messageContents[i] = msg.Content
		}

		// Send all messages to ChatGPT in a single request
		chatResponse, err := bc.ChatGPTClient.SendAllMessages(context.Background(), messageContents)
		if err != nil {
			bc.logger.Error("Failed to send messages to ChatGPT", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}

		// Process ChatGPT's response
		bc.logger.Info("ChatGPT response", zap.String("response", chatResponse))

		if strings.Contains(chatResponse, "success") {
			tx, err := bc.Instance.PayoutPrize(bc.Auth)
			if err != nil {
				bc.logger.Error("Failed to PayoutPrize", zap.Error(err))
				continue
			}
			bc.logger.Info("Reply transaction sent", zap.String("tx", tx.Hash().Hex()))
			break
		}

		// Optionally, call the smart contract's reply function with the ChatGPT response
		tx, err := bc.Instance.Reply(bc.Auth, chatResponse)
		if err != nil {
			bc.logger.Error("Failed to call reply function", zap.Error(err))
			continue
		}
		bc.logger.Info("Reply transaction sent", zap.String("tx", tx.Hash().Hex()))

		// Wait before checking for new messages
		time.Sleep(3 * time.Second)
	}
}

// CheckWalletBalance retrieves and prints the balance of the wallet associated with the Blockchain instance
func (bc *Blockchain) CheckWalletBalance() {
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

	chatGPTClient := chatgpt.NewChatGPTClient(env.OPEN_AI_KEY)

	accountManager, err := NewAccountManager("private_key.hex")
	if err != nil {
		logger.Fatal("Error creating account", zap.Error(err))
		//log.Fatalf("Error creating account: %v", err)
	}

	server := NewServer("serv1", logger)
	go server.Run("localhost:8080")

	for i := 0; i < 10; i++ {
		logger.Info("test", zap.Int("i", i))
	}

	blockchain := NewBlockchain(logger, *chatGPTClient, accountManager, server)
	//blockchain.DeployContract()
	//blockchain.StartGame()
	blockchain.MonitorMessages()
	//blockchain.CheckWalletBalance()
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
