package keeper_test

import "github.com/bianjieai/iritamod/modules/layer2/types"

func (s *TestSuite) TestCreateTokensForNFT() {
	err := s.keeper.CreateTokensForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsTokensForNFT, accAvata)
	s.Require().NoErrorf(err, "failed to create tokens")

	for _, token := range badKidsTokensForNFT {
		owner, err := s.keeper.GetTokenOwnerForNFT(s.ctx, avataSpaceId, badKidsClassId, token.Id)
		s.Require().NoErrorf(err, "failed to get token owner")
		s.Require().Equal(token.Owner, owner.String())
	}
}

func (s *TestSuite) TestUpdateTokensForNFT() {
	err := s.keeper.CreateTokensForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsTokensForNFT, accAvata)
	s.Require().NoErrorf(err, "failed to create tokens")

	// exchange owner
	badKidsTokensForNFT[0].Owner = accBob.String()
	badKidsTokensForNFT[1].Owner = accAlice.String()

	err = s.keeper.UpdateTokensForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsTokensForNFT, accAvata)
	s.Require().NoErrorf(err, "failed to update tokens")

	for _, token := range badKidsTokensForNFT {
		owner, err := s.keeper.GetTokenOwnerForNFT(s.ctx, avataSpaceId, badKidsClassId, token.Id)
		s.Require().NoErrorf(err, "failed to get token owner")
		s.Require().Equal(token.Owner, owner.String())
	}
}

func (s *TestSuite) TestDeleteTokensForNFT() {
	tokenId1 := badKidsTokensForNFT[0].Id
	err := s.keeper.DepositL1TokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId1, accAlice)
	s.Require().NoErrorf(err, "failed to deposit token")

	tokenId2 := badKidsTokensForNFT[1].Id
	err = s.keeper.DepositL1TokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId2, accBob)
	s.Require().NoErrorf(err, "failed to deposit token")

	err = s.keeper.UpdateTokensForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsTokensForNFT, accAvata)
	s.Require().NoErrorf(err, "failed to update tokens")

	nftIds := []string{badKidsTokensForNFT[0].Id, badKidsTokensForNFT[1].Id}
	err = s.keeper.DeleteTokensForNFT(s.ctx, avataSpaceId, badKidsClassId, nftIds, accAvata)
	s.Require().NoErrorf(err, "failed to delete tokens")

	for _, nftId := range nftIds {
		_, err := s.keeper.GetTokenOwnerForNFT(s.ctx, avataSpaceId, badKidsClassId, nftId)
		s.Require().Error(err)
	}
}

// TODO: table driven
func (s *TestSuite) TestDepositClassForNFT() {
	// alice deposit the class
	err := s.keeper.DepositL1ClassForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsClassUri, accAlice, accAlice)
	s.Require().NoErrorf(err, "failed to deposit class")

	class, err := s.keeper.GetClassForNFT(s.ctx, badKidsClassId)
	s.Require().NoErrorf(err, "failed to get class mapping")
	s.Require().Equal(badKidsClassId, class.Id)
	s.Require().Equal(badKidsClassUri, class.BaseUri)
	s.Require().Equal(accAlice.String(), class.Owner)

	badKids, err := s.keeper.GetNFTKeeper().GetClass(s.ctx, badKidsClassId)
	s.Require().NoErrorf(err, "failed to get layer1 class")
	s.Require().Equal(types.ModuleAccAddress.String(), badKids.GetCreator())
}

// TODO: table driven
func (s *TestSuite) TestUpdateClassForNFT() {
	// prepare: alice deposit the class
	err := s.keeper.DepositL1ClassForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsClassUri, accAlice, accAlice)
	s.Require().NoErrorf(err, "failed to deposit class")

	newUri := badKidsClassUri + "modified"
	classUpdates := []*types.UpdateClassForNFT{
		{
			Id:    badKidsClassId,
			Uri:   newUri,
			Owner: accBob.String(),
		},
	}

	err = s.keeper.UpdateL2ClassesForNFT(s.ctx, classUpdates, accAvata)
	s.Require().NoErrorf(err, "failed to update class")

	class, err := s.keeper.GetClassForNFT(s.ctx, badKidsClassId)
	s.Require().NoErrorf(err, "failed to get class")
	s.Require().Equal(badKidsClassId, class.Id)
	s.Require().Equal(newUri, class.BaseUri)
}

func (s *TestSuite) TestWithdrawClassForNFT() {
	// prepare: alice deposit the class
	err := s.keeper.DepositL1ClassForNFT(s.ctx, avataSpaceId, badKidsClassId, badKidsClassUri, accAlice, accAlice)
	s.Require().NoErrorf(err, "failed to deposit class")

	badKids, err := s.keeper.GetNFTKeeper().GetClass(s.ctx, badKidsClassId)
	s.Require().NoErrorf(err, "failed to get layer1 class")
	s.Require().Equal(types.ModuleAccAddress.String(), badKids.GetCreator())

	err = s.keeper.WithdrawL2ClassForNFT(s.ctx, badKidsClassId, accAlice, accAvata)
	s.Require().NoErrorf(err, "failed to withdraw class")

	badKids, err = s.keeper.GetNFTKeeper().GetClass(s.ctx, badKidsClassId)
	s.Require().NoErrorf(err, "failed to get layer1 class")
	s.Require().Equal(accAlice.String(), badKids.GetCreator())
}

func (s *TestSuite) TestDepositTokenForNFT() {
	tokenId := badKidsTokensForNFT[0].Id
	err := s.keeper.DepositL1TokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId, accAlice)
	s.Require().NoErrorf(err, "failed to deposit token")

	owner, err := s.keeper.GetTokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId)
	s.Require().NoErrorf(err, "failed to get token")
	s.Require().Equal(accAlice.String(), owner.String())

	nft, err := s.keeper.GetNFTKeeper().GetNFT(s.ctx, badKidsClassId, tokenId)
	s.Require().NoErrorf(err, "failed to get layer1 token")
	s.Require().Equal(types.ModuleAccAddress.String(), nft.GetOwner().String())
}

func (s *TestSuite) TestWithdrawTokenForNFT() {
	tokenId := badKidsTokensForNFT[0].Id
	err := s.keeper.DepositL1TokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId, accAlice)
	s.Require().NoErrorf(err, "failed to deposit token")

	owner, err := s.keeper.GetTokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId)
	s.Require().NoErrorf(err, "failed to get token")
	s.Require().Equal(accAlice.String(), owner.String())

	nft, err := s.keeper.GetNFTKeeper().GetNFT(s.ctx, badKidsClassId, tokenId)
	s.Require().NoErrorf(err, "failed to get layer1 token")
	s.Require().Equal(types.ModuleAccAddress.String(), nft.GetOwner().String())

	newTokenName := "new-name"
	newTokenUri := "new-uri"
	newTokenUriHash := "new-uri-hash"
	newTokenData := "new-data"

	err = s.keeper.WithdrawL2TokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId,
		newTokenName, newTokenUri, newTokenUriHash, newTokenData,
		accAlice, accAvata)
	s.Require().NoErrorf(err, "failed to withdraw token")

	exist := s.keeper.HasTokenForNFT(s.ctx, avataSpaceId, badKidsClassId, tokenId)
	s.Require().Equal(false, exist)

	nft, err = s.keeper.GetNFTKeeper().GetNFT(s.ctx, badKidsClassId, tokenId)
	s.Require().NoErrorf(err, "failed to get layer1 token")
	s.Require().Equal(accAlice, nft.GetOwner())
}
