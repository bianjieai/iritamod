package keeper_test

//func (s *TestSuite) TestCreateClassForNFT() {
//	err := s.keeper.CreateClassForNFT(s.ctx, badKidsClassId, badKidsClassUri, accAlice.String(), true)
//	s.Require().NoErrorf(err, "failed to create class mapping")
//
//	class, ok := s.keeper.GetClassForNFT(s.ctx, badKidsClassId)
//	s.Require().Equal(true, ok)
//	s.Require().Equal(badKidsClassId, class.Id)
//}
//
//func (s *TestSuite) TestUpdateClassForNFT() {
//	err := s.keeper.CreateClassForNFT(s.ctx, badKidsClassId, badKidsClassUri, accAlice.String(), true)
//	s.Require().NoErrorf(err, "failed to create class mapping")
//
//	err = s.keeper.UpdateClassForNFT(s.ctx, badKidsClassId, badKidsClassUri, accBob.String())
//	s.Require().NoErrorf(err, "failed to update class mapping")
//
//	class, ok := s.keeper.GetClassForNFT(s.ctx, badKidsClassId)
//	s.Require().Equal(true, ok)
//	s.Require().Equal(badKidsClassId, class.Id)
//	s.Require().Equal(accBob.String(), class.Owner)
//}
//
//func (s *TestSuite) TestCreateTokensForNFT() {
//	err := s.keeper.CreateTokensForNFT(s.ctx, defaultSpaceId, badKidsClassId, nfts, accL2User)
//	s.Require().NoErrorf(err, "failed to create tokens")
//
//	for _, nft := range nfts {
//		owner, ok := s.keeper.GetTokenOwnerForNFT(s.ctx, defaultSpaceId, badKidsClassId, nft.Id)
//		s.Require().Equal(true, ok)
//		s.Require().Equal(nft.Owner, owner.String())
//	}
//}
//
//func (s *TestSuite) TestUpdateTokensForNFT() {
//	err := s.keeper.CreateTokensForNFT(s.ctx, defaultSpaceId, badKidsClassId, nfts, accL2User)
//	s.Require().NoErrorf(err, "failed to create tokens")
//
//	// exchange owner
//	nfts[0].Owner = accBob.String()
//	nfts[1].Owner = accAlice.String()
//
//	err = s.keeper.UpdateTokensForNFT(s.ctx, defaultSpaceId, badKidsClassId, nfts, accL2User)
//	s.Require().NoErrorf(err, "failed to update tokens")
//
//	for _, nft := range nfts {
//		owner, ok := s.keeper.GetTokenOwnerForNFT(s.ctx, defaultSpaceId, badKidsClassId, nft.Id)
//		s.Require().Equal(true, ok)
//		s.Require().Equal(nft.Owner, owner.String())
//	}
//}
//
//func (s *TestSuite) TestDeleteTokensForNFT() {
//	err := s.keeper.CreateTokensForNFT(s.ctx, defaultSpaceId, badKidsClassId, nfts, accL2User)
//	s.Require().NoErrorf(err, "failed to create tokens")
//
//	nftIds := []string{nfts[0].Id, nfts[1].Id}
//
//	err = s.keeper.DeleteTokensForNFT(s.ctx, defaultSpaceId, badKidsClassId, nftIds, accL2User)
//	for _, nftId := range nftIds {
//		_, ok := s.keeper.GetTokenOwnerForNFT(s.ctx, defaultSpaceId, badKidsClassId, nftId)
//		s.Require().Equal(false, ok)
//	}
//}
//
//func (s *TestSuite) TestDepositClassForNFT() {
//
//}
//
//func (s *TestSuite) TestWithdrawClassForNFT() {
//	testCases := []struct {
//		msg      string
//		malleate func()
//	}{
//		{
//			msg: "class locked on layer1",
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		s.Run(tc.msg, func() {
//			tc.malleate()
//
//			s.Require()
//		})
//	}
//}
//
//func (s *TestSuite) TestDepositTokenForNFT() {
//	testCases := []struct {
//		msg      string
//		malleate func()
//	}{
//		{
//			msg: "",
//		},
//		{
//			msg: "",
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		s.Run(tc.msg, func() {
//			tc.malleate()
//
//			s.Require()
//		})
//	}
//}
//
//func (s *TestSuite) TestWithdrawTokenForNFT() {
//	testCases := []struct {
//		msg      string
//		malleate func()
//		expPass  bool
//	}{
//		{
//			msg:     "token locked on layer1",
//			expPass: true,
//		},
//		{
//			msg:     "token not exist on layer1",
//			expPass: true,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		s.Run(tc.msg, func() {
//			tc.malleate()
//
//			s.Require()
//		})
//	}
//}
