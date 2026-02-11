# Setup Instructions

## Prerequisites
- Go 1.24+ installed ([Download](https://go.dev/dl/))
- An Alchemy API key ([Get one here](https://www.alchemy.com/))

## Initial Setup

1. **Navigate to this directory:**
   ```bash
   cd clients-and-accounts
   ```

2. **Initialize Go module:**
   ```bash
   go mod init geth-workshop/clients-and-accounts
   ```

3. **Install dependencies:**
   ```bash
   go get github.com/ethereum/go-ethereum
   ```

4. **For wallet generation, also install:**
   ```bash
   go get golang.org/x/crypto/sha3
   ```

5. **Clean up dependencies (optional but recommended):**
   ```bash
   go mod tidy
   ```

## Running the Exercises

Each file can be run independently:

```bash
go run client.go
go run address.go
go run account_balance.go
go run wallet.go
```

**Note:** Make sure to:
- Uncomment the import statements in each file
- Replace `YOUR_ALCHEMY_RPC_URL_HERE` with your actual Alchemy RPC URL
- Replace `YOUR_ADDRESS_HERE` with a valid Ethereum address

## Getting Your Alchemy RPC URL

1. Sign up at [Alchemy](https://www.alchemy.com/)
2. Create a new app
3. Copy your HTTP URL (format: `https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY`)

## Troubleshooting

- **"go: cannot find main module"**: Run `go mod init geth-workshop/clients-and-accounts` first
- **"package not found"**: Run `go get github.com/ethereum/go-ethereum` (make sure you're in the `clients-and-accounts` directory)
- **Multiple main functions error**: Make sure to run files individually with `go run filename.go`
