package identity

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math/big"
	"encoding/hex"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/rarimo/ldif-sdk/ldif"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const smartChunking2BlockSize uint64 = 512
const brainpoolP256CurveOID = "1.2.840.10045.2.1"
const lowSMaxHex = "54fdabedd0f754de1f3305484ec1c6b9371dfb11ea9310141009a40e8fb729bb"
const nHex = "A9FB57DBA1EEA9BC3E660A909D838D718C397AA3B561A6F7901E0E82974856A7"

// SignMessageWithSecp256k1 signs a string message using a private key string (hex format) and the secp256k1 curve.
func SignMessageWithSecp256k1(privateKey string, message string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error decoding private key hex: %v", err)
	}

	messageBytes := []byte(message)

	hash := sha256.Sum256(messageBytes)

	signature, err := secp256k1.Sign(hash[:], privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("error signing the message: %v", err)
	}

	signatureHex := hex.EncodeToString(signature)

	return signatureHex, nil
}

// LoadMasterCertificatesPem loads the master certificates from an LDIF file in an S3 bucket.
func LoadMasterCertificatesPem(bucketName string, fileName string) ([]byte, error) {
	masters, err := ldif.FromS3Bucket(context.Background(), bucketName, fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading master certificates: %v", err)
	}

	mastersPemSlice := masters.ToPem()

	var mastersPem []byte
	for _, masterPem := range mastersPemSlice {
		mastersPem = append(mastersPem, []byte(masterPem)...)
	}

	return mastersPem, nil
}

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
func SmartChunking(x *big.Int, chunksNumber int) []string {
	var res []string

	mod := big.NewInt(1)
	for i := 0; i < 64; i++ {
		mod.Mul(mod, big.NewInt(2))
	}
	for i := 0; i < chunksNumber; i++ {
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

// pubKeyPemToRaw extracts the modulus from a RSA public key PEM.
func pubKeyPemToRaw(pubKeyPem []byte) ([]byte, bool, error) {
	block, _ := pem.Decode(pubKeyPem)
	if block == nil {
		return nil, false, fmt.Errorf("error decoding public key pem")
	}

	var info publicKeyInfo
	_, err := asn1.Unmarshal(block.Bytes, &info)
	if err == nil {
		if info.Algorithm.Algorithm.String() == brainpoolP256CurveOID {
			var raw []byte
			raw = append(raw, info.SubjectPublicKey.Bytes[1:]...)

			return raw, true, nil
		}
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, false, fmt.Errorf("error parsing public key: %v", err)
	}

	var isEcdsa bool
	var raw []byte
	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		isEcdsa = false
		raw = pub.N.Bytes()
	case *ecdsa.PublicKey:
		isEcdsa = true
		raw = pub.X.Bytes()
		raw = append(raw, pub.Y.Bytes()...)
	default:
		return nil, false, fmt.Errorf("unsupported public key type: %T", pub)
	}

	return raw, isEcdsa, nil
}

func pemToRsaPubKey(pubKeyPem []byte) (*rsa.PublicKey, error) {
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
		return nil, fmt.Errorf("error converting public key to RSA public key")
	}

	return rsaPubKey, nil
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

// HashKey computes the Poseidon hash of 5 elements.
func HashKey(x509Key []byte) (*big.Int, error) {
	var decomposed [5]*big.Int
	position := len(x509Key)

	for i := 0; i < 5; i++ {
		if position < 24 {
			return nil, fmt.Errorf("x509Key is too short")
		}

		element := new(big.Int).SetBytes(x509Key[position-24 : position])
		reversed := big.NewInt(0)

		for j := 0; j < 3; j++ {
			extracted := new(big.Int).Rsh(element, uint(j*64))
			extracted.And(extracted, new(big.Int).SetUint64(0xffffffffffffffff))
			reversed.Lsh(reversed, 64)
			reversed.Or(reversed, extracted)
		}

		decomposed[i] = reversed
		position -= 24
	}

	keyHash, err := poseidon.Hash(decomposed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to compute Poseidon hash: %v", err)
	}

	return keyHash, nil
}

// NormalizeSignature normalizes the signature.
func NormalizeSignature(signature []byte) ([]byte, error) {
	if len(signature) != 64 {
		return signature, nil
	}

	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	lowSMax, ok := new(big.Int).SetString(lowSMaxHex, 16)
	if !ok {
		return nil, fmt.Errorf("error converting lowSMaxHex to big int")
	}

	n, ok := new(big.Int).SetString(nHex, 16)
	if !ok {
		return nil, fmt.Errorf("error converting nHex to big int")
	}

	if s.Cmp(lowSMax) == 1 {
		s = s.Sub(n, s)
	}

	resR := make([]byte, 32)
	resS := make([]byte, 32)

	resultSignature := append(r.FillBytes(resR), s.FillBytes(resS)...)

	return resultSignature, nil
}

func calculateSmartChunkingNumber(bytesNumber int) int {
	if bytesNumber == 2048 {
		return 32
	}

	return 64
}

// CalculateHmacMessage calculates the HMAC message.
func CalculateHmacMessage(nullifierRaw string, country string, anonymousID []byte) ([]byte, error) {
	nullifier, ok := new(big.Int).SetString(nullifierRaw, 0)
	if !ok {
		return nil, fmt.Errorf("error converting nullifier hex to big int")
	}

	countryBytes := []byte(country)

	var result []byte
	result = append(nullifier.Bytes(), countryBytes...)
	result = append(result, anonymousID...)

	return result, nil
}

// CalculateAnonymousID calculates the anonymous ID.
func CalculateAnonymousID(dg1 []byte, eventID string) ([]byte, error) {
	eventIDInt, ok := new(big.Int).SetString(eventID, 0)
	if !ok {
		return nil, fmt.Errorf("error converting event ID hex to big int")
	}

	sha256Hash := sha256.New()
	sha256Hash.Write(dg1)
	sha256Hash.Write(eventIDInt.Bytes())

	return sha256Hash.Sum(nil), nil
}

type algorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters ecParameters
}

type publicKeyInfo struct {
	Algorithm        algorithmIdentifier
	SubjectPublicKey asn1.BitString
}

type ecParameters struct {
	Version *big.Int
	FieldID fieldID
	Curve   curve
}

type fieldID struct {
	FieldType asn1.ObjectIdentifier
	Data      *big.Int
}

type curve struct {
	Placeholder asn1.RawContent
	X           asn1.RawContent
	Y           asn1.RawContent
}
