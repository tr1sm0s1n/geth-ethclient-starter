# Geth ethclient Starter

A comprehensive guide to using the `ethclient` package from **Geth** ([go-ethereum](https://github.com/ethereum/go-ethereum)) to build Ethereum applications with **Go**.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)

## ⚙️ Run Locally

Clone the project:

```sh
git clone https://github.com/tr1sm0s1n/geth-ethclient-starter.git
cd geth-ethclient-starter
```

Use AetherGuild submodule for simulated blockchain:

```sh
git submodule update --init --recursive
cd aetherguild && make
```

Generate address-key pair:

```sh
go run cmd/accounts/generator.go
```

Visit [Druid Faucet](http://127.0.0.1:8580) to get test ether.

Export private key as env variable:

```sh
export PRIVATE_KEY=<private-key>
```

Install `abigen`:

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for the contract:

```bash
abigen --v2 --bin contract/output/Cert/Cert.bin --abi contract/output/Cert/Cert_abi.json --pkg contract --type Cert --out contract/Cert.go
```

Deploy the contract:

```sh
go run cmd/contracts/deploy.go
```

Export contract address as env variable:

```sh
export CONTRACT_ADDRESS=<contract-address>
```

Run event listener for the contract:

```sh
go run cmd/events/listener.go
```

Run CLI instigator for contract (export values if needed):

```sh
go run .
```

Send a blob transaction to the simulation (optional):

```sh
go run cmd/transactions/blob.go
```

Use [`graphql/queries`](graphql/queries) to get balance, blocks and receipts via [GraphQL UI](http://127.0.0.1:8545/graphql/ui).
