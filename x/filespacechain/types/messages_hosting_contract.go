package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHostingContract{}

func NewMsgCreateHostingContract(creator string, inquiryId uint64, offerId uint64) *MsgCreateHostingContract {
	return &MsgCreateHostingContract{
		Creator:   creator,
		InquiryId: inquiryId,
		OfferId:   offerId,
	}
}

func (msg *MsgCreateHostingContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHostingContract{}

func NewMsgUpdateHostingContract(creator string, id uint64, inquiryId uint64, offerId uint64) *MsgUpdateHostingContract {
	return &MsgUpdateHostingContract{
		Id:        id,
		Creator:   creator,
		InquiryId: inquiryId,
		OfferId:   offerId,
	}
}

func (msg *MsgUpdateHostingContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHostingContract{}

func NewMsgDeleteHostingContract(creator string, id uint64) *MsgDeleteHostingContract {
	return &MsgDeleteHostingContract{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteHostingContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
