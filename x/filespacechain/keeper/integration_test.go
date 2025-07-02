package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// TestEndToEndHostingFlow tests the complete flow from inquiry to contract completion
func TestEndToEndHostingFlow(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Step 1: Create a hosting inquiry
	inquiry := types.HostingInquiry{
		Id:               1,
		Creator:          "cosmos1creator",
		FileEntryCid:     "QmTest123",
		ReplicationRate:  1,
		EscrowAmount:     sdk.NewCoin("utoken", math.NewInt(2000)),
		EndTime:          3700,
		MaxPricePerBlock: 10,
	}
	k.SetHostingInquiry(ctx, inquiry)
	
	// Step 2: Set up escrow for the inquiry
	escrowAmount := sdk.NewCoin("utoken", math.NewInt(2000))
	k.SetEscrowRecord(ctx, inquiry.Id, escrowAmount, inquiry.Creator)
	
	// Verify escrow is set
	escrowRecord, found := k.GetEscrowRecord(ctx, inquiry.Id)
	require.True(t, found)
	require.Equal(t, escrowAmount, escrowRecord.Amount)
	
	// Step 3: Provider stakes tokens to participate
	provider := "cosmos1provider"
	stakeAmount := sdk.NewCoin("utoken", math.NewInt(5000))
	k.SetProviderStake(ctx, provider, stakeAmount, 50)
	
	// Verify stake is set
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	require.Equal(t, stakeAmount, stake.Amount)
	
	// Step 4: Provider creates a hosting offer
	offer := types.HostingOffer{
		Id:            1,
		Region:        "us-east",
		PricePerBlock: sdk.NewCoin("utoken", math.NewInt(1)),
		Creator:       provider,
		InquiryId:     inquiry.Id,
	}
	k.SetHostingOffer(ctx, offer)
	
	// Step 5: Create hosting contract
	contract := types.HostingContract{
		Id:         1,
		InquiryId:  inquiry.Id,
		OfferId:    offer.Id,
		Creator:    inquiry.Creator,
		StartBlock: 100,
		EndBlock:   3700,
	}
	k.SetHostingContract(ctx, contract)
	
	// Step 6: Initialize payment history
	paymentHistory := types.PaymentHistory{
		ContractId:          contract.Id,
		TotalPaid:           sdk.NewCoin("utoken", math.ZeroInt()),
		LastPaymentBlock:    contract.StartBlock,
		CompletionBonusPaid: false,
	}
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Step 7: Verify the complete setup
	// Check that all entities exist and are properly linked
	retrievedInquiry, found := k.GetHostingInquiry(ctx, inquiry.Id)
	require.True(t, found)
	require.Equal(t, inquiry.Creator, retrievedInquiry.Creator)
	
	retrievedOffer, found := k.GetHostingOffer(ctx, offer.Id)
	require.True(t, found)
	require.Equal(t, inquiry.Id, retrievedOffer.InquiryId)
	
	retrievedContract, found := k.GetHostingContract(ctx, contract.Id)
	require.True(t, found)
	require.Equal(t, inquiry.Id, retrievedContract.InquiryId)
	require.Equal(t, offer.Id, retrievedContract.OfferId)
	
	retrievedPayment, found := k.GetPaymentHistory(ctx, contract.Id)
	require.True(t, found)
	require.Equal(t, contract.Id, retrievedPayment.ContractId)
	
	// Step 8: Test business logic queries
	// Query contracts by provider
	providerContracts, err := k.QueryContractsByProvider(ctx, provider)
	require.NoError(t, err)
	require.Len(t, providerContracts, 1)
	require.Equal(t, contract.Id, providerContracts[0].Id)
	
	// Query contracts by inquiry creator
	creatorContracts, err := k.QueryContractsByInquiryCreator(ctx, inquiry.Creator)
	require.NoError(t, err)
	require.Len(t, creatorContracts, 1)
	require.Equal(t, contract.Id, creatorContracts[0].Id)
	
	// Query provider performance
	performance, err := k.QueryProviderPerformance(ctx, provider)
	require.NoError(t, err)
	require.Equal(t, 1, performance["total_contracts"])
	require.Equal(t, 1, performance["active_contracts"])
	require.Equal(t, 0, performance["completed_contracts"])
}

