package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xdadB0d80178819F2319190D340ce9A924f783711")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current balance (wei):", balance)

	blockNumber := big.NewInt(24430197)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance at block %d (wei): %s\n", blockNumber, balanceAt.String())

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("Balance at block %d (ETH): %s\n", blockNumber, ethValue.Text('f', 18))

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pending balance (wei):", pendingBalance)
}
