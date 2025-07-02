package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CleanupExpiredRecords performs comprehensive cleanup of expired records
func (k Keeper) CleanupExpiredRecords(ctx context.Context) error {
	k.Logger().Info("starting cleanup of expired records")
	
	// Cleanup expired inquiries and refund escrow
	err := k.ProcessExpiredInquiries(ctx)
	if err != nil {
		k.Logger().Error("failed to process expired inquiries", "error", err)
		return fmt.Errorf("failed to process expired inquiries: %w", err)
	}
	
	// Cleanup old payment history records
	err = k.CleanupOldPaymentHistory(ctx)
	if err != nil {
		k.Logger().Error("failed to cleanup old payment history", "error", err)
		return fmt.Errorf("failed to cleanup old payment history: %w", err)
	}
	
	// Cleanup abandoned escrow records
	err = k.CleanupAbandonedEscrow(ctx)
	if err != nil {
		k.Logger().Error("failed to cleanup abandoned escrow", "error", err)
		return fmt.Errorf("failed to cleanup abandoned escrow: %w", err)
	}
	
	k.Logger().Info("completed cleanup of expired records")
	return nil
}

// CleanupOldPaymentHistory removes payment history for contracts completed more than specified blocks ago
func (k Keeper) CleanupOldPaymentHistory(ctx context.Context) error {
	// Clean up payment history older than 100,000 blocks (approximately 1 week)
	olderThanBlocks := uint64(100000)
	return k.CleanupCompletedPaymentHistory(ctx, olderThanBlocks)
}

// CleanupAbandonedEscrow identifies and refunds escrow for inquiries with no activity
func (k Keeper) CleanupAbandonedEscrow(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	// Consider inquiries abandoned if they're past end time with no contracts
	allEscrowRecords := k.GetAllEscrowRecords(ctx)
	cleanedCount := 0
	
	for _, escrowRecord := range allEscrowRecords {
		inquiry, found := k.GetHostingInquiry(ctx, escrowRecord.InquiryId)
		if !found {
			// Inquiry doesn't exist, clean up orphaned escrow
			err := k.refundEscrowRecord(ctx, escrowRecord)
			if err != nil {
				k.Logger().Error("failed to refund orphaned escrow", 
					"inquiry_id", escrowRecord.InquiryId,
					"error", err)
				continue
			}
			cleanedCount++
			continue
		}
		
		// Check if inquiry is expired and has no active contracts
		if currentBlock > inquiry.EndTime {
			contracts := k.GetAllHostingContract(ctx)
			hasActiveContract := false
			
			for _, contract := range contracts {
				if contract.InquiryId == inquiry.Id && currentBlock < contract.EndBlock {
					hasActiveContract = true
					break
				}
			}
			
			if !hasActiveContract {
				// No active contracts, refund escrow
				err := k.refundEscrowRecord(ctx, escrowRecord)
				if err != nil {
					k.Logger().Error("failed to refund abandoned escrow", 
						"inquiry_id", escrowRecord.InquiryId,
						"error", err)
					continue
				}
				cleanedCount++
			}
		}
	}
	
	k.Logger().Info("cleaned up abandoned escrow records", "count", cleanedCount)
	return nil
}

// refundEscrowRecord handles the refund process for a single escrow record
func (k Keeper) refundEscrowRecord(ctx context.Context, escrowRecord EscrowRecord) error {
	// Convert creator address
	creatorAddr, err := sdk.AccAddressFromBech32(escrowRecord.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address %s: %w", escrowRecord.Creator, err)
	}
	
	// Refund the funds
	err = k.RefundFunds(ctx, creatorAddr, escrowRecord.Amount)
	if err != nil {
		return fmt.Errorf("failed to refund funds: %w", err)
	}
	
	// Remove escrow record
	k.RemoveEscrowRecord(ctx, escrowRecord.InquiryId)
	
	k.Logger().Info("refunded abandoned escrow", 
		"inquiry_id", escrowRecord.InquiryId,
		"amount", escrowRecord.Amount.String(),
		"creator", escrowRecord.Creator)
	
	return nil
}

