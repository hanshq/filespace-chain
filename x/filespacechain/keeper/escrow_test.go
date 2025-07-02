package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/keeper"
)

func TestSetGetEscrowRecord(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test data
	inquiryId := uint64(1)
	amount := sdk.NewCoin("utoken", math.NewInt(1000))
	creator := "cosmos1test"
	
	// Set escrow record
	k.SetEscrowRecord(ctx, inquiryId, amount, creator)
	
	// Get escrow record
	record, found := k.GetEscrowRecord(ctx, inquiryId)
	require.True(t, found)
	require.Equal(t, inquiryId, record.InquiryId)
	require.Equal(t, amount, record.Amount)
	require.Equal(t, creator, record.Creator)
}

func TestGetEscrowRecordNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Try to get non-existent escrow record
	_, found := k.GetEscrowRecord(ctx, 999)
	require.False(t, found)
}

func TestRemoveEscrowRecord(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	inquiryId := uint64(1)
	amount := sdk.NewCoin("utoken", math.NewInt(1000))
	creator := "cosmos1test"
	
	// Set and then remove escrow record
	k.SetEscrowRecord(ctx, inquiryId, amount, creator)
	
	// Verify it exists
	_, found := k.GetEscrowRecord(ctx, inquiryId)
	require.True(t, found)
	
	// Remove it
	k.RemoveEscrowRecord(ctx, inquiryId)
	
	// Verify it's gone
	_, found = k.GetEscrowRecord(ctx, inquiryId)
	require.False(t, found)
}

func TestGetAllEscrowRecords(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create multiple escrow records
	records := []struct {
		inquiryId uint64
		amount    sdk.Coin
		creator   string
	}{
		{1, sdk.NewCoin("utoken", math.NewInt(1000)), "cosmos1test1"},
		{2, sdk.NewCoin("utoken", math.NewInt(2000)), "cosmos1test2"},
		{3, sdk.NewCoin("stake", math.NewInt(3000)), "cosmos1test3"},
	}
	
	for _, record := range records {
		k.SetEscrowRecord(ctx, record.inquiryId, record.amount, record.creator)
	}
	
	// Get all records
	allRecords := k.GetAllEscrowRecords(ctx)
	require.Len(t, allRecords, 3)
	
	// Verify each record exists in the results
	for _, expectedRecord := range records {
		found := false
		for _, actualRecord := range allRecords {
			if actualRecord.InquiryId == expectedRecord.inquiryId {
				require.Equal(t, expectedRecord.amount, actualRecord.Amount)
				require.Equal(t, expectedRecord.creator, actualRecord.Creator)
				found = true
				break
			}
		}
		require.True(t, found, "Record with inquiry ID %d not found", expectedRecord.inquiryId)
	}
}

func TestGetEscrowRecordsByCreator(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	creator1 := "cosmos1test1"
	creator2 := "cosmos1test2"
	
	// Create records for different creators
	k.SetEscrowRecord(ctx, 1, sdk.NewCoin("utoken", math.NewInt(1000)), creator1)
	k.SetEscrowRecord(ctx, 2, sdk.NewCoin("utoken", math.NewInt(2000)), creator1)
	k.SetEscrowRecord(ctx, 3, sdk.NewCoin("utoken", math.NewInt(3000)), creator2)
	
	// Get records for creator1
	creator1Records := k.GetEscrowRecordsByCreator(ctx, creator1)
	require.Len(t, creator1Records, 2)
	
	// Get records for creator2
	creator2Records := k.GetEscrowRecordsByCreator(ctx, creator2)
	require.Len(t, creator2Records, 1)
	
	// Get records for non-existent creator
	nonExistentRecords := k.GetEscrowRecordsByCreator(ctx, "cosmos1nonexistent")
	require.Len(t, nonExistentRecords, 0)
}

