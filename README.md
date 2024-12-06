# Ethereum Transaction Parser

A Go application that parses Ethereum blockchain transactions and allows subscribing to specific addresses to monitor their transactions.

## Features

- Monitor the Ethereum blockchain for new blocks and transactions
- Subscribe to specific Ethereum addresses
- Retrieve transaction history for subscribed addresses
- Concurrent processing of blocks and transactions
- Configurable settings via YAML configuration

## Requirements

- Go 1.21 or higher
- Access to an Ethereum JSON-RPC endpoint

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ethereumparser.git
cd ethereumparser
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The application can be configured using a YAML file. Default configuration is in `configs/config.yaml`:

```yaml
ethereum_rpc: "https://ethereum-rpc.publicnode.com"
poll_interval: "15s"
worker_count: 4
```

- `ethereum_rpc`: Ethereum JSON-RPC endpoint URL
- `poll_interval`: Interval between checking for new blocks
- `worker_count`: Number of concurrent workers for processing blocks

## Building

```bash
go build -o eth-parser ./cmd/parser
```

## Running

```bash
./eth-parser -config configs/config.yaml
```

## Project Structure

```
.
├── cmd/
│   └── parser/           # Application entry point
├── internal/
│   ├── parser/          # Core parser implementation
│   ├── subscriptions/   # Subscription management
│   └── transactions/    # Transaction processing
├── pkg/
│   ├── ethereum/        # Ethereum RPC client
│   └── logger/          # Logging utilities
├── configs/             # Configuration files
├── docs/               # Documentation
└── scripts/            # Build and utility scripts
```

## Usage Example

```go
package main

import (
    "github.com/yourusername/ethereumparser/internal/parser"
)

func main() {
    config := &parser.Config{
        EthereumRPC:  "https://ethereum-rpc.publicnode.com",
        PollInterval: 15 * time.Second,
        WorkerCount:  4,
    }

    p := parser.NewEthParser(config)
    p.Start()

    // Subscribe to an address
    p.Subscribe("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")

    // Get transactions for the address
    txs := p.GetTransactions("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
    for _, tx := range txs {
        fmt.Printf("Transaction: %s\n", tx.Hash)
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.