# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Cosmos SDK-based blockchain called Filespace Chain, designed for file storage and hosting services. The chain implements custom modules for file entries, hosting inquiries, hosting offers, and hosting contracts.

**Built with Ignite CLI**: This project was scaffolded and is managed using the Ignite CLI, a powerful tool for creating sovereign blockchains built with Cosmos SDK.

## Development Commands

### Ignite CLI Commands (Preferred)

#### Building and Testing
- `ignite chain build` - Build the blockchain node binary (preferred over go build)
- `ignite chain serve` - Build, initialize, and start blockchain locally with hot-reload for development
- `ignite chain init` - Build binary and initialize a local validator node
- `ignite chain lint` - Lint codebase using golangci-lint
- `ignite chain simulate` - Run simulation testing for the blockchain

#### Development Tools
- `ignite chain faucet [address] [amount]` - Send tokens to an address from faucet account
- `ignite chain debug` - Launch debugger for blockchain app

#### Code Generation
- `ignite generate proto-go` - Compile protocol buffer files to Go source code
- `ignite generate openapi` - Generate OpenAPI spec for your chain
- `ignite generate composables` - Generate TypeScript frontend client and Vue 3 composables
- `ignite generate hooks` - Generate TypeScript frontend client and React hooks

#### Scaffolding New Components
- `ignite scaffold module [name]` - Create a new custom Cosmos SDK module
- `ignite scaffold message [name] [field:type]...` - Create a new message type for state transitions
- `ignite scaffold query [name] [field:type]...` - Create a new query for fetching data
- `ignite scaffold list [name] [field:type]...` - Create CRUD operations for array-stored data
- `ignite scaffold map [name] [field:type]...` - Create CRUD operations for key-value stored data
- `ignite scaffold single [name] [field:type]...` - Create CRUD operations for single-instance data
- `ignite scaffold type [name] [field:type]...` - Create a new protocol buffer type definition
- `ignite scaffold params [name]` - Create parameters for a custom module
- `ignite scaffold packet [name] [field:type]...` - Create IBC packet for inter-blockchain communication

#### Examples of Scaffolding Commands Used in This Project
```bash
# This project was likely created with:
ignite scaffold chain filespace-chain

# Module scaffolding (example):
ignite scaffold module filespacechain

# Message scaffolding examples:
ignite scaffold message create-file-entry cid replication-factor size creator
ignite scaffold message create-hosting-inquiry file-entry-cid replication-rate escrow-amount end-time creator
ignite scaffold message create-hosting-offer region price-per-block creator
ignite scaffold message create-hosting-contract inquiry-id offer-id creator

# Query scaffolding examples:
ignite scaffold query file-entry id:uint
ignite scaffold query hosting-inquiry id:uint
ignite scaffold query hosting-contract id:uint
ignite scaffold query hosting-offer id:uint
```

### Alternative Go Commands
- `go build ./cmd/filespace-chaind` - Build the blockchain daemon directly
- `go run ./cmd/filespace-chaind/main.go` - Run the daemon directly
- `go test ./...` - Run all tests
- `go mod tidy` - Clean up Go modules

### Docker Operations
- `./scripts/src/docker_push.sh <version>` - Build, run, and push Docker image with version tag
- `sudo docker build -t hanshq/filespace-chain:<tag> ./` - Build Docker image
- `sudo docker run -p 26656:26656 -p 26657:26657 -p 20080:1317 -p 29090:9090 -p 29091:9091 --name filespace-<tag> hanshq/filespace-chain:<tag>` - Run container with proper port mappings

### Blockchain Operations
- `filespace-chaind keys add owner --recover` - Add owner key (recovery mode)
- `./scripts/src/enable_validator.sh` - Enable validator functionality
- Use SDL files in `./data/deploy/` for Akash deployments

## Architecture

### Core Structure
- **Main binary**: `cmd/filespace-chaind/` - Contains the main blockchain daemon
- **App module**: `app/` - Core application setup, dependency injection, and module configuration
- **Custom module**: `x/filespacechain/` - Business logic for file storage and hosting

### Custom Blockchain Module (`x/filespacechain/`)
The module implements core entities and a hosting provider reward system:

#### Core Entities:
1. **FileEntry** - Represents files with CID, metadata, and size information
2. **HostingInquiry** - Requests for file hosting services with escrow locking
3. **HostingOffer** - Responses to hosting inquiries with terms and provider staking
4. **HostingContract** - Agreements between inquiry and offer parties with payment distribution

#### Hosting Provider Reward System:
- **Escrow Management**: Automatic locking of funds when inquiries are created
- **Provider Staking**: Providers must stake minimum amounts before creating offers
- **Periodic Payments**: 50% of escrowed funds distributed proportionally over contract duration
- **Completion Bonuses**: Remaining 50% paid when contracts complete successfully
- **Payment History**: Complete tracking of all payments per contract
- **Automatic Processing**: BeginBlock hooks handle payments and contract expiration

#### Module Structure
- `keeper/` - State management and business logic
  - `escrow.go` - Escrow fund management and state tracking
  - `staking.go` - Provider staking functionality
  - `payment_history.go` - Payment tracking and query helpers
  - `hosting_contract.go` - Contract lifecycle and completion logic
  - `keeper.go` - Core keeper with payment processing methods
- `types/` - Message types, parameters, and validation
- `module/` - Module definition, genesis handling, and BeginBlock payment processing
- `simulation/` - Simulation functions for testing

#### Key Parameters
- `base_price_per_byte_per_block` - Base pricing for file storage per byte per block
- `min_provider_stake` - Minimum stake required for hosting providers
- `slashing_fraction` - Fraction of stake slashed for provider service failures

### Protocol Buffers
- **Location**: `proto/filespacechain/filespacechain/`
- **Generated code**: `api/filespacechain/filespacechain/` (Pulsar) and `x/filespacechain/types/` (standard protobuf)
- **Build config**: `buf.*.yaml` files configure code generation

### Key Patterns
- Standard Cosmos SDK keeper pattern for state management
- Protobuf for all message definitions and state
- Auto-incremented IDs for all entities
- Creator field tracking for all user-generated content

## Testing
- Unit tests are co-located with source files (`*_test.go`)
- Integration tests in `testutil/`
- Simulation tests in `x/filespacechain/simulation/`

## Deployment
- Akash network deployment configurations in `data/deploy/`
- Genesis configuration in `data/genesis/`
- Docker-based deployment with environment variable configuration