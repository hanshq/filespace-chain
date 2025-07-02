package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func (k msgServer) CreateHostingOffer(goCtx context.Context, msg *types.MsgCreateHostingOffer) (*types.MsgCreateHostingOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert creator address for validation
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	// Validate that provider has sufficient stake before allowing offer creation
	err = k.ValidateProviderStake(goCtx, creatorAddr)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
			fmt.Sprintf("provider stake validation failed: %s", err.Error()))
	}

	var hostingOffer = types.HostingOffer{
		Creator:       msg.Creator,
		Region:        msg.Region,
		PricePerBlock: msg.PricePerBlock,
	}

	id := k.AppendHostingOffer(
		ctx,
		hostingOffer,
	)

	return &types.MsgCreateHostingOfferResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHostingOffer(goCtx context.Context, msg *types.MsgUpdateHostingOffer) (*types.MsgUpdateHostingOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingOffer = types.HostingOffer{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Region:        msg.Region,
		PricePerBlock: msg.PricePerBlock,
	}

	// Checks that the element exists
	val, found := k.GetHostingOffer(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHostingOffer(ctx, hostingOffer)

	return &types.MsgUpdateHostingOfferResponse{}, nil
}

func (k msgServer) DeleteHostingOffer(goCtx context.Context, msg *types.MsgDeleteHostingOffer) (*types.MsgDeleteHostingOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHostingOffer(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHostingOffer(ctx, msg.Id)

	return &types.MsgDeleteHostingOfferResponse{}, nil
}
