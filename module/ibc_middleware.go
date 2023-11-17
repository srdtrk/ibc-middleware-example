package module

import (
	"encoding/json"
	"errors"
	"strings"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	"github.com/cosmosregistry/example"
	"github.com/cosmosregistry/example/keeper"
)

var _ porttypes.Middleware = (*IBCMiddleware)(nil)

// IBCMiddleware implements the ICS26 callbacks for example given the
// example keeper and the underlying application.
type IBCMiddleware struct {
	app         porttypes.IBCModule
	ics4Wrapper porttypes.ICS4Wrapper

	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddlware given the keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, ics4Wrapper porttypes.ICS4Wrapper, k keeper.Keeper) IBCMiddleware {
	if app == nil {
		panic(errors.New("IBCModule cannot be nil"))
	}

	if ics4Wrapper == nil {
		panic(errors.New("ICS4Wrapper cannot be nil"))
	}

	return IBCMiddleware{
		app:         app,
		ics4Wrapper: ics4Wrapper,
		keeper:      k,
	}
}

// OnChanOpenInit implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	var versionMetadata example.Metadata

	if strings.TrimSpace(version) == "" {
		// default version
		versionMetadata = example.NewMetadata(example.Version, "")
	} else {
		metadata, err := example.MetadataFromVersion(version)
		if err != nil {
			// Since it is valid for example version to not be specified, the above middleware version may be for a middleware
			// lower down in the stack. Thus, if it is not an example version we pass the entire version string onto the underlying
			// application.
			return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID,
				chanCap, counterparty, version)
		}
		versionMetadata = metadata
	}

	if versionMetadata.ExampleVersion != example.Version {
		return "", errorsmod.Wrapf(example.ErrInvalidChannelVersion, "expected %s, got %s", example.Version, versionMetadata.ExampleVersion)
	}

	appVersion, err := im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, versionMetadata.AppVersion)
	if err != nil {
		return "", err
	}

	versionMetadata.AppVersion = appVersion
	versionBytes, err := json.Marshal(&versionMetadata)
	if err != nil {
		return "", err
	}

	err = im.keeper.MiddlewareEnabled.Set(ctx, collections.Join(portID, channelID))
	if err != nil {
		return "", err
	}

	// call underlying app's OnChanOpenInit callback with the appVersion
	return string(versionBytes), nil
}

// OnChanOpenTry implements the IBCMiddleware interface
// If the channel is not example enabled the underlying application version will be returned
// If the channel is example enabled we merge the underlying application version with the example version
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	versionMetadata, err := example.MetadataFromVersion(counterpartyVersion)
	if err != nil {
		// Since it is valid for example version to not be specified, the above middleware version may be for a middleware
		// lower down in the stack. Thus, if it is not an example version we pass the entire version string onto the underlying
		// application.
		return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
	}

	if versionMetadata.ExampleVersion != example.Version {
		return "", errorsmod.Wrapf(example.ErrInvalidChannelVersion, "expected %s, got %s", example.Version, versionMetadata.ExampleVersion)
	}

	err = im.keeper.MiddlewareEnabled.Set(ctx, collections.Join(portID, channelID))
	if err != nil {
		return "", err
	}

	// call underlying app's OnChanOpenTry callback with the app versions
	appVersion, err := im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, versionMetadata.AppVersion)
	if err != nil {
		return "", err
	}

	versionMetadata.AppVersion = appVersion
	versionBytes, err := json.Marshal(&versionMetadata)
	if err != nil {
		return "", err
	}

	return string(versionBytes), nil
}

