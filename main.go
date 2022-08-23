package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens any
}

func main() {
	// first we connect to the RPC provider for Ethereum mainnet (please include your own)
	client, err := ethclient.Dial("INFURA_LINK")
	if err != nil {
		log.Println(err)
	}
	//block number is already provided for the event log
	blockNum := big.NewInt(15384242)
	block, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		log.Println(err)
	}
	log.Println("Following is the block number:")
	log.Println(block.Number().Uint64())
	//looping through this will only tell us how the user interacted with the contract and not where the tokens were received.
	// Its required to loop through the logs of the SMART CONTRACT with which the interaction took place to know how token transaction took place.
	for _, tx := range block.Transactions() {
		if tx.Hash().Hex() == "0x8af2db010af2d9df698ffe0ef9d52a27bb00f079ba360f0e157cb19c0141c98f" {
			log.Println("Transaction found: Since it is a token interaction, the recepient is the contract address")
			// contract address where the transaction was sent
			log.Println("Contract address for USDT:", tx.To().Hex())
			log.Println()
			//fetching the event logs of ERC20
			contractAddress := common.HexToAddress(tx.To().Hex())
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(15384242),
				ToBlock:   big.NewInt(15384242),
				Addresses: []common.Address{
					contractAddress,
				},
			}
			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				log.Fatal(err)
			}
			//loading the ABI file to read the logs of Contract ABI
			contractAbi, err := abi.JSON(strings.NewReader(MainMetaData.ABI))
			if err != nil {
				log.Fatal(err)
			}
			//looping through the USDT contract's logs to search for the transfer with the specified receiver's address
			for _, vLog := range logs {
				if common.HexToAddress(vLog.Topics[2].Hex()) == common.HexToAddress("0x4b84c177ab5b6808c0017305b6dc88a81269f27b") {
					fmt.Println(" ")
					fmt.Println("Found USDT transaction")
					fmt.Printf("Log Name: Transfer\n")

					var transferEvent LogTransfer
					var res []interface{}
					res, err := contractAbi.Unpack("Transfer", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
					transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
					transferEvent.Tokens = res[0]

					fmt.Printf("From Address: %s\n", transferEvent.From.Hex())
					fmt.Printf("To Address: %s\n", transferEvent.To.Hex())
					fmt.Printf("Tokens Transferred: %s\n", transferEvent.Tokens)
				}
			}
		}
	}
}