// TestMultipleConcurrentContracts tests handling of multiple contracts simultaneously
func TestMultipleConcurrentContracts(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create multiple providers with stakes
	providers := []string{"cosmos1provider1", "cosmos1provider2", "cosmos1provider3"}
	for i, provider := range providers {
		stakeAmount := sdk.NewCoin("utoken", math.NewInt(int64(5000*(i+1))))
		k.SetProviderStake(ctx, provider, stakeAmount, uint64(50+i*10))
	}
	
	// Create multiple inquiries with escrow
	inquiries := make([]types.HostingInquiry, 3)
	for i := 0; i < 3; i++ {
		inquiry := types.HostingInquiry{
			Id:               uint64(i + 1),
			Creator:          "cosmos1creator" + fmt.Sprintf("%d", i+1),
			FileEntryCid:     "QmTest" + fmt.Sprintf("%d", i+1),
			ReplicationRate:  uint64(i + 1),
			EscrowAmount:     sdk.NewCoin("utoken", math.NewInt(int64(2000*(i+1)))),
			EndTime:          uint64(3700 + i*10),
			MaxPricePerBlock: uint64(10 * (i + 1)),
		}
		inquiries[i] = inquiry
		k.SetHostingInquiry(ctx, inquiry)
		
		// Set escrow for each inquiry
		escrowAmount := sdk.NewCoin("utoken", math.NewInt(int64(2000*(i+1))))
		k.SetEscrowRecord(ctx, inquiry.Id, escrowAmount, inquiry.Creator)
	}
	
	// Create offers and contracts for each inquiry
	contracts := make([]types.HostingContract, 3)
	for i := 0; i < 3; i++ {
		// Create offer
		offer := types.HostingOffer{
			Id:            uint64(i + 1),
			Region:        "region-" + fmt.Sprintf("%d", i+1),
			PricePerBlock: sdk.NewCoin("utoken", math.NewInt(int64(i+1))),
			Creator:       providers[i],
			InquiryId:     inquiries[i].Id,
		}
		k.SetHostingOffer(ctx, offer)
		
		// Create contract
		contract := types.HostingContract{
			Id:         uint64(i + 1),
			InquiryId:  inquiries[i].Id,
			OfferId:    offer.Id,
			Creator:    inquiries[i].Creator,
			StartBlock: uint64(100 + i*10),
			EndBlock:   inquiries[i].EndTime,
		}
		contracts[i] = contract
		k.SetHostingContract(ctx, contract)
		
		// Initialize payment history
		paymentHistory := types.PaymentHistory{
			ContractId:          contract.Id,
			TotalPaid:           sdk.NewCoin("utoken", math.ZeroInt()),
			LastPaymentBlock:    contract.StartBlock,
			CompletionBonusPaid: false,
		}
		k.SetPaymentHistory(ctx, paymentHistory)
	}
	
	// Verify all contracts exist
	allContracts := k.GetAllHostingContract(ctx)
	require.Len(t, allContracts, 3)
	
	// Verify all payment histories exist
	allPayments := k.GetAllPaymentHistory(ctx)
	require.Len(t, allPayments, 3)
	
	// Verify all escrow records exist
	allEscrow := k.GetAllEscrowRecords(ctx)
	require.Len(t, allEscrow, 3)
	
	// Test system statistics with multiple entities
	stats, err := k.QuerySystemStatistics(ctx)
	require.NoError(t, err)
	require.Equal(t, 3, stats["total_hosting_inquiries"])
	require.Equal(t, 3, stats["total_hosting_offers"])
	require.Equal(t, 3, stats["total_hosting_contracts"])
	require.Equal(t, 3, stats["total_payment_histories"])
	require.Equal(t, 3, stats["total_escrow_records"])
	require.Equal(t, 3, stats["total_providers"])
}

// TestProviderStakingWorkflow tests the complete provider staking workflow
func TestProviderStakingWorkflow(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	
	// Step 1: Initial staking
	initialStake := sdk.NewCoin("utoken", math.NewInt(5000))
	k.SetProviderStake(ctx, provider, initialStake, 100)
	
	// Verify initial stake
	stake, found := k.GetProviderStake(ctx, provider)
	require.True(t, found)
	require.Equal(t, initialStake, stake.Amount)
	
	// Step 2: Increment stake
	increment := sdk.NewCoin("utoken", math.NewInt(2000))
	err := k.IncrementProviderStake(ctx, provider, increment)
	require.NoError(t, err)
	
	// Verify incremented stake
	stake, found = k.GetProviderStake(ctx, provider)
	require.True(t, found)
	expectedAmount := sdk.NewCoin("utoken", math.NewInt(7000))
	require.Equal(t, expectedAmount, stake.Amount)
	
	// Step 3: Test minimum stake validation
	minStake := sdk.NewCoin("utoken", math.NewInt(6000))
	qualifiedProviders := k.GetProvidersByMinStake(ctx, minStake)
	require.Len(t, qualifiedProviders, 1)
	require.Equal(t, provider, qualifiedProviders[0].Provider)
	
	// Step 4: Partial decrement
	decrement := sdk.NewCoin("utoken", math.NewInt(1000))
	err = k.DecrementProviderStake(ctx, provider, decrement)
	require.NoError(t, err)
	
	// Verify decremented stake
	stake, found = k.GetProviderStake(ctx, provider)
	require.True(t, found)
	expectedAmount = sdk.NewCoin("utoken", math.NewInt(6000))
	require.Equal(t, expectedAmount, stake.Amount)
	
	// Step 5: Query stake statistics
	stats := k.GetProviderStakeStats(ctx, "utoken")
	require.Equal(t, 1, stats["count"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(6000)), stats["total"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(6000)), stats["average"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(6000)), stats["min"])
	require.Equal(t, sdk.NewCoin("utoken", math.NewInt(6000)), stats["max"])
	
	// Step 6: Complete unstaking (remove stake)
	fullDecrement := sdk.NewCoin("utoken", math.NewInt(6000))
	err = k.DecrementProviderStake(ctx, provider, fullDecrement)
	require.NoError(t, err)
	
	// Verify stake is removed
	_, found = k.GetProviderStake(ctx, provider)
	require.False(t, found)
}

