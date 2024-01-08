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

func createNHostingOffer(keeper keeper.Keeper, ctx context.Context, n int) []types.HostingOffer {
	items := make([]types.HostingOffer, n)
	for i := range items {
		items[i].Id = keeper.AppendHostingOffer(ctx, items[i])
	}
	return items
}

func TestHostingOfferGet(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingOffer(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHostingOffer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHostingOfferRemove(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingOffer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHostingOffer(ctx, item.Id)
		_, found := keeper.GetHostingOffer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHostingOfferGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingOffer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHostingOffer(ctx)),
	)
}

func TestHostingOfferCount(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingOffer(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHostingOfferCount(ctx))
}
