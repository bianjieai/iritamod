package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/bianjieai/iritamod/modules/layer2/keeper"
	"github.com/bianjieai/iritamod/modules/layer2/types"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
	permtypes "github.com/bianjieai/iritamod/modules/perm/types"
	"github.com/bianjieai/iritamod/simapp"
)

var (
	rootAdmin = sdk.AccAddress(tmhash.SumTruncated([]byte("rootAdmin")))
	accAvata  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_avata"))) // l2 user: cosmos1j0898zyz64cyxy2s2km99t2c3s6tn5tzfppw9h
	accXvata  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_xvata"))) // l2 user
	accAlice  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_alice"))) // cosmos16877jxzrdetmzsl3pntv4n402m8d0cpvwd74w7
	accBob    = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_bob")))   // cosmos1a53v8ksyd6x47sju572t48s3ynmqyan0n2c6kx

	avataSpaceId   = uint64(1) //
	avataSpaceName = "Avata Space"
	avataSpaceUri  = "https://space.avata.com"

	badKidsClassId  = "badKids"
	badKidsClassUri = "https://badkids.com"
	// TODO： build from mock_data/nfts.json
	badKidsTokensForNFT = []types.TokenForNFT{
		{
			Id:    "kid001",
			Owner: accAlice.String(),
		},
		{
			Id:    "kid002",
			Owner: accBob.String(),
		},
	}
)

type TestSuite struct {
	suite.Suite

	ctx        sdk.Context
	cdc        *codec.LegacyAmino
	permKeeper permkeeper.Keeper
	keeper     keeper.Keeper
	app        *simapp.SimApp
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// Note: After setting up, we have:
// 1. the spaceId(1) has been created and belongs to accAvata
// 2. layer1 nft holds class `badKids` and two nfts from mocked file `mock_data/nfts.json`
// 3. class `badKids` is owned by `alice`
func (s *TestSuite) SetupTest() {
	app := simapp.Setup(false)

	s.cdc = app.LegacyAmino()
	s.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	s.app = app
	s.permKeeper = app.PermKeeper
	s.keeper = app.Layer2Keeper

	s.prepareRoles()
	s.prepareLayer2()
}

func (s *TestSuite) prepareRoles() {
	err := s.permKeeper.Authorize(s.ctx, accAvata, rootAdmin, permtypes.RoleLayer2User)
	if err != nil {
		panic("failed to authorize role")
	}
	err = s.permKeeper.Authorize(s.ctx, accXvata, rootAdmin, permtypes.RoleLayer2User)
	if err != nil {
		panic("failed to authorize role")
	}
}

func (s *TestSuite) prepareLayer2() {
	id, err := s.keeper.CreateSpace(s.ctx, avataSpaceName, avataSpaceUri, accAvata)
	s.Require().NoError(err)
	s.Require().Equal(avataSpaceId, id)
}

//accL2UserStr = "cosmos1l6rgzjskese3dew3vkru3fk3gv8d0e7d63l5fd"
//accAliceStr  = "cosmos1l4vsfaujy5sy9rcgkggv4vnlwefhwupc8t8362"
//accBobStr    = "cosmos1y7ud35l6tt50aeycd4xcjmu6z844v9edj90kt0"
//func (s *TestSuite) setupAddrs() {
//	accL2User, _ = sdk.AccAddressFromBech32(accL2UserStr)
//	accAlice, _ = sdk.AccAddressFromBech32(accAliceStr)
//	accBob, _ = sdk.AccAddressFromBech32(accBobStr)
//}
//
//func (s *TestSuite) setupAddrs(numAddrs int) []sdk.AccAddress {
//	var addresses []sdk.AccAddress
//	var buffer bytes.Buffer
//
//	for i := 100; i < (numAddrs + 100); i++ {
//		numString := strconv.Itoa(i)
//		buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6")
//		buffer.WriteString(numString)
//		res, _ := sdk.AccAddressFromHex(buffer.String())
//		bech := res.String()
//		addresses = append(addresses, testAddr(buffer.String(), bech))
//		buffer.Reset()
//	}
//
//	return addresses
//}
//
//func testAddr(addr string, bech string) sdk.AccAddress {
//	res, err := sdk.AccAddressFromHex(addr)
//	if err != nil {
//		panic(err)
//	}
//	bechexpected := res.String()
//	if bech != bechexpected {
//		panic("Bech encoding doesn't match reference")
//	}
//
//	bechres, err := sdk.AccAddressFromBech32(bech)
//	if err != nil {
//		panic(err)
//	}
//	if !bytes.Equal(bechres, res) {
//		panic("Bech decode and hex decode don't match")
//	}
//
//	return res
//}
