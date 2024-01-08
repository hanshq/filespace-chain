package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/testutil/nullify"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestHostingOfferQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	msgs := createNHostingOffer(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetHostingOfferRequest
		response *types.QueryGetHostingOfferResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetHostingOfferRequest{Id: msgs[0].Id},
			response: &types.QueryGetHostingOfferResponse{HostingOffer: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetHostingOfferRequest{Id: msgs[1].Id},
			response: &types.QueryGetHostingOfferResponse{HostingOffer: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetHostingOfferRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.HostingOffer(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestHostingOfferQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	msgs := createNHostingOffer(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHostingOfferRequest {
		return &types.QueryAllHostingOfferRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HostingOfferAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HostingOffer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HostingOffer),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HostingOfferAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HostingOffer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HostingOffer),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HostingOfferAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HostingOffer),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HostingOfferAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
