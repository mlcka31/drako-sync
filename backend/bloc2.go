package main

import (
	"backend/contract"
	"backend/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type BlockChain struct {
	client   *ethclient.Client
	address  common.Address
	abi      abi.ABI
	auth     *bind.TransactOpts
	contract *bind.BoundContract

	instance *contract.Contract
}

func (b *BlockChain) Init() {
	var err error
	env := utils.LoadEnvs()

	// Connect to Ethereum client (e.g., Ganache, Infura, local node)
	b.client, err = ethclient.Dial(env.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(env.PRIVATE_KEY)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	chainID, err := b.client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transaction auth: %v", err)
	}

	// Set initial message price and coefficient increase (example values)
	initialMessagePrice := big.NewInt(1e15)          // 0.001 ETH
	coeffIncrease := big.NewInt(1100000000000000000) // 1.1 scaled by PRECISION (1e18)

	// Deploy the contract
	address, tx, instance, err := contract.DeployContract(auth, b.client, initialMessagePrice, coeffIncrease)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed at: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	// Wait for deployment confirmation (simplified, in production use proper receipt checking)
	_, err = bind.WaitDeployed(context.Background(), b.client, tx)
	if err != nil {
		log.Fatalf("Failed to confirm contract deployment: %v", err)
	}

	// Interact with contract: Set AI Agent Address (admin-only action)
	aiAgentAddr := common.HexToAddress(env.ADMIN_ADDRESS)
	tx, err = instance.SetAIAgentAddress(auth, aiAgentAddr)
	if err != nil {
		log.Fatalf("Failed to set AI agent address: %v", err)
	}
	fmt.Printf("Set AI Agent address TX: %s\n", tx.Hash().Hex())

	// Initialize Game
	tx, err = instance.InitGame(auth, aiAgentAddr)
	if err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}
	fmt.Printf("Game initialized TX: %s\n", tx.Hash().Hex())
}

func (b *BlockChain) PlayGame() {
	// User needs to send the exact message price; fetch it first
	msgPrice, err := b.instance.GetmessagePrice(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to fetch message price: %v", err)
	}

	// Example: Send message (user action)
	messageContent := "Hello, AI!"

	b.auth.Value = msgPrice // attach exact message price as ETH
	tx, err := b.instance.SendMessage(b.auth, messageContent)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Printf("Message sent TX: %s\n", tx.Hash().Hex())

	// Fetch and print game state
	state, err := b.instance.GetGameState(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve game state: %v", err)
	}
	fmt.Printf("Current Game State: %s\n", parseGameState(state))

	// Fetch prize pool
	prizePool, err := b.instance.GetPrizePool(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve prize pool: %v", err)
	}
	fmt.Printf("Current Prize Pool: %s wei\n", prizePool.String())
}

// Helper function to parse game state enum
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
