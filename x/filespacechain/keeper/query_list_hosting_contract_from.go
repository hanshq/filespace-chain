package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListHostingContractFrom(goCtx context.Context, req *types.QueryListHostingContractFromRequest) (*types.QueryListHostingContractFromResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostingContracts []types.HostingContract

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	hostingContractStore := prefix.NewStore(store, types.KeyPrefix(types.HostingContractKey))
	// Iterate with pagination
	pageRes, err := query.FilteredPaginate(hostingContractStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var hostingContract types.HostingContract
		if err := k.cdc.Unmarshal(value, &hostingContract); err != nil {
			return false, err
		}

		if hostingContract.Creator == req.Creator {
			if accumulate {
				hostingContracts = append(hostingContracts, hostingContract)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListHostingContractFromResponse{HostingContract: hostingContracts, Pagination: pageRes}, nil

}
