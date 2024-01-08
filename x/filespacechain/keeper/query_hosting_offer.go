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

func (k Keeper) HostingOfferAll(ctx context.Context, req *types.QueryAllHostingOfferRequest) (*types.QueryAllHostingOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostingOffers []types.HostingOffer

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	hostingOfferStore := prefix.NewStore(store, types.KeyPrefix(types.HostingOfferKey))

	pageRes, err := query.Paginate(hostingOfferStore, req.Pagination, func(key []byte, value []byte) error {
		var hostingOffer types.HostingOffer
		if err := k.cdc.Unmarshal(value, &hostingOffer); err != nil {
			return err
		}

		hostingOffers = append(hostingOffers, hostingOffer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHostingOfferResponse{HostingOffer: hostingOffers, Pagination: pageRes}, nil
}

func (k Keeper) HostingOffer(ctx context.Context, req *types.QueryGetHostingOfferRequest) (*types.QueryGetHostingOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	hostingOffer, found := k.GetHostingOffer(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHostingOfferResponse{HostingOffer: hostingOffer}, nil
}
