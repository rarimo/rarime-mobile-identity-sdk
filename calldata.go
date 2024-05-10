package identity

import (
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/ldif-sdk/mt"
)

// RsaSha12688Hex represents the register data type.
const RsaSha12688Hex = "ee72172757e0738b89f37b0c9d04d6b9056da936d5e0959e3e8829d8fb91e4eb"

// RsaSha12688TimestampHex represents the register data type.
const RsaSha12688TimestampHex = "b889e143461475f0ed26836f5d521b7960b904f6e1a14b06754abec6a3f326a2"

// EcdsaSha12704Hex represents the register data type
const EcdsaSha12704Hex = "2370a10b5e0b6f239856dd28c7b60a91591658322ef67976cf52cabd81a153bc"

// RegistrationMetaData contains all metadata for the Registration contract.
//
// Register(certificatesRoot_ [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints)
// RegisterCertificate(icaoMerkleProof_ [][32]byte, icaoMemberKey_ []byte, icaoMemberSignature_ []byte, x509SignedAttributes_ []byte, x509KeyOffset_ *big.Int, x509ExpirationOffset_ *big.Int)
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"certificateKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"name\":\"CertificateRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"certificateKey\",\"type\":\"bytes32\"}],\"name\":\"CertificateRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"ReissuedIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ICAO_PREFIX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVOKED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registrationSmt_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"certificatesSmt_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"icaoMasterTreeMerkleRoot_\",\"type\":\"bytes32\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dispatcherType_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"dispatcher_\",\"type\":\"address\"}],\"name\":\"addDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"certificatesSmt\",\"outputs\":[{\"internalType\":\"contractPoseidonSMT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"changeICAOMasterTreeRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificateKey_\",\"type\":\"bytes32\"}],\"name\":\"getCertificateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structRegistration.CertificateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportKey_\",\"type\":\"bytes32\"}],\"name\":\"getPassportInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"activeIdentity\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"identityReissueCounter\",\"type\":\"uint64\"}],\"internalType\":\"structRegistration.PassportInfo\",\"name\":\"passportInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"activePassport\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"issueTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structRegistration.IdentityInfo\",\"name\":\"identityInfo_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"icaoMasterTreeMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot_\",\"type\":\"bytes32\"}],\"name\":\"mockChangeICAOMasterTreeRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportDispatchers\",\"outputs\":[{\"internalType\":\"contractIPassportDispatcher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"icaoMerkleProof_\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"icaoMemberKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"icaoMemberSignature_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"x509SignedAttributes_\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"x509KeyOffset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x509ExpirationOffset_\",\"type\":\"uint256\"}],\"name\":\"registerCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationSmt\",\"outputs\":[{\"internalType\":\"contractPoseidonSMT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"reissueIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dispatcherType_\",\"type\":\"bytes32\"}],\"name\":\"removeDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificateKey_\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationPassport represents a registration passport.
type RegistrationPassport struct {
	DataType  [32]byte
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

func newRegistrationCoder() (*abi.ABI, error) {
	parsed, err := RegistrationMetaData.GetAbi()
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
	signature []byte,
	pubKeyPem []byte,
	certificatesRootRaw []byte,
) ([]byte, error) {
	signature, err := NormalizeSignature(signature)
	if err != nil {
		return nil, err
	}

	zkProof := new(ZkProof)
	if err := json.Unmarshal(proofJSON, zkProof); err != nil {
		return nil, err
	}

	var a [2]*big.Int
	for index, val := range zkProof.Proof.A[:2] {
		aI, ok := new(big.Int).SetString(val, 10)
		if !ok {
			return nil, fmt.Errorf("error setting a[%d]: %v", index, val)
		}

		a[index] = aI
	}

	var b [2][2]*big.Int
	for index, val := range zkProof.Proof.B[:2] {
		for index2, val2 := range val[:2] {
			bI, ok := new(big.Int).SetString(val2, 10)
			if !ok {
				return nil, fmt.Errorf("error setting b[%d][%d]: %v", index, index2, val2)
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
			return nil, fmt.Errorf("error setting c[%d]: %v", index, val)
		}

		c[index] = cI
	}

	dg1Commitment, ok := new(big.Int).SetString(zkProof.PubSignals[1], 10)
	if !ok {
		return nil, fmt.Errorf("error setting dg1Commitment: %v", zkProof.PubSignals[1])
	}

	pkIdentityHash, ok := new(big.Int).SetString(zkProof.PubSignals[2], 10)
	if !ok {
		return nil, fmt.Errorf("error setting pkIdentityHash: %v", zkProof.PubSignals[2])
	}

	proofPoints := &VerifierHelperProofPoints{
		A: a,
		B: b,
		C: c,
	}

	pubKey, isEcdsa, err := pubKeyPemToRaw(pubKeyPem)
	if err != nil {
		return nil, err
	}

	var datatypeBuf []byte
	if isEcdsa {
		datatypeBuf, err = hex.DecodeString(EcdsaSha12704Hex)
	} else {
		datatypeBuf, err = hex.DecodeString(RsaSha12688Hex)
	}

	datatype := [32]byte{}
	copy(datatype[:], datatypeBuf)

	passport := &RegistrationPassport{
		DataType:  datatype,
		PublicKey: pubKey,
		Signature: signature,
	}

	var certificatesRoot [32]byte
	copy(certificatesRoot[:], certificatesRootRaw)

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("register", certificatesRoot, pkIdentityHash, dg1Commitment, passport, proofPoints)
}

// BuildRegisterCertificateCalldata builds the calldata for the register certificate function.
func (s *CallDataBuilder) BuildRegisterCertificateCalldata(
	cosmosAddr string,
	slavePem []byte,
	mastersPem []byte,
) ([]byte, error) {
	icaoTree, err := mt.BuildFromCosmos(cosmosAddr, true)
	if err != nil {
		return nil, fmt.Errorf("failed to build tree from collection: %v", err)
	}

	fmt.Println("tree built, root: ", hex.EncodeToString(icaoTree.Root()))

	x := X509Util{}
	slaveCert, masterCert, err := x.GetMaster(slavePem, mastersPem)
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

	var icaoMemberKey []byte
	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		icaoMemberKey = pub.N.Bytes()
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	icaoMemberSignature := slaveCert.Signature
	x509SignedAttributes := slaveCert.RawTBSCertificate

	x509KeyOffset, err := x.FindKeyPositionInSignedAttributes(slaveCert)
	if err != nil {
		return nil, fmt.Errorf("failed to find key position in signed attributes: %v", err)
	}

	x509ExpirationOffset, err := x.FindExpirationPositionInSignedAttributes(slaveCert)
	if err != nil {
		return nil, fmt.Errorf("failed to find expiration position in signed attributes: %v", err)
	}

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

	return abi.Pack("registerCertificate", icaoMerkleProofSiblings, icaoMemberKey, icaoMemberSignature, x509SignedAttributes, x509KeyOffset, x509ExpirationOffset)
}
