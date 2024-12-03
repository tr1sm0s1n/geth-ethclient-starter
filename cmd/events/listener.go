package main

import (
	"context"
	"encoding/json"
	"example/fe/helpers"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	client, err := helpers.DialClient(true)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for events...")
	fmt.Println("-----------------------")

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			uLog, _ := json.Marshal(vLog)
			fmt.Println(string(uLog))
		}
	}
}
