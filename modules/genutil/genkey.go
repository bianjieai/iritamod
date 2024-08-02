package genutil

import (
	"crypto/ed25519"
	"errors"
	"iritamod.bianjie.ai/modules/genutil/utils/ca"

	"github.com/cometbft/cometbft/crypto"
	ed25519util "github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/crypto/sm2"
)

func Genkey(privKey crypto.PrivKey) ([]byte, error) {
	switch pk := privKey.(type) {
	case sm2.PrivKeySm2:
		privKey := privKey.(sm2.PrivKeySm2)
		return ca.Sm2Cert{PrivateKey: privKey.GetPrivateKey()}.WritePrivateKeytoMem()
	case ed25519util.PrivKey:
		priKey := make([]byte, ed25519.PrivateKeySize)
		copy(priKey, pk[:])
		return ca.X509Cert{PrivateKey: ed25519.PrivateKey(priKey)}.WritePrivateKeytoMem()
	default:
		return nil, errors.New("unsupported algorithm type")
	}

}
