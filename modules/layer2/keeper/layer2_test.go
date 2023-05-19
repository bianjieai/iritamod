package keeper_test

func (s *TestSuite) TestCreateSpace() {
	spaceName := "NewSpace"
	spaceUri := "NewSpaceUri"
	spaceId, err := s.keeper.CreateL2Space(s.ctx, spaceName, spaceUri, accAvata)
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

	err = s.keeper.TransferL2Space(s.ctx, avataSpaceId, accAvata, accXvata)
	s.Require().NoErrorf(err, "failed to transfer space")

	space, err = s.keeper.GetSpace(s.ctx, avataSpaceId)
	s.Require().NoErrorf(err, "failed to get space")
	s.Require().Equal(accXvata.String(), space.Owner)
}

func (s *TestSuite) TestCreateL2BlockHeader() {
	height := uint64(1000)
	header := "block header"
	err := s.keeper.CreateL2BlockHeader(s.ctx, avataSpaceId, height, header, accAvata)
	s.Require().NoErrorf(err, "failed to create block header")

	resHeader, err := s.keeper.GetL2BlockHeader(s.ctx, avataSpaceId, height)
	s.Require().NoErrorf(err, "failed to get block header")
	s.Require().Equal(header, resHeader)
}
