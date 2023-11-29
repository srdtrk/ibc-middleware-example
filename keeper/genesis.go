package keeper

import (
	"context"

	"cosmossdk.io/collections"

	"github.com/cosmosregistry/example"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *example.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, counter := range data.CallbackCounters {
		if err := k.CallbackCounter.Set(ctx, counter.ChannelId, counter); err != nil {
			return err
		}
	}

	for _, channel := range data.MiddlewareEnabledChannels {
		if err := k.MiddlewareEnabled.Set(ctx, collections.Join(channel.PortId, channel.ChannelId)); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*example.GenesisState, error) {
	genesisState := &example.GenesisState{}

	var err error
	genesisState.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	iter, err := k.CallbackCounter.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}

	genesisState.CallbackCounters, err = iter.Values()
	if err != nil {
		return nil, err
	}

	if err := k.MiddlewareEnabled.Walk(ctx, nil, func(key collections.Pair[string, string]) (bool, error) {
		genesisState.MiddlewareEnabledChannels = append(genesisState.MiddlewareEnabledChannels, example.MiddlewareEnabledChannel{
			PortId:    key.K1(),
			ChannelId: key.K2(),
		})

		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesisState, nil
}
