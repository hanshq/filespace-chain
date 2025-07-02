package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/keeper"
)

func TestSetGetProviderStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	amount := sdk.NewCoin("utoken", math.NewInt(5000))
	height := uint64(100)
	
	// Set provider stake
	k.SetProviderStake(ctx, provider, amount, height)
	
	// Get provider stake
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	require.Equal(t, provider, stake.Provider)
	require.Equal(t, amount, stake.Amount)
	require.Equal(t, height, stake.Height)
}

func TestGetProviderStakeNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Try to get non-existent provider stake
	_, found := k.GetProviderStake(ctx, "cosmos1nonexistent")
	require.False(t, found)
}

func TestRemoveProviderStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	amount := sdk.NewCoin("utoken", math.NewInt(5000))
	height := uint64(100)
	
	// Set and then remove provider stake
	k.SetProviderStake(ctx, provider, amount, height)
	
	// Verify it exists
	_, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	
	// Remove it
	k.RemoveProviderStake(ctx, provider)
	
	// Verify it's gone
	_, found = k.GetProviderStake(ctx, provider)
	require.False(t, found)
}

func TestGetAllProviderStakes(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create multiple provider stakes
	stakes := []struct {
		provider string
		amount   sdk.Coin
		height   uint64
	}{
		{"cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100},
		{"cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 200},
		{"cosmos1provider3", sdk.NewCoin("stake", math.NewInt(3000)), 300},
	}
	
	for _, stake := range stakes {
		k.SetProviderStake(ctx, stake.provider, stake.amount, stake.height)
	}
	
	// Get all stakes
	allStakes := k.GetAllProviderStakes(ctx)
	require.Len(t, allStakes, 3)
	
	// Verify each stake exists in the results
	for _, expectedStake := range stakes {
		found := false
		for _, actualStake := range allStakes {
			if actualStake.Provider == expectedStake.provider {
				require.Equal(t, expectedStake.amount, actualStake.Amount)
				require.Equal(t, expectedStake.height, actualStake.Height)
				found = true
				break
			}
		}
		require.True(t, found, "Stake for provider %s not found", expectedStake.provider)
	}
}

func TestUpdateProviderStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	originalAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	newAmount := sdk.NewCoin("utoken", math.NewInt(2000))
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, originalAmount, height)
	
	// Update stake
	err := k.UpdateProviderStake(ctx, provider, newAmount)
	require.NoError(t, err)
	
	// Verify update
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	require.Equal(t, newAmount, stake.Amount)
	require.Equal(t, height, stake.Height) // Height should remain unchanged
}

func TestUpdateProviderStakeNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	newAmount := sdk.NewCoin("utoken", math.NewInt(2000))
	
	// Try to update non-existent provider stake
	err := k.UpdateProviderStake(ctx, "cosmos1nonexistent", newAmount)
	require.Error(t, err)
	require.Contains(t, err.Error(), "provider stake not found")
}

func TestGetProvidersByMinStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create providers with different stake amounts
	k.SetProviderStake(ctx, "cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100)
	k.SetProviderStake(ctx, "cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 200)
	k.SetProviderStake(ctx, "cosmos1provider3", sdk.NewCoin("utoken", math.NewInt(3000)), 300)
	k.SetProviderStake(ctx, "cosmos1provider4", sdk.NewCoin("stake", math.NewInt(1500)), 400) // Different denom
	
	// Test with minimum stake of 1500 utoken
	minStake := sdk.NewCoin("utoken", math.NewInt(1500))
	qualifiedProviders := k.GetProvidersByMinStake(ctx, minStake)
	
	// Should return providers 2 and 3 (amounts 2000 and 3000)
	require.Len(t, qualifiedProviders, 2)
	
	providerNames := make([]string, len(qualifiedProviders))
	for i, provider := range qualifiedProviders {
		providerNames[i] = provider.Provider
	}
	require.Contains(t, providerNames, "cosmos1provider2")
	require.Contains(t, providerNames, "cosmos1provider3")
}

func TestGetTotalStakedAmount(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create stakes with different denominations
	k.SetProviderStake(ctx, "cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100)
	k.SetProviderStake(ctx, "cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 200)
	k.SetProviderStake(ctx, "cosmos1provider3", sdk.NewCoin("utoken", math.NewInt(3000)), 300)
	k.SetProviderStake(ctx, "cosmos1provider4", sdk.NewCoin("stake", math.NewInt(5000)), 400)
	
	// Test total for utoken
	totalUtoken := k.GetTotalStakedAmount(ctx, "utoken")
	expectedUtoken := sdk.NewCoin("utoken", math.NewInt(6000))
	require.Equal(t, expectedUtoken, totalUtoken)
	
	// Test total for stake
	totalStake := k.GetTotalStakedAmount(ctx, "stake")
	expectedStake := sdk.NewCoin("stake", math.NewInt(5000))
	require.Equal(t, expectedStake, totalStake)
	
	// Test total for non-existent denomination
	totalNonExistent := k.GetTotalStakedAmount(ctx, "nonexistent")
	expectedNonExistent := sdk.NewCoin("nonexistent", math.ZeroInt())
	require.Equal(t, expectedNonExistent, totalNonExistent)
}

func TestIncrementProviderStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	initialAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	increment := sdk.NewCoin("utoken", math.NewInt(500))
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, initialAmount, height)
	
	// Increment stake
	err := k.IncrementProviderStake(ctx, provider, increment)
	require.NoError(t, err)
	
	// Verify increment
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	expectedAmount := sdk.NewCoin("utoken", math.NewInt(1500))
	require.Equal(t, expectedAmount, stake.Amount)
}

func TestIncrementProviderStakeNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	increment := sdk.NewCoin("utoken", math.NewInt(500))
	
	// Try to increment non-existent provider stake
	err := k.IncrementProviderStake(ctx, "cosmos1nonexistent", increment)
	require.Error(t, err)
	require.Contains(t, err.Error(), "provider stake not found")
}

