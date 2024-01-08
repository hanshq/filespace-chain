package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHostingOffer{}

func NewMsgCreateHostingOffer(creator string, region string, pricePerBlock sdk.Coin) *MsgCreateHostingOffer {
	return &MsgCreateHostingOffer{
		Creator:       creator,
		Region:        region,
		PricePerBlock: pricePerBlock,
	}
}

func (msg *MsgCreateHostingOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHostingOffer{}

func NewMsgUpdateHostingOffer(creator string, id uint64, region string, pricePerBlock sdk.Coin) *MsgUpdateHostingOffer {
	return &MsgUpdateHostingOffer{
		Id:            id,
		Creator:       creator,
		Region:        region,
		PricePerBlock: pricePerBlock,
	}
}

func (msg *MsgUpdateHostingOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHostingOffer{}

func NewMsgDeleteHostingOffer(creator string, id uint64) *MsgDeleteHostingOffer {
	return &MsgDeleteHostingOffer{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteHostingOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
