package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryEscrowSummary returns comprehensive escrow statistics
func (k Keeper) QueryEscrowSummary(ctx context.Context) (map[string]interface{}, error) {
	summary := make(map[string]interface{})
	
	allRecords := k.GetAllEscrowRecords(ctx)
	activeRecords := k.GetActiveEscrowRecords(ctx)
	
	summary["total_escrow_records"] = len(allRecords)
	summary["active_escrow_records"] = len(activeRecords)
	summary["inactive_escrow_records"] = len(allRecords) - len(activeRecords)
	
	// Calculate total amounts by denomination
	denomTotals := make(map[string]math.Int)
	activeDenomTotals := make(map[string]math.Int)
	
	for _, record := range allRecords {
		if _, exists := denomTotals[record.Amount.Denom]; !exists {
			denomTotals[record.Amount.Denom] = math.ZeroInt()
		}
		denomTotals[record.Amount.Denom] = denomTotals[record.Amount.Denom].Add(record.Amount.Amount)
	}
	
	for _, record := range activeRecords {
		if _, exists := activeDenomTotals[record.Amount.Denom]; !exists {
			activeDenomTotals[record.Amount.Denom] = math.ZeroInt()
		}
		activeDenomTotals[record.Amount.Denom] = activeDenomTotals[record.Amount.Denom].Add(record.Amount.Amount)
	}
	
	summary["total_amounts_by_denom"] = denomTotals
	summary["active_amounts_by_denom"] = activeDenomTotals
	
	return summary, nil
}

// QueryEscrowByCreator returns all escrow records for a specific creator
func (k Keeper) QueryEscrowByCreator(ctx context.Context, creator string) ([]EscrowRecord, error) {
	return k.GetEscrowRecordsByCreator(ctx, creator), nil
}

// QueryEscrowByInquiry returns escrow record for a specific inquiry
func (k Keeper) QueryEscrowByInquiry(ctx context.Context, inquiryId uint64) (EscrowRecord, error) {
	record, found := k.GetEscrowRecord(ctx, inquiryId)
	if !found {
		return EscrowRecord{}, fmt.Errorf("escrow record not found for inquiry %d", inquiryId)
	}
	return record, nil
}

// QueryActiveEscrow returns all currently active escrow records
func (k Keeper) QueryActiveEscrow(ctx context.Context) ([]EscrowRecord, error) {
	return k.GetActiveEscrowRecords(ctx), nil
}

// QueryEscrowTotalsByDenom returns total escrowed amounts for each denomination
func (k Keeper) QueryEscrowTotalsByDenom(ctx context.Context) (map[string]sdk.Coin, error) {
	allRecords := k.GetAllEscrowRecords(ctx)
	totals := make(map[string]sdk.Coin)
	
	for _, record := range allRecords {
		if existing, exists := totals[record.Amount.Denom]; exists {
			totals[record.Amount.Denom] = existing.Add(record.Amount)
		} else {
			totals[record.Amount.Denom] = record.Amount
		}
	}
	
	return totals, nil
}

// QueryStakingSummary returns comprehensive staking statistics
func (k Keeper) QueryStakingSummary(ctx context.Context) (map[string]interface{}, error) {
	summary := make(map[string]interface{})
	
	allStakes := k.GetAllProviderStakes(ctx)
	summary["total_providers"] = len(allStakes)
	
	if len(allStakes) == 0 {
		summary["total_staked_by_denom"] = make(map[string]sdk.Coin)
		return summary, nil
	}
	
	// Calculate totals by denomination
	denomTotals := make(map[string]math.Int)
	denomCounts := make(map[string]int)
	
	for _, stake := range allStakes {
		if _, exists := denomTotals[stake.Amount.Denom]; !exists {
			denomTotals[stake.Amount.Denom] = math.ZeroInt()
			denomCounts[stake.Amount.Denom] = 0
		}
		denomTotals[stake.Amount.Denom] = denomTotals[stake.Amount.Denom].Add(stake.Amount.Amount)
		denomCounts[stake.Amount.Denom]++
	}
	
	// Convert to coins and add statistics
	totalsByDenom := make(map[string]sdk.Coin)
	for denom, total := range denomTotals {
		totalsByDenom[denom] = sdk.NewCoin(denom, total)
	}
	
	summary["total_staked_by_denom"] = totalsByDenom
	summary["provider_count_by_denom"] = denomCounts
	
	return summary, nil
}

