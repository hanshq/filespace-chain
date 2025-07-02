package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// EscrowSummary implements the EscrowSummary gRPC method
func (k Keeper) EscrowSummary(goCtx context.Context, req *types.QueryEscrowSummaryRequest) (*types.QueryEscrowSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	summary := k.QueryEscrowSummary(ctx)

	var escrowInfos []*types.EscrowInfo
	for _, escrow := range summary.Escrows {
		escrowInfos = append(escrowInfos, &types.EscrowInfo{
			InquiryId:   escrow.InquiryId,
			Creator:     escrow.Creator,
			Amount:      escrow.Amount,
			LockedBlock: escrow.LockedBlock,
			IsActive:    escrow.IsActive,
		})
	}

	return &types.QueryEscrowSummaryResponse{
		Escrows:        escrowInfos,
		TotalEscrowed:  summary.TotalEscrowed,
		ActiveEscrows:  summary.ActiveEscrows,
	}, nil
}

// EscrowByCreator implements the EscrowByCreator gRPC method
func (k Keeper) EscrowByCreator(goCtx context.Context, req *types.QueryEscrowByCreatorRequest) (*types.QueryEscrowByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	escrows := k.QueryEscrowByCreator(ctx, req.Creator)

	var escrowInfos []*types.EscrowInfo
	for _, escrow := range escrows {
		escrowInfos = append(escrowInfos, &types.EscrowInfo{
			InquiryId:   escrow.InquiryId,
			Creator:     escrow.Creator,
			Amount:      escrow.Amount,
			LockedBlock: escrow.LockedBlock,
			IsActive:    escrow.IsActive,
		})
	}

	return &types.QueryEscrowByCreatorResponse{
		Escrows: escrowInfos,
	}, nil
}

// EscrowByInquiry implements the EscrowByInquiry gRPC method
func (k Keeper) EscrowByInquiry(goCtx context.Context, req *types.QueryEscrowByInquiryRequest) (*types.QueryEscrowByInquiryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	escrow := k.QueryEscrowByInquiry(ctx, req.InquiryId)

	return &types.QueryEscrowByInquiryResponse{
		Escrow: &types.EscrowInfo{
			InquiryId:   escrow.InquiryId,
			Creator:     escrow.Creator,
			Amount:      escrow.Amount,
			LockedBlock: escrow.LockedBlock,
			IsActive:    escrow.IsActive,
		},
	}, nil
}

// ActiveEscrow implements the ActiveEscrow gRPC method
func (k Keeper) ActiveEscrow(goCtx context.Context, req *types.QueryActiveEscrowRequest) (*types.QueryActiveEscrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	escrows := k.QueryActiveEscrow(ctx)

	var escrowInfos []*types.EscrowInfo
	for _, escrow := range escrows {
		escrowInfos = append(escrowInfos, &types.EscrowInfo{
			InquiryId:   escrow.InquiryId,
			Creator:     escrow.Creator,
			Amount:      escrow.Amount,
			LockedBlock: escrow.LockedBlock,
			IsActive:    escrow.IsActive,
		})
	}

	return &types.QueryActiveEscrowResponse{
		Escrows: escrowInfos,
	}, nil
}

// EscrowTotalsByDenom implements the EscrowTotalsByDenom gRPC method
func (k Keeper) EscrowTotalsByDenom(goCtx context.Context, req *types.QueryEscrowTotalsByDenomRequest) (*types.QueryEscrowTotalsByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	totalAmount, escrowCount := k.QueryEscrowTotalsByDenom(ctx, req.Denom)

	return &types.QueryEscrowTotalsByDenomResponse{
		TotalAmount: totalAmount,
		EscrowCount: escrowCount,
	}, nil
}

