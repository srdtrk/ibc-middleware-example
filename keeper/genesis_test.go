package keeper_test

import (
	"fmt"
	"testing"

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

				for i := 0; i < 10; i++ {
					counters = append(counters, example.CallbackCounter{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               fmt.Sprintf("channel-%d", i),
					})
				}

				data.Counters = counters
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f = initFixture(t)

			data = &example.GenesisState{
				Counters: []example.CallbackCounter{},
				Params:   example.DefaultParams(),
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
			}
		})
	}
}

// func TestExportGenesis(t *testing.T) {
// 	fixture := initFixture(t)
//
// 	_, err := fixture.msgServer.IncrementCounter(fixture.ctx, &example.MsgIncrementCounter{
// 		Sender: fixture.addrs[0].String(),
// 	})
// 	require.NoError(t, err)
//
// 	out, err := fixture.k.ExportGenesis(fixture.ctx)
// 	require.NoError(t, err)
//
// 	require.Equal(t, example.DefaultParams(), out.Params)
// 	require.Equal(t, uint64(1), out.Counters[0].Count)
// }
