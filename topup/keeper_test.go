package topup_test

import (
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/heimdall/app"
	"github.com/maticnetwork/heimdall/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/types/simulation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app *app.HeimdallApp
	ctx sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app, suite.ctx, _ = createTestApp(false)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// Tests

func (suite *KeeperTestSuite) TestTopupSequenceSet() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	topupSequence := strconv.Itoa(simulation.RandIntBetween(r1, 1000, 100000))

	app.TopupKeeper.SetTopupSequence(ctx, topupSequence)

	actualResult := app.TopupKeeper.HasTopupSequence(ctx, topupSequence)
	require.Equal(t, true, actualResult)
}

// tests setter/getters for Dividend account
func (suite *KeeperTestSuite) TestDividendAccount() {
	t, app, ctx := suite.T(), suite.app, suite.ctx

	dividendAccount := types.DividendAccount{
		User:      hmTypes.BytesToHeimdallAddress([]byte("some-address")),
		FeeAmount: big.NewInt(0).String(),
	}
	app.TopupKeeper.AddDividendAccount(ctx, dividendAccount)
	ok := app.TopupKeeper.CheckIfDividendAccountExists(ctx, dividendAccount.User)
	require.Equal(t, ok, true)
}
