package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// ActiveContracts implements the ActiveContracts gRPC method
func (k Keeper) ActiveContracts(goCtx context.Context, req *types.QueryActiveContractsRequest) (*types.QueryActiveContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryActiveContracts(ctx)

	return &types.QueryActiveContractsResponse{
		Contracts: contracts,
	}, nil
}

// ExpiredContracts implements the ExpiredContracts gRPC method
func (k Keeper) ExpiredContracts(goCtx context.Context, req *types.QueryExpiredContractsRequest) (*types.QueryExpiredContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryExpiredContracts(ctx)

	return &types.QueryExpiredContractsResponse{
		Contracts: contracts,
	}, nil
}

// ContractsByProvider implements the ContractsByProvider gRPC method
func (k Keeper) ContractsByProvider(goCtx context.Context, req *types.QueryContractsByProviderRequest) (*types.QueryContractsByProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryContractsByProvider(ctx, req.Provider)

	return &types.QueryContractsByProviderResponse{
		Contracts: contracts,
	}, nil
}

// ContractsByInquiryCreator implements the ContractsByInquiryCreator gRPC method
func (k Keeper) ContractsByInquiryCreator(goCtx context.Context, req *types.QueryContractsByInquiryCreatorRequest) (*types.QueryContractsByInquiryCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryContractsByInquiryCreator(ctx, req.Creator)

	return &types.QueryContractsByInquiryCreatorResponse{
		Contracts: contracts,
	}, nil
}

// OffersByProvider implements the OffersByProvider gRPC method
func (k Keeper) OffersByProvider(goCtx context.Context, req *types.QueryOffersByProviderRequest) (*types.QueryOffersByProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	offers := k.QueryOffersByProvider(ctx, req.Provider)

	return &types.QueryOffersByProviderResponse{
		Offers: offers,
	}, nil
}

// InquiriesByCreator implements the InquiriesByCreator gRPC method
func (k Keeper) InquiriesByCreator(goCtx context.Context, req *types.QueryInquiriesByCreatorRequest) (*types.QueryInquiriesByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	inquiries := k.QueryInquiriesByCreator(ctx, req.Creator)

	return &types.QueryInquiriesByCreatorResponse{
		Inquiries: inquiries,
	}, nil
}

// ProviderEarnings implements the ProviderEarnings gRPC method
func (k Keeper) ProviderEarnings(goCtx context.Context, req *types.QueryProviderEarningsRequest) (*types.QueryProviderEarningsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	earnings := k.QueryProviderEarnings(ctx, req.Provider)

	return &types.QueryProviderEarningsResponse{
		Provider:        earnings.Provider,
		TotalEarnings:   earnings.TotalEarnings,
		PendingEarnings: earnings.PendingEarnings,
		TotalContracts:  earnings.TotalContracts,
	}, nil
}

// SystemStatistics implements the SystemStatistics gRPC method
func (k Keeper) SystemStatistics(goCtx context.Context, req *types.QuerySystemStatisticsRequest) (*types.QuerySystemStatisticsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	stats := k.QuerySystemStatistics(ctx)

	return &types.QuerySystemStatisticsResponse{
		TotalFileEntries:   stats.TotalFileEntries,
		TotalInquiries:     stats.TotalInquiries,
		TotalOffers:        stats.TotalOffers,
		TotalContracts:     stats.TotalContracts,
		ActiveContracts:    stats.ActiveContracts,
		CompletedContracts: stats.CompletedContracts,
		TotalValueLocked:   stats.TotalValueLocked,
		TotalPayments:      stats.TotalPayments,
		UniqueProviders:    stats.UniqueProviders,
		UniqueClients:      stats.UniqueClients,
	}, nil
}

// ContractDetails implements the ContractDetails gRPC method
func (k Keeper) ContractDetails(goCtx context.Context, req *types.QueryContractDetailsRequest) (*types.QueryContractDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	details := k.QueryContractDetails(ctx, req.ContractId)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range details.Payments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryContractDetailsResponse{
		Details: &types.ContractDetail{
			Contract:        details.Contract,
			Inquiry:         details.Inquiry,
			Offer:          details.Offer,
			Payments:        paymentInfos,
			TotalPaid:       details.TotalPaid,
			RemainingEscrow: details.RemainingEscrow,
			IsActive:        details.IsActive,
			IsCompleted:     details.IsCompleted,
		},
	}, nil
}

// ProviderPerformance implements the ProviderPerformance gRPC method
func (k Keeper) ProviderPerformance(goCtx context.Context, req *types.QueryProviderPerformanceRequest) (*types.QueryProviderPerformanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	performance := k.QueryProviderPerformance(ctx, req.Provider)

	return &types.QueryProviderPerformanceResponse{
		Provider:           performance.Provider,
		TotalOffers:        performance.TotalOffers,
		AcceptedOffers:     performance.AcceptedOffers,
		CompletedContracts: performance.CompletedContracts,
		FailedContracts:    performance.FailedContracts,
		TotalEarnings:      performance.TotalEarnings,
		CurrentStake:       performance.CurrentStake,
		ReputationScore:    performance.ReputationScore,
	}, nil
}