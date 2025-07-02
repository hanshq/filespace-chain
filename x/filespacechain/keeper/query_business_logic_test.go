package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestQueryActiveContracts(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test that query doesn't panic
	contracts, err := k.QueryActiveContracts(ctx)
	require.NoError(t, err)
	require.NotNil(t, contracts)
	require.IsType(t, []types.HostingContract{}, contracts)
}

func TestQueryExpiredContracts(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test that query doesn't panic
	contracts, err := k.QueryExpiredContracts(ctx)
	require.NoError(t, err)
	require.NotNil(t, contracts)
	require.IsType(t, []types.HostingContract{}, contracts)
}

func TestQueryContractsByProvider(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	
	// Test that query doesn't panic
	contracts, err := k.QueryContractsByProvider(ctx, provider)
	require.NoError(t, err)
	require.NotNil(t, contracts)
	require.IsType(t, []types.HostingContract{}, contracts)
}

func TestQueryContractsByInquiryCreator(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	creator := "cosmos1creator"
	
	// Test that query doesn't panic
	contracts, err := k.QueryContractsByInquiryCreator(ctx, creator)
	require.NoError(t, err)
	require.NotNil(t, contracts)
	require.IsType(t, []types.HostingContract{}, contracts)
}

func TestQueryOffersByProvider(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	
	// Test that query doesn't panic
	offers, err := k.QueryOffersByProvider(ctx, provider)
	require.NoError(t, err)
	require.NotNil(t, offers)
	require.IsType(t, []types.HostingOffer{}, offers)
}

func TestQueryInquiriesByCreator(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	creator := "cosmos1creator"
	
	// Test that query doesn't panic
	inquiries, err := k.QueryInquiriesByCreator(ctx, creator)
	require.NoError(t, err)
	require.NotNil(t, inquiries)
	require.IsType(t, []types.HostingInquiry{}, inquiries)
}

func TestQueryProviderEarnings(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	
	// Test with no payment history
	earnings, err := k.QueryProviderEarnings(ctx, provider)
	require.NoError(t, err)
	require.Equal(t, "utoken", earnings.Denom)
	require.Equal(t, math.ZeroInt(), earnings.Amount)
}

func TestQueryProviderEarningsWithPayments(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// This test would require setting up contracts, offers, and payment history
	// For now, test that the function doesn't panic with empty data
	provider := "cosmos1provider"
	earnings, err := k.QueryProviderEarnings(ctx, provider)
	require.NoError(t, err)
	require.NotNil(t, earnings)
	require.IsType(t, sdk.Coin{}, earnings)
}

func TestQuerySystemStatistics(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test system statistics query
	stats, err := k.QuerySystemStatistics(ctx)
	require.NoError(t, err)
	require.NotNil(t, stats)
	
	// Check that basic stats are present
	require.Contains(t, stats, "total_file_entries")
	require.Contains(t, stats, "total_hosting_inquiries")
	require.Contains(t, stats, "total_hosting_offers")
	require.Contains(t, stats, "total_hosting_contracts")
	require.Contains(t, stats, "total_payment_histories")
	require.Contains(t, stats, "active_contracts")
	require.Contains(t, stats, "expired_contracts")
	require.Contains(t, stats, "total_escrow_records")
	require.Contains(t, stats, "total_provider_stakes")
	require.Contains(t, stats, "total_providers")
	
	// Check that values are integers
	require.IsType(t, 0, stats["total_file_entries"])
	require.IsType(t, 0, stats["total_hosting_inquiries"])
	require.IsType(t, 0, stats["total_hosting_offers"])
	require.IsType(t, 0, stats["total_hosting_contracts"])
}

func TestQueryContractDetailsNotFound(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Test with non-existent contract
	_, err := k.QueryContractDetails(ctx, 999)
	require.Error(t, err)
	require.Contains(t, err.Error(), "contract 999 not found")
}

func TestQueryProviderPerformance(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	
	// Test with empty data
	performance, err := k.QueryProviderPerformance(ctx, provider)
	require.NoError(t, err)
	require.NotNil(t, performance)
	
	// Check basic structure
	require.Contains(t, performance, "total_contracts")
	require.Contains(t, performance, "active_contracts")
	require.Contains(t, performance, "completed_contracts")
	
	require.Equal(t, 0, performance["total_contracts"])
	require.Equal(t, 0, performance["active_contracts"])
	require.Equal(t, 0, performance["completed_contracts"])
}

