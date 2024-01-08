package filespacechain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hanshq/filespace-chain/x/filespacechain/keeper"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the fileEntry
	for _, elem := range genState.FileEntryList {
		k.SetFileEntry(ctx, elem)
	}

	// Set fileEntry count
	k.SetFileEntryCount(ctx, genState.FileEntryCount)
	// Set all the hostingInquiry
	for _, elem := range genState.HostingInquiryList {
		k.SetHostingInquiry(ctx, elem)
	}

	// Set hostingInquiry count
	k.SetHostingInquiryCount(ctx, genState.HostingInquiryCount)
	// Set all the hostingContract
	for _, elem := range genState.HostingContractList {
		k.SetHostingContract(ctx, elem)
	}

	// Set hostingContract count
	k.SetHostingContractCount(ctx, genState.HostingContractCount)
	// Set all the hostingOffer
	for _, elem := range genState.HostingOfferList {
		k.SetHostingOffer(ctx, elem)
	}

	// Set hostingOffer count
	k.SetHostingOfferCount(ctx, genState.HostingOfferCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FileEntryList = k.GetAllFileEntry(ctx)
	genesis.FileEntryCount = k.GetFileEntryCount(ctx)
	genesis.HostingInquiryList = k.GetAllHostingInquiry(ctx)
	genesis.HostingInquiryCount = k.GetHostingInquiryCount(ctx)
	genesis.HostingContractList = k.GetAllHostingContract(ctx)
	genesis.HostingContractCount = k.GetHostingContractCount(ctx)
	genesis.HostingOfferList = k.GetAllHostingOffer(ctx)
	genesis.HostingOfferCount = k.GetHostingOfferCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
