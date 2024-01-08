package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateHostingInquiry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHostingInquiry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHostingInquiry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHostingInquiry{
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

func TestMsgUpdateHostingInquiry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHostingInquiry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHostingInquiry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHostingInquiry{
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

func TestMsgDeleteHostingInquiry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHostingInquiry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHostingInquiry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHostingInquiry{
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
