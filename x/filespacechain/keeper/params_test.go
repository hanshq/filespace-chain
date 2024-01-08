package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FilespacechainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
