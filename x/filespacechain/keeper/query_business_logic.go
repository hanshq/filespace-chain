package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// QueryActiveContracts returns all currently active hosting contracts
func (k Keeper) QueryActiveContracts(ctx context.Context) ([]types.HostingContract, error) {
	activeContracts := k.GetActiveContracts(ctx)
	return activeContracts, nil
}

// QueryExpiredContracts returns all expired hosting contracts
func (k Keeper) QueryExpiredContracts(ctx context.Context) ([]types.HostingContract, error) {
	expiredContracts := k.GetExpiredContracts(ctx)
	return expiredContracts, nil
}

// QueryContractsByProvider returns all contracts where the provider has an offer
func (k Keeper) QueryContractsByProvider(ctx context.Context, provider string) ([]types.HostingContract, error) {
	allContracts := k.GetAllHostingContract(ctx)
	var providerContracts []types.HostingContract

	for _, contract := range allContracts {
		// Get offers for this contract's inquiry
		offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
		
		// Check if this provider has an offer
		for _, offer := range offers {
			if offer.Creator == provider {
				providerContracts = append(providerContracts, contract)
				break
			}
		}
	}

	return providerContracts, nil
}

// QueryContractsByInquiryCreator returns all contracts for inquiries created by a specific address
func (k Keeper) QueryContractsByInquiryCreator(ctx context.Context, creator string) ([]types.HostingContract, error) {
	allContracts := k.GetAllHostingContract(ctx)
	var creatorContracts []types.HostingContract

	for _, contract := range allContracts {
		inquiry, found := k.GetHostingInquiry(ctx, contract.InquiryId)
		if found && inquiry.Creator == creator {
			creatorContracts = append(creatorContracts, contract)
		}
	}

	return creatorContracts, nil
}

// QueryOffersByProvider returns all hosting offers created by a provider
func (k Keeper) QueryOffersByProvider(ctx context.Context, provider string) ([]types.HostingOffer, error) {
	allOffers := k.GetAllHostingOffer(ctx)
	var providerOffers []types.HostingOffer

	for _, offer := range allOffers {
		if offer.Creator == provider {
			providerOffers = append(providerOffers, offer)
		}
	}

	return providerOffers, nil
}

// QueryInquiriesByCreator returns all hosting inquiries created by a specific address
func (k Keeper) QueryInquiriesByCreator(ctx context.Context, creator string) ([]types.HostingInquiry, error) {
	allInquiries := k.GetAllHostingInquiry(ctx)
	var creatorInquiries []types.HostingInquiry

	for _, inquiry := range allInquiries {
		if inquiry.Creator == creator {
			creatorInquiries = append(creatorInquiries, inquiry)
		}
	}

	return creatorInquiries, nil
}

// QueryProviderEarnings calculates total earnings for a provider from payment history
func (k Keeper) QueryProviderEarnings(ctx context.Context, provider string) (sdk.Coin, error) {
	paymentHistories := k.GetPaymentHistoryByProvider(ctx, provider)
	
	var totalEarnings sdk.Coin
	if len(paymentHistories) == 0 {
		// Return zero coin with default denomination
		return sdk.NewCoin("utoken", math.ZeroInt()), nil
	}

	// Initialize with first payment's denomination
	totalEarnings = sdk.NewCoin(paymentHistories[0].TotalPaid.Denom, math.ZeroInt())
	
	for _, payment := range paymentHistories {
		// Get contract to find number of providers to split payment
		contract, found := k.GetHostingContract(ctx, payment.ContractId)
		if !found {
			continue
		}
		
		// Get offers for this contract to count providers
		offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
		if len(offers) == 0 {
			continue
		}
		
		// Calculate this provider's share (equal split among all providers)
		providerShare := payment.TotalPaid.Amount.Quo(math.NewInt(int64(len(offers))))
		totalEarnings = totalEarnings.Add(sdk.NewCoin(payment.TotalPaid.Denom, providerShare))
	}

	return totalEarnings, nil
}

