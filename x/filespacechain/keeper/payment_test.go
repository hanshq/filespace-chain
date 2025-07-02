package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestSetGetPaymentHistory(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test data
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Get payment history
	retrieved, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	require.Equal(t, paymentHistory.ContractId, retrieved.ContractId)
	require.Equal(t, paymentHistory.TotalPaid, retrieved.TotalPaid)
	require.Equal(t, paymentHistory.LastPaymentBlock, retrieved.LastPaymentBlock)
	require.Equal(t, paymentHistory.CompletionBonusPaid, retrieved.CompletionBonusPaid)
}

func TestGetPaymentHistoryNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Try to get non-existent payment history
	_, found := k.GetPaymentHistory(ctx, 999)
	require.False(t, found)
}

func TestRemovePaymentHistory(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set and then remove payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Verify it exists
	_, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	
	// Remove it
	k.RemovePaymentHistory(ctx, contractId)
	
	// Verify it's gone
	_, found = k.GetPaymentHistory(ctx, contractId)
	require.False(t, found)
}

func TestGetAllPaymentHistory(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create multiple payment histories
	histories := []types.PaymentHistory{
		{
			ContractId:          1,
			TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
			LastPaymentBlock:    100,
			CompletionBonusPaid: false,
		},
		{
			ContractId:          2,
			TotalPaid:           sdk.NewCoin("utoken", math.NewInt(2000)),
			LastPaymentBlock:    200,
			CompletionBonusPaid: true,
		},
		{
			ContractId:          3,
			TotalPaid:           sdk.NewCoin("stake", math.NewInt(3000)),
			LastPaymentBlock:    300,
			CompletionBonusPaid: false,
		},
	}
	
	for _, history := range histories {
		k.SetPaymentHistory(ctx, history)
	}
	
	// Get all histories
	allHistories := k.GetAllPaymentHistory(ctx)
	require.Len(t, allHistories, 3)
	
	// Verify each history exists in the results
	for _, expectedHistory := range histories {
		found := false
		for _, actualHistory := range allHistories {
			if actualHistory.ContractId == expectedHistory.ContractId {
				require.Equal(t, expectedHistory.TotalPaid, actualHistory.TotalPaid)
				require.Equal(t, expectedHistory.LastPaymentBlock, actualHistory.LastPaymentBlock)
				require.Equal(t, expectedHistory.CompletionBonusPaid, actualHistory.CompletionBonusPaid)
				found = true
				break
			}
		}
		require.True(t, found, "Payment history for contract %d not found", expectedHistory.ContractId)
	}
}

func TestGetPaymentHistoryForContract(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Get payment history using helper function
	retrieved, err := k.GetPaymentHistoryForContract(ctx, contractId)
	require.NoError(t, err)
	require.Equal(t, paymentHistory.ContractId, retrieved.ContractId)
	require.Equal(t, paymentHistory.TotalPaid, retrieved.TotalPaid)
}

func TestGetPaymentHistoryForContractNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Try to get non-existent payment history
	_, err := k.GetPaymentHistoryForContract(ctx, 999)
	require.Error(t, err)
	require.Contains(t, err.Error(), "payment history not found")
}

func TestUpdateLastPaymentBlock(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set initial payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Update last payment block
	newBlockHeight := uint64(200)
	err := k.UpdateLastPaymentBlock(ctx, contractId, newBlockHeight)
	require.NoError(t, err)
	
	// Verify update
	updated, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	require.Equal(t, newBlockHeight, updated.LastPaymentBlock)
	require.Equal(t, paymentHistory.TotalPaid, updated.TotalPaid) // Other fields unchanged
}

func TestUpdateLastPaymentBlockNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Try to update non-existent payment history
	err := k.UpdateLastPaymentBlock(ctx, 999, 200)
	require.Error(t, err)
	require.Contains(t, err.Error(), "payment history not found")
}

func TestAddPaymentAmount(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set initial payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Add payment amount
	additionalAmount := sdk.NewCoin("utoken", math.NewInt(500))
	err := k.AddPaymentAmount(ctx, contractId, additionalAmount)
	require.NoError(t, err)
	
	// Verify update
	updated, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	expectedTotal := sdk.NewCoin("utoken", math.NewInt(1500))
	require.Equal(t, expectedTotal, updated.TotalPaid)
	require.True(t, updated.LastPaymentBlock > paymentHistory.LastPaymentBlock) // Should be updated to current block
}

func TestAddPaymentAmountDenomMismatch(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	
	// Set initial payment history
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Try to add amount with different denomination
	additionalAmount := sdk.NewCoin("stake", math.NewInt(500))
	err := k.AddPaymentAmount(ctx, contractId, additionalAmount)
	require.Error(t, err)
	require.Contains(t, err.Error(), "denomination mismatch")
}

