package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFileEntry{},
		&MsgUpdateFileEntry{},
		&MsgDeleteFileEntry{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHostingInquiry{},
		&MsgUpdateHostingInquiry{},
		&MsgDeleteHostingInquiry{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHostingContract{},
		&MsgUpdateHostingContract{},
		&MsgDeleteHostingContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHostingOffer{},
		&MsgUpdateHostingOffer{},
		&MsgDeleteHostingOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakeForHosting{},
		&MsgUnstakeFromHosting{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
