package keeper

import (
	"context"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PaymentHistory returns payment history for a specific contract
func (k Keeper) PaymentHistory(goCtx context.Context, req *types.QueryPaymentHistoryRequest) (*types.QueryPaymentHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	paymentHistory, err := k.GetPaymentHistoryForContract(ctx, req.ContractId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryPaymentHistoryResponse{
		PaymentHistory: paymentHistory,
	}, nil
}

// PaymentHistoryAll returns all payment history records
func (k Keeper) PaymentHistoryAll(goCtx context.Context, req *types.QueryAllPaymentHistoryRequest) (*types.QueryAllPaymentHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	paymentHistories := k.GetAllPaymentHistories(ctx)

	return &types.QueryAllPaymentHistoryResponse{
		PaymentHistory: paymentHistories,
	}, nil
}

// EscrowRecord returns escrow record for a specific inquiry
func (k Keeper) EscrowRecord(goCtx context.Context, req *types.QueryEscrowRecordRequest) (*types.QueryEscrowRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Get the keeper's EscrowRecord
	keeperRecord, err := k.GetEscrowStatusForInquiry(ctx, req.InquiryId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Convert keeper's EscrowRecord to protobuf EscrowRecord
	protoRecord := types.EscrowRecord{
		InquiryId: keeperRecord.InquiryId,
		Amount:    keeperRecord.Amount,
		Creator:   keeperRecord.Creator,
	}

	return &types.QueryEscrowRecordResponse{
		EscrowRecord: protoRecord,
	}, nil
}

// EscrowRecordAll returns all escrow records
func (k Keeper) EscrowRecordAll(goCtx context.Context, req *types.QueryAllEscrowRecordRequest) (*types.QueryAllEscrowRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Get the keeper's EscrowRecords
	keeperRecords := k.GetAllEscrowStatuses(ctx)
	
	// Convert to protobuf EscrowRecords
	var protoRecords []types.EscrowRecord
	for _, keeperRecord := range keeperRecords {
		protoRecord := types.EscrowRecord{
			InquiryId: keeperRecord.InquiryId,
			Amount:    keeperRecord.Amount,
			Creator:   keeperRecord.Creator,
		}
		protoRecords = append(protoRecords, protoRecord)
	}

	return &types.QueryAllEscrowRecordResponse{
		EscrowRecord: protoRecords,
	}, nil
}

// ProviderStake returns provider stake for a specific address
func (k Keeper) ProviderStake(goCtx context.Context, req *types.QueryProviderStakeRequest) (*types.QueryProviderStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Get the keeper's ProviderStake
	keeperStake, err := k.GetProviderStakeByAddress(ctx, req.Provider)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Convert keeper's ProviderStake to protobuf ProviderStake
	protoStake := types.ProviderStake{
		Provider: keeperStake.Provider,
		Amount:   keeperStake.Amount,
		Height:   keeperStake.Height,
	}

	return &types.QueryProviderStakeResponse{
		ProviderStake: protoStake,
	}, nil
}

// ProviderStakeAll returns all provider stakes
func (k Keeper) ProviderStakeAll(goCtx context.Context, req *types.QueryAllProviderStakeRequest) (*types.QueryAllProviderStakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Get the keeper's ProviderStakes
	keeperStakes := k.GetAllProviderStakes(ctx)
	
	// Convert to protobuf ProviderStakes
	var protoStakes []types.ProviderStake
	for _, keeperStake := range keeperStakes {
		protoStake := types.ProviderStake{
			Provider: keeperStake.Provider,
			Amount:   keeperStake.Amount,
			Height:   keeperStake.Height,
		}
		protoStakes = append(protoStakes, protoStake)
	}

	return &types.QueryAllProviderStakeResponse{
		ProviderStake: protoStakes,
	}, nil
}