func TestUpdateEscrowAmount(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	inquiryId := uint64(1)
	originalAmount := sdk.NewCoin("utoken", math.NewInt(1000))
	newAmount := sdk.NewCoin("utoken", math.NewInt(2000))
	creator := "cosmos1test"
	
	// Set initial escrow record
	k.SetEscrowRecord(ctx, inquiryId, originalAmount, creator)
	
	// Update amount
	err := k.UpdateEscrowAmount(ctx, inquiryId, newAmount)
	require.NoError(t, err)
	
	// Verify update
	record, found := k.GetEscrowRecord(ctx, inquiryId)
	require.True(t, found)
	require.Equal(t, newAmount, record.Amount)
	require.Equal(t, creator, record.Creator) // Creator should remain unchanged
}

func TestUpdateEscrowAmountNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	newAmount := sdk.NewCoin("utoken", math.NewInt(2000))
	
	// Try to update non-existent escrow record
	err := k.UpdateEscrowAmount(ctx, 999, newAmount)
	require.Error(t, err)
	require.Contains(t, err.Error(), "escrow record not found")
}

func TestGetTotalEscrowedAmount(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create multiple escrow records with different denominations
	k.SetEscrowRecord(ctx, 1, sdk.NewCoin("utoken", math.NewInt(1000)), "cosmos1test1")
	k.SetEscrowRecord(ctx, 2, sdk.NewCoin("utoken", math.NewInt(2000)), "cosmos1test2")
	k.SetEscrowRecord(ctx, 3, sdk.NewCoin("utoken", math.NewInt(3000)), "cosmos1test3")
	k.SetEscrowRecord(ctx, 4, sdk.NewCoin("stake", math.NewInt(5000)), "cosmos1test4")
	
	// Test total for utoken
	totalUtoken := k.GetTotalEscrowedAmount(ctx, "utoken")
	expectedUtoken := sdk.NewCoin("utoken", math.NewInt(6000))
	require.Equal(t, expectedUtoken, totalUtoken)
	
	// Test total for stake
	totalStake := k.GetTotalEscrowedAmount(ctx, "stake")
	expectedStake := sdk.NewCoin("stake", math.NewInt(5000))
	require.Equal(t, expectedStake, totalStake)
	
	// Test total for non-existent denomination
	totalNonExistent := k.GetTotalEscrowedAmount(ctx, "nonexistent")
	expectedNonExistent := sdk.NewCoin("nonexistent", math.ZeroInt())
	require.Equal(t, expectedNonExistent, totalNonExistent)
}

func TestGetActiveEscrowRecords(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// This test would require actual hosting inquiries to be set up
	// For now, we'll test the function doesn't panic and returns empty slice
	activeRecords := k.GetActiveEscrowRecords(ctx)
	require.NotNil(t, activeRecords)
	require.IsType(t, []keeper.EscrowRecord{}, activeRecords)
}

func TestEscrowRecordStructure(t *testing.T) {
	// Test the EscrowRecord structure
	record := keeper.EscrowRecord{
		InquiryId: 123,
		Amount:    sdk.NewCoin("utoken", math.NewInt(1000)),
		Creator:   "cosmos1test",
	}
	
	require.Equal(t, uint64(123), record.InquiryId)
	require.Equal(t, "utoken", record.Amount.Denom)
	require.Equal(t, math.NewInt(1000), record.Amount.Amount)
	require.Equal(t, "cosmos1test", record.Creator)
}

func TestEscrowKey(t *testing.T) {
	// Test the escrow key generation
	inquiryId := uint64(12345)
	key := keeper.EscrowKey(inquiryId)
	
	require.NotNil(t, key)
	require.True(t, len(key) > 0)
	
	// Keys for different inquiry IDs should be different
	key2 := keeper.EscrowKey(54321)
	require.NotEqual(t, key, key2)
	
	// Keys for same inquiry ID should be same
	key3 := keeper.EscrowKey(inquiryId)
	require.Equal(t, key, key3)
}