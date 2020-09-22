package utils

import (
	"errors"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

func PubKeySm2From(pubkey []byte) (*sm2.PublicKey, error) {
	if len(pubkey) == 0 {
		return nil, errors.New("empty pubkey")
	}

	switch {
	case (pubkey[0] == 2 || pubkey[0] == 3) && len(pubkey) == 33:
		return sm2.Decompress(pubkey), nil
	case pubkey[0] == 4 && len(pubkey) == 65:
		return &sm2.PublicKey{
			Curve: sm2.P256Sm2(),
			X:     new(big.Int).SetBytes(pubkey[1:33]),
			Y:     new(big.Int).SetBytes(pubkey[33:64]),
		}, nil
	case pubkey[0] != 4 && len(pubkey) == 64:
		return &sm2.PublicKey{
			Curve: sm2.P256Sm2(),
			X:     new(big.Int).SetBytes(pubkey[0:32]),
			Y:     new(big.Int).SetBytes(pubkey[32:64]),
		}, nil
	}
	return nil, errors.New("invalid pubkey format")
}
