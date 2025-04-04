package identity

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/rarimo/certificate-transparency-go/x509"
	"github.com/rarimo/ldif-sdk/mt"
)

// ECMaxSizeInBits represents the maximum size in bits for an encapsulated content
const ECMaxSizeInBits = 2688

// ZKTypePrefix represerts the circuit zk type prefix
const ZKTypePrefix = "Z_PER_PASSPORT"

// ZKNoirTypePrefix represents the circuit zk type prefix length
const ZKNoirTypePrefix = "Z_NOIR_PASSPORT"

// RegistrationMetaData contains all metadata for the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCallVerifyProof\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"}],\"name\":\"InvalidCircomProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubSignals\",\"type\":\"bytes32[]\"}],\"name\":\"InvalidNoirProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P_NO_AA\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateKeeper_\",\"type\":\"address\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certificateDispatchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportDispatchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signedAttributes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"keyOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationOffset\",\"type\":\"uint256\"}],\"internalType\":\"structRegistration2.Certificate\",\"name\":\"certificate_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration2.ICAOMember\",\"name\":\"icaoMember_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"icaoMerkleProof_\",\"type\":\"bytes32[]\"}],\"name\":\"registerCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"registerViaNoir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"reissueIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"reissueIdentityViaNoir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificateKey_\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateKeeper\",\"outputs\":[{\"internalType\":\"contractStateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRegistration2.MethodId\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"updateDependency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistrationSimplePassport is an auto generated low-level Go binding around an user-defined struct.
type RegistrationSimplePassport struct {
	DgCommit     *big.Int
	Dg1Hash      [32]byte
	PublicKey    [32]byte
	PassportHash [32]byte
	Verifier     common.Address
}

// RegistrationSimpleMetaData contains all meta data concerning the RegistrationSimple contract.
var RegistrationSimpleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newOwners\",\"type\":\"address[]\"}],\"name\":\"OwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"removedOwners\",\"type\":\"address[]\"}],\"name\":\"OwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRATION_SIMPLE_PREFIX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssSigner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"stateKeeper_\",\"type\":\"address\"}],\"name\":\"__RegistrationSimple_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOwners_\",\"type\":\"address[]\"}],\"name\":\"addOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dgCommit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dg1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"internalType\":\"structRegistrationSimple.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"registerSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldOwners_\",\"type\":\"address[]\"}],\"name\":\"removeOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateKeeper\",\"outputs\":[{\"internalType\":\"contractStateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationPassport represents a registration passport.
type RegistrationPassport struct {
	DataType     [32]byte
	ZkType       [32]byte
	Signature    []byte
	PublicKey    []byte
	PassportHash [32]byte
}

// RegistrationPassportData represents a registration passport data.
type RegistrationPassportData struct {
	AADataType  []byte
	AASignature []byte
	AAPublicKey []byte
}

// RegistrationCertificate is an auto generated low-level Go binding around an user-defined struct.
type RegistrationCertificate struct {
	DataType         [32]byte
	SignedAttributes []byte
	KeyOffset        *big.Int
	ExpirationOffset *big.Int
}

// RegistrationICAOMember is an auto generated low-level Go binding around an user-defined struct.
type RegistrationICAOMember struct {
	Signature []byte
	PublicKey []byte
}

// VerifierHelperProofPoints represents the proof points for the VerifierHelper contract.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// Proof represents a proof.
type Proof struct {
	A []string   `json:"pi_a"`
	B [][]string `json:"pi_b"`
	C []string   `json:"pi_c"`
}

// PubSignals represents the public signals.
type PubSignals []string

// ZkProof represents a zk proof.
type ZkProof struct {
	Proof      Proof      `json:"proof"`
	PubSignals PubSignals `json:"pub_signals"`
}

// RegisterCalldataResult represents the result of the register calldata.
type RegisterCalldataResult struct {
	Calldata       []byte
	DispatcherName string
}

