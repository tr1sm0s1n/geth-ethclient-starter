package main

import (
	"context"
	"fmt"
	"github.com/tr1sm0s1n/geth-ethclient-starter/contract"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

	certABI, err := abi.JSON(strings.NewReader(contract.CertMetaData.ABI))
	if err != nil {
		panic(err)
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
			var IssuedEvent contract.CertIssued

			certABI.UnpackIntoInterface(&IssuedEvent, "Issued", vLog.Data)
			uLog, _ := vLog.MarshalJSON()
			fmt.Println("Certificate issued!!")
			fmt.Println("--------------------")
			fmt.Printf("Course: \033[34m%s\033[0m\n", IssuedEvent.Course)
			fmt.Printf("ID: \033[34m%d\033[0m\n", IssuedEvent.Id.Int64())
			fmt.Printf("Date: \033[34m%s\033[0m\n", IssuedEvent.Date)
			fmt.Printf("Raw log: \033[32m%s\033[0m\n", string(uLog))
			fmt.Println("--------------------")
		}
	}
}
