package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func (k msgServer) CreateFileEntry(goCtx context.Context, msg *types.MsgCreateFileEntry) (*types.MsgCreateFileEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var fileEntry = types.FileEntry{
		Creator:   msg.Creator,
		Cid:       msg.Cid,
		RootCid:   msg.RootCid,
		ParentCid: msg.ParentCid,
		MetaData:  msg.MetaData,
		FileSize:  msg.FileSize,
	}

	id := k.AppendFileEntry(
		ctx,
		fileEntry,
	)

	return &types.MsgCreateFileEntryResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateFileEntry(goCtx context.Context, msg *types.MsgUpdateFileEntry) (*types.MsgUpdateFileEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var fileEntry = types.FileEntry{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Cid:       msg.Cid,
		RootCid:   msg.RootCid,
		ParentCid: msg.ParentCid,
		MetaData:  msg.MetaData,
		FileSize:  msg.FileSize,
	}

	// Checks that the element exists
	val, found := k.GetFileEntry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetFileEntry(ctx, fileEntry)

	return &types.MsgUpdateFileEntryResponse{}, nil
}

func (k msgServer) DeleteFileEntry(goCtx context.Context, msg *types.MsgDeleteFileEntry) (*types.MsgDeleteFileEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetFileEntry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFileEntry(ctx, msg.Id)

	return &types.MsgDeleteFileEntryResponse{}, nil
}
