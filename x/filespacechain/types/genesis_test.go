package types_test

import (
	"testing"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated fileEntry",
			genState: &types.GenesisState{
				FileEntryList: []types.FileEntry{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid fileEntry count",
			genState: &types.GenesisState{
				FileEntryList: []types.FileEntry{
					{
						Id: 1,
					},
				},
				FileEntryCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated hostingInquiry",
			genState: &types.GenesisState{
				HostingInquiryList: []types.HostingInquiry{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid hostingInquiry count",
			genState: &types.GenesisState{
				HostingInquiryList: []types.HostingInquiry{
					{
						Id: 1,
					},
				},
				HostingInquiryCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated hostingContract",
			genState: &types.GenesisState{
				HostingContractList: []types.HostingContract{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid hostingContract count",
			genState: &types.GenesisState{
				HostingContractList: []types.HostingContract{
					{
						Id: 1,
					},
				},
				HostingContractCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated hostingOffer",
			genState: &types.GenesisState{
				HostingOfferList: []types.HostingOffer{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid hostingOffer count",
			genState: &types.GenesisState{
				HostingOfferList: []types.HostingOffer{
					{
						Id: 1,
					},
				},
				HostingOfferCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
