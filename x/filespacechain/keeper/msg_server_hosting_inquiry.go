package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func (k msgServer) CreateHostingInquiry(goCtx context.Context, msg *types.MsgCreateHostingInquiry) (*types.MsgCreateHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingInquiry = types.HostingInquiry{
		Creator:         msg.Creator,
		FileEntryCid:    msg.FileEntryCid,
		ReplicationRate: msg.ReplicationRate,
		EscrowAmount:    msg.EscrowAmount,
		EndTime:         msg.EndTime,
	}

	id := k.AppendHostingInquiry(
		ctx,
		hostingInquiry,
	)

	return &types.MsgCreateHostingInquiryResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHostingInquiry(goCtx context.Context, msg *types.MsgUpdateHostingInquiry) (*types.MsgUpdateHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingInquiry = types.HostingInquiry{
		Creator:         msg.Creator,
		Id:              msg.Id,
		FileEntryCid:    msg.FileEntryCid,
		ReplicationRate: msg.ReplicationRate,
		EscrowAmount:    msg.EscrowAmount,
		EndTime:         msg.EndTime,
	}

	// Checks that the element exists
	val, found := k.GetHostingInquiry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHostingInquiry(ctx, hostingInquiry)

	return &types.MsgUpdateHostingInquiryResponse{}, nil
}

func (k msgServer) DeleteHostingInquiry(goCtx context.Context, msg *types.MsgDeleteHostingInquiry) (*types.MsgDeleteHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHostingInquiry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHostingInquiry(ctx, msg.Id)

	return &types.MsgDeleteHostingInquiryResponse{}, nil
}
