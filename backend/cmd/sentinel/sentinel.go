package main

import (
	"context"
	"fmt"
	"log"

	"price-onchain-monitoring/backend/contracts" // Ensure this path matches your go.mod

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. Connect to Anvil
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 2. Configuration (Matches your Foundry deploy)
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	contractAddr  := "0x5FbDB2315678afecb367f032d93F642f64180aa3"

	fmt.Println("Attempting to trigger circuit breaker...")
	triggerCircuitBreaker(client, privateKeyHex, contractAddr)
}

func triggerCircuitBreaker(client *ethclient.Client, privateKeyHex string, contractAddr string) {
	// 1. Setup credentials
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Point to your deployed TradingGuard
	address := common.HexToAddress(contractAddr)
	instance, err := contracts.NewTradingGuard(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 3. PULL THE BRAKE
	tx, err := instance.TriggerCircuitBreaker(auth, "AI detected price manipulation")
	if err != nil {
		log.Fatalf("FAILED TO STOP TRADING: %v", err)
	}

	fmt.Printf("🛑 EMERGENCY: Trading Paused. TX Hash: %s\n", tx.Hash().Hex())
}