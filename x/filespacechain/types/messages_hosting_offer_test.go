package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateHostingOffer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHostingOffer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHostingOffer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHostingOffer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateHostingOffer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHostingOffer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHostingOffer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHostingOffer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteHostingOffer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHostingOffer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHostingOffer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHostingOffer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
