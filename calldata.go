package identity

import (
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/rarimo/ldif-sdk/mt"
)

// PNoAaHex represents the register data type.
const PNoAaHex = "9a0f175c44aa7c405c6ab99fbc9aa9e2cdc8971d1fea806282f7746205e8b807"

// PRsaSha12688Hex represents the register data type.
const PRsaSha12688Hex = "42ec2321c321a7ac25cb817a4e29d805e64817ab3a18b50127e0ccae327a111b"

// PEcdsaSha12704Hex represents the register data type.
const PEcdsaSha12704Hex = "b8abd3b1d40edd7da5ec75a9204661d45e63c046b5a19626a760482ca842fa1d"

// PRsaSha126883Hex represents the register data type.
const PRsaSha126883Hex = "8dc1b8b03716166cd99f5b390f2c6924085b150659e0cb3ca421ab47a1e65e09"

// CRsa4096Hex represents the register certificate data type.
const CRsa4096Hex = "16a6ebfa039c78163278bbd4ec27579c8c8939b01f1b3f6fb029c9682991cb5b"

// CRsa2048Hex represents the register certificate data type.
const CRsa2048Hex = "bf09b046e1fd32abb843f6ee4422c076a6fb365390d5be71020535c149781da1"

// ZUniversal4096Hex represents the register certificate data type.
const ZUniversal4096Hex = "fdd39d1855d1c9f04b1caa605196f28d42ffbe1673ecbfaf256c0b92e2aae9b5"

// ZUniversal2048Hex represents the register certificate data type.
const ZUniversal2048V3Hex = "77556bf7987c465a2feaca4a9277cdb9f98c9872275d283dea9a3854c8b692f1"

// ZInternalHex represents the register certificate data type.
const ZInternalHex = "6caaf1a07b99c61d7eab067e1af8e43fdeda473f5d537cf8250c8b6154121d21"

// RSAEXPONENT3 represents the RSA exponent 3.
const RSAEXPONENT3 = 3

// RegistrationMetaData contains all metadata for the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"stateKeeper_\",\"type\":\"address\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certificateDispatchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportDispatchers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signedAttributes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"keyOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationOffset\",\"type\":\"uint256\"}],\"internalType\":\"structRegistration2.Certificate\",\"name\":\"certificate_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration2.ICAOMember\",\"name\":\"icaoMember_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"icaoMerkleProof_\",\"type\":\"bytes32[]\"}],\"name\":\"registerCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"reissueIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRegistration2.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificateKey_\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateKeeper\",\"outputs\":[{\"internalType\":\"contractStateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRegistration2.MethodId\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"updateDependency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationPassport represents a registration passport.
type RegistrationPassport struct {
	DataType     [32]byte
	ZkType       [32]byte
	Signature    []byte
	PublicKey    []byte
	PassportHash [32]byte
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
	isRevoked bool,
	circuitName string,
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

	proofPoints := &VerifierHelperProofPoints{
		A: a,
		B: b,
		C: c,
	}

	var datatypeBuf []byte
	var pubKey []byte
	if len(pubKeyPem) == 0 {
		datatypeBuf, err = hex.DecodeString(PNoAaHex)
		if err != nil {
			return nil, err
		}

		pubKey = []byte{}
	} else {
		pubKeyBuf, isEcdsa, err := pubKeyPemToRaw(pubKeyPem)
		if err != nil {
			return nil, err
		}

		pubKey = pubKeyBuf

		if isEcdsa {
			datatypeBuf, err = hex.DecodeString(PEcdsaSha12704Hex)
		} else {
			pubKeyRsa, err := pemToRsaPubKey(pubKeyPem)
			if err != nil {
				return nil, err
			}

			if pubKeyRsa.E == RSAEXPONENT3 {
				datatypeBuf, err = hex.DecodeString(PRsaSha126883Hex)
			} else {
				datatypeBuf, err = hex.DecodeString(PRsaSha12688Hex)
			}
		}
	}

	datatype := [32]byte{}
	copy(datatype[:], datatypeBuf)

	zkTypeBuf := keccak256.Hash([]byte(circuitName))

	var zkType [32]byte
	copy(zkType[:], zkTypeBuf)

	passportHashBytes32 := make([]byte, 32)
	passportHash.FillBytes(passportHashBytes32)

	passport := &RegistrationPassport{
		DataType:     datatype,
		ZkType:       zkType,
		PublicKey:    pubKey,
		Signature:    signature,
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

// BuildRegisterCertificateCalldata builds the calldata for the register certificate function.
func (s *CallDataBuilder) BuildRegisterCertificateCalldata(
	cosmosAddr string,
	slavePem []byte,
	masterCertificatesBucketname string,
	masterCertificatesFilename string,
) ([]byte, error) {
	mastersPem, err := LoadMasterCertificatesPem(masterCertificatesBucketname, masterCertificatesFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load master certificates pem: %v", err)
	}

	icaoTree, err := mt.BuildFromCosmos(cosmosAddr, true)
	if err != nil {
		return nil, fmt.Errorf("failed to build tree from collection: %v", err)
	}

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

	var slaveMemberKey []byte
	switch pub := slaveCert.PublicKey.(type) {
	case *rsa.PublicKey:
		slaveMemberKey = pub.N.Bytes()
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	var dataType [32]byte
	if len(slaveMemberKey)*8 == 4096 {
		dataTypeBuf, err := hex.DecodeString(CRsa4096Hex)
		if err != nil {
			return nil, err
		}

		copy(dataType[:], dataTypeBuf)
	} else {
		dataTypeBuf, err := hex.DecodeString(CRsa2048Hex)
		if err != nil {
			return nil, err
		}

		copy(dataType[:], dataTypeBuf)
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

	return abi.Pack("registerCertificate", registrationCertificate, icaoMember, icaoMerkleProofSiblings)
}

// BuildRevoceCalldata builds the calldata for the revoke function.
func (s *CallDataBuilder) BuildRevoceCalldata(
	identityKey []byte,
	signature []byte,
	pubKeyPem []byte,
) ([]byte, error) {
	signature, err := NormalizeSignature(signature)
	if err != nil {
		return nil, err
	}

	var datatypeBuf []byte
	var pubKey []byte
	if len(pubKeyPem) == 0 {
		var err error
		datatypeBuf, err = hex.DecodeString(PNoAaHex)
		if err != nil {
			return nil, err
		}

	} else {
		pubKeyBuf, isEcdsa, err := pubKeyPemToRaw(pubKeyPem)
		if err != nil {
			return nil, err
		}

		pubKey = pubKeyBuf

		if isEcdsa {
			datatypeBuf, err = hex.DecodeString(PEcdsaSha12704Hex)
		} else {
			pubKeyRsa, err := pemToRsaPubKey(pubKeyPem)
			if err != nil {
				return nil, err
			}

			if pubKeyRsa.E == RSAEXPONENT3 {
				datatypeBuf, err = hex.DecodeString(PRsaSha126883Hex)
			} else {
				datatypeBuf, err = hex.DecodeString(PRsaSha12688Hex)
			}
		}
	}

	datatype := [32]byte{}
	copy(datatype[:], datatypeBuf)

	identityKeyInt := new(big.Int).SetBytes(identityKey)

	passport := &RegistrationPassport{
		DataType:  datatype,
		PublicKey: pubKey,
		Signature: signature,
	}

	abi, err := newRegistrationCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("revoke", identityKeyInt, passport)
}
