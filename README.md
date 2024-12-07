# Ethereum Transaction Parser

## Overview

The Ethereum Transaction Parser is a Go-based application designed to parse and store Ethereum transactions. It includes functionalities for subscribing to addresses, validating transactions, and fetching the latest block numbers from the Ethereum blockchain.

## Features

- Subscribe to Ethereum addresses
- Fetch the latest block numbers
- API endpoints for current block, subscriptions, and transactions

## Setup

### Prerequisites

- Go 1.23.4 or later
- Docker (optional, for containerized deployment)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/KRR19/EthereumParser.git
    cd EthereumParser
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Build the project:
    ```sh
    make build
    ```

### Running the Application

#### Using Go

1. Run the application:
    ```sh
    make run
    ```

#### Using Docker

1. Build and run the Docker container:
    ```sh
    make docker
    ```

### API Endpoints

- **Get Current Block**
  - **URL:** `/api/v1/block`
  - **Method:** `GET`
  - **Description:** Fetches the current block number.

- **Subscribe to Address**
  - **URL:** `/api/v1/subscribe`
  - **Method:** `POST`
  - **Description:** Subscribes to an Ethereum address.
  - **Request Body:**
    ```json
    {
      "address": "0xYourEthereumAddress"
    }
    ```

- **Get Transactions**
  - **URL:** `/api/v1/transactions`
  - **Method:** `GET`
  - **Description:** Fetches transactions for a subscribed address.
  - **Query Parameter:** `address`

## Project Structure

- `cmd/txparser`: Main application entry point
- `internal/core`: Core business logic
- `internal/infrastructure`: Infrastructure-related code (e.g., API handlers, Ethereum client, logger)
- `internal/models`: Data models
- `pkg/hex`: Utility functions for hex conversion

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
