package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProviderStake tracks staked funds for a hosting provider
type ProviderStake struct {
	Provider string   `json:"provider"`
	Amount   sdk.Coin `json:"amount"`
	Height   uint64   `json:"height"` // Block height when stake was created
}

var (
	ProviderStakeKeyPrefix = []byte{0x02} // Prefix for provider stake records
)

// ProviderStakeKey returns the store key for provider stake records
func ProviderStakeKey(provider string) []byte {
	key := make([]byte, len(ProviderStakeKeyPrefix)+len(provider))
	copy(key, ProviderStakeKeyPrefix)
	copy(key[len(ProviderStakeKeyPrefix):], []byte(provider))
	return key
}

// SetProviderStake stores a provider stake record
func (k Keeper) SetProviderStake(ctx context.Context, provider string, amount sdk.Coin, height uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	stake := ProviderStake{
		Provider: provider,
		Amount:   amount,
		Height:   height,
	}
	
	bz, err := json.Marshal(stake)
	if err != nil {
		panic(err)
	}
	storeAdapter.Set(ProviderStakeKey(provider), bz)
}

// GetProviderStake retrieves a provider stake record
func (k Keeper) GetProviderStake(ctx context.Context, provider string) (ProviderStake, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	bz := storeAdapter.Get(ProviderStakeKey(provider))
	if bz == nil {
		return ProviderStake{}, false
	}
	
	var stake ProviderStake
	err := json.Unmarshal(bz, &stake)
	if err != nil {
		return ProviderStake{}, false
	}
	return stake, true
}

// RemoveProviderStake removes a provider stake record
func (k Keeper) RemoveProviderStake(ctx context.Context, provider string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storeAdapter.Delete(ProviderStakeKey(provider))
}

// GetAllProviderStakes returns all provider stake records
func (k Keeper) GetAllProviderStakes(ctx context.Context) []ProviderStake {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	var stakes []ProviderStake
	iterator := storetypes.KVStorePrefixIterator(storeAdapter, ProviderStakeKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var stake ProviderStake
		err := json.Unmarshal(iterator.Value(), &stake)
		if err != nil {
			continue // Skip invalid records
		}
		stakes = append(stakes, stake)
	}
	
	return stakes
}

// UpdateProviderStake updates an existing provider stake amount
func (k Keeper) UpdateProviderStake(ctx context.Context, provider string, newAmount sdk.Coin) error {
	stake, found := k.GetProviderStake(ctx, provider)
	if !found {
		return fmt.Errorf("provider stake not found")
	}
	
	// Keep original height, don't update it
	k.SetProviderStake(ctx, provider, newAmount, stake.Height)
	return nil
}

// GetProvidersByMinStake returns all providers with stake above the minimum threshold
func (k Keeper) GetProvidersByMinStake(ctx context.Context, minAmount sdk.Coin) []ProviderStake {
	allStakes := k.GetAllProviderStakes(ctx)
	var qualifiedStakes []ProviderStake
	
	for _, stake := range allStakes {
		if stake.Amount.IsGTE(minAmount) {
			qualifiedStakes = append(qualifiedStakes, stake)
		}
	}
	
	return qualifiedStakes
}

// GetTotalStakedAmount returns the total amount staked across all providers for a specific denomination
func (k Keeper) GetTotalStakedAmount(ctx context.Context, denom string) sdk.Coin {
	allStakes := k.GetAllProviderStakes(ctx)
	totalAmount := math.ZeroInt()
	
	for _, stake := range allStakes {
		if stake.Amount.Denom == denom {
			totalAmount = totalAmount.Add(stake.Amount.Amount)
		}
	}
	
	return sdk.NewCoin(denom, totalAmount)
}

// IncrementProviderStake increases a provider's stake by the specified amount
func (k Keeper) IncrementProviderStake(ctx context.Context, provider string, increment sdk.Coin) error {
	stake, found := k.GetProviderStake(ctx, provider)
	if !found {
		return fmt.Errorf("provider stake not found for provider %s", provider)
	}
	
	if stake.Amount.Denom != increment.Denom {
		return fmt.Errorf("denomination mismatch: stake has %s, increment has %s", 
			stake.Amount.Denom, increment.Denom)
	}
	
	newAmount := stake.Amount.Add(increment)
	return k.UpdateProviderStake(ctx, provider, newAmount)
}

// DecrementProviderStake decreases a provider's stake by the specified amount
func (k Keeper) DecrementProviderStake(ctx context.Context, provider string, decrement sdk.Coin) error {
	stake, found := k.GetProviderStake(ctx, provider)
	if !found {
		return fmt.Errorf("provider stake not found for provider %s", provider)
	}
	
	if stake.Amount.Denom != decrement.Denom {
		return fmt.Errorf("denomination mismatch: stake has %s, decrement has %s", 
			stake.Amount.Denom, decrement.Denom)
	}
	
	if stake.Amount.IsLT(decrement) {
		return fmt.Errorf("insufficient stake: current %s, requested decrement %s", 
			stake.Amount.String(), decrement.String())
	}
	
	newAmount := stake.Amount.Sub(decrement)
	
	// If stake becomes zero, remove the record entirely
	if newAmount.IsZero() {
		k.RemoveProviderStake(ctx, provider)
		return nil
	}
	
	return k.UpdateProviderStake(ctx, provider, newAmount)
}

// GetProvidersStakedInRange returns providers with stakes within a specific block height range
func (k Keeper) GetProvidersStakedInRange(ctx context.Context, startHeight, endHeight uint64) []ProviderStake {
	allStakes := k.GetAllProviderStakes(ctx)
	var rangeStakes []ProviderStake
	
	for _, stake := range allStakes {
		if stake.Height >= startHeight && stake.Height <= endHeight {
			rangeStakes = append(rangeStakes, stake)
		}
	}
	
	return rangeStakes
}

// GetProviderStakeStats returns statistical information about provider stakes
func (k Keeper) GetProviderStakeStats(ctx context.Context, denom string) map[string]interface{} {
	allStakes := k.GetAllProviderStakes(ctx)
	stats := make(map[string]interface{})
	
	var denomStakes []ProviderStake
	for _, stake := range allStakes {
		if stake.Amount.Denom == denom {
			denomStakes = append(denomStakes, stake)
		}
	}
	
	if len(denomStakes) == 0 {
		stats["count"] = 0
		stats["total"] = sdk.NewCoin(denom, math.ZeroInt())
		stats["average"] = sdk.NewCoin(denom, math.ZeroInt())
		stats["min"] = sdk.NewCoin(denom, math.ZeroInt())
		stats["max"] = sdk.NewCoin(denom, math.ZeroInt())
		return stats
	}
	
	// Calculate statistics
	totalAmount := math.ZeroInt()
	minAmount := denomStakes[0].Amount.Amount
	maxAmount := denomStakes[0].Amount.Amount
	
	for _, stake := range denomStakes {
		totalAmount = totalAmount.Add(stake.Amount.Amount)
		if stake.Amount.Amount.LT(minAmount) {
			minAmount = stake.Amount.Amount
		}
		if stake.Amount.Amount.GT(maxAmount) {
			maxAmount = stake.Amount.Amount
		}
	}
	
	averageAmount := totalAmount.Quo(math.NewInt(int64(len(denomStakes))))
	
	stats["count"] = len(denomStakes)
	stats["total"] = sdk.NewCoin(denom, totalAmount)
	stats["average"] = sdk.NewCoin(denom, averageAmount)
	stats["min"] = sdk.NewCoin(denom, minAmount)
	stats["max"] = sdk.NewCoin(denom, maxAmount)
	
	return stats
}