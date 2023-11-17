package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmosregistry/example"

	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Params(f.ctx, &example.QueryParamsRequest{})
	require.NoError(err)
	require.Equal(example.Params{}, resp.Params)
}

func TestQueryCallbackCounter(t *testing.T) {
	var (
		f *testFixture

		expCallbackCounter *example.CallbackCounter
	)

	testCases := []struct {
		name     string
		malleate func()
		expErr   bool
	}{
		{
			"success",
			func() {
				expCallbackCounter = &example.CallbackCounter{
					OnRecvPacket:            1,
					OnAcknowledgementPacket: 1,
					OnTimeoutPacket:         1,
					SendPacket:              1,
					ChannelId:               defaultTestChanID,
				}

				err := f.k.CallbackCounter.Set(f.ctx, collections.Join(defaultTestPortID, defaultTestChanID), *expCallbackCounter)
				require.NoError(t, err)
			},
			false,
		},
		{
			"failure: counter not found",
			func() {},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f = initFixture(t)

			expCallbackCounter = nil
			tc.malleate()

			resp, err := f.queryServer.CallbackCounter(f.ctx, &example.QueryCallbackCounterRequest{PortId: defaultTestPortID, ChannelId: defaultTestChanID})

			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.Equal(t, expCallbackCounter, resp.Counter)
			}
		})
	}
}

func TestQueryCallbackCounters(t *testing.T) {
	var (
		f *testFixture

		pagination *query.PageRequest

		expCallbackCounters []*example.CallbackCounter
	)

	testCases := []struct {
		name     string
		malleate func()
		expErr   bool
	}{
		{
			"success: empty counters",
			func() {
				expCallbackCounters = ([]*example.CallbackCounter)(nil)
			},
			false,
		},
		{
			"success: one counter",
			func() {
				expCallbackCounters = []*example.CallbackCounter{
					{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               defaultTestChanID,
					},
				}

				err := f.k.CallbackCounter.Set(f.ctx, collections.Join(defaultTestPortID, defaultTestChanID), *expCallbackCounters[0])
				require.NoError(t, err)
			},
			false,
		},
		{
			"success: many counters",
			func() {
				expCallbackCounters = []*example.CallbackCounter{}

				for i := 0; i < 10; i++ {
					counter := example.CallbackCounter{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               fmt.Sprintf("channel-%d", i),
					}

					expCallbackCounters = append(expCallbackCounters, &counter)

					err := f.k.CallbackCounter.Set(f.ctx, collections.Join(defaultTestPortID, counter.ChannelId), counter)
					require.NoError(t, err)
				}
			},
			false,
		},
		{
			"success: with pagination",
			func() {
				pagination = &query.PageRequest{Limit: 5}

				expCallbackCounters = []*example.CallbackCounter{}

				for i := 0; i < 10; i++ {
					counter := example.CallbackCounter{
						OnRecvPacket:            1,
						OnAcknowledgementPacket: 1,
						OnTimeoutPacket:         1,
						SendPacket:              1,
						ChannelId:               fmt.Sprintf("channel-%d", i),
					}

					if i < 5 {
						expCallbackCounters = append(expCallbackCounters, &counter)
					}

					err := f.k.CallbackCounter.Set(f.ctx, collections.Join(defaultTestPortID, counter.ChannelId), counter)
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

			expCallbackCounters = nil
			tc.malleate()

			resp, err := f.queryServer.CallbackCounters(f.ctx, &example.QueryCallbackCountersRequest{Pagination: pagination})

			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.Equal(t, expCallbackCounters, resp.Counters)
			}
		})
	}
}

// func TestQueryCounters(t *testing.T) {
// 	f := initFixture(t)
// 	require := require.New(t)
//
// 	resp, err := f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
// 	require.NoError(err)
// 	require.Equal(0, len(resp.Counters))
//
// 	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[0].String()})
// 	require.NoError(err)
//
// 	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
// 	require.NoError(err)
// 	require.Equal(1, len(resp.Counters))
// 	require.Equal(uint64(1), resp.Counters[0].Count)
// 	require.Equal(f.addrs[0].String(), resp.Counters[0].Address)
// }
//
// func TestQueryCountersPaginated(t *testing.T) {
// 	f := initFixture(t)
// 	require := require.New(t)
//
// 	resp, err := f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{Pagination: &query.PageRequest{Limit: 1}})
// 	require.NoError(err)
// 	require.Equal(0, len(resp.Counters))
//
// 	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[0].String()})
// 	require.NoError(err)
// 	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[1].String()})
// 	require.NoError(err)
//
// 	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{Pagination: &query.PageRequest{Limit: 1}})
// 	require.NoError(err)
// 	require.Equal(1, len(resp.Counters))
// 	require.Equal(uint64(1), resp.Counters[0].Count)
// 	require.Equal(f.addrs[1].String(), resp.Counters[0].Address)
//
// 	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
// 	require.NoError(err)
// 	require.Equal(2, len(resp.Counters))
// }
