package keeper_test

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

func (s *TestSuite) TestCreateSpace() {
	spaceName := "NewSpace"
	spaceUri := "NewSpaceUri"
	spaceId, err := s.keeper.CreateSpace(s.ctx, spaceName, spaceUri, accAvata)
	s.Require().NoErrorf(err, "failed to create space")
	s.Require().Equal(spaceId, uint64(2))

	space, err := s.keeper.GetSpace(s.ctx, spaceId)
	s.Require().NoErrorf(err, "failed to get space")
	s.Require().Equal(spaceName, space.Name)
	s.Require().Equal(spaceUri, space.Uri)
	s.Require().Equal(spaceId, space.Id)
	s.Require().Equal(accAvata.String(), space.Owner)
}

func (s *TestSuite) TestTransferSpace() {
	space, err := s.keeper.GetSpace(s.ctx, avataSpaceId)
	s.Require().NoErrorf(err, "failed to get space")
	s.Require().Equal(accAvata.String(), space.Owner)

	err = s.keeper.TransferSpace(s.ctx, avataSpaceId, accAvata, accXvata)
	s.Require().NoErrorf(err, "failed to transfer space")

	space, err = s.keeper.GetSpace(s.ctx, avataSpaceId)
	s.Require().NoErrorf(err, "failed to get space")
	s.Require().Equal(accXvata.String(), space.Owner)
}

func (s *TestSuite) TestCreateBlockHeader() {
	height := uint64(1000)
	header := "block header"
	err := s.keeper.CreateBlockHeader(s.ctx, avataSpaceId, height, header, accAvata)
	s.Require().NoErrorf(err, "failed to create block header")

	resHeader, err := s.keeper.GetBlockHeader(s.ctx, avataSpaceId, height)
	s.Require().NoErrorf(err, "failed to get block header")
	s.Require().Equal(header, resHeader)

	h, err := s.keeper.GetSpaceLatestHeight(s.ctx, avataSpaceId)
	s.Require().NoErrorf(err, "failed to get block header latest height")
	s.Require().Equal(height, h)

	var expected tmbytes.HexBytes = tmhash.Sum(s.ctx.TxBytes())
	txHash, err := s.keeper.GetBlockHeaderTxHash(s.ctx, avataSpaceId, height)
	s.Require().NoErrorf(err, "failed to get block header tx hash")
	s.Require().Equal(expected.String(), txHash)

	fmt.Println(txHash)
}
