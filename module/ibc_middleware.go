package module

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	"github.com/cosmosregistry/example/keeper"
)

var _ porttypes.Middleware = (*IBCMiddleware)(nil)

// IBCMiddleware implements the ICS26 callbacks for linked-packets given the
// linked-packets keeper and the underlying application.
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
