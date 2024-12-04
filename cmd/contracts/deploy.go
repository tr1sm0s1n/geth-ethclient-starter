package main

import (
	"context"
	"github.com/tr1sm0s1n/geth-ethclient-starter/contract"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := helpers.DialClient(false)
	if err != nil {
		panic(err)
	}

	contractAddress, trx, err := deployContract(client)
	if err != nil {
		panic(err)
	}

	for {
		_, pending, err := client.TransactionByHash(context.Background(), trx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				continue
			} else {
				panic(err)
			}
		}

		if !pending {
			fmt.Println("Transaction has been committed!!")
			fmt.Println("--------------------------------")
			break
		}

		fmt.Println("Transaction is pending...")
		time.Sleep(time.Second)
	}

	fmt.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
	fmt.Println("-----------------")
	fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
	fmt.Println("-----------------")
}

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	auth, _, err := helpers.AuthGenerator(client, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return common.Address{}, nil, err
	}
	contract, transaction, _, err := contract.DeployCert(auth, client)
	return contract, transaction, err
}
