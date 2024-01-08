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

		FileEntryList: []types.FileEntry{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FileEntryCount: 2,
		HostingInquiryList: []types.HostingInquiry{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HostingInquiryCount: 2,
		HostingContractList: []types.HostingContract{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HostingContractCount: 2,
		HostingOfferList: []types.HostingOffer{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HostingOfferCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FilespacechainKeeper(t)
	filespacechain.InitGenesis(ctx, k, genesisState)
	got := filespacechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FileEntryList, got.FileEntryList)
	require.Equal(t, genesisState.FileEntryCount, got.FileEntryCount)
	require.ElementsMatch(t, genesisState.HostingInquiryList, got.HostingInquiryList)
	require.Equal(t, genesisState.HostingInquiryCount, got.HostingInquiryCount)
	require.ElementsMatch(t, genesisState.HostingContractList, got.HostingContractList)
	require.Equal(t, genesisState.HostingContractCount, got.HostingContractCount)
	require.ElementsMatch(t, genesisState.HostingOfferList, got.HostingOfferList)
	require.Equal(t, genesisState.HostingOfferCount, got.HostingOfferCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
