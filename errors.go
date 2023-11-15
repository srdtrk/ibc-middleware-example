package example

import errorsmod "cosmossdk.io/errors"

var (
	// ErrDuplicateAddress error if there is a duplicate address
	ErrDuplicateAddress = errorsmod.Register(ModuleName, 2, "duplicate address")
	// ErrInvalidVersion error if the channel version is invalid
	ErrInvalidChannelVersion = errorsmod.Register(ModuleName, 3, "invalid example middleware version")
)
