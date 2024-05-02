package identity

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

const smartChunking2BlockSize uint64 = 512

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

// SmartChunking2 does some weird stuff.
//
// For more details contact the den4ik.
func SmartChunking2(bits []int64, blockNumber uint64) []int64 {
	dataBitsNumber := uint64(len(bits) + 1 + 64)
	dataBlockNumber := dataBitsNumber/smartChunking2BlockSize + 1
	zeroDataBitsNumber := dataBlockNumber*smartChunking2BlockSize - dataBitsNumber

	var result []int64
	result = append(result, bits...)
	result = append(result, 1)

	for i := uint64(0); i < zeroDataBitsNumber; i++ {
		result = append(result, 0)
	}

	bitsNumberBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bitsNumberBytes, uint64(len(bits)))

	bitsNumber := ByteArrayToBits(bitsNumberBytes)

	result = append(result, bitsNumber...)

	restBlocksNumber := blockNumber - dataBlockNumber

	for i := uint64(0); i < restBlocksNumber*smartChunking2BlockSize; i++ {
		result = append(result, 0)
	}

	return result
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

// CalculateProofIndex calculates the proof index.
func CalculateProofIndex(passportKey string, identityKey string) ([]byte, error) {
	passportKeyInt, ok := new(big.Int).SetString(passportKey, 10)
	if !ok {
		return nil, fmt.Errorf("passport key is not int")
	}

	identityKeyInt, ok := new(big.Int).SetString(identityKey, 10)
	if !ok {
		return nil, fmt.Errorf("identity key is not int")
	}

	hash, err := poseidon.Hash([]*big.Int{passportKeyInt, identityKeyInt})
	if err != nil {
		return nil, fmt.Errorf("error hashing passport and identity key: %v", err)
	}

	return hash.Bytes(), nil
}

// BigIntToBytes converts a big integer to a byte array.
func BigIntToBytes(x string) ([]byte, error) {
	bigInt, ok := new(big.Int).SetString(x, 10)
	if !ok {
		return nil, fmt.Errorf("error converting string to big int")
	}

	return bigInt.Bytes(), nil
}