// QuerySystemStatistics returns comprehensive system statistics
func (k Keeper) QuerySystemStatistics(ctx context.Context) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Basic counts
	stats["total_file_entries"] = len(k.GetAllFileEntry(ctx))
	stats["total_hosting_inquiries"] = len(k.GetAllHostingInquiry(ctx))
	stats["total_hosting_offers"] = len(k.GetAllHostingOffer(ctx))
	stats["total_hosting_contracts"] = len(k.GetAllHostingContract(ctx))
	
	// Payment and state statistics
	paymentSummary := k.QueryPaymentSummary(ctx)
	for key, value := range paymentSummary {
		stats[key] = value
	}

	// Provider statistics
	providerStakes := k.GetAllProviderStakes(ctx)
	stats["total_providers"] = len(providerStakes)
	
	// Escrow statistics
	escrowRecords := k.GetAllEscrowRecords(ctx)
	stats["total_escrow_records"] = len(escrowRecords)

	// Active vs expired contracts
	activeContracts := k.GetActiveContracts(ctx)
	expiredContracts := k.GetExpiredContracts(ctx)
	stats["active_contracts"] = len(activeContracts)
	stats["expired_contracts"] = len(expiredContracts)

	return stats, nil
}

// QueryContractDetails returns detailed information about a specific contract
func (k Keeper) QueryContractDetails(ctx context.Context, contractId uint64) (map[string]interface{}, error) {
	contract, found := k.GetHostingContract(ctx, contractId)
	if !found {
		return nil, fmt.Errorf("contract %d not found", contractId)
	}

	details := make(map[string]interface{})
	details["contract"] = contract

	// Get associated inquiry
	inquiry, found := k.GetHostingInquiry(ctx, contract.InquiryId)
	if found {
		details["inquiry"] = inquiry
	}

	// Get associated offer
	offer, found := k.GetHostingOffer(ctx, contract.OfferId)
	if found {
		details["offer"] = offer
	}

	// Get payment history
	paymentHistory, found := k.GetPaymentHistory(ctx, contractId)
	if found {
		details["payment_history"] = paymentHistory
	}

	// Get escrow record for the inquiry
	escrowRecord, found := k.GetEscrowRecord(ctx, contract.InquiryId)
	if found {
		details["escrow_record"] = escrowRecord
	}

	// Calculate contract status
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	if currentBlock >= contract.EndBlock {
		details["status"] = "expired"
	} else if currentBlock >= contract.StartBlock {
		details["status"] = "active"
	} else {
		details["status"] = "pending"
	}

	return details, nil
}

// QueryProviderPerformance returns performance metrics for a provider
func (k Keeper) QueryProviderPerformance(ctx context.Context, provider string) (map[string]interface{}, error) {
	performance := make(map[string]interface{})

	// Get provider stake
	stake, found := k.GetProviderStake(ctx, provider)
	if found {
		performance["stake"] = stake
	}

	// Get all contracts for this provider
	contracts, err := k.QueryContractsByProvider(ctx, provider)
	if err != nil {
		return nil, err
	}
	performance["total_contracts"] = len(contracts)

	// Count active vs completed contracts
	activeCount := 0
	completedCount := 0
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())

	for _, contract := range contracts {
		if currentBlock >= contract.EndBlock {
			completedCount++
		} else {
			activeCount++
		}
	}
	performance["active_contracts"] = activeCount
	performance["completed_contracts"] = completedCount

	// Get total earnings
	earnings, err := k.QueryProviderEarnings(ctx, provider)
	if err == nil {
		performance["total_earnings"] = earnings
	}

	// Get all offers by this provider
	offers, err := k.QueryOffersByProvider(ctx, provider)
	if err == nil {
		performance["total_offers"] = len(offers)
	}

	return performance, nil
}