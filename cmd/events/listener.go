package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tr1sm0s1n/geth-ethclient-starter/contract"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"
)

func main() {
	client, err := helpers.DialClient(true)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	filterer, err := contract.NewCertFilterer(contractAddress, client)
	if err != nil {
		panic(err)
	}

	logs := make(chan *contract.CertIssued)
	sub, err := filterer.WatchIssued(nil, logs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for events...")
	fmt.Println("-----------------------")

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case log := <-logs:
			rw, _ := log.Raw.MarshalJSON()
			fmt.Println("Certificate issued!!")
			fmt.Println("--------------------")
			fmt.Printf("Course: \033[34m%s\033[0m\n", log.Course)
			fmt.Printf("ID: \033[34m%d\033[0m\n", log.Id.Int64())
			fmt.Printf("Date: \033[34m%s\033[0m\n", log.Date)
			fmt.Printf("Raw log: \033[32m%s\033[0m\n", string(rw))
			fmt.Println("--------------------")
		}
	}
}