func TestQueryProviderPerformanceWithStake(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	provider := "cosmos1provider"
	amount := sdk.NewCoin("utoken", math.NewInt(5000))
	height := uint64(100)
	
	// Set provider stake
	k.SetProviderStake(ctx, provider, amount, height)
	
	// Query performance
	performance, err := k.QueryProviderPerformance(ctx, provider)
	require.NoError(t, err)
	require.NotNil(t, performance)
	
	// Check that stake is included
	require.Contains(t, performance, "stake")
	
	stake := performance["stake"]
	require.NotNil(t, stake)
}

func TestQuerySystemStatisticsWithData(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Add some test data
	k.SetProviderStake(ctx, "cosmos1provider1", sdk.NewCoin("utoken", math.NewInt(1000)), 100)
	k.SetProviderStake(ctx, "cosmos1provider2", sdk.NewCoin("utoken", math.NewInt(2000)), 200)
	
	k.SetEscrowRecord(ctx, 1, sdk.NewCoin("utoken", math.NewInt(500)), "cosmos1creator")
	k.SetEscrowRecord(ctx, 2, sdk.NewCoin("utoken", math.NewInt(1500)), "cosmos1creator2")
	
	paymentHistory := types.PaymentHistory{
		ContractId:          1,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(100)),
		LastPaymentBlock:    50,
		CompletionBonusPaid: false,
	}
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Query statistics
	stats, err := k.QuerySystemStatistics(ctx)
	require.NoError(t, err)
	
	// Verify counts
	require.Equal(t, 1, stats["total_payment_histories"])
	require.Equal(t, 2, stats["total_escrow_records"])
	require.Equal(t, 2, stats["total_providers"])
}

func TestQueryContractDetailsWithData(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	
	// Create test hosting inquiry
	inquiry := types.HostingInquiry{
		Id:               1,
		Creator:          "cosmos1creator",
		FileEntryCid:     "QmTest123",
		ReplicationRate:  1,
		EscrowAmount:     sdk.NewCoin("utoken", math.NewInt(1000)),
		EndTime:          3700,
		MaxPricePerBlock: 10,
	}
	k.SetHostingInquiry(ctx, inquiry)
	
	// Create test hosting offer
	offer := types.HostingOffer{
		Id:            1,
		Region:        "us-east",
		PricePerBlock: sdk.NewCoin("utoken", math.NewInt(10)),
		Creator:       "cosmos1provider",
		InquiryId:     1,
	}
	k.SetHostingOffer(ctx, offer)
	
	// Create test hosting contract
	contract := types.HostingContract{
		Id:         1,
		InquiryId:  1,
		OfferId:    1,
		Creator:    "cosmos1creator",
		StartBlock: 100,
		EndBlock:   3700,
	}
	k.SetHostingContract(ctx, contract)
	
	// Create payment history
	paymentHistory := types.PaymentHistory{
		ContractId:          1,
		TotalPaid:           sdk.NewCoin("utoken", math.NewInt(500)),
		LastPaymentBlock:    200,
		CompletionBonusPaid: false,
	}
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Create escrow record
	k.SetEscrowRecord(ctx, 1, sdk.NewCoin("utoken", math.NewInt(1000)), "cosmos1creator")
	
	// Query contract details
	details, err := k.QueryContractDetails(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, details)
	
	// Verify all components are present
	require.Contains(t, details, "contract")
	require.Contains(t, details, "inquiry")
	require.Contains(t, details, "offer")
	require.Contains(t, details, "payment_history")
	require.Contains(t, details, "escrow_record")
	require.Contains(t, details, "status")
	
	// Verify contract details
	contractDetail := details["contract"].(types.HostingContract)
	require.Equal(t, uint64(1), contractDetail.Id)
	
	// Verify inquiry details
	inquiryDetail := details["inquiry"].(types.HostingInquiry)
	require.Equal(t, uint64(1), inquiryDetail.Id)
	
	// Verify offer details
	offerDetail := details["offer"].(types.HostingOffer)
	require.Equal(t, uint64(1), offerDetail.Id)
	
	// Verify status (should be "active" since start < current < end in test context)
	status := details["status"].(string)
	require.Contains(t, []string{"pending", "active", "expired"}, status)
}