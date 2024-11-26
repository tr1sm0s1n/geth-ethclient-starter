package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Printf("Private Key: \033[32m%s\033[0m\n", hexutil.Encode(crypto.FromECDSA(key)))
	fmt.Printf("Address: \033[32m%s\033[0m\n", address.String())
}
