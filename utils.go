package identity

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

// NewBJJSecretKey generates a new secret key for the Baby JubJub curve.
func NewBJJSecretKey() []byte {
	secretKey := babyjub.NewRandPrivKey()

	return secretKey.Scalar().BigInt().Bytes()
}

// ByteArrayToBits converts a byte array to a bit array.
func ByteArrayToBits(bytes []byte) []int64 {
	var bits []int64
	for _, b := range bytes {
		for i := 0; i < 8; i++ {
			bits = append(bits, int64(b>>uint(7-i)&1))
		}
	}

	return bits
}

// SmartChunking splits a big.Int into chunks of 8 bytes.
//
// It does its best to split the big.Int into chunks of 8 bytes,
// but it may not be perfect though it smart (I heard... do not believe everything you hear).
func SmartChunking(x *big.Int) []string {
	var res []string

	mod := big.NewInt(1)
	for i := 0; i < 64; i++ {
		mod.Mul(mod, big.NewInt(2))
	}
	for i := 0; i < 64; i++ {
		chunk := new(big.Int).Mod(x, mod)

		res = append(res, chunk.String())
		x.Div(x, mod)
	}

	return res
}

// RsaPubKeyPemToN extracts the modulus from a RSA public key PEM.
func RsaPubKeyPemToN(pubKeyPem []byte) (*big.Int, error) {
	block, _ := pem.Decode(pubKeyPem)
	if block == nil {
		return nil, fmt.Errorf("error decoding public key pem")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to rsa public key")
	}

	return rsaPubKey.N, nil
}
