# Requirements Specification - Hosting Provider Reward System

## Problem Statement

The Filespace Chain currently lacks any payment or reward mechanisms for hosting providers. While the blockchain supports hosting inquiries, offers, and contracts, providers receive no compensation for storage services, and inquiry creators' escrow amounts are not locked or distributed. This prevents the platform from functioning as a viable storage marketplace.

## Solution Overview

Implement a comprehensive reward system that:
- Locks escrow funds when hosting inquiries are created
- Distributes 50% of payments periodically during active storage
- Distributes 50% as completion bonus when contracts finish successfully
- Requires provider staking to prevent spam and enable slashing
- Uses platform-wide pricing based on file size, duration, and replication factor

## Functional Requirements

### FR1: Escrow Management
- **FR1.1**: Lock escrow funds immediately when hosting inquiry is created
- **FR1.2**: Validate sufficient balance before allowing inquiry creation
- **FR1.3**: Calculate total escrow based on: `file_size × duration_blocks × replication_factor × base_price_per_byte_per_block`
- **FR1.4**: Refund unused escrow if inquiry expires without contracts
- **FR1.5**: Hold escrowed funds in dedicated module account separate from staking

### FR2: Provider Staking
- **FR2.1**: Require providers to stake tokens before creating hosting offers
- **FR2.2**: Validate minimum stake amount during offer creation
- **FR2.3**: Reject offer creation if provider has insufficient stake
- **FR2.4**: Hold staked funds in dedicated staking module account
- **FR2.5**: Enable slashing of staked amounts for service failures

### FR3: Payment Distribution
- **FR3.1**: Distribute 50% of contract value as periodic payments during active storage
- **FR3.2**: Distribute 50% as completion bonus when contract ends successfully
- **FR3.3**: Calculate payments per provider: `total_contract_value / replication_factor`
- **FR3.4**: Process periodic payments in BeginBlock hook each block
- **FR3.5**: Release completion bonus only when contract reaches end time

### FR4: Platform Pricing
- **FR4.1**: Maintain platform-wide base price per byte per block parameter
- **FR4.2**: Enable governance-based updates to pricing parameter
- **FR4.3**: Calculate contract costs as: `file_size × duration × replication × base_price`
- **FR4.4**: Support multiple denominations through Cosmos SDK Coin types

### FR5: Slashing Mechanism
- **FR5.1**: Enable slashing of provider stakes for service failures
- **FR5.2**: Define slashing conditions (implementation deferred - out of scope for MVP)
- **FR5.3**: Architecture must support future proof-of-storage integration
- **FR5.4**: Maintain slashing history for provider reputation

## Technical Requirements

### TR1: Keeper Architecture
- **File**: `x/filespacechain/keeper/keeper.go`
- **Change**: Add BankKeeper field to keeper struct
- **Change**: Update NewKeeper constructor to accept BankKeeper parameter
- **Change**: Implement escrow and payment methods using bank operations

### TR2: Bank Module Integration
- **File**: `x/filespacechain/types/expected_keepers.go`
- **Change**: Expand BankKeeper interface with required methods:
  - `SendCoinsFromAccountToModule`
  - `SendCoinsFromModuleToAccount`
  - `GetBalance`
  - `GetAllBalances`

### TR3: Module Account Configuration
- **File**: `app/app_config.go`
- **Change**: Add module account permissions:
  - `filespacechain` account for escrow funds
  - `hosting_bonded_pool` account for provider staking

### TR4: Parameter Management
- **File**: `x/filespacechain/types/params.go`
- **Change**: Add to Params struct:
  - `base_price_per_byte_per_block` (Cosmos SDK Dec type)
  - `min_provider_stake` (Cosmos SDK Coin type)
  - `slashing_fraction` (Cosmos SDK Dec type)

### TR5: Message Handler Updates
- **File**: `x/filespacechain/keeper/msg_server_hosting_inquiry.go`
- **Change**: Add escrow locking in CreateHostingInquiry
- **Validation**: Check sender balance before locking funds

- **File**: `x/filespacechain/keeper/msg_server_hosting_offer.go`
- **Change**: Add stake validation in CreateHostingOffer
- **Validation**: Verify provider has sufficient stake

