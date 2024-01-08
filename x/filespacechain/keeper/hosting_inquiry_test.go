package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/testutil/nullify"
	"github.com/hanshq/filespace-chain/x/filespacechain/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	"github.com/stretchr/testify/require"
)

func createNHostingInquiry(keeper keeper.Keeper, ctx context.Context, n int) []types.HostingInquiry {
	items := make([]types.HostingInquiry, n)
	for i := range items {
		items[i].Id = keeper.AppendHostingInquiry(ctx, items[i])
	}
	return items
}

func TestHostingInquiryGet(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingInquiry(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHostingInquiry(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHostingInquiryRemove(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingInquiry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHostingInquiry(ctx, item.Id)
		_, found := keeper.GetHostingInquiry(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHostingInquiryGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingInquiry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHostingInquiry(ctx)),
	)
}

func TestHostingInquiryCount(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingInquiry(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHostingInquiryCount(ctx))
}
