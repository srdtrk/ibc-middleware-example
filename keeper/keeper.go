package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmosregistry/example"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema  collections.Schema
	Params  collections.Item[example.Params]
	// MiddlewareEnabled is a KeySet of (portID, channelID) that indicates whether this middleware is enabled
	// for a given portID and channelID.
	MiddlewareEnabled collections.KeySet[collections.Pair[string, string]]
	// CallbackCounter is a map of (portID, channelID) to a counter that is incremented every time an expected
	// channel or packet lifecycle callback is called.
	CallbackCounter collections.Map[collections.Pair[string, string], example.CallbackCounter]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, example.ParamsKey, "params", codec.CollValue[example.Params](cdc)),
		MiddlewareEnabled: collections.NewKeySet(
			sb, example.MiddlewareEnabledKey, "mw_enabled", collections.PairKeyCodec(collections.StringKey, collections.StringKey),
		),
		CallbackCounter: collections.NewMap(
			sb, example.CallbackCounterKey, "callback_counter", collections.PairKeyCodec(collections.StringKey, collections.StringKey),
			codec.CollValue[example.CallbackCounter](cdc),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