func TestGetPaymentHistoryByBlockRange(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create payment histories with different block heights
	histories := []types.PaymentHistory{
		{ContractId: 1, TotalPaid: sdk.NewCoin("utoken", math.NewInt(1000)), LastPaymentBlock: 50},
		{ContractId: 2, TotalPaid: sdk.NewCoin("utoken", math.NewInt(2000)), LastPaymentBlock: 150},
		{ContractId: 3, TotalPaid: sdk.NewCoin("utoken", math.NewInt(3000)), LastPaymentBlock: 250},
		{ContractId: 4, TotalPaid: sdk.NewCoin("utoken", math.NewInt(4000)), LastPaymentBlock: 350},
	}
	
	for _, history := range histories {
		k.SetPaymentHistory(ctx, history)
	}
	
	// Get histories in range 100-300
	rangeHistories := k.GetPaymentHistoryByBlockRange(ctx, 100, 300)
	require.Len(t, rangeHistories, 2)
	
	// Should include contracts 2 and 3 (blocks 150 and 250)
	contractIds := make([]uint64, len(rangeHistories))
	for i, history := range rangeHistories {
		contractIds[i] = history.ContractId
	}
	require.Contains(t, contractIds, uint64(2))
	require.Contains(t, contractIds, uint64(3))
}

func TestCleanupCompletedPaymentHistory(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create mix of completed and pending payment histories
	histories := []types.PaymentHistory{
		{ContractId: 1, TotalPaid: sdk.NewCoin("utoken", math.NewInt(1000)), LastPaymentBlock: 50, CompletionBonusPaid: true},
		{ContractId: 2, TotalPaid: sdk.NewCoin("utoken", math.NewInt(2000)), LastPaymentBlock: 150, CompletionBonusPaid: false},
		{ContractId: 3, TotalPaid: sdk.NewCoin("utoken", math.NewInt(3000)), LastPaymentBlock: 250, CompletionBonusPaid: true},
		{ContractId: 4, TotalPaid: sdk.NewCoin("utoken", math.NewInt(4000)), LastPaymentBlock: 350, CompletionBonusPaid: true},
	}
	
	for _, history := range histories {
		k.SetPaymentHistory(ctx, history)
	}
	
	// Cleanup histories older than 100 blocks (should remove contracts 1 and 3)
	err := k.CleanupCompletedPaymentHistory(ctx, 100)
	require.NoError(t, err)
	
	// Verify remaining histories
	remaining := k.GetAllPaymentHistory(ctx)
	
	// Should have contracts 2 and 4 remaining
	contractIds := make([]uint64, len(remaining))
	for i, history := range remaining {
		contractIds[i] = history.ContractId
	}
	require.Contains(t, contractIds, uint64(2)) // Not completed
	require.Contains(t, contractIds, uint64(4)) // Recent completed
	require.NotContains(t, contractIds, uint64(1)) // Old completed - should be cleaned
	require.NotContains(t, contractIds, uint64(3)) // Old completed - should be cleaned
}

func TestQueryPaymentSummary(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create some payment histories
	histories := []types.PaymentHistory{
		{ContractId: 1, TotalPaid: sdk.NewCoin("utoken", math.NewInt(1000)), CompletionBonusPaid: true},
		{ContractId: 2, TotalPaid: sdk.NewCoin("utoken", math.NewInt(2000)), CompletionBonusPaid: false},
	}
	
	for _, history := range histories {
		k.SetPaymentHistory(ctx, history)
	}
	
	// Get payment summary
	summary := k.QueryPaymentSummary(ctx)
	
	require.Equal(t, 2, summary["total_payment_histories"])
	require.NotNil(t, summary["active_contracts"])
	require.NotNil(t, summary["expired_contracts"])
	require.NotNil(t, summary["total_escrow_records"])
	require.NotNil(t, summary["total_provider_stakes"])
}

func TestFormatPaymentHistoryInfo(t *testing.T) {
	k, _ := keepertest.FilespacechainKeeper(t)
	
	paymentHistory := types.PaymentHistory{
		ContractId:          123,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    456,
		CompletionBonusPaid: true,
	}
	
	formatted := k.FormatPaymentHistoryInfo(paymentHistory)
	
	require.Equal(t, "123", formatted["contract_id"])
	require.Equal(t, "1000utoken", formatted["total_paid"])
	require.Equal(t, "456", formatted["last_payment_block"])
	require.Equal(t, "true", formatted["completion_bonus_paid"])
}

func TestGetHostingOffersByInquiry(t *testing.T) {
	k, _ := keepertest.FilespacechainKeeper(t)
	
	// This function depends on hosting offers existing
	// For now, test that it doesn't panic and returns empty slice
	offers := k.GetHostingOffersByInquiry(context.Background(), 123)
	require.NotNil(t, offers)
	require.IsType(t, []types.HostingOffer{}, offers)
}