package identity

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	asn1Crypto "golang.org/x/crypto/cryptobyte/asn1"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/rarimo/ldif-sdk/ldif"
	"github.com/rarimo/rarime-mobile-identity-sdk/brainpool"
	"golang.org/x/crypto/cryptobyte"
)

const smartChunking2BlockSize uint64 = 512
const lowSMaxHex = "54fdabedd0f754de1f3305484ec1c6b9371dfb11ea9310141009a40e8fb729bb"
const nHex = "A9FB57DBA1EEA9BC3E660A909D838D718C397AA3B561A6F7901E0E82974856A7"

type Call3 struct {
	Target       *common.Address `json:"target"`
	AllowFailure bool            `json:"allowFailure"`
	CallData     []byte          `json:"callData"`
}

func CalculateAggregate3Calldata(calls3Raw []byte) ([]byte, error) {
	var call3s []Call3
	err := json.Unmarshal(calls3Raw, &call3s)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling calls3: %v", err)
	}

	callsType, _ := abi.NewType("tuple[]", "struct Multicall3.Call3[]", []abi.ArgumentMarshaling{
		{
			Name: "target",
			Type: "address",
		},
		{
			Name: "allowFailure",
			Type: "bool",
		},
		{
			Name: "callData",
			Type: "bytes",
		},
	})

	args := abi.Arguments{
		{
			Type: callsType,
		},
	}

	selectorBytes, err := hex.DecodeString("82ad56cb")
	if err != nil {
		return nil, fmt.Errorf("error decoding selector: %v", err)
	}

	call3sBytes, err := args.Pack(call3s)
	if err != nil {
		return nil, fmt.Errorf("error packing call3s: %v", err)
	}

	return append(selectorBytes, call3sBytes...), nil
}

// SignPubSignalsWithSecp256k1 signs a public signals using a private key string (hex format) and the secp256k1 curve.
func SignPubSignalsWithSecp256k1(privateKey string, pubSignalsJSON []byte) (string, error) {
	var pubSignals []string
	if err := json.Unmarshal(pubSignalsJSON, &pubSignals); err != nil {
		return "", fmt.Errorf("error decoding  pub  signals: %v", err)
	}

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error decoding private key hex: %v", err)
	}

	var hash = sha256.New()
	for _, pubSignalByte := range pubSignals {
		if len(pubSignalByte) > 1 && pubSignalByte[:2] == "0x" {
			pubSignalBytes, convertErr := hex.DecodeString(pubSignalByte[2:])
			if convertErr != nil {
				return "", fmt.Errorf("error setting pubSignalHex: %v", pubSignalByte)
			}
			hash.Write(pubSignalBytes)
		} else {
			pubSignalDecimal, ok := new(big.Int).SetString(pubSignalByte, 10)
			if !ok {
				return "", fmt.Errorf("error setting pubSignal: %v", pubSignalByte)
			}
			hash.Write(pubSignalDecimal.Bytes())
		}
	}
	messageHash := hash.Sum(nil)

	signature, err := secp256k1.Sign(messageHash, privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("error signing the message: %v", err)
	}

	signatureHex := hex.EncodeToString(signature)

	return signatureHex, nil
}

// SignMessageWithSecp256k1 signs a string message using a private key string (hex format) and the secp256k1 curve.
func SignMessageWithSecp256k1(privateKey string, message string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error decoding private key hex: %v", err)
	}

	hash := sha256.New()
	hash.Write([]byte(message))
	messageHash := hash.Sum(nil)

	signature, err := secp256k1.Sign(messageHash, privateKeyBytes)
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

