package identity

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

type Profile struct {
	secretKey   *babyjub.PrivateKey
	publicKey   *babyjub.PublicKey
	dataBlinder *big.Int
}

func NewProfile(secretKeyHex string) (*Profile, error) {
	rawSecretKey, err := hex.DecodeString(secretKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error decoding secret key: %v", err)
	}

	secretKey := babyjub.PrivateKey(rawSecretKey)

	secretKeyInt := secretKey.Scalar().BigInt()

	dataBlinder, err := poseidon.Hash([]*big.Int{secretKeyInt})
	if err != nil {
		return nil, fmt.Errorf("error hashing secret key: %v", err)
	}

	return &Profile{
		secretKey:   &secretKey,
		publicKey:   secretKey.Public(),
		dataBlinder: dataBlinder,
	}, nil
}

func (p *Profile) GetRegistrationChallenge() ([]byte, error) {
	publicKeyHash, err := poseidon.Hash([]*big.Int{p.publicKey.X, p.publicKey.Y})
	if err != nil {
		return nil, fmt.Errorf("error hashing public key: %v", err)
	}

	return publicKeyHash.Bytes()[:8], nil
}
