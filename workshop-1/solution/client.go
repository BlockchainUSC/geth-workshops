package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/")
	//client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latest Block Number: %d\n", blockNum)
}