// ParsePemToPubKey parses a public key PEM to a public key.
func ParsePemToPubKey(pubKeyPem []byte) (interface{}, error) {
	block, _ := pem.Decode(pubKeyPem)
	if block == nil {
		return nil, fmt.Errorf("error decoding public key pem")
	}

	brainpoolPubKey, _ := brainpool.GetPublicKeyFromPem(block)
	if brainpoolPubKey != nil {
		return brainpoolPubKey, nil
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	return pubKey, nil
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

func Hash512P512(key []byte) (*big.Int, error) {
	if len(key) != 128 {
		return nil, fmt.Errorf("key is not 128 bytes long, got %d", len(key))
	}

	modulus := new(big.Int).Exp(big.NewInt(2), big.NewInt(248), nil)

	X := new(big.Int).SetBytes(key[0:64])
	Y := new(big.Int).SetBytes(key[64:128])

	lowerX := new(big.Int).Mod(X, modulus)

	upperX := new(big.Int).Rsh(X, 256)
	upperX.Mod(upperX, modulus)

	lowerY := new(big.Int).Mod(Y, modulus)
	upperY := new(big.Int).Rsh(Y, 256)
	upperY.Mod(upperY, modulus)

	decomposed := []*big.Int{lowerX, upperX, lowerY, upperY}

	keyHash, err := poseidon.Hash(decomposed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to compute Poseidon hash: %v", err)
	}

	return keyHash, nil
}

// Hash512 applies poseidon2 to [32, 32] bytes long integers mod 2 ** 248
func Hash512(key []byte) (*big.Int, error) {
	if len(key) != 64 {
		return nil, fmt.Errorf("key is not 64 bytes long")
	}

	var decomposed [2]*big.Int
	for i := 0; i < 2; i++ {
		element := new(big.Int).SetBytes(key[i*32 : (i+1)*32])
		decomposed[i] = new(big.Int).Mod(element, new(big.Int).Exp(big.NewInt(2), big.NewInt(248), nil))
	}

	keyHash, err := poseidon.Hash(decomposed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to compute Poseidon hash: %v", err)
	}

	return keyHash, nil
}

// Hash1024 applies poseidon5 to [25, 25, 25, 25, 28] bytes long integers
func Hash1024(key []byte) (*big.Int, error) {
	if len(key) != 128 {
		return nil, fmt.Errorf("key is not 128 bytes long")
	}

	var decomposed [5]*big.Int
	position := len(key)

	for i := 0; i < 5; i++ {
		if position < 25 {
			return nil, fmt.Errorf("key is too short")
		}

		element := new(big.Int).SetBytes(key[position-25 : position])
		decomposed[i] = element
		position -= 25
	}

	keyHash, err := poseidon.Hash(decomposed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to compute Poseidon hash: %v", err)
	}

	return keyHash, nil
}

// HashPacked computes the Poseidon hash of 5 elements.
func HashPacked(x509Key []byte) (*big.Int, error) {
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

// NormalizeSignatureWithCurve normalizes the signature with a curve.
func NormalizeSignatureWithCurve(signature []byte, curve elliptic.Curve) ([]byte, error) {
	pointSize := len(signature) / 2

	r := new(big.Int).SetBytes(signature[:pointSize])
	s := new(big.Int).SetBytes(signature[pointSize:])

	n := curve.Params().N
	lowSMax := getLowSMax(curve)

	if s.Cmp(lowSMax) == 1 {
		s = s.Sub(n, s)
	}

	resR := make([]byte, pointSize)
	resS := make([]byte, pointSize)

	resultSignature := append(r.FillBytes(resR), s.FillBytes(resS)...)

	return resultSignature, nil
}

func getLowSMax(curve elliptic.Curve) *big.Int {
	n := curve.Params().N
	lowSMax := new(big.Int).Rsh(n, 1) // lowSMax = N / 2

	return lowSMax
}

func calculateSmartChunkingNumber(bytesNumber int) int {
	if bytesNumber == 2048 {
		return 32
	}

	return 64
}

func parseECDSASignature(sig []byte) (r, s []byte, err error) {
	var inner cryptobyte.String
	input := cryptobyte.String(sig)
	if !input.ReadASN1(&inner, asn1Crypto.SEQUENCE) ||
		!input.Empty() ||
		!inner.ReadASN1Integer(&r) ||
		!inner.ReadASN1Integer(&s) ||
		!inner.Empty() {
		return nil, nil, errors.New("invalid ASN.1")
	}

	maxLen := len(r)
	if len(s) > maxLen {
		maxLen = len(s)
	}

	resR := make([]byte, maxLen)
	resS := make([]byte, maxLen)

	copy(resR[maxLen-len(r):], r)
	copy(resS[maxLen-len(s):], s)

	return resR, resS, nil
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

// RSAPublicDecrypt decrypts an RSA signature
func RSAPublicDecrypt(pub *rsa.PublicKey, sig []byte) ([]byte, error) {
	k := pub.Size()
	if k < 11 || k != len(sig) {
		return nil, fmt.Errorf("invalid signature length")
	}

	c := new(big.Int).SetBytes(sig)
	e := big.NewInt(int64(pub.E))
	m := new(big.Int).Exp(c, e, pub.N)

	return m.Bytes(), nil
}

func prepareZKProofForEVMVerification(proofJSON []byte) (*ZkProof, *VerifierHelperProofPoints, error) {
	zkProof := new(ZkProof)
	if err := json.Unmarshal(proofJSON, zkProof); err != nil {
		return nil, nil, err
	}

	var a [2]*big.Int
	for index, val := range zkProof.Proof.A[:2] {
		aI, ok := new(big.Int).SetString(val, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error setting a[%d]: %v", index, val)
		}

		a[index] = aI
	}

	var b [2][2]*big.Int
	for index, val := range zkProof.Proof.B[:2] {
		for index2, val2 := range val[:2] {
			bI, ok := new(big.Int).SetString(val2, 10)
			if !ok {
				return nil, nil, fmt.Errorf("error setting b[%d][%d]: %v", index, index2, val2)
			}

			b[index][index2] = bI
		}
	}

	b[0][0], b[0][1] = b[0][1], b[0][0]
	b[1][0], b[1][1] = b[1][1], b[1][0]

	var c [2]*big.Int
	for index, val := range zkProof.Proof.C[:2] {
		cI, ok := new(big.Int).SetString(val, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error setting c[%d]: %v", index, val)
		}

		c[index] = cI
	}

	proofPoints := &VerifierHelperProofPoints{
		A: a,
		B: b,
		C: c,
	}

	return zkProof, proofPoints, nil
}

// GetCurveNameFromECDSAPublicKeyPEM gets the curve name from an ECDSA public key PEM.
func GetCurveNameFromECDSAPublicKeyPEM(pubKeyPem []byte) (string, error) {
	publicKey, err := ParsePemToPubKey(pubKeyPem)
	if err != nil {
		return "", fmt.Errorf("error parsing public key: %v", err)
	}

	switch pub := publicKey.(type) {
	case *ecdsa.PublicKey:
		return pub.Curve.Params().Name, nil
	default:
		return "", fmt.Errorf("unsupported public key type: %T", pub)
	}
}

// IsBrainpoolPublicKey checks if a public key is a Brainpool public key.
func IsBrainpoolPublicKey(pubKeyPem []byte) bool {
	publicKey, err := ParsePemToPubKey(pubKeyPem)
	if err != nil {
		return false
	}

	switch pub := publicKey.(type) {
	case *ecdsa.PublicKey:
		return brainpool.IsBrainpool(pub)
	default:
		return false
	}
}
