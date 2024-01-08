package keeper

import (
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

var _ types.QueryServer = Keeper{}
