package identity

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// RegisterIdentityExp represents RSA exponent.
var RegisterIdentityExp = []string{"65537", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// IcaoMerkleInclusionBranches represents the ICAO merkle inclusion branches.
var IcaoMerkleInclusionBranches = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// IcaoMerkleInclusionOrder represents the ICAO merkle inclusion order.
var IcaoMerkleInclusionOrder = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// Profile represents a user profile.
type Profile struct {
	secretKey   *babyjub.PrivateKey
	publicKey   *babyjub.PublicKey
	dataBlinder *big.Int
}

// NewProfile creates a new profile.
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

// GetRegistrationChallenge returns the registration challenge.
func (p *Profile) GetRegistrationChallenge() ([]byte, error) {
	publicKeyHash, err := poseidon.Hash([]*big.Int{p.publicKey.X, p.publicKey.Y})
	if err != nil {
		return nil, fmt.Errorf("error hashing public key: %v", err)
	}

	return publicKeyHash.Bytes()[:8], nil
}

// BuildRegisterIdentityInputs builds the inputs for the registerIdentity circuit.
func BuildRegisterIdentityInputs(
	secretKeyHex string,
	encapsulatedContent []byte,
	signedAttributes []byte,
	dg1 []byte,
	dg15 []byte,
	pubKeyPem []byte,
	signature []byte,
) ([]byte, error) {
	rawSecretKey, err := hex.DecodeString(secretKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error decoding secret key: %v", err)
	}

	secretKey := babyjub.PrivateKey(rawSecretKey)

	secretKeyInt := new(big.Int).SetBytes(secretKey[:])

	pemBlock, _ := pem.Decode(pubKeyPem)
	if pemBlock == nil {
		return nil, fmt.Errorf("error decoding public key pem")
	}

	pubKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not an RSA public key")
	}

	signatureInt := new(big.Int).SetBytes(signature)

	inputs := &RegisterIdentityInputs{
		SkIdentity:                  secretKeyInt.String(),
		EncapsulatedContent:         ByteArrayToBits(encapsulatedContent),
		SignedAttributes:            ByteArrayToBits(signedAttributes),
		Sign:                        SmartChunking(signatureInt),
		Exp:                         RegisterIdentityExp,
		Modulus:                     SmartChunking(rsaPubKey.N),
		Dg1:                         ByteArrayToBits(dg1),
		Dg15:                        ByteArrayToBits(dg15),
		IcaoMerkleRoot:              "0",
		IcaoMerkleInclusionBranches: IcaoMerkleInclusionBranches,
		IcaoMerkleInclusionOrder:    IcaoMerkleInclusionOrder,
	}

	return inputs.Marshal()
}
