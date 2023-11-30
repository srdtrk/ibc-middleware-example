package example_test

import (
	"encoding/json"
	"errors"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/suite"

	"cosmossdk.io/collections"
	"cosmossdk.io/log"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"

	"github.com/cosmosregistry/example"
	"github.com/cosmosregistry/example/testing/simapp"
)

func init() {
	ibctesting.DefaultTestingAppInit = SetupTestingApp
}

// SetupTestingApp provides the duplicated simapp which is specific to the callbacks module on chain creation.
func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	app := simapp.NewSimApp(log.NewNopLogger(), db, nil, true, simtestutil.EmptyAppOptions{})
	return app, app.DefaultGenesis()
}

// GetSimApp returns the duplicated SimApp from within the callbacks directory.
// This must be used instead of chain.GetSimApp() for tests within this directory.
func GetSimApp(chain *ibctesting.TestChain) *simapp.SimApp {
	app, ok := chain.App.(*simapp.SimApp)
	if !ok {
		panic(errors.New("chain is not a simapp.SimApp"))
	}
	return app
}

// MiddlewareTestSuite defines the needed instances and methods to test callbacks
type MiddlewareTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	path *ibctesting.Path
}

// setupChains sets up a coordinator with 2 test chains.
func (s *MiddlewareTestSuite) setupChains() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 2)
	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(1))
	s.chainB = s.coordinator.GetChain(ibctesting.GetChainID(2))
	s.path = ibctesting.NewPath(s.chainA, s.chainB)
}

// SetupTransferTest sets up a non-middleware transfer channel between chainA and chainB
func (s *MiddlewareTestSuite) SetupTransferTest() {
	s.setupChains()

	s.path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	s.path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	s.path.EndpointA.ChannelConfig.Version = transfertypes.Version
	s.path.EndpointB.ChannelConfig.Version = transfertypes.Version

	s.coordinator.Setup(s.path)
}

// SetupMiddlewareTransferTest sets up an example middleware enabled transfer channel between chainA and chainB
func (s *MiddlewareTestSuite) SetupMiddlewareTransferTest() {
	s.setupChains()

	byteVersion, err := json.Marshal(example.Metadata{ExampleVersion: example.Version, AppVersion: transfertypes.Version})
	s.Require().NoError(err)

	mwTransferVersion := string(byteVersion)
	s.path.EndpointA.ChannelConfig.Version = mwTransferVersion
	s.path.EndpointB.ChannelConfig.Version = mwTransferVersion
	s.path.EndpointA.ChannelConfig.PortID = transfertypes.PortID
	s.path.EndpointB.ChannelConfig.PortID = transfertypes.PortID

	s.coordinator.Setup(s.path)

	isEnabled, err := GetSimApp(s.chainA).MiddlewareKeeper.MiddlewareEnabled.Has(
		s.chainA.GetContext(),
		collections.Join(s.path.EndpointA.ChannelConfig.PortID, s.path.EndpointA.ChannelID),
	)
	s.Require().NoError(err)
	s.Require().True(isEnabled)
}
