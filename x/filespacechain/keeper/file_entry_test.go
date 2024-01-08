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

func createNFileEntry(keeper keeper.Keeper, ctx context.Context, n int) []types.FileEntry {
	items := make([]types.FileEntry, n)
	for i := range items {
		items[i].Id = keeper.AppendFileEntry(ctx, items[i])
	}
	return items
}

func TestFileEntryGet(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNFileEntry(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFileEntry(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFileEntryRemove(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNFileEntry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFileEntry(ctx, item.Id)
		_, found := keeper.GetFileEntry(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFileEntryGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNFileEntry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFileEntry(ctx)),
	)
}

func TestFileEntryCount(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNFileEntry(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFileEntryCount(ctx))
}
