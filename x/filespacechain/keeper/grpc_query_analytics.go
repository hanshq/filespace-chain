package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// PaymentAnalytics implements the PaymentAnalytics gRPC method
func (k Keeper) PaymentAnalytics(goCtx context.Context, req *types.QueryPaymentAnalyticsRequest) (*types.QueryPaymentAnalyticsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	analytics := k.QueryPaymentAnalytics(ctx)

	return &types.QueryPaymentAnalyticsResponse{
		TotalContracts:    analytics.TotalContracts,
		TotalPayments:     analytics.TotalPayments, 
		TotalAmount:       analytics.TotalAmount,
		ActiveContracts:   analytics.ActiveContracts,
		CompletedContracts: analytics.CompletedContracts,
		TotalEscrow:       analytics.TotalEscrow,
	}, nil
}

// PaymentsByContract implements the PaymentsByContract gRPC method
func (k Keeper) PaymentsByContract(goCtx context.Context, req *types.QueryPaymentsByContractRequest) (*types.QueryPaymentsByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	payments := k.QueryPaymentsByContract(ctx, req.ContractId)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range payments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryPaymentsByContractResponse{
		Payments: paymentInfos,
	}, nil
}

// PaymentsByProvider implements the PaymentsByProvider gRPC method
func (k Keeper) PaymentsByProvider(goCtx context.Context, req *types.QueryPaymentsByProviderRequest) (*types.QueryPaymentsByProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	payments, totalEarnings := k.QueryPaymentsByProvider(ctx, req.Provider)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range payments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryPaymentsByProviderResponse{
		Payments:      paymentInfos,
		TotalEarnings: totalEarnings,
	}, nil
}

// PaymentsByBlockRange implements the PaymentsByBlockRange gRPC method
func (k Keeper) PaymentsByBlockRange(goCtx context.Context, req *types.QueryPaymentsByBlockRangeRequest) (*types.QueryPaymentsByBlockRangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	payments := k.QueryPaymentsByBlockRange(ctx, req.StartBlock, req.EndBlock)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range payments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryPaymentsByBlockRangeResponse{
		Payments: paymentInfos,
	}, nil
}

// RecentPayments implements the RecentPayments gRPC method
func (k Keeper) RecentPayments(goCtx context.Context, req *types.QueryRecentPaymentsRequest) (*types.QueryRecentPaymentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	payments := k.QueryRecentPayments(ctx, req.Limit)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range payments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryRecentPaymentsResponse{
		Payments: paymentInfos,
	}, nil
}

// PendingPayments implements the PendingPayments gRPC method
func (k Keeper) PendingPayments(goCtx context.Context, req *types.QueryPendingPaymentsRequest) (*types.QueryPendingPaymentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryPendingPayments(ctx)

	return &types.QueryPendingPaymentsResponse{
		Contracts: contracts,
	}, nil
}

// CompletedPayments implements the CompletedPayments gRPC method
func (k Keeper) CompletedPayments(goCtx context.Context, req *types.QueryCompletedPaymentsRequest) (*types.QueryCompletedPaymentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryCompletedPayments(ctx)

	return &types.QueryCompletedPaymentsResponse{
		Contracts: contracts,
	}, nil
}

// PaymentDistribution implements the PaymentDistribution gRPC method
func (k Keeper) PaymentDistribution(goCtx context.Context, req *types.QueryPaymentDistributionRequest) (*types.QueryPaymentDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	distribution := k.QueryPaymentDistribution(ctx)

	return &types.QueryPaymentDistributionResponse{
		TotalDistributed:  distribution.TotalDistributed,
		TotalPending:      distribution.TotalPending,
		DistributionCount: distribution.DistributionCount,
	}, nil
}

// ProviderPaymentSummary implements the ProviderPaymentSummary gRPC method
func (k Keeper) ProviderPaymentSummary(goCtx context.Context, req *types.QueryProviderPaymentSummaryRequest) (*types.QueryProviderPaymentSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	summary := k.QueryProviderPaymentSummary(ctx, req.Provider)

	return &types.QueryProviderPaymentSummaryResponse{
		Provider:           summary.Provider,
		TotalEarnings:      summary.TotalEarnings,
		TotalContracts:     summary.TotalContracts,
		ActiveContracts:    summary.ActiveContracts,
		CompletedContracts: summary.CompletedContracts,
	}, nil
}

// PaymentTrends implements the PaymentTrends gRPC method
func (k Keeper) PaymentTrends(goCtx context.Context, req *types.QueryPaymentTrendsRequest) (*types.QueryPaymentTrendsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	trends := k.QueryPaymentTrends(ctx, req.Blocks)

	var paymentInfos []*types.PaymentInfo
	for _, payment := range trends.RecentPayments {
		paymentInfos = append(paymentInfos, &types.PaymentInfo{
			ContractId:  payment.ContractId,
			Provider:    payment.Provider,
			Amount:      payment.Amount,
			BlockHeight: payment.BlockHeight,
			PaymentType: payment.PaymentType,
		})
	}

	return &types.QueryPaymentTrendsResponse{
		RecentPayments: paymentInfos,
		TrendTotal:     trends.TrendTotal,
	}, nil
}

// UnpaidContracts implements the UnpaidContracts gRPC method
func (k Keeper) UnpaidContracts(goCtx context.Context, req *types.QueryUnpaidContractsRequest) (*types.QueryUnpaidContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts := k.QueryUnpaidContracts(ctx)

	return &types.QueryUnpaidContractsResponse{
		Contracts: contracts,
	}, nil
}