### TR6: BeginBlock Processing
- **File**: `x/filespacechain/module/module.go`
- **Change**: Implement BeginBlock hook for periodic payments
- **Logic**: Process all active contracts and distribute 50% payments

### TR7: New Keeper Methods Required
```go
// Escrow management
func (k Keeper) EscrowFunds(ctx context.Context, sender sdk.AccAddress, amount sdk.Coin) error
func (k Keeper) ReleaseFunds(ctx context.Context, recipient sdk.AccAddress, amount sdk.Coin) error
func (k Keeper) RefundFunds(ctx context.Context, recipient sdk.AccAddress, amount sdk.Coin) error

// Payment processing
func (k Keeper) ProcessPeriodicPayments(ctx context.Context) error
func (k Keeper) ProcessCompletionBonus(ctx context.Context, contractId uint64) error

// Provider staking
func (k Keeper) StakeForHosting(ctx context.Context, provider sdk.AccAddress, amount sdk.Coin) error
func (k Keeper) SlashProvider(ctx context.Context, provider sdk.AccAddress, fraction sdk.Dec) error
```

## Implementation Hints and Patterns

### Pattern 1: Module Account Operations
Follow existing Cosmos SDK staking module patterns:
```go
// Lock funds in module account
k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(amount))

// Release funds from module account
k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(amount))
```

### Pattern 2: Parameter Access
Follow existing parameter patterns in the module:
```go
params := k.GetParams(ctx)
basePrice := params.BasePricePerBytePerBlock
totalCost := fileSize.Mul(duration).Mul(replication).Mul(basePrice)
```

### Pattern 3: BeginBlock Processing
Follow Cosmos SDK module patterns:
```go
func (am AppModule) BeginBlock(ctx context.Context) error {
    return am.keeper.ProcessPeriodicPayments(ctx)
}
```

## Acceptance Criteria

### AC1: Escrow Functionality
- [ ] Creating hosting inquiry locks specified escrow amount
- [ ] Insufficient balance prevents inquiry creation
- [ ] Unused escrow refunds correctly when inquiry expires
- [ ] Escrowed funds held in correct module account

### AC2: Provider Staking
- [ ] Providers must stake before creating offers
- [ ] Insufficient stake prevents offer creation
- [ ] Staked funds held in dedicated staking account
- [ ] Slashing reduces provider stake amounts

### AC3: Payment Distribution
- [ ] Active contracts receive 50% periodic payments per block
- [ ] Completed contracts receive 50% completion bonus
- [ ] Multiple providers receive equal shares per replication
- [ ] Payment calculations match: file_size × duration × replication × base_price

### AC4: Parameter Management
- [ ] Base price parameter exists and can be queried
- [ ] Governance can update pricing parameters
- [ ] Parameter validation prevents invalid values
- [ ] Default parameters set appropriate values

### AC5: Integration Testing
- [ ] End-to-end flow: inquiry → escrow → contract → periodic payments → completion bonus
- [ ] Multiple concurrent contracts process payments correctly
- [ ] Provider staking integrates with offer creation
- [ ] Bank module operations complete successfully

## Assumptions

### A1: Trust-Based MVP
For the initial implementation, assume providers are honest and actually store files. Proof-of-storage verification will be added in a future version.

### A2: Simple Slashing
Initially, slashing will be manually triggered rather than automated based on service failures. Automated monitoring will be added later.

### A3: Single Token
The system will initially support a single native token denomination. Multi-token support can be added later if needed.

### A4: Fixed Pricing
The base price will be a simple per-byte-per-block rate. More sophisticated pricing models (e.g., regional pricing, quality tiers) can be added later.

### A5: No Partial Refunds
If a contract is partially completed, the system will not implement partial refunds in the MVP. Either the full completion bonus is paid or none at all.

## Out of Scope

- Proof-of-storage verification mechanisms
- Automated service level monitoring
- Complex pricing models (regional, tiered, etc.)
- Multi-token denomination support
- Provider reputation systems
- Automated slashing based on performance metrics
- Partial completion and prorated payments