package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestHostingInquiryMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHostingInquiry(wctx, &types.MsgCreateHostingInquiry{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHostingInquiryMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateHostingInquiry
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHostingInquiry{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingInquiry{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingInquiry{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingInquiry(wctx, &types.MsgCreateHostingInquiry{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateHostingInquiry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHostingInquiryMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteHostingInquiry
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHostingInquiry{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHostingInquiry{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHostingInquiry{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingInquiry(wctx, &types.MsgCreateHostingInquiry{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteHostingInquiry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
