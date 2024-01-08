package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHostingInquiry{}

func NewMsgCreateHostingInquiry(creator string, fileEntryCid string, replicationRate uint64, escrowAmount sdk.Coin, endTime uint64) *MsgCreateHostingInquiry {
	return &MsgCreateHostingInquiry{
		Creator:         creator,
		FileEntryCid:    fileEntryCid,
		ReplicationRate: replicationRate,
		EscrowAmount:    escrowAmount,
		EndTime:         endTime,
	}
}

func (msg *MsgCreateHostingInquiry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHostingInquiry{}

func NewMsgUpdateHostingInquiry(creator string, id uint64, fileEntryCid string, replicationRate uint64, escrowAmount sdk.Coin, endTime uint64) *MsgUpdateHostingInquiry {
	return &MsgUpdateHostingInquiry{
		Id:              id,
		Creator:         creator,
		FileEntryCid:    fileEntryCid,
		ReplicationRate: replicationRate,
		EscrowAmount:    escrowAmount,
		EndTime:         endTime,
	}
}

func (msg *MsgUpdateHostingInquiry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHostingInquiry{}

func NewMsgDeleteHostingInquiry(creator string, id uint64) *MsgDeleteHostingInquiry {
	return &MsgDeleteHostingInquiry{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteHostingInquiry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