// CleanupExpiredContracts removes expired contracts and processes final payments
func (k Keeper) CleanupExpiredContracts(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	allContracts := k.GetAllHostingContract(ctx)
	cleanedCount := 0
	
	for _, contract := range allContracts {
		if currentBlock >= contract.EndBlock {
			// Process completion bonus if not already paid
			paymentHistory, found := k.GetPaymentHistory(ctx, contract.Id)
			if found && !paymentHistory.CompletionBonusPaid {
				err := k.ProcessCompletionBonus(ctx, contract.Id)
				if err != nil {
					k.Logger().Error("failed to process completion bonus for expired contract",
						"contract_id", contract.Id,
						"error", err)
				}
			}
			
			// Note: We don't remove the contract itself as it serves as historical record
			// The payment history cleanup will handle old payment records
			cleanedCount++
		}
	}
	
	k.Logger().Info("processed expired contracts", "count", cleanedCount)
	return nil
}

// CleanupOrphanedRecords removes records that reference non-existent entities
func (k Keeper) CleanupOrphanedRecords(ctx context.Context) error {
	k.Logger().Info("starting cleanup of orphaned records")
	
	// Check for payment histories with invalid contract IDs
	allPayments := k.GetAllPaymentHistory(ctx)
	orphanedPayments := 0
	
	for _, payment := range allPayments {
		_, found := k.GetHostingContract(ctx, payment.ContractId)
		if !found {
			k.RemovePaymentHistory(ctx, payment.ContractId)
			orphanedPayments++
			k.Logger().Info("removed orphaned payment history", 
				"contract_id", payment.ContractId)
		}
	}
	
	// Check for escrow records with invalid inquiry IDs
	allEscrowRecords := k.GetAllEscrowRecords(ctx)
	orphanedEscrow := 0
	
	for _, escrowRecord := range allEscrowRecords {
		_, found := k.GetHostingInquiry(ctx, escrowRecord.InquiryId)
		if !found {
			err := k.refundEscrowRecord(ctx, escrowRecord)
			if err != nil {
				k.Logger().Error("failed to refund orphaned escrow record",
					"inquiry_id", escrowRecord.InquiryId,
					"error", err)
			} else {
				orphanedEscrow++
			}
		}
	}
	
	k.Logger().Info("completed cleanup of orphaned records",
		"orphaned_payments", orphanedPayments,
		"orphaned_escrow", orphanedEscrow)
	
	return nil
}

// PerformMaintenanceCleanup runs all cleanup operations - designed to be called periodically
func (k Keeper) PerformMaintenanceCleanup(ctx context.Context) error {
	k.Logger().Info("starting scheduled maintenance cleanup")
	
	// Run all cleanup operations
	cleanupOps := []struct {
		name string
		fn   func(context.Context) error
	}{
		{"expired_records", k.CleanupExpiredRecords},
		{"expired_contracts", k.CleanupExpiredContracts},
		{"orphaned_records", k.CleanupOrphanedRecords},
	}
	
	for _, op := range cleanupOps {
		k.Logger().Info("running cleanup operation", "operation", op.name)
		err := op.fn(ctx)
		if err != nil {
			k.Logger().Error("cleanup operation failed", 
				"operation", op.name, 
				"error", err)
			// Continue with other operations even if one fails
		}
	}
	
	k.Logger().Info("completed scheduled maintenance cleanup")
	return nil
}

// GetCleanupStatus returns information about records that need cleanup
func (k Keeper) GetCleanupStatus(ctx context.Context) map[string]interface{} {
	status := make(map[string]interface{})
	
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	// Count expired inquiries
	allInquiries := k.GetAllHostingInquiry(ctx)
	expiredInquiries := 0
	for _, inquiry := range allInquiries {
		if currentBlock > inquiry.EndTime {
			expiredInquiries++
		}
	}
	status["expired_inquiries"] = expiredInquiries
	
	// Count expired contracts
	allContracts := k.GetAllHostingContract(ctx)
	expiredContracts := 0
	for _, contract := range allContracts {
		if currentBlock >= contract.EndBlock {
			expiredContracts++
		}
	}
	status["expired_contracts"] = expiredContracts
	
	// Count old payment histories
	allPayments := k.GetAllPaymentHistory(ctx)
	oldPayments := 0
	cutoffBlock := currentBlock - 100000 // Same threshold as cleanup
	
	for _, payment := range allPayments {
		if payment.CompletionBonusPaid && payment.LastPaymentBlock < cutoffBlock {
			oldPayments++
		}
	}
	status["old_payment_histories"] = oldPayments
	
	// Count total records
	status["total_inquiries"] = len(allInquiries)
	status["total_contracts"] = len(allContracts)
	status["total_payments"] = len(allPayments)
	status["total_escrow_records"] = len(k.GetAllEscrowRecords(ctx))
	status["total_provider_stakes"] = len(k.GetAllProviderStakes(ctx))
	
	return status
}