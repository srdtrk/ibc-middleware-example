package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmosregistry/example"
)

var _ example.QueryServer = (*queryServer)(nil)

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) example.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// MiddlewareEnabledChannel defines the handler for the Query/MiddlewareEnabledChannel RPC method.
func (qs queryServer) MiddlewareEnabledChannel(ctx context.Context, req *example.QueryMiddlewareEnabledChannelRequest) (*example.QueryMiddlewareEnabledChannelResponse, error) {
	enabled, err := qs.k.MiddlewareEnabled.Has(ctx, collections.Join(req.PortId, req.ChannelId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &example.QueryMiddlewareEnabledChannelResponse{MiddlewareEnabled: enabled}, nil
}

// MiddlewareEnabledChannels defines the handler for the Query/MiddlewareEnabledChannels RPC method.
func (qs queryServer) MiddlewareEnabledChannels(ctx context.Context, req *example.QueryMiddlewareEnabledChannelsRequest) (*example.QueryMiddlewareEnabledChannelsResponse, error) {
	channels, pageRes, err := query.CollectionPaginate(
		ctx,
		qs.k.MiddlewareEnabled,
		req.Pagination,
		func(key collections.Pair[string, string], _ collections.NoValue) (*example.MiddlewareEnabledChannel, error) {
			return &example.MiddlewareEnabledChannel{PortId: key.K1(), ChannelId: key.K2()}, nil
		})
	if err != nil {
		return nil, err
	}

	return &example.QueryMiddlewareEnabledChannelsResponse{MiddlewareEnabledChannels: channels, Pagination: pageRes}, nil
}

// CallbackCounter defines the handler for the Query/CallbackCounter RPC method.
func (qs queryServer) CallbackCounter(ctx context.Context, req *example.QueryCallbackCounterRequest) (*example.QueryCallbackCounterResponse, error) {
	counter, err := qs.k.CallbackCounter.Get(ctx, req.ChannelId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &example.QueryCallbackCounterResponse{Counter: &counter}, nil
}

// Counters defines the handler for the Query/Counters RPC method.
func (qs queryServer) CallbackCounters(ctx context.Context, req *example.QueryCallbackCountersRequest) (*example.QueryCallbackCountersResponse, error) {
	counters, pageRes, err := query.CollectionPaginate(
		ctx,
		qs.k.CallbackCounter,
		req.Pagination,
		func(_ string, value example.CallbackCounter) (*example.CallbackCounter, error) {
			return &value, nil
		})
	if err != nil {
		return nil, err
	}

	return &example.QueryCallbackCountersResponse{Counters: counters, Pagination: pageRes}, nil
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *example.QueryParamsRequest) (*example.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &example.QueryParamsResponse{Params: example.Params{}}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &example.QueryParamsResponse{Params: params}, nil
}
