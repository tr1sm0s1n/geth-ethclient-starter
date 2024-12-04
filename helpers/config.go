package helpers

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClient(ws bool) (*ethclient.Client, error) {
	url := "http://127.0.0.1:8545"
	if ws {
		url = "ws://127.0.0.1:8546"
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return client, nil
}