func TestIncrementProviderStakeDenomMismatch(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	initialAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	increment := sdk.NewCoin("stake", math.NewInt(500)) // Different denom
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, initialAmount, height)
	
	// Try to increment with different denomination
	err := k.IncrementProviderStake(ctx, provider, increment)
	require.Error(t, err)
	require.Contains(t, err.Error(), "denomination mismatch")
}

func TestDecrementProviderStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	initialAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	decrement := sdk.NewCoin("utoken", math.NewInt(300))
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, initialAmount, height)
	
	// Decrement stake
	err := k.DecrementProviderStake(ctx, provider, decrement)
	require.NoError(t, err)
	
	// Verify decrement
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	expectedAmount := sdk.NewCoin("utoken", math.NewInt(700))
	require.Equal(t, expectedAmount, stake.Amount)
}

func TestDecrementProviderStakeToZero(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	initialAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	decrement := sdk.NewCoin("utoken", math.NewInt(1000)) // Full amount
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, initialAmount, height)
	
	// Decrement to zero
	err := k.DecrementProviderStake(ctx, provider, decrement)
	require.NoError(t, err)
	
	// Verify stake is removed
	_, found := k.GetProviderStake(ctx, provider)
	require.False(t, found)
}

func TestDecrementProviderStakeInsufficientFunds(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	initialAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	decrement := sdk.NewCoin("utoken", math.NewInt(1500)) // More than available
	height := uint64(100)
	
	// Set initial stake
	k.SetProviderStake(ctx, provider, initialAmount, height)
	
	// Try to decrement more than available
	err := k.DecrementProviderStake(ctx, provider, decrement)
	require.Error(t, err)
	require.Contains(t, err.Error(), "insufficient stake")
}

func TestGetProvidersStakedInRange(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create stakes at different heights
	k.SetProviderStake(ctx, "cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100)
	k.SetProviderStake(ctx, "cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 150)
	k.SetProviderStake(ctx, "cosmos1provider3", sdk.NewCoin("utoken", math.NewInt(3000)), 200)
	k.SetProviderStake(ctx, "cosmos1provider4", sdk.NewCoin("utoken", math.NewInt(4000)), 250)
	
	// Get stakes in range 120-220
	rangeStakes := k.GetProvidersStakedInRange(ctx, 120, 220)
	require.Len(t, rangeStakes, 2)
	
	// Should include providers 2 and 3 (heights 150 and 200)
	providerNames := make([]string, len(rangeStakes))
	for i, provider := range rangeStakes {
		providerNames[i] = provider.Provider
	}
	require.Contains(t, providerNames, "cosmos1provider2")
	require.Contains(t, providerNames, "cosmos1provider3")
}

func TestGetProviderStakeStats(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create stakes with utoken denomination
	k.SetProviderStake(ctx, "cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100)
	k.SetProviderStake(ctx, "cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 200)
	k.SetProviderStake(ctx, "cosmos1provider3", sdk.NewCoin("utoken", math.NewInt(3000)), 300)
	k.SetProviderStake(ctx, "cosmos1provider4", sdk.NewCoin("stake", math.NewInt(5000)), 400) // Different denom
	
	// Get stats for utoken
	stats := k.GetProviderStakeStats(ctx, "utoken")
	
	require.Equal(t, 3, stats["count"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(6000)), stats["total"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(2000)), stats["average"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(1000)), stats["min"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(3000)), stats["max"])
}

func TestGetProviderStakeStatsEmpty(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Get stats for non-existent denomination
	stats := k.GetProviderStakeStats(ctx, "nonexistent")
	
	require.Equal(t, 0, stats["count"])
	require.Equal(t, sdk.NewCoin("nonexistent", math.ZeroInt()), stats["total"])
	require.Equal(t, sdk.NewCoin("nonexistent", math.ZeroInt()), stats["average"])
	require.Equal(t, sdk.NewCoin("nonexistent", math.ZeroInt()), stats["min"])
	require.Equal(t, sdk.NewCoin("nonexistent", math.ZeroInt()), stats["max"])
}

func TestProviderStakeStructure(t *testing.T) {
	// Test the ProviderStake structure
	stake := keeper.ProviderStake{
		Provider: "cosmos1test",
		Amount:   sdk.NewCoin("utoken", math.NewInt(1000)),
		Height:   uint64(123),
	}
	
	require.Equal(t, "cosmos1test", stake.Provider)
	require.Equal(t, "utoken", stake.Amount.Denom)
	require.Equal(t, math.NewInt(1000), stake.Amount.Amount)
	require.Equal(t, uint64(123), stake.Height)
}

func TestProviderStakeKey(t *testing.T) {
	// Test the provider stake key generation
	provider := "cosmos1testprovider"
	key := keeper.ProviderStakeKey(provider)
	
	require.NotNil(t, key)
	require.True(t, len(key) > 0)
	
	// Keys for different providers should be different
	key2 := keeper.ProviderStakeKey("cosmos1otherprovider")
	require.NotEqual(t, key, key2)
	
	// Keys for same provider should be same
	key3 := keeper.ProviderStakeKey(provider)
	require.Equal(t, key, key3)
}