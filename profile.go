package identity

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"cosmossdk.io/errors"
	"github.com/decred/dcrd/bech32"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/rarimo/zkp-iden3-exposer/client"
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// AddressPrefix represents the cosmos address prefix.
const AddressPrefix = "rarimo"

// EventID represents the event ID.
var EventID, _ = new(big.Int).SetString("ac42d1a986804618c7a793fbe814d9b31e47be51e082806363dca6958f3062", 16)

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
	secretKey *babyjub.PrivKeyScalar
	publicKey *babyjub.PublicKey
	wallet    *wallet.Wallet
}

// NewProfile creates a new profile.
func (p *Profile) NewProfile(secretKey []byte) (*Profile, error) {
	secretKeyInt := new(big.Int).SetBytes(secretKey)

	secretKeyScalar := babyjub.NewPrivKeyScalar(secretKeyInt)

	wallet, err := wallet.NewWallet(hex.EncodeToString(secretKeyInt.Bytes()), AddressPrefix)
	if err != nil {
		return nil, fmt.Errorf("error creating wallet: %v", err)
	}

	return &Profile{
		secretKey: secretKeyScalar,
		publicKey: secretKeyScalar.Public(),
		wallet:    wallet,
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

// GetRarimoAddress returns the Rarimo cosmos address
func (p *Profile) GetRarimoAddress() string {
	return p.wallet.Address
}

// BuildRegisterIdentityInputs builds the inputs for the registerIdentity circuit.
func (p *Profile) BuildRegisterIdentityInputs(
	encapsulatedContent []byte,
	signedAttributes []byte,
	dg1 []byte,
	dg15 []byte,
	pubKeyPem []byte,
	signature []byte,
	isEcdsaActiveAuthentication bool,
	certificatesSMTProofJSON []byte,
) ([]byte, error) {
	var certificatesSMTProof SMTProof
	if err := json.Unmarshal(certificatesSMTProofJSON, &certificatesSMTProof); err != nil {
		return nil, fmt.Errorf("error unmarshalling certificates SMT proof: %v", err)
	}

	rsaPubKeyN, err := RsaPubKeyPemToN(pubKeyPem)
	if err != nil {
		return nil, fmt.Errorf("error parsing rsa public key: %v", err)
	}

	signatureInt := new(big.Int).SetBytes(signature)

	signedAttributesSize := len(signedAttributes) * 8

	var SaTimestampEnabled string
	if signedAttributesSize == 832 {
		SaTimestampEnabled = "1"
	} else {
		SaTimestampEnabled = "0"
	}

	var EcdsaShiftEnabled string
	if isEcdsaActiveAuthentication {
		EcdsaShiftEnabled = "1"
	} else {
		EcdsaShiftEnabled = "0"
	}

	slaveMerleRoot := new(big.Int).SetBytes(certificatesSMTProof.Root).String()

	var slaveMerkleInclusionBranches []string
	for _, sibling := range certificatesSMTProof.Siblings {
		slaveMerkleInclusionBranches = append(slaveMerkleInclusionBranches, new(big.Int).SetBytes(sibling).String())
	}

	inputs := &RegisterIdentityInputs{
		SkIdentity:                   p.secretKey.BigInt().String(),
		EncapsulatedContent:          SmartChunking2(ByteArrayToBits(encapsulatedContent), 6),
		SignedAttributes:             SmartChunking2(ByteArrayToBits(signedAttributes), 2),
		Sign:                         SmartChunking(signatureInt),
		Modulus:                      SmartChunking(rsaPubKeyN),
		Exp:                          RegisterIdentityExp,
		Dg1:                          SmartChunking2(ByteArrayToBits(dg1), 2),
		Dg15:                         SmartChunking2(ByteArrayToBits(dg15), 6),
		SlaveMerleRoot:               slaveMerleRoot,
		SlaveMerkleInclusionBranches: slaveMerkleInclusionBranches,
		EcdsaShiftEnabled:            EcdsaShiftEnabled,
		SaTimestampEnabled:           SaTimestampEnabled,
	}

	return inputs.Marshal()
}

// BuildQueryIdentityInputs builds the inputs for the queryIdentity circuit.
func (p *Profile) BuildQueryIdentityInputs(
	dg1 []byte,
	smtProofJSON []byte,
	selector string,
	pkPassportHash string,
	issueTimestamp string,
	identityCounter string,
	timestampLowerbound string,
	timestampUpperbound string,
	identityCounterLowerbound string,
	identityCounterUpperbound string,
) ([]byte, error) {
	var smtProof SMTProof
	if err := json.Unmarshal(smtProofJSON, &smtProof); err != nil {
		return nil, fmt.Errorf("error unmarshalling id state siblings: %v", err)
	}

	idStateRoot := new(big.Int).SetBytes(smtProof.Root).String()

	var idStateSiblings []string
	for _, sibling := range smtProof.Siblings {
		idStateSiblings = append(idStateSiblings, new(big.Int).SetBytes(sibling).String())
	}

	_, decodedAddress, err := bech32.Decode(p.wallet.Address)
	if err != nil {
		return nil, fmt.Errorf("error decoding address: %v", err)
	}

	decodedAddressInt := new(big.Int).SetBytes(decodedAddress)

	inputs := &QueryIdentityInputs{
		Dg1:                       ByteArrayToBits(dg1),
		EventID:                   EventID.String(),
		EventData:                 decodedAddressInt.String(),
		IDStateRoot:               idStateRoot,
		IDStateSiblings:           idStateSiblings,
		PkPassportHash:            pkPassportHash,
		Selector:                  selector,
		SkIdentity:                p.secretKey.BigInt().String(),
		Timestamp:                 issueTimestamp,
		IdentityCounter:           identityCounter,
		TimestampLowerbound:       "0",
		TimestampUpperbound:       "0",
		IdentityCounterLowerbound: identityCounterLowerbound,
		IdentityCounterUpperbound: identityCounterUpperbound,
		BirthDateLowerbound:       "0x303030303030",
		BirthDateUpperbound:       "0x303030303030",
		ExpirationDateLowerbound:  "0x303030303030",
		ExpirationDateUpperbound:  "0x303030303030",
		CitizenshipMask:           "0x303030303030",
	}

	json, err := inputs.Marshal()
	if err != nil {
		return nil, fmt.Errorf("error marshalling inputs: %v", err)
	}

	return json, nil
}

// WalletSend sends tokens to desrination via Cosmos
func (p *Profile) WalletSend(toAddr string, amount string, chainID string, denom string, rpcIP string) ([]byte, error) {
	chainConfig := client.ChainConfig{
		ChainId:     chainID,
		Denom:       denom,
		MinGasPrice: 0,
		GasLimit:    uint64(1_000_000),
	}

	grpcClient, err := grpc.Dial(
		rpcIP,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second, // wait time before ping if no activity
			Timeout: 20 * time.Second, // ping timeout
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error dialing grpc")
	}

	rarimoClient, err := client.NewClient(
		grpcClient,
		chainConfig,
		*p.wallet,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating client")
	}

	sendAmount, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse amount: ")
	}

	txResp, err := rarimoClient.Send(
		p.wallet.Address,
		toAddr,
		sendAmount,
		denom,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error sending tx")
	}

	return txResp, nil
}
