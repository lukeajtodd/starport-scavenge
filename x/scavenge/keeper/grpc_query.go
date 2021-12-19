package keeper

import (
	"github.com/lukeajtodd/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