// StakingSummary implements the StakingSummary gRPC method
func (k Keeper) StakingSummary(goCtx context.Context, req *types.QueryStakingSummaryRequest) (*types.QueryStakingSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	summary := k.QueryStakingSummary(ctx)

	var stakeInfos []*types.StakeInfo
	for _, stake := range summary.Stakes {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryStakingSummaryResponse{
		Stakes:         stakeInfos,
		TotalStaked:    summary.TotalStaked,
		ValidProviders: summary.ValidProviders,
	}, nil
}

// StakeByProvider implements the StakeByProvider gRPC method
func (k Keeper) StakeByProvider(goCtx context.Context, req *types.QueryStakeByProviderRequest) (*types.QueryStakeByProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	stake := k.QueryStakeByProvider(ctx, req.Provider)

	return &types.QueryStakeByProviderResponse{
		Stake: &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		},
	}, nil
}

// ProvidersWithMinStake implements the ProvidersWithMinStake gRPC method
func (k Keeper) ProvidersWithMinStake(goCtx context.Context, req *types.QueryProvidersWithMinStakeRequest) (*types.QueryProvidersWithMinStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	providers := k.QueryProvidersWithMinStake(ctx)

	var stakeInfos []*types.StakeInfo
	for _, stake := range providers {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryProvidersWithMinStakeResponse{
		Providers: stakeInfos,
	}, nil
}

// ProvidersByStakeRange implements the ProvidersByStakeRange gRPC method
func (k Keeper) ProvidersByStakeRange(goCtx context.Context, req *types.QueryProvidersByStakeRangeRequest) (*types.QueryProvidersByStakeRangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	providers := k.QueryProvidersByStakeRange(ctx, req.StartBlock, req.EndBlock)

	var stakeInfos []*types.StakeInfo
	for _, stake := range providers {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryProvidersByStakeRangeResponse{
		Providers: stakeInfos,
	}, nil
}

// StakeStatistics implements the StakeStatistics gRPC method
func (k Keeper) StakeStatistics(goCtx context.Context, req *types.QueryStakeStatisticsRequest) (*types.QueryStakeStatisticsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	stats := k.QueryStakeStatistics(ctx, req.Denom)

	return &types.QueryStakeStatisticsResponse{
		TotalStaked:   stats.TotalStaked,
		AverageStake:  stats.AverageStake,
		MinStake:      stats.MinStake,
		MaxStake:      stats.MaxStake,
		ProviderCount: stats.ProviderCount,
	}, nil
}

// TopProvidersByStake implements the TopProvidersByStake gRPC method
func (k Keeper) TopProvidersByStake(goCtx context.Context, req *types.QueryTopProvidersByStakeRequest) (*types.QueryTopProvidersByStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	providers := k.QueryTopProvidersByStake(ctx, req.Limit)

	var stakeInfos []*types.StakeInfo
	for _, stake := range providers {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryTopProvidersByStakeResponse{
		Providers: stakeInfos,
	}, nil
}

// ProviderStakeHistory implements the ProviderStakeHistory gRPC method
func (k Keeper) ProviderStakeHistory(goCtx context.Context, req *types.QueryProviderStakeHistoryRequest) (*types.QueryProviderStakeHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	history := k.QueryProviderStakeHistory(ctx, req.Provider)

	var stakeInfos []*types.StakeInfo
	for _, stake := range history {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryProviderStakeHistoryResponse{
		History: stakeInfos,
	}, nil
}

// ValidProviders implements the ValidProviders gRPC method
func (k Keeper) ValidProviders(goCtx context.Context, req *types.QueryValidProvidersRequest) (*types.QueryValidProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	providers := k.QueryValidProviders(ctx)

	var stakeInfos []*types.StakeInfo
	for _, stake := range providers {
		stakeInfos = append(stakeInfos, &types.StakeInfo{
			Provider:      stake.Provider,
			StakedAmount:  stake.StakedAmount,
			StakeBlock:    stake.StakeBlock,
			MeetsMinimum:  stake.MeetsMinimum,
		})
	}

	return &types.QueryValidProvidersResponse{
		Providers: stakeInfos,
	}, nil
}