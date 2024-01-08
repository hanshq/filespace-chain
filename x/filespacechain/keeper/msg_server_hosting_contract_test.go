package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestHostingContractMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHostingContract(wctx, &types.MsgCreateHostingContract{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHostingContractMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateHostingContract
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHostingContract{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingContract{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingContract{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingContract(wctx, &types.MsgCreateHostingContract{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateHostingContract(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHostingContractMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteHostingContract
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHostingContract{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHostingContract{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHostingContract{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingContract(wctx, &types.MsgCreateHostingContract{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteHostingContract(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
