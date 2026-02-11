package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0xB3966a54341fDB5AabEC33B9A29Fb3896047df99")

	fmt.Println(address.Hex())                             // returns a string of the address in hex format *MOST IMPORTANT AND COMMONLY USED*
	fmt.Println(common.BytesToHash(address.Bytes()).Hex()) // Convertes a 20 byte addy to hash format (32 bytes), some smart contract functions expect 32-byte hashes
	fmt.Println(address.Bytes())                           // returns a slice (array) of bytes
}
