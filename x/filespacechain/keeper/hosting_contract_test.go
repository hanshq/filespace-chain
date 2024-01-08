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

func createNHostingContract(keeper keeper.Keeper, ctx context.Context, n int) []types.HostingContract {
	items := make([]types.HostingContract, n)
	for i := range items {
		items[i].Id = keeper.AppendHostingContract(ctx, items[i])
	}
	return items
}

func TestHostingContractGet(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingContract(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHostingContract(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHostingContractRemove(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHostingContract(ctx, item.Id)
		_, found := keeper.GetHostingContract(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHostingContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHostingContract(ctx)),
	)
}

func TestHostingContractCount(t *testing.T) {
	keeper, ctx := keepertest.FilespacechainKeeper(t)
	items := createNHostingContract(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHostingContractCount(ctx))
}
