package filespacechain_test

import (
	"testing"

	keepertest "github.com/hanshq/filespace-chain/testutil/keeper"
	"github.com/hanshq/filespace-chain/testutil/nullify"
	"github.com/hanshq/filespace-chain/x/filespacechain/module"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FilespacechainKeeper(t)
	filespacechain.InitGenesis(ctx, k, genesisState)
	got := filespacechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
