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

func TestFileEntryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	msgs := createNFileEntry(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetFileEntryRequest
		response *types.QueryGetFileEntryResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetFileEntryRequest{Id: msgs[0].Id},
			response: &types.QueryGetFileEntryResponse{FileEntry: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetFileEntryRequest{Id: msgs[1].Id},
			response: &types.QueryGetFileEntryResponse{FileEntry: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetFileEntryRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.FileEntry(ctx, tc.request)
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

func TestFileEntryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	msgs := createNFileEntry(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFileEntryRequest {
		return &types.QueryAllFileEntryRequest{
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
			resp, err := keeper.FileEntryAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FileEntry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FileEntry),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FileEntryAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FileEntry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FileEntry),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FileEntryAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.FileEntry),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FileEntryAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
