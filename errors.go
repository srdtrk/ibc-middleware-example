package example

import errorsmod "cosmossdk.io/errors"

var (
	// ErrDuplicateChannelID error if there is a duplicate channel id
	ErrDuplicateChannelID = errorsmod.Register(ModuleName, 2, "duplicate channel id")
	// ErrInvalidVersion error if the channel version is invalid
	ErrInvalidChannelVersion = errorsmod.Register(ModuleName, 3, "invalid example middleware version")
)
