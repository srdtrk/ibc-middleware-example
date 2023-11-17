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

	for _, counter := range data.Counters {
		if err := k.CallbackCounter.Set(ctx, collections.Join(counter.ChannelId, "TODO"), counter); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*example.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	iter, err := k.CallbackCounter.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}

	counters, err := iter.Values()
	if err != nil {
		return nil, err
	}

	return &example.GenesisState{
		Params:   params,
		Counters: counters,
	}, nil
}
