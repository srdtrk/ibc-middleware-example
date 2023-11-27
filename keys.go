package example

import "cosmossdk.io/collections"

const (
	ModuleName = "example"

	StoreKey = ModuleName

	Version = "example-1"
)

var (
	ParamsKey            = collections.NewPrefix(0)
	CounterKey           = collections.NewPrefix(1)
	MiddlewareEnabledKey = collections.NewPrefix(2)
	CallbackCounterKey   = collections.NewPrefix(3)
)
