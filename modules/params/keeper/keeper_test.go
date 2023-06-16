package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	ctmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	opbtypes "github.com/bianjieai/iritamod/modules/opb/types"
	"github.com/bianjieai/iritamod/modules/params/keeper"
	"github.com/bianjieai/iritamod/modules/params/types"
	cslashingtypes "github.com/bianjieai/iritamod/modules/slashing/types"
	"github.com/bianjieai/iritamod/simapp"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	msgsrv types.MsgServer
	keeper *keeper.Keeper
	app    *simapp.SimApp
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	app := simapp.Setup(suite.T(), false)
	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(false, ctmproto.Header{})
	suite.app = app
	suite.msgsrv = keeper.NewMsgServerImpl(app.CParamsKeeper)
	suite.keeper = &app.CParamsKeeper
}

func (suite *KeeperTestSuite) TestUpdateParams() {
	suite.reset()

	authority := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName).String()
	cases := map[string]struct {
		preRun func() []sdk.Msg
		expErr bool
	}{
		"update slashing": {
			preRun: func() []sdk.Msg {
				param := slashingtypes.DefaultParams()

				return []sdk.Msg{
					&cslashingtypes.MsgUpdateParams{
						Authority: authority,
						Params: cslashingtypes.Params{
							SignedBlocksWindow:      param.SignedBlocksWindow,
							MinSignedPerWindow:      param.MinSignedPerWindow,
							DowntimeJailDuration:    param.DowntimeJailDuration,
							SlashFractionDoubleSign: param.SlashFractionDoubleSign,
							SlashFractionDowntime:   param.SlashFractionDowntime,
						},
					},
				}
			},
			expErr: true,
		},
		"update node": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&nodetypes.MsgUpdateParams{
						Authority: authority,
						Params:    nodetypes.Params{HistoricalEntries: 110},
					},
				}
			},
			expErr: true,
		},
		"update opb": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&opbtypes.MsgUpdateParams{
						Authority: authority,
						Params:    opbtypes.DefaultParams(),
					},
				}
			},
			expErr: true,
		},
		"update mint": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&minttypes.MsgUpdateParams{
						Authority: authority,
						Params:    minttypes.DefaultParams(),
					},
				}
			},
			expErr: true,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msgs := tc.preRun()
			updateMsg, err := types.NewMsgUpdateParams(msgs, authority)
			suite.Require().NoError(err)
			err = updateMsg.ValidateBasic()
			if tc.expErr {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestValidateUpdateParamsMsgs() {
	suite.reset()

	authority := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName).String()
	cases := map[string]struct {
		preRun func() []sdk.Msg
		expErr bool
	}{
		"all messages are invalid": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&nodetypes.MsgGrantNode{
						Name:        "",
						Certificate: "",
						Operator:    "",
					},
					&types.MsgUpdateParams{
						Authority: authority,
						Messages:  nil,
					},
				}
			},
			expErr: true,
		},
		"containing invalid messages": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&nodetypes.MsgUpdateParams{
						Authority: authority,
						Params:    nodetypes.Params{HistoricalEntries: 110},
					},
					&nodetypes.MsgGrantNode{
						Name:        "",
						Certificate: "",
						Operator:    "",
					},
				}
			},
			expErr: true,
		},
		"containing multiple same message": {
			preRun: func() []sdk.Msg {
				return []sdk.Msg{
					&nodetypes.MsgUpdateParams{
						Authority: authority,
						Params:    nodetypes.Params{HistoricalEntries: 110},
					},
					&nodetypes.MsgUpdateParams{
						Authority: authority,
						Params:    nodetypes.Params{HistoricalEntries: 111},
					},
				}
			},
			expErr: false,
		},
		// 	TODO: uncomment this when proto package is renamed.
		//"containing cparam update param message": {
		//	preRun: func() []sdk.Msg {
		//		return []sdk.Msg{
		//			&nodetypes.MsgUpdateParams{
		//				Authority: authority,
		//				Params:    nodetypes.Params{HistoricalEntries: 110},
		//			},
		//
		//			&types.MsgUpdateParams{
		//				Authority: authority,
		//				Messages:  nil,
		//			},
		//			&minttypes.MsgUpdateParams{
		//				Authority: authority,
		//				Params:    minttypes.DefaultParams(),
		//			},
		//		}
		//	},
		//	expErr: true,
		//},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msgs := tc.preRun()
			updateMsg, err := types.NewMsgUpdateParams(msgs, authority)
			suite.Require().NoError(err)
			err = updateMsg.ValidateBasic()
			if tc.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
