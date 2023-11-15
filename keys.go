package example

import "cosmossdk.io/collections"

const (
	ModuleName = "example"

	Version = "example-1"
)

var (
	ParamsKey            = collections.NewPrefix(0)
	CounterKey           = collections.NewPrefix(1)
	MiddlewareEnabledKey = collections.NewPrefix(2)
)
