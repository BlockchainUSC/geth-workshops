package main

// TODO: Import the necessary packages
// You'll need:
// - "context" - for context management
// - "fmt" - for formatted printing
// - "log" - for logging errors
// - "math" - for math operations (like Pow10)
// - "math/big" - for handling large numbers (Ethereum uses big.Int and big.Float)
// - "github.com/ethereum/go-ethereum/common" - for Address type
// - "github.com/ethereum/go-ethereum/ethclient" - the Ethereum client library
//
// Resources:
// - Go packages: https://go.dev/learn/#packages
// - math/big package: https://pkg.go.dev/math/big
import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	// Add common and ethclient imports here
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// TODO: Create an Ethereum client connection
	// Use ethclient.Dial() to connect to Alchemy
	// Format: "https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY"
	//
	// Resources:
	// - ethclient.Dial: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Dial
	// - Alchemy docs: https://docs.alchemy.com/

	// Create client connection here
	// Fill in your Alchemy RPC URL
	client, err := ethclient.Dial("YOUR_ALCHEMY_RPC_URL_HERE")
	// Handle error here

	// TODO: Convert address string to common.Address type
	// Use common.HexToAddress() with an Ethereum address string
	// Example: "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	// Go syntax hint:
    // - Functions are called: result := functionName(argument)
    // - Variables are declared: var name type or name := value
	//
	// Resources:
	// - common.HexToAddress: https://pkg.go.dev/github.com/ethereum/go-ethereum/common#HexToAddress

	account := // Convert address string here

	// TODO: Get the current balance
	// Use client.BalanceAt() method
	// Parameters: context, address, blockNumber (use nil for latest)
	// Returns: balance (*big.Int), error
	//
	// Go syntax hint:
	// - nil means "no value" - use it when you want the latest block
	// - Methods return multiple values: result, err := object.Method()
	//
	// Resources:
	// - BalanceAt: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Client.BalanceAt

	balance, err := // Complete the BalanceAt call here
	// Handle error
	fmt.Println("Current balance (wei):", balance)

	// TODO: Get balance at a specific block
	// First create a big.Int for block number 5532993
	// Then call BalanceAt with that block number
	//
	// Go syntax hint:
	// - big.NewInt(5532993) creates a new big integer
	// - You can reuse the same account variable
	//
	// Resources:
	// - big.NewInt: https://pkg.go.dev/math/big#NewInt

	blockNumber := // Create big.Int here
	balanceAt, err := // Get balance at specific block
	// Handle error
	fmt.Println("Balance at block", blockNumber, "(wei):", balanceAt)

	// TODO: Convert wei to ETH
	// 1 ETH = 10^18 wei
	// Steps:
	// 1. Create a new big.Float
	// 2. Convert balanceAt to string and set it
	// 3. Divide by 10^18 (use math.Pow10(18))
	//
	// Go syntax hint:
	// - new(big.Float) creates a new big float
	// - SetString() converts a string to big.Float
	// - Quo() divides two big.Float values
	//
	// Resources:
	// - big.Float: https://pkg.go.dev/math/big#Float
	// - big.Float.Quo: https://pkg.go.dev/math/big#Float.Quo
	// - math.Pow10: https://pkg.go.dev/math#Pow10
     //convert eth to wei
	fbalance := new(/*something*/)
	fbalance.SetString(/*something*/)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance in ETH:", /* var */)

	// TODO: Get pending balance
	// Use client.PendingBalanceAt() method
	// Similar to BalanceAt but doesn't need block number
	//
	// Resources:
	// - PendingBalanceAt: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Client.PendingBalanceAt

	pendingBalance, err := // Get pending balance here
	// Handle error
	fmt.Println("Pending balance (wei):", pendingBalance)
}
