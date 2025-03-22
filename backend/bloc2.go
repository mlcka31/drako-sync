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
	"log"
	"math/big"
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

	ChatGPTClient chatgpt.ChatGPTClient
}

func NewBlockchain() *Blockchain {
	env := utils.LoadEnvs()

	client, err := ethclient.Dial(env.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(env.PRIVATE_KEY)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transaction auth: %v", err)
	}

	contractAddress := common.HexToAddress(env.GAME_CONTRACT_ADDRESS)

	// Create an instance of the contract
	instance, err := contract.NewContract(contractAddress, client)

	chatGPTClient := chatgpt.NewChatGPTClient(env.OPEN_AI_KEY)

	return &Blockchain{
		Client:              client,
		Auth:                auth,
		ChainID:             chainID,
		Instance:            instance,
		ChatGPTClient:       *chatGPTClient,
		InitialMessagePrice: big.NewInt(1e15),
		CoeffIncrease:       big.NewInt(1100000000000000000),
		AiAgentAddress:      common.HexToAddress(env.ADMIN_ADDRESS),
	}
}

func (bc *Blockchain) DeployContract() {
	address, tx, instance, err := contract.DeployContract(bc.Auth, bc.Client, bc.InitialMessagePrice, bc.CoeffIncrease)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed at: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	_, err = bind.WaitDeployed(context.Background(), bc.Client, tx)
	if err != nil {
		log.Fatalf("Failed to confirm contract deployment: %v", err)
	}

	bc.ContractAddress = address
	bc.Instance = instance
}

func (bc *Blockchain) StartGame() {
	tx, err := bc.Instance.SetAIAgentAddress(bc.Auth, bc.AiAgentAddress)
	if err != nil {
		log.Fatalf("Failed to set AI agent address: %v", err)
	}
	fmt.Printf("Set AI Agent address TX: %s\n", tx.Hash().Hex())

	tx, err = bc.Instance.InitGame(bc.Auth, bc.AiAgentAddress)
	if err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}
	fmt.Printf("Game initialized TX: %s\n", tx.Hash().Hex())

	messageContent := "Hello, AI!"
	msgPrice, err := bc.Instance.GetmessagePrice(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to fetch message price: %v", err)
	}

	bc.Auth.Value = msgPrice
	tx, err = bc.Instance.SendMessage(bc.Auth, messageContent)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Printf("Message sent TX: %s\n", tx.Hash().Hex())

	state, err := bc.Instance.GetGameState(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve game state: %v", err)
	}
	fmt.Printf("Current Game State: %s\n", parseGameState(state))

	prizePool, err := bc.Instance.GetPrizePool(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve prize pool: %v", err)
	}
	fmt.Printf("Current Prize Pool: %s wei\n", prizePool.String())
}

func (bc *Blockchain) MonitorMessages() {
	for {
		stateInt, err := bc.Instance.GetGameState(&bind.CallOpts{})
		state := parseGameState(stateInt)
		if state != "AgentAction" {
			time.Sleep(10 * time.Second)
			continue
		}

		// Retrieve all messages from the contract
		messages, err := bc.Instance.GetMessages(&bind.CallOpts{})
		if err != nil {
			log.Printf("Failed to retrieve messages: %v", err)
			time.Sleep(10 * time.Second)
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
			log.Printf("Failed to send messages to ChatGPT: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// Process ChatGPT's response
		log.Printf("ChatGPT response:\n%s", chatResponse)

		if strings.Contains(chatResponse, "success") {
			tx, err := bc.Instance.PayoutPrize(bc.Auth)
			if err != nil {
				log.Printf("Failed to PayoutPrize: %v", err)
				continue
			}
			log.Printf("Reply transaction sent: %s", tx.Hash().Hex())
			break
		}

		// Optionally, call the smart contract's reply function with the ChatGPT response
		tx, err := bc.Instance.Reply(bc.Auth, chatResponse)
		if err != nil {
			log.Printf("Failed to call reply function: %v", err)
			continue
		}
		log.Printf("Reply transaction sent: %s", tx.Hash().Hex())

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
	blockchain := NewBlockchain()
	//blockchain.DeployContract()
	//blockchain.StartGame()
	blockchain.MonitorMessages()
	//blockchain.CheckWalletBalance()
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
