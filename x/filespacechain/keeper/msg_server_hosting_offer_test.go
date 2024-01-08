package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func TestHostingOfferMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHostingOffer(wctx, &types.MsgCreateHostingOffer{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHostingOfferMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateHostingOffer
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHostingOffer{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingOffer{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHostingOffer{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingOffer(wctx, &types.MsgCreateHostingOffer{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateHostingOffer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHostingOfferMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteHostingOffer
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHostingOffer{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHostingOffer{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHostingOffer{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateHostingOffer(wctx, &types.MsgCreateHostingOffer{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteHostingOffer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
