package module

import (
	"encoding/json"
	"errors"
	"strings"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	"github.com/cosmosregistry/example"
	"github.com/cosmosregistry/example/keeper"
)

// var _ porttypes.Middleware = (*IBCMiddleware)(nil)

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
			return errorsmod.Wrapf(example.ErrInvalidChannelVersion, "expected counterparty linked packets version: %s, got: %s", example.Version, versionMetadata.ExampleVersion)
		}

		// call underlying app's OnChanOpenAck callback with the counterparty app version.
		return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, versionMetadata.AppVersion)
	}

	// call underlying app's OnChanOpenAck callback with the counterparty app version.
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
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
