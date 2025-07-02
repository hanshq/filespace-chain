# Context Findings - Hosting Provider Reward System

## Files That Need Modification

Based on the codebase analysis, the following files will need modifications:

### Core Keeper Changes
- `x/filespacechain/keeper/keeper.go` - Add BankKeeper and StakingKeeper dependencies
- `x/filespacechain/types/expected_keepers.go` - Expand BankKeeper interface and add StakingKeeper interface
- `x/filespacechain/module/module.go` - Add keeper dependencies and BeginBlock/EndBlock logic

### App Configuration
- `app/app_config.go` - Add module account permissions for escrow and staking pools
- `app/app.go` - Update keeper initialization with additional dependencies

### Message Handlers
- `x/filespacechain/keeper/msg_server_hosting_inquiry.go` - Add escrow locking on creation
- `x/filespacechain/keeper/msg_server_hosting_contract.go` - Add payment distribution logic

### New Components Needed
- New escrow and staking management functions in keeper
- New message types for staking operations
- BeginBlock/EndBlock hooks for periodic payments and slashing
- Parameter updates for pricing and staking requirements

## Exact Patterns to Follow

### 1. BankKeeper Integration Pattern
**File**: `x/filespacechain/types/expected_keepers.go`

**Current minimal interface**:
```go
type BankKeeper interface {
    SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
}
```

**Required expansion** (following Cosmos SDK patterns):
```go
type BankKeeper interface {
    SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
    SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
    SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
    GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
    GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}
```

### 2. Module Account Pattern
**File**: `app/app_config.go`

**Current structure** has validator staking pools:
```go
{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
```

**Pattern to add** for hosting provider escrow/staking:
```go
{Account: filespacechaintypes.ModuleName}, // For escrow
{Account: "hosting_bonded_pool", Permissions: []string{authtypes.Burner, "hosting"}}, // For provider staking
```

### 3. Keeper Initialization Pattern
**File**: `x/filespacechain/keeper/keeper.go`

**Current structure**:
```go
type Keeper struct {
    cdc          codec.BinaryCodec
    storeService store.KVStoreService
    logger       log.Logger
    authority    string
}
```

**Required expansion**:
```go
type Keeper struct {
    cdc          codec.BinaryCodec
    storeService store.KVStoreService
    logger       log.Logger
    authority    string
    
    accountKeeper types.AccountKeeper
    bankKeeper    types.BankKeeper
    stakingKeeper types.StakingKeeper // For hosting provider staking
}
```

### 4. BeginBlock/EndBlock Pattern
**File**: `x/filespacechain/module/module.go`

**Current empty hooks**:
```go
func (am AppModule) BeginBlock(_ context.Context) error {
    return nil
}

func (am AppModule) EndBlock(_ context.Context) error {
    return nil
}
```

**Pattern for periodic processing**:
```go
func (am AppModule) BeginBlock(ctx context.Context) error {
    return am.keeper.ProcessPeriodicPayments(ctx)
}

func (am AppModule) EndBlock(ctx context.Context) error {
    return am.keeper.ProcessSlashing(ctx)
}
```

## Similar Features Analyzed

### 1. Existing Coin Handling
The codebase already uses `sdk.Coin` in:
- **HostingInquiry.EscrowAmount** - Field exists but not used for actual escrow
- **HostingOffer.PricePerBlock** - Pricing field defined in protobuf

### 2. Module Account Usage
Standard Cosmos SDK staking module provides examples:
- **Bonded Pool**: `stakingtypes.BondedPoolName` with burner permissions
- **Not Bonded Pool**: `stakingtypes.NotBondedPoolName` with burner permissions
- **Pattern**: Module accounts hold staked/escrowed tokens safely

### 3. Dependency Injection Pattern
**File**: `app/app.go` (line 242)
Shows how BankKeeper is injected via depinject framework - same pattern needed for additional keepers.

## Technical Constraints and Considerations

### 1. Cosmos SDK Version Compatibility
- Uses Cosmos SDK v0.50+ with dependency injection
- Must follow new keeper patterns with context.Context (not legacy sdk.Context)
- Module accounts managed through depinject configuration

### 2. Token Economics Design
- **Global Pricing**: Single parameter for base price per byte per block
- **Calculation**: `total_cost = file_size × duration × replication_factor × base_price`
- **Split Payments**: 50% periodic, 50% completion bonus
- **Escrow Flow**: Lock full amount upfront, release portions as earned

### 3. State Management
- **Escrow Tracking**: Need to track locked amounts per inquiry/contract
- **Staking Records**: Track provider stakes and slashing history
- **Payment Status**: Track periodic payment history and completion status

### 4. Integration Points
- **Bank Module**: For token transfers and escrow
- **Staking Module**: For understanding delegation patterns (not direct usage)
- **Distribution Module**: For potential reward distribution patterns
- **Gov Module**: For parameter governance of pricing and staking requirements

## Implementation Complexity Assessment

### High Priority (MVP)
1. **Escrow Integration** - Critical for preventing fund availability issues
2. **Basic Payment Distribution** - Core value proposition
3. **Platform Pricing Parameters** - Economic foundation

### Medium Priority
1. **Provider Staking** - Important for security but can start with trust-based
2. **Periodic Payment Processing** - Can start with manual triggers
3. **Completion Bonus Logic** - Enhances incentives

### Lower Priority (Future)
1. **Slashing Implementation** - Requires proof-of-storage verification
2. **Complex Stake Management** - Unbonding periods, delegation
3. **Governance Integration** - Parameter updates through governance