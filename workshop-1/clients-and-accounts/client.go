package main

// TODO: Import the necessary packages
// You'll need:
// - "context" - for context management (used in Go for cancellation/timeouts)
// - "fmt" - for formatted printing
// - "log" - for logging errors
// - "github.com/ethereum/go-ethereum/ethclient" - the Ethereum client library
//
// Go syntax hint: Use the import keyword followed by parentheses for multiple imports
// Example: import ( "package1" "package2" )
//
// Resources:
// - Go packages: https://go.dev/learn/#packages
import (
	"fmt"
	// Add the ethclient import here
)

func main() {
	// TODO: Create an Ethereum client connection
	// ethclient.Dial() returns a client to connect to an Ethereum node
	// You'll need a connection URL - try using a public RPC endpoint
	// Example URL format: "https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY"
	//
	// Go syntax hint:
	// - Functions can return multiple values: result, err := function()
	// - The := operator declares and assigns in one step
	// - Check the error immediately after the function call
	//
	// Resources:
	// - ethclient.Dial documentation: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Dial
	// - Go error handling: https://go.dev/doc/effective_go#errors

	// Create client connection here
	// Fill in the URL - you'll need an Ethereum RPC endpoint
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")

	// TODO: Handle the error from Dial()
	// If err is not nil, use log.Fatal(err) to print and exit
	//
	// Go syntax hint: if err != nil { ... } is the standard error checking pattern
	// Resources:
	// - ethclient.Dial: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Dial

	client, err := // Complete the Dial() call here
	// Handle error here

	// TODO: Print a success message to confirm the connection
	// Use fmt.Println() to print a message like "we have a connection"
	//
	// Go syntax hint: fmt.Println("your message here")
	// Resources:
	// - fmt package: https://pkg.go.dev/fmt

	// Print message here

	// TODO: Get the latest block number
	// Use client.BlockNumber() method
	// This method requires a context - use context.Background() for a basic context
	//
	// Go syntax hint:
	// - calling methods: object.MethodName()
	// - context.Background() creates a new context, it's a basic context that can be used to pass to a function
	// - The method returns (unsigned 64 bit integer, error), handle both
	//
	// Resources:
	// - BlockNumber: https://pkg.go.dev/github.com/ethereum/go-ethereum/ethclient#Client.BlockNumber
	// - context.Background: https://pkg.go.dev/context#Background

	// Get block number here
	// Complete the method call - you need to pass a context
	blockNum, err := client.BlockNumber( /* pass context here */ ) // not required but good practice as it gives us a way to cancel the request if there is an issue

	// TODO: Handle the error from BlockNumber()
	// Use the same error handling pattern as before
	if err != nil {
		// log an error
	}

	// TODO: Print the block number
	// Use fmt.Printf() with format string %d for integer
	fmt.Printf("Latest Block Number: %d\n", /* add variable */)

	// check info about the block here https://etherscan.io/
}
