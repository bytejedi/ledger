package keeper

import (
	"github.com/bytejedi/ledger/x/ledger/types"
)

var _ types.QueryServer = Keeper{}
