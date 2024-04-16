package identity

import (
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// RegisterIdentityExp represents RSA exponent.
var RegisterIdentityExp = []string{"65537", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// IcaoMerkleRoot represents the ICAO merkle root.
var IcaoMerkleRoot, _ = new(big.Int).SetString("2c50ce3aa92bc3dd0351a89970b02630415547ea83c487befbc8b1795ea90c45", 16)

// IcaoMerkleInclusionBranches represents the ICAO merkle inclusion branches.
var IcaoMerkleInclusionBranches = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// IcaoMerkleInclusionOrder represents the ICAO merkle inclusion order.
var IcaoMerkleInclusionOrder = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

// Profile represents a user profile.
type Profile struct {
	secretKey   *babyjub.PrivKeyScalar
	publicKey   *babyjub.PublicKey
	dataBlinder *big.Int
}

// NewProfile creates a new profile.
func NewProfile(secretKey []byte) (*Profile, error) {
	secretKeyInt := new(big.Int).SetBytes(secretKey)

	dataBlinder, err := poseidon.Hash([]*big.Int{secretKeyInt})
	if err != nil {
		return nil, fmt.Errorf("error hashing secret key: %v", err)
	}

	secretKeyScalar := babyjub.NewPrivKeyScalar(secretKeyInt)

	return &Profile{
		secretKey:   secretKeyScalar,
		publicKey:   secretKeyScalar.Public(),
		dataBlinder: dataBlinder,
	}, nil
}

// GetRegistrationChallenge returns the registration challenge.
func (p *Profile) GetRegistrationChallenge() ([]byte, error) {
	publicKeyHash, err := p.GetPublicKeyHash()
	if err != nil {
		return nil, fmt.Errorf("error getting public key hash: %v", err)
	}

	return publicKeyHash[len(publicKeyHash)-8:], nil
}

// GetPublicKeyHash returns the public key hash.
func (p *Profile) GetPublicKeyHash() ([]byte, error) {
	publicKeyHash, err := poseidon.Hash([]*big.Int{p.publicKey.X, p.publicKey.Y})
	if err != nil {
		return nil, fmt.Errorf("error hashing public key: %v", err)
	}

	return publicKeyHash.Bytes(), nil
}

// BuildRegisterIdentityInputs builds the inputs for the registerIdentity circuit.
func BuildRegisterIdentityInputs(
	secretKey []byte,
	encapsulatedContent []byte,
	signedAttributes []byte,
	dg1 []byte,
	dg15 []byte,
	pubKeyPem []byte,
	signature []byte,
) ([]byte, error) {
	secretKeyInt := new(big.Int).SetBytes(secretKey[:])

	rsaPubKeyN, err := RsaPubKeyPemToN(pubKeyPem)
	if err != nil {
		return nil, fmt.Errorf("error parsing rsa public key: %v", err)
	}

	signatureInt := new(big.Int).SetBytes(signature)

	inputs := &RegisterIdentityInputs{
		SkIdentity:                  secretKeyInt.String(),
		EncapsulatedContent:         ByteArrayToBits(encapsulatedContent),
		SignedAttributes:            ByteArrayToBits(signedAttributes),
		Sign:                        SmartChunking(signatureInt),
		Exp:                         RegisterIdentityExp,
		Modulus:                     SmartChunking(rsaPubKeyN),
		Dg1:                         ByteArrayToBits(dg1),
		Dg15:                        ByteArrayToBits(dg15),
		IcaoMerkleRoot:              IcaoMerkleRoot.String(),
		IcaoMerkleInclusionBranches: IcaoMerkleInclusionBranches,
		IcaoMerkleInclusionOrder:    IcaoMerkleInclusionOrder,
	}

	return inputs.Marshal()
}