// TestPaymentProcessingWorkflow tests payment processing and tracking
func TestPaymentProcessingWorkflow(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Set up a contract with payment history
	contractId := uint64(1)
	paymentHistory := types.PaymentHistory{
		ContractId:          contractId,
		TotalPaid:           sdk.NewCoin("utoken", math.ZeroInt()),
		LastPaymentBlock:    100,
		CompletionBonusPaid: false,
	}
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Step 1: Add periodic payments
	payment1 := sdk.NewCoin("utoken", math.NewInt(500))
	err := k.AddPaymentAmount(ctx, contractId, payment1)
	require.NoError(t, err)
	
	// Verify first payment
	updated, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	require.Equal(t, payment1, updated.TotalPaid)
	require.False(t, updated.CompletionBonusPaid)
	
	// Step 2: Add more payments
	payment2 := sdk.NewCoin("utoken", math.NewInt(300))
	err = k.AddPaymentAmount(ctx, contractId, payment2)
	require.NoError(t, err)
	
	// Verify accumulated payments
	updated, found = k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	expectedTotal := sdk.NewCoin("utoken", math.NewInt(800))
	require.Equal(t, expectedTotal, updated.TotalPaid)
	
	// Step 3: Update last payment block
	newBlock := uint64(200)
	err = k.UpdateLastPaymentBlock(ctx, contractId, newBlock)
	require.NoError(t, err)
	
	// Verify block update
	updated, found = k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	require.Equal(t, newBlock, updated.LastPaymentBlock)
	
	// Step 4: Mark completion bonus as paid
	updated.CompletionBonusPaid = true
	k.SetPaymentHistory(ctx, updated)
	
	// Verify completion bonus status
	final, found := k.GetPaymentHistory(ctx, contractId)
	require.True(t, found)
	require.True(t, final.CompletionBonusPaid)
	
	// Step 5: Test payment queries
	// Query by block range
	rangePayments := k.GetPaymentHistoryByBlockRange(ctx, 150, 250)
	require.Len(t, rangePayments, 1)
	require.Equal(t, contractId, rangePayments[0].ContractId)
	
	// Test format payment info
	formatted := k.FormatPaymentHistoryInfo(final)
	require.Equal(t, "1", formatted["contract_id"])
	require.Equal(t, "800utoken", formatted["total_paid"])
	require.Equal(t, "true", formatted["completion_bonus_paid"])
}

// TestStateCleanupWorkflow tests the state cleanup functionality
func TestStateCleanupWorkflow(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create old completed payment history
	oldPayment := types.PaymentHistory{
		ContractId:          1,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(1000)),
		LastPaymentBlock:    50, // Old block
		CompletionBonusPaid: true,
	}
	k.SetPaymentHistory(ctx, oldPayment)
	
	// Create recent pending payment history
	recentPayment := types.PaymentHistory{
		ContractId:          2,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(2000)),
		LastPaymentBlock:    200, // Recent block
		CompletionBonusPaid: false,
	}
	k.SetPaymentHistory(ctx, recentPayment)
	
	// Create recent completed payment history
	recentCompleted := types.PaymentHistory{
		ContractId:          3,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(3000)),
		LastPaymentBlock:    300, // Recent block
		CompletionBonusPaid: true,
	}
	k.SetPaymentHistory(ctx, recentCompleted)
	
	// Verify all payments exist before cleanup
	allPayments := k.GetAllPaymentHistory(ctx)
	require.Len(t, allPayments, 3)
	
	// Perform cleanup (older than 100 blocks)
	err := k.CleanupCompletedPaymentHistory(ctx, 100)
	require.NoError(t, err)
	
	// Verify cleanup results
	remainingPayments := k.GetAllPaymentHistory(ctx)
	require.Len(t, remainingPayments, 2) // Should have 2 remaining (recent ones)
	
	// Verify the old completed payment was removed
	_, found := k.GetPaymentHistory(ctx, 1)
	require.False(t, found)
	
	// Verify recent payments remain
	_, found = k.GetPaymentHistory(ctx, 2)
	require.True(t, found)
	_, found = k.GetPaymentHistory(ctx, 3)
	require.True(t, found)
	
	// Test cleanup status
	status := k.GetCleanupStatus(ctx)
	require.Contains(t, status, "total_payments")
	require.Contains(t, status, "old_payment_histories")
	require.Equal(t, 2, status["total_payments"])
}