func newRegistrationCoder() (*abi.ABI, error) {
	parsed, err := RegistrationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func newRegistrationSimpleCoder() (*abi.ABI, error) {
	parsed, err := RegistrationSimpleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

// CallDataBuilder builds the calldata for the register function.
type CallDataBuilder struct{}

// BuildRegisterCalldata builds the calldata for the register function.
func (s *CallDataBuilder) BuildRegisterCalldata(
	proofJSON []byte,
	aaSignature []byte,
	aaPubKeyPem []byte,
	ecSizeInBits int,
	certificatesRootRaw []byte,
	isRevoked bool,
	circuitName string,
) ([]byte, error) {
	zkProof, proofPoints, err := prepareZKProofForEVMVerification(proofJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare zk proof for evm verification: %v", err)
	}

	passportKey, ok := new(big.Int).SetString(zkProof.PubSignals[0], 10)
	if !ok {
		return nil, fmt.Errorf("error setting passportKey: %v", zkProof.PubSignals[0])
	}

	passportHash, ok := new(big.Int).SetString(zkProof.PubSignals[1], 10)
	if !ok {
		return nil, fmt.Errorf("error setting passportHash: %v", zkProof.PubSignals[1])
	}

	dg1Commitment, ok := new(big.Int).SetString(zkProof.PubSignals[2], 10)
	if !ok {
		return nil, fmt.Errorf("error setting dg1Commitment: %v", zkProof.PubSignals[2])
	}

	pkIdentityHash, ok := new(big.Int).SetString(zkProof.PubSignals[3], 10)
	if !ok {
		return nil, fmt.Errorf("error setting pkIdentityHash: %v", zkProof.PubSignals[3])
	}

	registrationPassportData, err := retriveRegistrationPassportData(
		aaSignature,
		aaPubKeyPem,
		ecSizeInBits,
		passportKey.Bytes(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive registration passport data: %v", err)
	}

	datatype := [32]byte{}
	copy(datatype[:], registrationPassportData.AADataType)

	_, zkTypeSuffix, wasCircuitNameCut := strings.Cut(circuitName, "_")
	if !wasCircuitNameCut {
		return nil, fmt.Errorf("circuit name is in invalid format")
	}

	var zkTypeName = fmt.Sprintf("%v_%v", ZKTypePrefix, zkTypeSuffix)

	zkTypeBuf := keccak256.Hash([]byte(zkTypeName))

	var zkType [32]byte
	copy(zkType[:], zkTypeBuf)

	passportHashBytes32 := make([]byte, 32)
	passportHash.FillBytes(passportHashBytes32)

	passport := &RegistrationPassport{
		DataType:     datatype,
		ZkType:       zkType,
		PublicKey:    registrationPassportData.AAPublicKey,
		Signature:    registrationPassportData.AASignature,
		PassportHash: [32]byte(passportHashBytes32),
	}

	var certificatesRoot [32]byte
	copy(certificatesRoot[:], certificatesRootRaw)

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	var methodName string
	if isRevoked {
		methodName = "reissueIdentity"
	} else {
		methodName = "register"
	}

	return abi.Pack(methodName, certificatesRoot, pkIdentityHash, dg1Commitment, passport, proofPoints)
}

func (s *CallDataBuilder) BuildNoirRegisterCalldata(
	fullProof []byte,
	aaSignature []byte,
	aaPubKeyPem []byte,
	ecSizeInBits int,
	certificatesRootRaw []byte,
	isRevoked bool,
	circuitName string,
) ([]byte, error) {
	proof, err := newNoirRegistrationProof(fullProof)
	if err != nil {
		return nil, fmt.Errorf("failed to create noir registration proof: %v", err)
	}

	registrationPassportData, err := retriveRegistrationPassportData(
		aaSignature,
		aaPubKeyPem,
		ecSizeInBits,
		proof.passportKey.Bytes(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive registration passport data: %v", err)
	}

	datatype := [32]byte{}
	copy(datatype[:], registrationPassportData.AADataType)

	_, zkTypeSuffix, wasCircuitNameCut := strings.Cut(circuitName, "_")
	if !wasCircuitNameCut {
		return nil, fmt.Errorf("circuit name is in invalid format")
	}

	var zkTypeName = fmt.Sprintf("%v_%v", ZKNoirTypePrefix, zkTypeSuffix)

	zkTypeBuf := keccak256.Hash([]byte(zkTypeName))

	var zkType [32]byte
	copy(zkType[:], zkTypeBuf)

	passportHashBytes32 := make([]byte, 32)
	proof.passportHash.FillBytes(passportHashBytes32)

	passport := &RegistrationPassport{
		DataType:     datatype,
		ZkType:       zkType,
		PublicKey:    registrationPassportData.AAPublicKey,
		Signature:    registrationPassportData.AASignature,
		PassportHash: [32]byte(passportHashBytes32),
	}

	var certificatesRoot [32]byte
	copy(certificatesRoot[:], certificatesRootRaw)

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	var methodName string
	if isRevoked {
		methodName = "reissueIdentityViaNoir"
	} else {
		methodName = "registerViaNoir"
	}

	return abi.Pack(methodName, certificatesRoot, proof.identityKey, proof.dgCommit, passport, proof.points)
}

// BuildRegisterCertificateCalldata builds the calldata for the register certificate function.
func (s *CallDataBuilder) BuildRegisterCertificateCalldata(masterCertificatesPem []byte, slavePem []byte) (*RegisterCalldataResult, error) {
	icaoTree, err := mt.BuildTreeFromCollection(masterCertificatesPem)
	if err != nil {
		return nil, fmt.Errorf("failed to build tree from collection: %v", err)
	}

	x := X509Util{}
	slaveCert, masterCert, err := x.GetMaster(slavePem, masterCertificatesPem)
	if err != nil {
		return nil, fmt.Errorf("failed to get master: %v", err)
	}

	masterCertPem, err := x.CertificateToPem(masterCert)
	if err != nil {
		return nil, fmt.Errorf("failed to convert certificate to pem: %v", err)
	}

	icaoMerkleProof, err := icaoTree.GenerateInclusionProof(string(masterCertPem))
	if err != nil {
		return nil, fmt.Errorf("failed to generate inclusion proof: %v", err)
	}

	if len(icaoMerkleProof.Siblings) == 0 {
		return nil, fmt.Errorf("failed to generate inclusion proof: no siblings")
	}

	var icaoMemberSignature []byte
	var icaoMemberKey []byte

	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		icaoMemberKey = pub.N.Bytes()

		icaoMemberSignature = slaveCert.Signature
	case *ecdsa.PublicKey:
		icaoMemberKey = pub.X.Bytes()

		icaoMemberKey = append(icaoMemberKey, pub.Y.Bytes()...)

		slaveCertSignatureR, slaveCertSignaturS, err := parseECDSASignature(slaveCert.Signature)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ECDSA signature: %v", err)
		}

		icaoMemberSignature = append(slaveCertSignatureR, slaveCertSignaturS...)
		icaoMemberSignature, err = NormalizeSignatureWithCurve(icaoMemberSignature, pub.Curve)
		if err != nil {
			return nil, fmt.Errorf("failed to normalize signature with curve: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	x509SignedAttributes := slaveCert.RawTBSCertificate

	x509KeyOffset, err := x.FindKeyPositionInSignedAttributes(slaveCert)
	if err != nil {
		return nil, fmt.Errorf("failed to find key position in signed attributes: %v", err)
	}

	x509ExpirationOffset, err := x.FindExpirationPositionInSignedAttributes(slaveCert)
	if err != nil {
		return nil, fmt.Errorf("failed to find expiration position in signed attributes: %v", err)
	}

	dispatcher, dispatcherName, err := retriveCertificateRegistrationDispatcher(masterCert, slaveCert)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve certificate registration dispatcher and slave key: %v", err)
	}

	var dataType [32]byte
	copy(dataType[:], dispatcher)

	var icaoMerkleProofSiblings [][32]byte
	for _, sibling := range icaoMerkleProof.Siblings {
		var siblingBuf [32]byte
		copy(siblingBuf[:], sibling)
		icaoMerkleProofSiblings = append(icaoMerkleProofSiblings, siblingBuf)
	}

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	registrationCertificate := RegistrationCertificate{
		DataType:         dataType,
		SignedAttributes: x509SignedAttributes,
		KeyOffset:        x509KeyOffset,
		ExpirationOffset: x509ExpirationOffset,
	}

	icaoMember := RegistrationICAOMember{
		Signature: icaoMemberSignature,
		PublicKey: icaoMemberKey,
	}

	calldata, err := abi.Pack("registerCertificate", registrationCertificate, icaoMember, icaoMerkleProofSiblings)
	if err != nil {
		return nil, fmt.Errorf("failed to pack register certificate: %v", err)
	}

	return &RegisterCalldataResult{
		Calldata:       calldata,
		DispatcherName: dispatcherName,
	}, nil
}

// BuildRegisterSimpleCalldata builds the calldata for the registerSimple function.
//
// RegisterSimple(identityKey_ *big.Int, passport_ RegistrationSimplePassport, signature_ []byte, zkPoints_ VerifierHelperProofPoints)
func (s *CallDataBuilder) BuildRegisterSimpleCalldata(
	registerIdentityLightProofJSON []byte,
	signature []byte,
	passportHash []byte,
	publicKey []byte,
	verifierAddress string,
) ([]byte, error) {
	zkProof, proofPoints, err := prepareZKProofForEVMVerification(registerIdentityLightProofJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare zk proof for evm verification: %v", err)
	}

	dg1Hash, ok := new(big.Int).SetString(zkProof.PubSignals[0], 10)
	if !ok {
		return nil, fmt.Errorf("error setting dg1Hash: %v", zkProof.PubSignals[0])
	}

	dg1Commit, ok := new(big.Int).SetString(zkProof.PubSignals[1], 10)
	if !ok {
		return nil, fmt.Errorf("error setting dg1Commit: %v", zkProof.PubSignals[1])
	}

	identityKey, ok := new(big.Int).SetString(zkProof.PubSignals[2], 10)
	if !ok {
		return nil, fmt.Errorf("error setting identityKey: %v", zkProof.PubSignals[2])
	}

	var dg1HashBuf [32]byte
	dg1Hash.FillBytes(dg1HashBuf[:])

	var publicKeyBuf [32]byte
	copy(publicKeyBuf[:], publicKey)

	var passportHashBuf [32]byte
	copy(passportHashBuf[:], passportHash)

	passport := &RegistrationSimplePassport{
		DgCommit:     dg1Commit,
		Dg1Hash:      dg1HashBuf,
		PublicKey:    publicKeyBuf,
		PassportHash: passportHashBuf,
		Verifier:     common.HexToAddress(verifierAddress),
	}

	abi, err := newRegistrationSimpleCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("registerSimple", identityKey, passport, signature, proofPoints)
}

// BuildRevoceCalldata builds the calldata for the revoke function.
func (s *CallDataBuilder) BuildRevoceCalldata(
	identityKey []byte,
	aaSignature []byte,
	aaPubKeyPem []byte,
	ecSizeInBits int,
) ([]byte, error) {
	registrationPassportData, err := retriveRegistrationPassportData(
		aaSignature,
		aaPubKeyPem,
		ecSizeInBits,
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive registration passport data: %v", err)
	}

	datatype := [32]byte{}
	copy(datatype[:], registrationPassportData.AADataType)

	identityKeyInt := new(big.Int).SetBytes(identityKey)

	passport := &RegistrationPassport{
		DataType:  datatype,
		PublicKey: registrationPassportData.AAPublicKey,
		Signature: registrationPassportData.AASignature,
	}

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("revoke", identityKeyInt, passport)
}

func retriveRegistrationPassportData(
	aaSignature []byte,
	aaPubKeyPem []byte,
	ecSizeInBits int,
	passportKey []byte,
) (*RegistrationPassportData, error) {
	registrationPassportData := &RegistrationPassportData{}
	if len(aaPubKeyPem) == 0 {
		registrationPassportData.AADataType = keccak256.Hash([]byte("P_NO_AA"))
		registrationPassportData.AAPublicKey = []byte{}
		registrationPassportData.AASignature = []byte{}

		return registrationPassportData, nil
	}

	aaPubKey, err := ParsePemToPubKey(aaPubKeyPem)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pem to pubkey: %v", err)
	}

	switch pub := aaPubKey.(type) {
	case *rsa.PublicKey:
		aaSignatureHashAlgo, err := figureOutRSAAAHashAlgorithm(pub, aaSignature)
		if err != nil {
			return nil, fmt.Errorf("failed to figure out rsa aa hash algorithm: %v", err)
		}

		if aaSignatureHashAlgo == "" {
			registrationPassportData.AADataType = keccak256.Hash([]byte("P_NO_AA"))
			registrationPassportData.AAPublicKey = passportKey

			return registrationPassportData, nil
		}

		if ECMaxSizeInBits > ecSizeInBits {
			ecSizeInBits = ECMaxSizeInBits
		}

		dispatcherName := fmt.Sprintf("P_RSA_%v_%v", aaSignatureHashAlgo, ecSizeInBits)
		if pub.E == 3 {
			dispatcherName += "_3"
		}

		registrationPassportData.AAPublicKey = pub.N.Bytes()
		registrationPassportData.AASignature = aaSignature
		registrationPassportData.AADataType = keccak256.Hash([]byte(dispatcherName))

		return registrationPassportData, nil
	case *ecdsa.PublicKey:
		pubKeyX := make([]byte, pub.Params().BitSize/8)
		pubKeyY := make([]byte, pub.Params().BitSize/8)

		copy(pubKeyX[len(pubKeyX)-len(pub.X.Bytes()):], pub.X.Bytes())
		copy(pubKeyY[len(pubKeyY)-len(pub.Y.Bytes()):], pub.Y.Bytes())

		registrationPassportData.AAPublicKey = append(pubKeyX, pubKeyY...)

		registrationPassportData.AASignature, err = NormalizeSignatureWithCurve(aaSignature, pub.Curve)
		if err != nil {
			return nil, fmt.Errorf("failed to normalize signature with curve: %v", err)
		}

		dispatcherName := fmt.Sprintf("P_ECDSA_SHA1_%v", ecSizeInBits)
		registrationPassportData.AADataType = keccak256.Hash([]byte(dispatcherName))

		return registrationPassportData, nil
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}
}

func figureOutRSAAAHashAlgorithm(aaPubKey *rsa.PublicKey, aaSignature []byte) (string, error) {
	decyptedAASig, err := RSAPublicDecrypt(aaPubKey, aaSignature)
	if err != nil {
		return "", nil
	}

	if len(decyptedAASig) < 2 {
		return "", fmt.Errorf("invalid aa signature")
	}

	// See ISO/IEC9796-2 and ISO/IEC 10118-3 to understand this woodoo
	flagBit := decyptedAASig[len(decyptedAASig)-1]
	if flagBit == 0xcc {
		flagBit = decyptedAASig[len(decyptedAASig)-2]
	}

	switch flagBit {
	case 0xbc, 0x33:
		return "SHA1", nil
	case 0x34:
		return "SHA256", nil
	case 0x35:
		return "SHA512", nil
	case 0x36:
		return "SHA384", nil
	case 0x38:
		return "SHA224", nil
	default:
		return "SHA256", nil
	}
}

func retriveCertificateRegistrationDispatcher(
	masterCert *x509.Certificate,
	slaveCert *x509.Certificate,
) ([]byte, string, error) {
	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		return retriveCertificateRegistrationDispatcherForRSAFamily(masterCert, slaveCert)
	case *ecdsa.PublicKey:
		return retriveCertificateRegistrationDispatcherForECDSAFamily(masterCert, slaveCert)
	default:
		return nil, "", fmt.Errorf("unsupported public key type: %T", pub)
	}
}

func retriveCertificateRegistrationDispatcherForRSAFamily(
	masterCert *x509.Certificate,
	slaveCert *x509.Certificate,
) ([]byte, string, error) {
	var slavePubKeySizeInBits string
	switch pub := slaveCert.PublicKey.(type) {
	case *rsa.PublicKey:
		slavePubKeySizeInBits = fmt.Sprintf("%v", len(pub.N.Bytes())*8)
	default:
		return nil, "", fmt.Errorf("unsupported public key type: %T", pub)
	}

	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		var dispatcherName string
		switch slaveCert.SignatureAlgorithm {
		case x509.SHA1WithRSA:
			dispatcherName = "C_RSA_SHA1_" + slavePubKeySizeInBits
		case x509.SHA256WithRSA:
			dispatcherName = "C_RSA_" + slavePubKeySizeInBits
		case x509.SHA384WithRSA:
			dispatcherName = "C_RSA_SHA384_" + slavePubKeySizeInBits
		case x509.SHA512WithRSA:
			dispatcherName = "C_RSA_SHA512_" + slavePubKeySizeInBits
		case x509.SHA256WithRSAPSS:
			dispatcherName = "C_RSAPSS_SHA2_" + slavePubKeySizeInBits
		case x509.SHA384WithRSAPSS:
			dispatcherName = "C_RSAPSS_SHA384_" + slavePubKeySizeInBits
		case x509.SHA512WithRSAPSS:
			dispatcherName = "C_RSAPSS_SHA512_" + slavePubKeySizeInBits
		default:
			return nil, "", fmt.Errorf("unsupported certificate signature algorithm: %v", slaveCert.SignatureAlgorithm.String())
		}

		return keccak256.Hash([]byte(dispatcherName)), dispatcherName, nil
	default:
		return nil, "", fmt.Errorf("unsupported public key type: %T", pub)
	}
}

func retriveCertificateRegistrationDispatcherForECDSAFamily(
	masterCert *x509.Certificate,
	slaveCert *x509.Certificate,
) ([]byte, string, error) {
	var slavePubKeySizeInBits string
	switch pub := slaveCert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		rawPubKeyData := append(pub.X.Bytes(), pub.Y.Bytes()...)
		slavePubKeySizeInBits = fmt.Sprintf("%v", len(rawPubKeyData)*8)
	default:
		return nil, "", fmt.Errorf("unsupported public key type: %T", pub)
	}

	switch pub := masterCert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		var curveName string
		switch pub.Curve.Params().Name {
		case "P-160":
			curveName = "SECP160R1"
		case "P-192":
			curveName = "SECP192R1"
		case "P-224":
			curveName = "SECP224R1"
		case "P-256":
			curveName = "SECP256R1"
		case "P-384":
			curveName = "SECP384R1"
		case "P-521":
			curveName = "SECP521R1"
		case "brainpoolP160r1":
			curveName = "BRAINPOOLP160R1"
		case "brainpoolP192r1":
			curveName = "BRAINPOOLP192R1"
		case "brainpoolP224r1":
			curveName = "BRAINPOOLP224R1"
		case "brainpoolP256r1":
			curveName = "BRAINPOOLP256R1"
		case "brainpoolP384r1":
			curveName = "BRAINPOOLP384R1"
		case "brainpoolP512r1":
			curveName = "BRAINPOOLP512R1"
		default:
			return nil, "", fmt.Errorf("unsupported curve: %v", pub.Curve.Params().Name)
		}

		var dispatcherName string
		switch slaveCert.SignatureAlgorithm {
		case x509.ECDSAWithSHA1:
			dispatcherName = "C_ECDSA_" + curveName + "_SHA1_" + slavePubKeySizeInBits
		case x509.ECDSAWithSHA256:
			dispatcherName = "C_ECDSA_" + curveName + "_SHA2_" + slavePubKeySizeInBits
		case x509.ECDSAWithSHA384:
			dispatcherName = "C_ECDSA_" + curveName + "_SHA384_" + slavePubKeySizeInBits
		case x509.ECDSAWithSHA512:
			dispatcherName = "C_ECDSA_" + curveName + "_SHA512_" + slavePubKeySizeInBits
		default:
			return nil, "", fmt.Errorf("unsupported certificate signature algorithm: %v", slaveCert.SignatureAlgorithm.String())
		}

		return keccak256.Hash([]byte(dispatcherName)), dispatcherName, nil
	default:
		return nil, "", fmt.Errorf("unsupported public key type: %T", pub)
	}
}

type noirRegistrationProof struct {
	points           []byte
	passportKey      *big.Int
	passportHash     *big.Int
	dgCommit         *big.Int
	identityKey      *big.Int
	certificatesRoot *big.Int
}

func newNoirRegistrationProof(fullProof []byte) (*noirRegistrationProof, error) {
	var proof noirRegistrationProof
	proof.points = fullProof[32*5:]

	pubSignalsPart := fullProof[:32*5]

	proof.passportKey = new(big.Int).SetBytes(pubSignalsPart[:32])
	proof.passportHash = new(big.Int).SetBytes(pubSignalsPart[32:64])
	proof.dgCommit = new(big.Int).SetBytes(pubSignalsPart[64:96])
	proof.identityKey = new(big.Int).SetBytes(pubSignalsPart[96:128])
	proof.certificatesRoot = new(big.Int).SetBytes(pubSignalsPart[128:160])

	return &proof, nil
}
