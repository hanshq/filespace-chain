package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FileEntryAll(ctx context.Context, req *types.QueryAllFileEntryRequest) (*types.QueryAllFileEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fileEntrys []types.FileEntry

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	fileEntryStore := prefix.NewStore(store, types.KeyPrefix(types.FileEntryKey))

	pageRes, err := query.Paginate(fileEntryStore, req.Pagination, func(key []byte, value []byte) error {
		var fileEntry types.FileEntry
		if err := k.cdc.Unmarshal(value, &fileEntry); err != nil {
			return err
		}

		fileEntrys = append(fileEntrys, fileEntry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFileEntryResponse{FileEntry: fileEntrys, Pagination: pageRes}, nil
}

func (k Keeper) FileEntry(ctx context.Context, req *types.QueryGetFileEntryRequest) (*types.QueryGetFileEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	fileEntry, found := k.GetFileEntry(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFileEntryResponse{FileEntry: fileEntry}, nil
}
