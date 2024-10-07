package helpers

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}

	return client, nil
}
