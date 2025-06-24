package identity

import (
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
	"time"

	"github.com/decred/dcrd/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/rarimo/zkp-iden3-exposer/client"
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// AddressPrefix represents the cosmos address prefix.
const AddressPrefix = "rarimo"

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

type PollResult struct {
	QuestionIndex int `json:"questionIndex"`
	AnswerIndex   int `json:"answerIndex,omitempty"`
}

type ProposalRules struct {
	CitizenshipWhitelist                []*big.Int `json:"citizenshipWhitelist"`
	IdentityCreationTimestampUpperBound *big.Int   `json:"identityCreationTimestampUpperBound"`
	IdentityCounterUpperBound           *big.Int   `json:"identityCounterUpperBound"`
	BirthDateUpperbound                 *big.Int   `json:"birthDateUpperbound"`
	ExpirationDateLowerBound            *big.Int   `json:"expirationDateLowerBound"`
}

func makeByteWithBits(n int) byte {
	if n < 0 || n > 8 {
		panic("n must be between 0 and 8")
	}

	return byte((1 << n))
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

// BuildAirdropQueryIdentityInputs builds the inputs for the queryIdentity circuit.
func (p *Profile) BuildAirdropQueryIdentityInputs(
	dg1 []byte,
	smtProofJSON []byte,
	selector string,
	pkPassportHash string,
	issueTimestamp string,
	identityCounter string,
	eventID string,
	startedAt int64,
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

	issueTimestampInt, err := strconv.ParseInt(issueTimestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing issue timestamp: %v", err)
	}

	var timestampLowerbound int64
	var timestampUpperbound int64
	if issueTimestampInt > startedAt {
		timestampLowerbound = 0
		timestampUpperbound = issueTimestampInt + 1
	} else {
		timestampLowerbound = issueTimestampInt
		timestampUpperbound = startedAt
	}

	identityCounterInt, err := strconv.ParseInt(identityCounter, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing identity counter: %v", err)
	}

	identityCounterUpperbound := identityCounterInt + 1

	currentDate := time.Now().UTC()
	currentDateHex := "0x" + hex.EncodeToString([]byte(currentDate.Format("060102")))

	birthDateLowerbound := "0x" + hex.EncodeToString([]byte(currentDate.AddDate(-100, 0, 0).Format("060102")))
	birthDateUpperbound := "0x" + hex.EncodeToString([]byte(currentDate.AddDate(-18, 0, 0).Format("060102")))

	inputs := &QueryIdentityInputs{
		Dg1:                       ByteArrayToBits(dg1),
		EventID:                   eventID,
		EventData:                 decodedAddressInt.String(),
		IDStateRoot:               idStateRoot,
		IDStateSiblings:           idStateSiblings,
		PkPassportHash:            pkPassportHash,
		Selector:                  selector,
		SkIdentity:                p.secretKey.BigInt().String(),
		Timestamp:                 issueTimestamp,
		CurrentDate:               currentDateHex,
		IdentityCounter:           identityCounter,
		TimestampLowerbound:       strconv.FormatInt(timestampLowerbound, 10),
		TimestampUpperbound:       strconv.FormatInt(timestampUpperbound, 10),
		IdentityCounterLowerbound: "0",
		IdentityCounterUpperbound: strconv.FormatInt(identityCounterUpperbound, 10),
		ExpirationDateLowerbound:  currentDateHex,
		ExpirationDateUpperbound:  "0x303030303030",
		BirthDateLowerbound:       birthDateLowerbound,
		BirthDateUpperbound:       birthDateUpperbound,
		CitizenshipMask:           "0x303030303030",
	}

	json, err := inputs.Marshal()
	if err != nil {
		return nil, fmt.Errorf("error marshalling inputs: %v", err)
	}

	return json, nil
}

// TransformVote transforms the vote.
func TransformVote(vote []PollResult) ([]byte, error) {
	uint256Type, _ := abi.NewType("uint256[]", "", nil)

	args := abi.Arguments{
		{
			Type: uint256Type,
		},
	}

	var values []*big.Int
	for _, v := range vote {
		value := makeByteWithBits(v.AnswerIndex)

		values = append(values, big.NewInt(int64(value)))
	}

	packedVote, err := args.Pack(values)
	if err != nil {
		return nil, fmt.Errorf("failed to pack vote: %v", err)
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(packedVote)
	hashBytes := hash.Sum(nil)
	uint248Bytes := hashBytes[1:32]

	return uint248Bytes, nil
}

// BuildQueryIdentityInputs builds the inputs for the queryIdentity circuit
func (p *Profile) BuildQueryIdentityInputs(
	dg1 []byte,
	smtProofJSON []byte,
	selector string,
	pkPassportHash string,
	issueTimestamp string,
	identityCounter string,
	eventID string,
	eventData string,
	TimestampLowerbound string,
	TimestampUpperbound string,
	IdentityCounterLowerbound string,
	IdentityCounterUpperbound string,
	ExpirationDateLowerbound string,
	ExpirationDateUpperbound string,
	BirthDateLowerbound string,
	BirthDateUpperbound string,
	CitizenshipMask string,
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

	currentDate := time.Now().UTC()
	currentDateHex := "0x" + hex.EncodeToString([]byte(currentDate.Format("060102")))

	inputs := &QueryIdentityInputs{
		Dg1:                       ByteArrayToBits(dg1),
		EventID:                   eventID,
		EventData:                 eventData,
		IDStateRoot:               idStateRoot,
		IDStateSiblings:           idStateSiblings,
		PkPassportHash:            pkPassportHash,
		Selector:                  selector,
		SkIdentity:                p.secretKey.BigInt().String(),
		Timestamp:                 issueTimestamp,
		CurrentDate:               currentDateHex,
		IdentityCounter:           identityCounter,
		TimestampLowerbound:       TimestampLowerbound,
		TimestampUpperbound:       TimestampUpperbound,
		IdentityCounterLowerbound: IdentityCounterLowerbound,
		IdentityCounterUpperbound: IdentityCounterUpperbound,
		ExpirationDateLowerbound:  ExpirationDateLowerbound,
		ExpirationDateUpperbound:  ExpirationDateUpperbound,
		BirthDateLowerbound:       BirthDateLowerbound,
		BirthDateUpperbound:       BirthDateUpperbound,
		CitizenshipMask:           CitizenshipMask,
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
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS13,
		})),
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

// calculateEventNullifier calculates the event nullifier.
func (p *Profile) calculateEventNullifier(eventID string) (*big.Int, error) {
	secretKey := p.secretKey.BigInt()

	secretKeyHash, err := poseidon.Hash([]*big.Int{secretKey})
	if err != nil {
		return nil, fmt.Errorf("error hashing secret key: %v", err)
	}

	eventIDInt, ok := new(big.Int).SetString(eventID, 0)
	if !ok {
		return nil, fmt.Errorf("error parsing event ID: %v", err)
	}

	eventNullifier, err := poseidon.Hash([]*big.Int{secretKey, secretKeyHash, eventIDInt})
	if err != nil {
		return nil, fmt.Errorf("error hashing event: %v", err)
	}

	return eventNullifier, nil
}

// CalculateEventNullifierHex calculates the event nullifier in hex.
func (p *Profile) CalculateEventNullifierHex(eventID string) (string, error) {
	eventNullifier, err := p.calculateEventNullifier(eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	eventNullifierHex := hex.EncodeToString(eventNullifier.Bytes())
	if len(eventNullifierHex) < 64 {
		eventNullifierHex = fmt.Sprintf("%0*s", 64, eventNullifierHex)
	}

	return fmt.Sprintf("0x%s", eventNullifierHex), nil
}

// CalculateEventNullifierInt calculates the event nullifier in hex.
func (p *Profile) CalculateEventNullifierInt(eventID string) (string, error) {
	eventNullifier, err := p.calculateEventNullifier(eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	return eventNullifier.String(), nil
}

// CalculateVotingEventData calculates the voting event data.
func (p *Profile) CalculateVotingEventData(voteJson []byte) ([]byte, error) {
	var pollResults []PollResult
	if err := json.Unmarshal(voteJson, &pollResults); err != nil {
		return nil, fmt.Errorf("error unmarshalling vote: %v", err)
	}

	return TransformVote(pollResults)
}
