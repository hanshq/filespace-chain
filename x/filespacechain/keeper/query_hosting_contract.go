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

func (k Keeper) HostingContractAll(ctx context.Context, req *types.QueryAllHostingContractRequest) (*types.QueryAllHostingContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostingContracts []types.HostingContract

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	hostingContractStore := prefix.NewStore(store, types.KeyPrefix(types.HostingContractKey))

	pageRes, err := query.Paginate(hostingContractStore, req.Pagination, func(key []byte, value []byte) error {
		var hostingContract types.HostingContract
		if err := k.cdc.Unmarshal(value, &hostingContract); err != nil {
			return err
		}

		hostingContracts = append(hostingContracts, hostingContract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHostingContractResponse{HostingContract: hostingContracts, Pagination: pageRes}, nil
}

func (k Keeper) HostingContract(ctx context.Context, req *types.QueryGetHostingContractRequest) (*types.QueryGetHostingContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	hostingContract, found := k.GetHostingContract(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHostingContractResponse{HostingContract: hostingContract}, nil
}
