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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b3eec657d195488b902512ae0f6e7654")
	if err != nil {
		log.Println(err)
	}
	blockNum := big.NewInt(15384242)
	block, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		log.Println(err)
	}
	log.Println("Following is the block number:")
	log.Println(block.Number().Uint64())
	for _, tx := range block.Transactions() {
		if tx.Hash().Hex() == "0x8af2db010af2d9df698ffe0ef9d52a27bb00f079ba360f0e157cb19c0141c98f" {
			log.Println("Transaction found: Since it is a token interaction, the recepient is the contract address")
			// contract address
			log.Println("Contract address:")
			log.Println(tx.To().Hex())
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
			fmt.Println(logs)
			contractAbi, err := abi.JSON(strings.NewReader(string(TokenABI)))
			if err != nil {
				log.Fatal(err)
			}
			logTransferSig := []byte("Transfer(address,address,uint256)")
			LogApprovalSig := []byte("Approval(address,address,uint256)")
			logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
			logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

		}
	}
}
