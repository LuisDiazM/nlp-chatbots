package app

import (
	"github.com/google/wire"
)

var appProvider = wire.NewSet(NewApplication)
