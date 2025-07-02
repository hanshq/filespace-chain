package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetPaymentHistoryForContract returns payment history for a specific contract ID
// This is a helper function that can be used by CLI or other query mechanisms
func (k Keeper) GetPaymentHistoryForContract(ctx context.Context, contractId uint64) (types.PaymentHistory, error) {
	paymentHistory, found := k.GetPaymentHistory(ctx, contractId)
	if !found {
		return types.PaymentHistory{}, fmt.Errorf("payment history not found for contract %d", contractId)
	}
	return paymentHistory, nil
}

// GetAllPaymentHistories returns all payment history records
func (k Keeper) GetAllPaymentHistories(ctx context.Context) []types.PaymentHistory {
	return k.GetAllPaymentHistory(ctx)
}

// GetEscrowStatusForInquiry returns escrow status for a specific inquiry ID
func (k Keeper) GetEscrowStatusForInquiry(ctx context.Context, inquiryId uint64) (EscrowRecord, error) {
	escrowRecord, found := k.GetEscrowRecord(ctx, inquiryId)
	if !found {
		return EscrowRecord{}, fmt.Errorf("escrow record not found for inquiry %d", inquiryId)
	}
	return escrowRecord, nil
}

// GetAllEscrowStatuses returns all escrow records
func (k Keeper) GetAllEscrowStatuses(ctx context.Context) []EscrowRecord {
	return k.GetAllEscrowRecords(ctx)
}

// GetProviderStakeByAddress returns provider stake for a specific address
func (k Keeper) GetProviderStakeByAddress(ctx context.Context, provider string) (ProviderStake, error) {
	providerStake, found := k.GetProviderStake(ctx, provider)
	if !found {
		return ProviderStake{}, fmt.Errorf("provider stake not found for address %s", provider)
	}
	return providerStake, nil
}

// QueryPaymentSummary returns a summary of payments for debugging/monitoring
func (k Keeper) QueryPaymentSummary(ctx context.Context) map[string]interface{} {
	summary := make(map[string]interface{})
	
	// Count total payment histories
	paymentHistories := k.GetAllPaymentHistory(ctx)
	summary["total_payment_histories"] = len(paymentHistories)
	
	// Count active contracts
	activeContracts := k.GetActiveContracts(ctx)
	summary["active_contracts"] = len(activeContracts)
	
	// Count expired contracts  
	expiredContracts := k.GetExpiredContracts(ctx)
	summary["expired_contracts"] = len(expiredContracts)
	
	// Count escrow records
	escrowRecords := k.GetAllEscrowRecords(ctx)
	summary["total_escrow_records"] = len(escrowRecords)
	
	// Count provider stakes
	providerStakes := k.GetAllProviderStakes(ctx)
	summary["total_provider_stakes"] = len(providerStakes)
	
	return summary
}

// FormatPaymentHistoryInfo formats payment history for display
func (k Keeper) FormatPaymentHistoryInfo(paymentHistory types.PaymentHistory) map[string]string {
	return map[string]string{
		"contract_id":            strconv.FormatUint(paymentHistory.ContractId, 10),
		"total_paid":             paymentHistory.TotalPaid.String(),
		"last_payment_block":     strconv.FormatUint(paymentHistory.LastPaymentBlock, 10),
		"completion_bonus_paid":  strconv.FormatBool(paymentHistory.CompletionBonusPaid),
	}
}

// FormatEscrowInfo formats escrow record for display  
func (k Keeper) FormatEscrowInfo(escrowRecord EscrowRecord) map[string]string {
	return map[string]string{
		"inquiry_id":     strconv.FormatUint(escrowRecord.InquiryId, 10),
		"amount":         escrowRecord.Amount.String(),
		"creator":        escrowRecord.Creator,
	}
}

// FormatProviderStakeInfo formats provider stake for display
func (k Keeper) FormatProviderStakeInfo(providerStake ProviderStake) map[string]string {
	return map[string]string{
		"provider":      providerStake.Provider,
		"amount":        providerStake.Amount.String(),
		"staked_block":  strconv.FormatUint(providerStake.Height, 10),
	}
}

// GetPaymentHistoryByProvider returns all payment histories where the provider received payments
func (k Keeper) GetPaymentHistoryByProvider(ctx context.Context, provider string) []types.PaymentHistory {
	var providerPayments []types.PaymentHistory
	allPayments := k.GetAllPaymentHistory(ctx)
	
	for _, payment := range allPayments {
		// Get the contract to find associated offers/providers
		contract, found := k.GetHostingContract(ctx, payment.ContractId)
		if !found {
			continue
		}
		
		// Get offers for this contract's inquiry
		offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
		
		// Check if this provider has an offer for this contract
		for _, offer := range offers {
			if offer.Creator == provider {
				providerPayments = append(providerPayments, payment)
				break
			}
		}
	}
	
	return providerPayments
}

// GetPaymentHistoryByBlockRange returns payment histories with last payment in specified block range
func (k Keeper) GetPaymentHistoryByBlockRange(ctx context.Context, startBlock, endBlock uint64) []types.PaymentHistory {
	allPayments := k.GetAllPaymentHistory(ctx)
	var rangePayments []types.PaymentHistory
	
	for _, payment := range allPayments {
		if payment.LastPaymentBlock >= startBlock && payment.LastPaymentBlock <= endBlock {
			rangePayments = append(rangePayments, payment)
		}
	}
	
	return rangePayments
}

// UpdateLastPaymentBlock updates only the last payment block for a contract
func (k Keeper) UpdateLastPaymentBlock(ctx context.Context, contractId uint64, blockHeight uint64) error {
	payment, found := k.GetPaymentHistory(ctx, contractId)
	if !found {
		return fmt.Errorf("payment history not found for contract %d", contractId)
	}
	
	payment.LastPaymentBlock = blockHeight
	k.SetPaymentHistory(ctx, payment)
	
	return nil
}

// AddPaymentAmount increments the total paid amount for a contract
func (k Keeper) AddPaymentAmount(ctx context.Context, contractId uint64, amount sdk.Coin) error {
	payment, found := k.GetPaymentHistory(ctx, contractId)
	if !found {
		return fmt.Errorf("payment history not found for contract %d", contractId)
	}
	
	if payment.TotalPaid.Denom != amount.Denom {
		return fmt.Errorf("denomination mismatch: existing %s, new %s", 
			payment.TotalPaid.Denom, amount.Denom)
	}
	
	payment.TotalPaid = payment.TotalPaid.Add(amount)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	payment.LastPaymentBlock = uint64(sdkCtx.BlockHeight())
	
	k.SetPaymentHistory(ctx, payment)
	
	return nil
}

// CleanupCompletedPaymentHistory removes payment history for completed contracts older than specified blocks
func (k Keeper) CleanupCompletedPaymentHistory(ctx context.Context, olderThanBlocks uint64) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	cutoffHeight := currentHeight - olderThanBlocks
	
	allPayments := k.GetAllPaymentHistory(ctx)
	cleanedCount := 0
	
	for _, payment := range allPayments {
		// Only clean up completed contracts
		if payment.CompletionBonusPaid && payment.LastPaymentBlock < cutoffHeight {
			k.RemovePaymentHistory(ctx, payment.ContractId)
			cleanedCount++
		}
	}
	
	k.Logger().Info("cleaned up old payment history records",
		"cleaned_count", cleanedCount,
		"cutoff_height", cutoffHeight,
	)
	
	return nil
}