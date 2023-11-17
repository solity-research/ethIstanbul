package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func main() {
	log.Println("Code started...")
	client, err := ethclient.Dial("wss://arbitrum-goerli.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	// Get the latest block number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve the latest block header: %v", err)
	}

	// Start scanning from the latest block
	latestBlockNumber := header.Number

	//Insert our Arbitrum Smart Contract Address
	//Temp SC 0x650f547c84b12458186c002e5df58b9cdb1f23c0
	contractAddress := common.HexToAddress("0x4e6C88BD2C76DFDA0ce867EcEd0D4cF82a00D17F")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: latestBlockNumber,
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Scanning started...")
	for {
		select {
		case subErr := <-sub.Err():
			log.Fatal(subErr)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			if strings.Contains(vLog.Topics[0].Hex(), "0xdc4b8b75577d2547483852294c9ed357a0b46adecd2b69d6882c5a27ef9fe16d") {
				log.Println("Buldum")
				//TODO: Decode log data

			}
		}
	}
}
