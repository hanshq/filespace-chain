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

func (k Keeper) HostingInquiryAll(ctx context.Context, req *types.QueryAllHostingInquiryRequest) (*types.QueryAllHostingInquiryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostingInquirys []types.HostingInquiry

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	hostingInquiryStore := prefix.NewStore(store, types.KeyPrefix(types.HostingInquiryKey))

	pageRes, err := query.Paginate(hostingInquiryStore, req.Pagination, func(key []byte, value []byte) error {
		var hostingInquiry types.HostingInquiry
		if err := k.cdc.Unmarshal(value, &hostingInquiry); err != nil {
			return err
		}

		hostingInquirys = append(hostingInquirys, hostingInquiry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHostingInquiryResponse{HostingInquiry: hostingInquirys, Pagination: pageRes}, nil
}

func (k Keeper) HostingInquiry(ctx context.Context, req *types.QueryGetHostingInquiryRequest) (*types.QueryGetHostingInquiryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	hostingInquiry, found := k.GetHostingInquiry(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHostingInquiryResponse{HostingInquiry: hostingInquiry}, nil
}
