package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// query endpoints supported by the NFT Querier
const (
	QuerySupply     = "supply"
	QueryOwner      = "owner"
	QueryCollection = "collection"
	QueryDenoms     = "denoms"
	QueryDenom      = "denom"
	QueryNFT        = "nft"
)

// QuerySupplyParams defines the params for queries:
type QuerySupplyParams struct {
	Denom string
	Owner sdk.AccAddress
}

// NewQuerySupplyParams creates a new instance of QuerySupplyParams
func NewQuerySupplyParams(denom string, owner sdk.AccAddress) QuerySupplyParams {
	return QuerySupplyParams{
		Denom: denom,
		Owner: owner,
	}
}

// Bytes exports the Denom as bytes
func (q QuerySupplyParams) Bytes() []byte {
	return []byte(q.Denom)
}

// QueryOwnerParams defines the params for queries:
type QueryOwnerParams struct {
	Denom string
	Owner sdk.AccAddress
}

// NewQuerySupplyParams creates a new instance of QuerySupplyParams
func NewQueryOwnerParams(denom string, owner sdk.AccAddress) QueryOwnerParams {
	return QueryOwnerParams{
		Denom: denom,
		Owner: owner,
	}
}

// QuerySupplyParams defines the params for queries:
type QueryCollectionParams struct {
	Denom string
}

// NewQueryCollectionParams creates a new instance of QueryCollectionParams
func NewQueryCollectionParams(denom string) QueryCollectionParams {
	return QueryCollectionParams{
		Denom: denom,
	}
}

// QueryDenomParams defines the params for queries:
type QueryDenomParams struct {
	ID string
}

// NewQueryDenomParams creates a new instance of QueryDenomParams
func NewQueryDenomParams(id string) QueryDenomParams {
	return QueryDenomParams{
		ID: id,
	}
}

// QueryNFTParams params for query 'custom/nfts/nft'
type QueryNFTParams struct {
	Denom   string
	TokenID string
}

// NewQueryNFTParams creates a new instance of QueryNFTParams
func NewQueryNFTParams(denom, id string) QueryNFTParams {
	return QueryNFTParams{
		Denom:   denom,
		TokenID: id,
	}
}
