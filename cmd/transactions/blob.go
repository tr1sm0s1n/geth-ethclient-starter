package main

import (
	"context"
	"fmt"
	"github.com/tr1sm0s1n/geth-ethclient-starter/helpers"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

func blobTx() {
	myBlob := new(kzg4844.Blob)
	copy(myBlob[:], "Hello, World!")

	myBlobCommit, _ := kzg4844.BlobToCommitment(myBlob)
	myBlobProof, _ := kzg4844.ComputeBlobProof(myBlob, myBlobCommit)

	client, err := helpers.DialClient(false)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, key, err := helpers.AuthGenerator(client, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		panic(err)
	}

	to := common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7")

	sidecar := &types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{*myBlob},
		Commitments: []kzg4844.Commitment{myBlobCommit},
		Proofs:      []kzg4844.Proof{myBlobProof},
	}

	signedTx, _ := types.SignNewTx(key, types.LatestSignerForChainID(chainID), &types.BlobTx{
		Nonce:      nonce,
		To:         to,
		GasTipCap:  uint256.NewInt(1000000),
		GasFeeCap:  uint256.NewInt(1000000000),
		Gas:        21000,
		BlobFeeCap: uint256.NewInt(15),
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    sidecar,
	})

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", signedTx.Hash())
	fmt.Println("-----------------")

	for {
		r, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
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

	btx, _, err := client.TransactionByHash(context.Background(), signedTx.Hash())
	if err != nil {
		panic(err)
	}

	if btx.BlobHashes()[0] != sidecar.BlobHashes()[0] {
		panic("Failed to verify blob hashes")
	}
}

func main() {
	blobTx()
}
