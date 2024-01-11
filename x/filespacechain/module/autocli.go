package filespacechain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/hanshq/filespace-chain/api/filespacechain/filespacechain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "FileEntryAll",
					Use:       "list-file-entry",
					Short:     "List all FileEntry",
				},
				{
					RpcMethod:      "FileEntry",
					Use:            "show-file-entry [id]",
					Short:          "Shows a FileEntry by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "HostingInquiryAll",
					Use:       "list-hosting-inquiry",
					Short:     "List all HostingInquiry",
				},
				{
					RpcMethod:      "HostingInquiry",
					Use:            "show-hosting-inquiry [id]",
					Short:          "Shows a HostingInquiry by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "HostingContractAll",
					Use:       "list-hosting-contract",
					Short:     "List all HostingContract",
				},
				{
					RpcMethod:      "HostingContract",
					Use:            "show-hosting-contract [id]",
					Short:          "Shows a HostingContract by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "HostingOfferAll",
					Use:       "list-hosting-offer",
					Short:     "List all HostingOffer",
				},
				{
					RpcMethod:      "HostingOffer",
					Use:            "show-hosting-offer [id]",
					Short:          "Shows a HostingOffer by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ListHostingContractFrom",
					Use:            "list-hosting-contract-from [creator]",
					Short:          "Query list-hosting-contract-from",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "creator"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateFileEntry",
					Use:            "create-file-entry [cid] [rootCid] [parentCid] [metaData] [fileSize]",
					Short:          "Create FileEntry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cid"}, {ProtoField: "rootCid"}, {ProtoField: "parentCid"}, {ProtoField: "metaData"}, {ProtoField: "fileSize"}},
				},
				{
					RpcMethod:      "UpdateFileEntry",
					Use:            "update-file-entry [id] [cid] [rootCid] [parentCid] [metaData] [fileSize]",
					Short:          "Update FileEntry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "cid"}, {ProtoField: "rootCid"}, {ProtoField: "parentCid"}, {ProtoField: "metaData"}, {ProtoField: "fileSize"}},
				},
				{
					RpcMethod:      "DeleteFileEntry",
					Use:            "delete-file-entry [id]",
					Short:          "Delete FileEntry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateHostingInquiry",
					Use:            "create-hosting-inquiry [fileEntryCid] [replicationRate] [escrowAmount] [endTime]",
					Short:          "Create HostingInquiry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "fileEntryCid"}, {ProtoField: "replicationRate"}, {ProtoField: "escrowAmount"}, {ProtoField: "endTime"}},
				},
				{
					RpcMethod:      "UpdateHostingInquiry",
					Use:            "update-hosting-inquiry [id] [fileEntryCid] [replicationRate] [escrowAmount] [endTime]",
					Short:          "Update HostingInquiry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "fileEntryCid"}, {ProtoField: "replicationRate"}, {ProtoField: "escrowAmount"}, {ProtoField: "endTime"}},
				},
				{
					RpcMethod:      "DeleteHostingInquiry",
					Use:            "delete-hosting-inquiry [id]",
					Short:          "Delete HostingInquiry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateHostingContract",
					Use:            "create-hosting-contract [inquiryId] [offerId]",
					Short:          "Create HostingContract",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "inquiryId"}, {ProtoField: "offerId"}},
				},
				{
					RpcMethod:      "UpdateHostingContract",
					Use:            "update-hosting-contract [id] [inquiryId] [offerId]",
					Short:          "Update HostingContract",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "inquiryId"}, {ProtoField: "offerId"}},
				},
				{
					RpcMethod:      "DeleteHostingContract",
					Use:            "delete-hosting-contract [id]",
					Short:          "Delete HostingContract",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateHostingOffer",
					Use:            "create-hosting-offer [region] [pricePerBlock]",
					Short:          "Create HostingOffer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "region"}, {ProtoField: "pricePerBlock"}},
				},
				{
					RpcMethod:      "UpdateHostingOffer",
					Use:            "update-hosting-offer [id] [region] [pricePerBlock]",
					Short:          "Update HostingOffer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "region"}, {ProtoField: "pricePerBlock"}},
				},
				{
					RpcMethod:      "DeleteHostingOffer",
					Use:            "delete-hosting-offer [id]",
					Short:          "Delete HostingOffer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
