package main

// TODO: Import the necessary packages
// You'll need:
// - "fmt" - for formatted printing
// - "github.com/ethereum/go-ethereum/common" - for Address and Hash types
//
// Resources:
// - Go packages: https://go.dev/learn/#packages
// - common package: https://pkg.go.dev/github.com/ethereum/go-ethereum/common
import (
	"fmt"
	// Add common import here
	// "github.com/ethereum/go-ethereum/common"
)

func main() {
	// TODO: Convert a hex address string to common.Address type
	// Ethereum addresses are strings like "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	// Use common.HexToAddress() to convert the string to an Address type
	//
	// Why convert to Address type?
	// - go-ethereum methods require Address type, not strings
	// - Provides type safety and validation
	// - Used when querying balances, sending transactions, interacting with contracts
	//
	// Go syntax hint:
	// - Functions are called: result := functionName(argument)
	//
	// Resources:
	// - common.HexToAddress: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#HexToAddress
	// - Address type: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Address

	address := common.HexToAddress("YOUR_ADDRESS_HERE")

	// Convert Address back to hex string
	// Use address.Hex() to get the string representation
	// This is useful for:
	// - Displaying addresses to users
	// - Logging and debugging
	// - Converting back to string format for APIs
	//
	// Go syntax hint:
	// - Methods are called on objects: object.Method()
	//
	// Resources:
	// - Address.Hex: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Address.Hex

	fmt.Println(address.Hex())

	// Convert Address to Hash (32-byte format)
	// Sometimes you need addresses in 32-byte format (like in smart contract storage)
	// Steps:
	// 1. Get the address bytes using address.Bytes()
	// 2. Convert to Hash using common.BytesToHash()
	// 3. Get hex string with .Hex()
	//
	// When is this needed?
	// - Smart contract storage uses 32-byte slots
	// - Event logs and topics are 32 bytes
	// - Some contract functions expect 32-byte parameters
	//
	// Go syntax hint:
	// - Method chaining: object.Method1().Method2()
	// - BytesToHash pads the 20-byte address to 32 bytes (adds zeros on the left)
	//
	// Resources:
	// - Address.Bytes: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Address.Bytes
	// - common.BytesToHash: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#BytesToHash
	// - Hash type: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Hash

	fmt.Println(common.BytesToHash(address.Bytes()).Hex())

	// Get the raw bytes of the address
	// Use address.Bytes() to get the 20-byte array
	// This is useful for:
	// - Low-level operations
	// - Custom encoding/decoding
	// - Cryptographic operations
	//
	// Go syntax hint:
	// - Returns a byte slice ([]byte)
	//
	// Resources:
	// - Address.Bytes: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Address.Bytes
	// - Go bytes: https://pkg.go.dev/bytes

	fmt.Println(address.Bytes())
}
