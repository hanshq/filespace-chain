package keeper

import (
	"context"
	"fmt"
	"sort"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// QueryPaymentAnalytics returns comprehensive payment analytics
func (k Keeper) QueryPaymentAnalytics(ctx context.Context) (map[string]interface{}, error) {
	analytics := make(map[string]interface{})
	
	allPayments := k.GetAllPaymentHistory(ctx)
	analytics["total_payment_records"] = len(allPayments)
	
	if len(allPayments) == 0 {
		return analytics, nil
	}
	
	// Calculate totals by denomination
	denomTotals := make(map[string]math.Int)
	denomCounts := make(map[string]int)
	completedPayments := 0
	
	for _, payment := range allPayments {
		if _, exists := denomTotals[payment.TotalPaid.Denom]; !exists {
			denomTotals[payment.TotalPaid.Denom] = math.ZeroInt()
			denomCounts[payment.TotalPaid.Denom] = 0
		}
		denomTotals[payment.TotalPaid.Denom] = denomTotals[payment.TotalPaid.Denom].Add(payment.TotalPaid.Amount)
		denomCounts[payment.TotalPaid.Denom]++
		
		if payment.CompletionBonusPaid {
			completedPayments++
		}
	}
	
	// Convert to readable format
	totalsByDenom := make(map[string]sdk.Coin)
	for denom, total := range denomTotals {
		totalsByDenom[denom] = sdk.NewCoin(denom, total)
	}
	
	analytics["total_paid_by_denom"] = totalsByDenom
	analytics["payment_count_by_denom"] = denomCounts
	analytics["completed_payments"] = completedPayments
	analytics["pending_payments"] = len(allPayments) - completedPayments
	
	return analytics, nil
}

// QueryPaymentsByContract returns payment history for a specific contract
func (k Keeper) QueryPaymentsByContract(ctx context.Context, contractId uint64) (types.PaymentHistory, error) {
	return k.GetPaymentHistoryForContract(ctx, contractId)
}

// QueryPaymentsByProvider returns all payments received by a specific provider
func (k Keeper) QueryPaymentsByProvider(ctx context.Context, provider string) ([]types.PaymentHistory, error) {
	return k.GetPaymentHistoryByProvider(ctx, provider), nil
}

// QueryPaymentsByBlockRange returns payments within a specific block range
func (k Keeper) QueryPaymentsByBlockRange(ctx context.Context, startBlock, endBlock uint64) ([]types.PaymentHistory, error) {
	return k.GetPaymentHistoryByBlockRange(ctx, startBlock, endBlock), nil
}

// QueryRecentPayments returns the most recent N payments
func (k Keeper) QueryRecentPayments(ctx context.Context, limit int) ([]types.PaymentHistory, error) {
	allPayments := k.GetAllPaymentHistory(ctx)
	
	// Sort by last payment block (descending)
	sort.Slice(allPayments, func(i, j int) bool {
		return allPayments[i].LastPaymentBlock > allPayments[j].LastPaymentBlock
	})
	
	// Apply limit
	if limit > 0 && limit < len(allPayments) {
		allPayments = allPayments[:limit]
	}
	
	return allPayments, nil
}

// QueryPendingPayments returns contracts that haven't received completion bonus
func (k Keeper) QueryPendingPayments(ctx context.Context) ([]types.PaymentHistory, error) {
	allPayments := k.GetAllPaymentHistory(ctx)
	var pendingPayments []types.PaymentHistory
	
	for _, payment := range allPayments {
		if !payment.CompletionBonusPaid {
			pendingPayments = append(pendingPayments, payment)
		}
	}
	
	return pendingPayments, nil
}

// QueryCompletedPayments returns contracts that have received completion bonus
func (k Keeper) QueryCompletedPayments(ctx context.Context) ([]types.PaymentHistory, error) {
	allPayments := k.GetAllPaymentHistory(ctx)
	var completedPayments []types.PaymentHistory
	
	for _, payment := range allPayments {
		if payment.CompletionBonusPaid {
			completedPayments = append(completedPayments, payment)
		}
	}
	
	return completedPayments, nil
}

// QueryPaymentDistribution returns payment distribution statistics
func (k Keeper) QueryPaymentDistribution(ctx context.Context) (map[string]interface{}, error) {
	distribution := make(map[string]interface{})
	
	allPayments := k.GetAllPaymentHistory(ctx)
	if len(allPayments) == 0 {
		return distribution, nil
	}
	
	// Collect all payment amounts for statistical analysis
	var amounts []math.Int
	totalAmount := math.ZeroInt()
	
	for _, payment := range allPayments {
		amounts = append(amounts, payment.TotalPaid.Amount)
		totalAmount = totalAmount.Add(payment.TotalPaid.Amount)
	}
	
	// Sort amounts for percentile calculations
	sort.Slice(amounts, func(i, j int) bool {
		return amounts[i].LT(amounts[j])
	})
	
	// Calculate statistics
	count := len(amounts)
	distribution["count"] = count
	distribution["total"] = totalAmount
	distribution["average"] = totalAmount.Quo(math.NewInt(int64(count)))
	distribution["min"] = amounts[0]
	distribution["max"] = amounts[count-1]
	
	// Calculate percentiles
	if count >= 4 {
		distribution["25th_percentile"] = amounts[count/4]
		distribution["50th_percentile"] = amounts[count/2] // median
		distribution["75th_percentile"] = amounts[count*3/4]
	}
	
	return distribution, nil
}

// QueryProviderPaymentSummary returns payment summary for each provider
func (k Keeper) QueryProviderPaymentSummary(ctx context.Context) (map[string]interface{}, error) {
	summary := make(map[string]interface{})
	
	// Get all providers with stakes
	allStakes := k.GetAllProviderStakes(ctx)
	
	for _, stake := range allStakes {
		provider := stake.Provider
		payments := k.GetPaymentHistoryByProvider(ctx, provider)
		
		providerSummary := map[string]interface{}{
			"payment_count":       len(payments),
			"total_earned":        math.ZeroInt(),
			"completed_contracts": 0,
			"pending_contracts":   0,
		}
		
		totalEarned := math.ZeroInt()
		completedContracts := 0
		pendingContracts := 0
		
		for _, payment := range payments {
			// Get contract to calculate provider share
			contract, found := k.GetHostingContract(ctx, payment.ContractId)
			if !found {
				continue
			}
			
			// Get offers to calculate share
			offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
			if len(offers) == 0 {
				continue
			}
			
			// Calculate provider's share (equal split)
			providerShare := payment.TotalPaid.Amount.Quo(math.NewInt(int64(len(offers))))
			totalEarned = totalEarned.Add(providerShare)
			
			if payment.CompletionBonusPaid {
				completedContracts++
			} else {
				pendingContracts++
			}
		}
		
		providerSummary["total_earned"] = totalEarned
		providerSummary["completed_contracts"] = completedContracts
		providerSummary["pending_contracts"] = pendingContracts
		
		summary[provider] = providerSummary
	}
	
	return summary, nil
}

// QueryPaymentTrends returns payment trends over time
func (k Keeper) QueryPaymentTrends(ctx context.Context, blockWindow uint64) (map[string]interface{}, error) {
	trends := make(map[string]interface{})
	
	allPayments := k.GetAllPaymentHistory(ctx)
	if len(allPayments) == 0 {
		return trends, nil
	}
	
	// Group payments by block ranges
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	// Create time windows
	windows := make(map[string][]types.PaymentHistory)
	windowTotals := make(map[string]math.Int)
	
	for _, payment := range allPayments {
		// Calculate which window this payment belongs to
		windowStart := (payment.LastPaymentBlock / blockWindow) * blockWindow
		windowKey := fmt.Sprintf("blocks_%d-%d", windowStart, windowStart+blockWindow-1)
		
		windows[windowKey] = append(windows[windowKey], payment)
		
		if _, exists := windowTotals[windowKey]; !exists {
			windowTotals[windowKey] = math.ZeroInt()
		}
		windowTotals[windowKey] = windowTotals[windowKey].Add(payment.TotalPaid.Amount)
	}
	
	trends["payment_windows"] = windows
	trends["window_totals"] = windowTotals
	trends["block_window_size"] = blockWindow
	trends["current_block"] = currentBlock
	
	return trends, nil
}

// QueryUnpaidContracts returns contracts that should have payments but don't
func (k Keeper) QueryUnpaidContracts(ctx context.Context) ([]types.HostingContract, error) {
	allContracts := k.GetAllHostingContract(ctx)
	var unpaidContracts []types.HostingContract
	
	for _, contract := range allContracts {
		// Check if payment history exists
		_, found := k.GetPaymentHistory(ctx, contract.Id)
		if !found {
			// This contract has no payment history, check if it should
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			currentBlock := uint64(sdkCtx.BlockHeight())
			
			// If contract has started, it should have payment history
			if currentBlock >= contract.StartBlock {
				unpaidContracts = append(unpaidContracts, contract)
			}
		}
	}
	
	return unpaidContracts, nil
}