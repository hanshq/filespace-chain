package filespacechain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/hanshq/filespace-chain/testutil/sample"
	filespacechainsimulation "github.com/hanshq/filespace-chain/x/filespacechain/simulation"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// avoid unused import issue
var (
	_ = filespacechainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateFileEntry = "op_weight_msg_file_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFileEntry int = 100

	opWeightMsgUpdateFileEntry = "op_weight_msg_file_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFileEntry int = 100

	opWeightMsgDeleteFileEntry = "op_weight_msg_file_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFileEntry int = 100

	opWeightMsgCreateHostingInquiry = "op_weight_msg_hosting_inquiry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHostingInquiry int = 100

	opWeightMsgUpdateHostingInquiry = "op_weight_msg_hosting_inquiry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHostingInquiry int = 100

	opWeightMsgDeleteHostingInquiry = "op_weight_msg_hosting_inquiry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHostingInquiry int = 100

	opWeightMsgCreateHostingContract = "op_weight_msg_hosting_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHostingContract int = 100

	opWeightMsgUpdateHostingContract = "op_weight_msg_hosting_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHostingContract int = 100

	opWeightMsgDeleteHostingContract = "op_weight_msg_hosting_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHostingContract int = 100

	opWeightMsgCreateHostingOffer = "op_weight_msg_hosting_offer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHostingOffer int = 100

	opWeightMsgUpdateHostingOffer = "op_weight_msg_hosting_offer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHostingOffer int = 100

	opWeightMsgDeleteHostingOffer = "op_weight_msg_hosting_offer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHostingOffer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	filespacechainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		FileEntryList: []types.FileEntry{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		FileEntryCount: 2,
		HostingInquiryList: []types.HostingInquiry{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HostingInquiryCount: 2,
		HostingContractList: []types.HostingContract{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HostingContractCount: 2,
		HostingOfferList: []types.HostingOffer{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HostingOfferCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&filespacechainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateFileEntry int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFileEntry, &weightMsgCreateFileEntry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFileEntry = defaultWeightMsgCreateFileEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFileEntry,
		filespacechainsimulation.SimulateMsgCreateFileEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFileEntry int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFileEntry, &weightMsgUpdateFileEntry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFileEntry = defaultWeightMsgUpdateFileEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFileEntry,
		filespacechainsimulation.SimulateMsgUpdateFileEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFileEntry int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFileEntry, &weightMsgDeleteFileEntry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFileEntry = defaultWeightMsgDeleteFileEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFileEntry,
		filespacechainsimulation.SimulateMsgDeleteFileEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHostingInquiry int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateHostingInquiry, &weightMsgCreateHostingInquiry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHostingInquiry = defaultWeightMsgCreateHostingInquiry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHostingInquiry,
		filespacechainsimulation.SimulateMsgCreateHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHostingInquiry int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateHostingInquiry, &weightMsgUpdateHostingInquiry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHostingInquiry = defaultWeightMsgUpdateHostingInquiry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHostingInquiry,
		filespacechainsimulation.SimulateMsgUpdateHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHostingInquiry int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteHostingInquiry, &weightMsgDeleteHostingInquiry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHostingInquiry = defaultWeightMsgDeleteHostingInquiry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHostingInquiry,
		filespacechainsimulation.SimulateMsgDeleteHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHostingContract int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateHostingContract, &weightMsgCreateHostingContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHostingContract = defaultWeightMsgCreateHostingContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHostingContract,
		filespacechainsimulation.SimulateMsgCreateHostingContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHostingContract int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateHostingContract, &weightMsgUpdateHostingContract, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHostingContract = defaultWeightMsgUpdateHostingContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHostingContract,
		filespacechainsimulation.SimulateMsgUpdateHostingContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHostingContract int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteHostingContract, &weightMsgDeleteHostingContract, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHostingContract = defaultWeightMsgDeleteHostingContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHostingContract,
		filespacechainsimulation.SimulateMsgDeleteHostingContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHostingOffer int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateHostingOffer, &weightMsgCreateHostingOffer, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHostingOffer = defaultWeightMsgCreateHostingOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHostingOffer,
		filespacechainsimulation.SimulateMsgCreateHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHostingOffer int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateHostingOffer, &weightMsgUpdateHostingOffer, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHostingOffer = defaultWeightMsgUpdateHostingOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHostingOffer,
		filespacechainsimulation.SimulateMsgUpdateHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHostingOffer int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteHostingOffer, &weightMsgDeleteHostingOffer, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHostingOffer = defaultWeightMsgDeleteHostingOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHostingOffer,
		filespacechainsimulation.SimulateMsgDeleteHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFileEntry,
			defaultWeightMsgCreateFileEntry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgCreateFileEntry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFileEntry,
			defaultWeightMsgUpdateFileEntry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgUpdateFileEntry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFileEntry,
			defaultWeightMsgDeleteFileEntry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgDeleteFileEntry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHostingInquiry,
			defaultWeightMsgCreateHostingInquiry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgCreateHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHostingInquiry,
			defaultWeightMsgUpdateHostingInquiry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgUpdateHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHostingInquiry,
			defaultWeightMsgDeleteHostingInquiry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgDeleteHostingInquiry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHostingContract,
			defaultWeightMsgCreateHostingContract,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgCreateHostingContract(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHostingContract,
			defaultWeightMsgUpdateHostingContract,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgUpdateHostingContract(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHostingContract,
			defaultWeightMsgDeleteHostingContract,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgDeleteHostingContract(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHostingOffer,
			defaultWeightMsgCreateHostingOffer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgCreateHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHostingOffer,
			defaultWeightMsgUpdateHostingOffer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgUpdateHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHostingOffer,
			defaultWeightMsgDeleteHostingOffer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filespacechainsimulation.SimulateMsgDeleteHostingOffer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
