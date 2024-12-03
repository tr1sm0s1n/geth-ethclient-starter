package main

import (
	"example/fe/contract"
	"example/fe/helpers"
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
	auth, _, err := helpers.AuthGenerator(client, "0x2ceea081f9765bc358f9f0ae1cd462b706a8c8ff6eb1897a921ca8f3ef01ce28")
	if err != nil {
		return common.Address{}, nil, err
	}
	contract, transaction, _, err := contract.DeployCert(auth, client)
	return contract, transaction, err
}
