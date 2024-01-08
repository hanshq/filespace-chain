package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFileEntry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFileEntry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFileEntry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateFileEntry{
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

func TestMsgUpdateFileEntry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFileEntry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateFileEntry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateFileEntry{
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

func TestMsgDeleteFileEntry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteFileEntry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteFileEntry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteFileEntry{
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