// QueryStakeByProvider returns stake information for a specific provider
func (k Keeper) QueryStakeByProvider(ctx context.Context, provider string) (ProviderStake, error) {
	stake, found := k.GetProviderStake(ctx, provider)
	if !found {
		return ProviderStake{}, fmt.Errorf("provider stake not found for %s", provider)
	}
	return stake, nil
}

// QueryProvidersWithMinStake returns all providers with stake above minimum threshold
func (k Keeper) QueryProvidersWithMinStake(ctx context.Context, minAmount sdk.Coin) ([]ProviderStake, error) {
	return k.GetProvidersByMinStake(ctx, minAmount), nil
}

// QueryProvidersByStakeRange returns providers staked within a specific block height range
func (k Keeper) QueryProvidersByStakeRange(ctx context.Context, startHeight, endHeight uint64) ([]ProviderStake, error) {
	return k.GetProvidersStakedInRange(ctx, startHeight, endHeight), nil
}

// QueryStakeStatistics returns detailed statistics for a specific denomination
func (k Keeper) QueryStakeStatistics(ctx context.Context, denom string) (map[string]interface{}, error) {
	return k.GetProviderStakeStats(ctx, denom), nil
}

// QueryTopProvidersByStake returns the top N providers by stake amount for a specific denomination
func (k Keeper) QueryTopProvidersByStake(ctx context.Context, denom string, limit int) ([]ProviderStake, error) {
	allStakes := k.GetAllProviderStakes(ctx)
	var denomStakes []ProviderStake
	
	// Filter by denomination
	for _, stake := range allStakes {
		if stake.Amount.Denom == denom {
			denomStakes = append(denomStakes, stake)
		}
	}
	
	// Sort by stake amount (descending) - simple bubble sort for small datasets
	for i := 0; i < len(denomStakes)-1; i++ {
		for j := 0; j < len(denomStakes)-i-1; j++ {
			if denomStakes[j].Amount.Amount.LT(denomStakes[j+1].Amount.Amount) {
				denomStakes[j], denomStakes[j+1] = denomStakes[j+1], denomStakes[j]
			}
		}
	}
	
	// Apply limit
	if limit > 0 && limit < len(denomStakes) {
		denomStakes = denomStakes[:limit]
	}
	
	return denomStakes, nil
}

// QueryProviderStakeHistory returns historical information about when providers staked
func (k Keeper) QueryProviderStakeHistory(ctx context.Context) (map[string]interface{}, error) {
	allStakes := k.GetAllProviderStakes(ctx)
	history := make(map[string]interface{})
	
	// Group by block height
	blockHeights := make(map[uint64][]ProviderStake)
	for _, stake := range allStakes {
		blockHeights[stake.Height] = append(blockHeights[stake.Height], stake)
	}
	
	history["stakes_by_block_height"] = blockHeights
	history["total_unique_heights"] = len(blockHeights)
	
	// Find earliest and latest staking heights
	if len(allStakes) > 0 {
		earliestHeight := allStakes[0].Height
		latestHeight := allStakes[0].Height
		
		for _, stake := range allStakes {
			if stake.Height < earliestHeight {
				earliestHeight = stake.Height
			}
			if stake.Height > latestHeight {
				latestHeight = stake.Height
			}
		}
		
		history["earliest_stake_height"] = earliestHeight
		history["latest_stake_height"] = latestHeight
	}
	
	return history, nil
}

// QueryValidProviders returns providers that meet minimum stake requirements
func (k Keeper) QueryValidProviders(ctx context.Context, minStakeAmount sdk.Coin) ([]string, error) {
	qualifiedStakes := k.GetProvidersByMinStake(ctx, minStakeAmount)
	var providers []string
	
	for _, stake := range qualifiedStakes {
		providers = append(providers, stake.Provider)
	}
	
	return providers, nil
}