// OnChanOpenAck implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	// If handshake was initialized with middleware enabled it must complete with middleware enabled.
	// If handshake was initialized with middleware disabled it must complete with middleware disabled.
	isMiddlewareEnabled, err := im.keeper.MiddlewareEnabled.Has(ctx, collections.Join(portID, channelID))
	if err != nil {
		return err
	}
	if isMiddlewareEnabled {
		versionMetadata, err := example.MetadataFromVersion(counterpartyVersion)
		if err != nil {
			return errorsmod.Wrapf(err, "failed to unmarshal ICS29 counterparty version metadata: %s", counterpartyVersion)
		}

		if versionMetadata.ExampleVersion != example.Version {
			return errorsmod.Wrapf(example.ErrInvalidChannelVersion, "expected counterparty middleware version: %s, got: %s", example.Version, versionMetadata.ExampleVersion)
		}

		// call underlying app's OnChanOpenAck callback with the counterparty app version.
		return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, versionMetadata.AppVersion)
	}

	// call underlying app's OnChanOpenAck callback with the counterparty app version.
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// call underlying app's OnChanOpenConfirm callback.
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// call underlying app's OnChanCloseInit callback.
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// call underlying app's OnChanCloseConfirm callback.
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// SendPacket implements the ICS4 Wrapper interface
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	middlewareEnabled, err := im.keeper.MiddlewareEnabled.Has(ctx, collections.Join(sourcePort, sourceChannel))
	if err != nil {
		middlewareEnabled = false
	}

	if middlewareEnabled {
		counter, err := im.keeper.CallbackCounter.Get(ctx, sourceChannel)
		if err != nil {
			counter = example.NewCallbackCounter(sourceChannel)
		}

		counter.SendPacket++

		if err = im.keeper.CallbackCounter.Set(ctx, sourceChannel, counter); err != nil {
			return 0, err
		}

		// do custom logic here
	}

	// call underlying app's SendPacket callback.
	return im.ics4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// OnRecvPacket implements the IBCMiddleware interface.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// In OnRecvPacket, the destination port and channel are the (port, channel) of the receiving chain.
	middlewareEnabled, err := im.keeper.MiddlewareEnabled.Has(ctx, collections.Join(packet.DestinationPort, packet.DestinationChannel))
	if err != nil {
		middlewareEnabled = false
	}

	if middlewareEnabled {
		counter, err := im.keeper.CallbackCounter.Get(ctx, packet.DestinationChannel)
		if err != nil {
			counter = example.NewCallbackCounter(packet.DestinationChannel)
		}

		counter.OnRecvPacket++

		if err = im.keeper.CallbackCounter.Set(ctx, packet.DestinationChannel, counter); err != nil {
			panic(err)
		}

		// do custom logic here
		// you may also wrap the acknowledgement in a custom acknowledgement type to add additional data to the acknowledgement
	}

	// call underlying app's OnRecvPacket callback.
	return im.app.OnRecvPacket(ctx, packet, relayer)
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	middlewareEnabled, err := im.keeper.MiddlewareEnabled.Has(ctx, collections.Join(packet.SourcePort, packet.SourceChannel))
	if err != nil {
		middlewareEnabled = false
	}

	if middlewareEnabled {
		counter, err := im.keeper.CallbackCounter.Get(ctx, packet.SourceChannel)
		if err != nil {
			counter = example.NewCallbackCounter(packet.SourceChannel)
		}

		counter.OnAcknowledgementPacket++

		if err = im.keeper.CallbackCounter.Set(ctx, packet.SourceChannel, counter); err != nil {
			return err
		}

		// do custom logic here
	}

	// call underlying app's OnAcknowledgementPacket callback.
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface
// If fees are not enabled, this callback will default to the ibc-core packet callback
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	middlewareEnabled, err := im.keeper.MiddlewareEnabled.Has(ctx, collections.Join(packet.SourcePort, packet.SourceChannel))
	if err != nil {
		middlewareEnabled = false
	}

	if middlewareEnabled {
		counter, err := im.keeper.CallbackCounter.Get(ctx, packet.SourceChannel)
		if err != nil {
			counter = example.NewCallbackCounter(packet.SourceChannel)
		}

		counter.OnTimeoutPacket++

		if err = im.keeper.CallbackCounter.Set(ctx, packet.SourceChannel, counter); err != nil {
			return err
		}

		// do custom logic here
	}

	// call underlying app's OnTimeoutPacket callback.
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return im.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

// GetAppVersion returns the application version of the underlying application
func (im IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}
