package main

import (
	"example/fe/helpers"
	"example/fe/lib"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := helpers.DialClient()
	if err != nil {
		panic(err)
	}

	contractAddress, trx, err := deployContract(client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction has been committed!!")
	fmt.Println("--------------------------------")
	fmt.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
	fmt.Println("-----------------")
	fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
	fmt.Println("-----------------")
}

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	auth, _, err := helpers.AuthGenerator(client, "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		return common.Address{}, nil, err
	}
	contract, transaction, _, err := lib.DeployCert(auth, client)
	return contract, transaction, err
}
