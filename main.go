package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/tr1sm0s1n/geth-ethclient-starter/contract"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

var (
	contractAddress = common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	id              big.Int
	certificate     struct {
		name   string
		course string
		grade  string
		date   string
	}
)

func main() {
	client, err := helpers.DialClient(false)
	if err != nil {
		panic(err)
	}

	auth, _, err := helpers.AuthGenerator(client, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	cert := contract.NewCert()
	instance := cert.Instance(client, contractAddress)

	fmt.Println("Enter student details:")
	fmt.Println("----------------------")
	fmt.Print("ID: ")
	fmt.Scanln(&id)
	fmt.Print("Name: ")
	fmt.Scanln(&certificate.name)
	fmt.Print("Course: ")
	fmt.Scanln(&certificate.course)
	fmt.Print("Grade: ")
	fmt.Scanln(&certificate.grade)
	fmt.Print("Date: ")
	fmt.Scanln(&certificate.date)

	trx, err := bind.Transact(instance, auth, cert.PackIssueCertificate(&id, certificate.name, certificate.course, certificate.grade, certificate.date))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nTransaction Hash: \033[32m%s\033[0m\n", trx.Hash())
	fmt.Println("-----------------")

	for {
		r, err := client.TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				continue
			} else {
				panic(err)
			}
		}

		if r.Status == 1 {
			fmt.Println("Transaction has been committed!!")
			fmt.Println("--------------------------------")
			break
		}

		fmt.Println("Status not committed")
		time.Sleep(time.Second)
	}

	fmt.Print("\nEnter ID to fetch: ")
	fmt.Scanln(&id)

	c, err := bind.Call(instance, nil, cert.PackFetchCertificate(&id), cert.UnpackFetchCertificate)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nCertificate Details:")
	fmt.Println("--------------------")
	fmt.Printf("ID: \033[33m%d\033[0m\n", id.Int64())
	fmt.Printf("Name: \033[33m%s\033[0m\n", c.Name)
	fmt.Printf("Course: \033[33m%s\033[0m\n", c.Course)
	fmt.Printf("Grade: \033[33m%s\033[0m\n", c.Grade)
	fmt.Printf("Date: \033[33m%s\033[0m\n", c.Date)
}
