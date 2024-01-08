package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func (k msgServer) CreateHostingContract(goCtx context.Context, msg *types.MsgCreateHostingContract) (*types.MsgCreateHostingContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingContract = types.HostingContract{
		Creator:   msg.Creator,
		InquiryId: msg.InquiryId,
		OfferId:   msg.OfferId,
	}

	id := k.AppendHostingContract(
		ctx,
		hostingContract,
	)

	return &types.MsgCreateHostingContractResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHostingContract(goCtx context.Context, msg *types.MsgUpdateHostingContract) (*types.MsgUpdateHostingContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingContract = types.HostingContract{
		Creator:   msg.Creator,
		Id:        msg.Id,
		InquiryId: msg.InquiryId,
		OfferId:   msg.OfferId,
	}

	// Checks that the element exists
	val, found := k.GetHostingContract(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHostingContract(ctx, hostingContract)

	return &types.MsgUpdateHostingContractResponse{}, nil
}

func (k msgServer) DeleteHostingContract(goCtx context.Context, msg *types.MsgDeleteHostingContract) (*types.MsgDeleteHostingContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHostingContract(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHostingContract(ctx, msg.Id)

	return &types.MsgDeleteHostingContractResponse{}, nil
}
