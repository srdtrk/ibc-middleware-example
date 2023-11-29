package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmosregistry/example"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	var (
		f *testFixture

		data *example.GenesisState
	)

	testCases := []struct {
		name     string
		malleate func()
		expErr   bool
	}{
		{
			"success: no genesis state",
			func() {},
			false,
		},
		{
			"success: some genesis state",
			func() {
				counters := []example.CallbackCounter{}
				channels := []example.MiddlewareEnabledChannel{}

				for i := 0; i < 10; i++ {
					counters = append(counters, example.CallbackCounter{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               fmt.Sprintf("channel-%d", i),
					})

					channels = append(channels, example.MiddlewareEnabledChannel{
						PortId:    fmt.Sprintf("port-%d", i),
						ChannelId: fmt.Sprintf("channel-%d", i),
					})
				}

				data.Counters = counters
				data.MiddlewareEnabledChannels = channels
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f = initFixture(t)

			data = &example.GenesisState{
				Params: example.DefaultParams(),
			}

			tc.malleate()

			err := f.k.InitGenesis(f.ctx, data)

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				params, err := f.k.Params.Get(f.ctx)
				require.NoError(t, err)

				require.Equal(t, data.Params, params)

				iter, err := f.k.CallbackCounter.Iterate(f.ctx, nil)
				require.NoError(t, err)

				counters, err := iter.Values()
				require.NoError(t, err)

				if len(data.Counters) == 0 {
					require.Len(t, counters, 0)
				} else {
					require.Equal(t, data.Counters, counters)
				}

				keyIter, err := f.k.MiddlewareEnabled.Iterate(f.ctx, nil)
				require.NoError(t, err)

				keys, err := keyIter.Keys()
				require.NoError(t, err)

				if len(data.MiddlewareEnabledChannels) == 0 {
					require.Len(t, keys, 0)
				} else {
					expectedKeys := []collections.Pair[string, string]{}
					for _, channel := range data.MiddlewareEnabledChannels {
						expectedKeys = append(expectedKeys, collections.Join(channel.PortId, channel.ChannelId))
					}

					require.Equal(t, expectedKeys, keys)
				}
			}
		})
	}
}

func TestExportGenesis(t *testing.T) {
	var (
		f *testFixture

		expGenesisState *example.GenesisState
	)

	testCases := []struct {
		name     string
		malleate func()
		expErr   bool
	}{
		{
			"success: no genesis state",
			func() {
				expGenesisState = &example.GenesisState{
					Params: example.DefaultParams(),
				}
			},
			false,
		},
		{
			"success: some genesis state",
			func() {
				expGenesisState = &example.GenesisState{
					Params:                    example.DefaultParams(),
					Counters:                  []example.CallbackCounter{},
					MiddlewareEnabledChannels: []example.MiddlewareEnabledChannel{},
				}

				for i := 0; i < 10; i++ {
					cc := example.CallbackCounter{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               fmt.Sprintf("channel-%d", i),
					}
					expGenesisState.Counters = append(expGenesisState.Counters, cc)
					err := f.k.CallbackCounter.Set(f.ctx, cc.ChannelId, cc)
					require.NoError(t, err)

					expGenesisState.MiddlewareEnabledChannels = append(expGenesisState.MiddlewareEnabledChannels, example.MiddlewareEnabledChannel{
						PortId:    fmt.Sprintf("port-%d", i),
						ChannelId: fmt.Sprintf("channel-%d", i),
					})
					err = f.k.MiddlewareEnabled.Set(f.ctx, collections.Join(fmt.Sprintf("port-%d", i), fmt.Sprintf("channel-%d", i)))
					require.NoError(t, err)
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f = initFixture(t)

			tc.malleate()

			out, err := f.k.ExportGenesis(f.ctx)

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, expGenesisState, out)
			}
		})
	}
}
