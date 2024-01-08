package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFileEntry{}

func NewMsgCreateFileEntry(creator string, cid string, rootCid string, parentCid string, metaData string, fileSize uint64) *MsgCreateFileEntry {
	return &MsgCreateFileEntry{
		Creator:   creator,
		Cid:       cid,
		RootCid:   rootCid,
		ParentCid: parentCid,
		MetaData:  metaData,
		FileSize:  fileSize,
	}
}

func (msg *MsgCreateFileEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFileEntry{}

func NewMsgUpdateFileEntry(creator string, id uint64, cid string, rootCid string, parentCid string, metaData string, fileSize uint64) *MsgUpdateFileEntry {
	return &MsgUpdateFileEntry{
		Id:        id,
		Creator:   creator,
		Cid:       cid,
		RootCid:   rootCid,
		ParentCid: parentCid,
		MetaData:  metaData,
		FileSize:  fileSize,
	}
}

func (msg *MsgUpdateFileEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFileEntry{}

func NewMsgDeleteFileEntry(creator string, id uint64) *MsgDeleteFileEntry {
	return &MsgDeleteFileEntry{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteFileEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
