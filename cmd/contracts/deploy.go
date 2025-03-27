package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tr1sm0s1n/geth-ethclient-starter/contract"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"
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

	deployParams := bind.DeploymentParams{
		Contracts: []*bind.MetaData{&contract.CertMetaData},
	}

	deployer := bind.DefaultDeployer(auth, client)
	result, err := bind.LinkAndDeploy(&deployParams, deployer)

	if _, err := bind.WaitDeployed(context.Background(), client, result.Txs[contract.CertMetaData.ID].Hash()); err != nil {
		return common.Address{}, nil, err
	}
	return result.Addresses[contract.CertMetaData.ID], result.Txs[contract.CertMetaData.ID], err
}
