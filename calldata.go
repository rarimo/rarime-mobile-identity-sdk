package identity

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"math/big"
	"strings"

	"encoding/json"

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

// VotingMetaData contains all metadata for the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\": false, \"inputs\": [{\"indexed\": false, \"internalType\": \"address\", \"name\": \"previousAdmin\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"address\", \"name\": \"newAdmin\", \"type\": \"address\"}], \"name\": \"AdminChanged\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beacon\", \"type\": \"address\"}], \"name\": \"BeaconUpgraded\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": false, \"internalType\": \"uint8\", \"name\": \"version\", \"type\": \"uint8\"}], \"name\": \"Initialized\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"previousOwner\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"OwnershipTransferred\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"implementation\", \"type\": \"address\"}], \"name\": \"Upgraded\", \"type\": \"event\"}, {\"inputs\": [], \"name\": \"IDENTITY_LIMIT\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"PROOF_SIGNALS_COUNT\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"SELECTOR\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"ZERO_DATE\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"registrationSMT_\", \"type\": \"address\"}, {\"internalType\": \"address\", \"name\": \"proposalsState_\", \"type\": \"address\"}, {\"internalType\": \"address\", \"name\": \"votingVerifier_\", \"type\": \"address\"}], \"name\": \"__BioPassportVoting_init\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"implementation\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"owner\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"proposalsState\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"proxiableUUID\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"registrationSMT\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"renounceOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"transferOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"newImplementation\", \"type\": \"address\"}], \"name\": \"upgradeTo\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"newImplementation\", \"type\": \"address\"}, {\"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"upgradeToAndCall\", \"outputs\": [], \"stateMutability\": \"payable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"registrationRoot_\", \"type\": \"bytes32\"}, {\"internalType\": \"uint256\", \"name\": \"currentDate_\", \"type\": \"uint256\"}, {\"internalType\": \"uint256\", \"name\": \"proposalId_\", \"type\": \"uint256\"}, {\"internalType\": \"uint256[]\", \"name\": \"vote_\", \"type\": \"uint256[]\"}, {\"components\": [{\"internalType\": \"uint256\", \"name\": \"nullifier\", \"type\": \"uint256\"}, {\"internalType\": \"uint256\", \"name\": \"citizenship\", \"type\": \"uint256\"}, {\"internalType\": \"uint256\", \"name\": \"identityCreationTimestamp\", \"type\": \"uint256\"}], \"internalType\": \"struct BaseVoting.UserData\", \"name\": \"userData_\", \"type\": \"tuple\"}, {\"components\": [{\"internalType\": \"uint256[2]\", \"name\": \"a\", \"type\": \"uint256[2]\"}, {\"internalType\": \"uint256[2][2]\", \"name\": \"b\", \"type\": \"uint256[2][2]\"}, {\"internalType\": \"uint256[2]\", \"name\": \"c\", \"type\": \"uint256[2]\"}], \"internalType\": \"struct VerifierHelper.ProofPoints\", \"name\": \"zkPoints_\", \"type\": \"tuple\"}], \"name\": \"vote\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"votingVerifier\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\"}]",
}

var FaceRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCallVerifyProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"featureHash\",\"type\":\"uint256\"}],\"name\":\"FeatureHashAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"}],\"name\":\"InvalidCircomProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"KeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maxDepth\",\"type\":\"uint32\"}],\"name\":\"MaxDepthExceedsHardCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxDepthIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxDepthReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"currentDepth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"newDepth\",\"type\":\"uint32\"}],\"name\":\"NewMaxDepthMustBeLarger\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotAnOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeIsNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress\",\"type\":\"uint256\"}],\"name\":\"UserAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress\",\"type\":\"uint256\"}],\"name\":\"UserNotRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"MinThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newOwners\",\"type\":\"address[]\"}],\"name\":\"OwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"removedOwners\",\"type\":\"address[]\"}],\"name\":\"OwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAddress\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"}],\"name\":\"RulesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"RulesVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAddress\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"featureHash\",\"type\":\"uint256\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"VerifierUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVENT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACE_PROOF_SIGNALS_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_VALIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RULES_PROOF_SIGNALS_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evidenceRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"faceVerifier_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rulesVerifier_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minThreshold_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"}],\"name\":\"__FaceRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"}],\"name\":\"addOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newOwners_\",\"type\":\"address[]\"}],\"name\":\"addOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evidenceRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"}],\"name\":\"getFeatureHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"}],\"name\":\"getRule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"address_\",\"type\":\"uint256\"}],\"name\":\"getVerificationNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"featureHash_\",\"type\":\"uint256\"}],\"name\":\"isFeatureHashUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"}],\"name\":\"isOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootLatest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"}],\"name\":\"isUserRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"featureHash_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"}],\"name\":\"removeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oldOwners_\",\"type\":\"address[]\"}],\"name\":\"removeOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"featureHash\",\"type\":\"uint256\"}],\"name\":\"rules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rulesVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier_\",\"type\":\"address\"}],\"name\":\"setFaceVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold_\",\"type\":\"uint256\"}],\"name\":\"setMinThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier_\",\"type\":\"address\"}],\"name\":\"setRulesVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structGroth16VerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"updateRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"featureHash\",\"type\":\"uint256\"}],\"name\":\"usedFeatureHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAddress\",\"type\":\"uint256\"}],\"name\":\"userRegistryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"featureHash\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// VotingUserData is an auto generated low-level Go binding around an user-defined struct.
type VotingUserData struct {
	Nullifier                 *big.Int
	Citizenship               *big.Int
	IdentityCreationTimestamp *big.Int
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

func newVotingCoder() (*abi.ABI, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func newFaceRegistryCoder() (*abi.ABI, error) {
	parsed, err := FaceRegistryMetaData.GetAbi()
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

// BuildVoteCalldata builds the calldata for the vote function.
//
// Vote(registrationRoot_ [32]byte, currentDate_ *big.Int, proposalId_ *big.Int, vote_ []*big.Int, userData_ VotingUserData, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
func (s *CallDataBuilder) BuildVoteCalldata(
	queryZkProofJSON []byte,
	proposalID int64,
	pollResultsJSON []byte,
	citizenship string,
	isRegisteredAfterVoting bool,
) ([]byte, error) {
	zkProof := new(ZkProof)
	if err := json.Unmarshal(queryZkProofJSON, zkProof); err != nil {
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

	nullifier, ok := new(big.Int).SetString(zkProof.PubSignals[0], 10)
	if !ok {
		return nil, fmt.Errorf("error setting nullifier: %v", zkProof.PubSignals[0])
	}

	var identityCreationTimestamp = new(big.Int).SetInt64(0)
	if isRegisteredAfterVoting {
		identityCreationTimestamp, ok = new(big.Int).SetString(zkProof.PubSignals[15], 10)
		if !ok {
			return nil, fmt.Errorf("error setting identityCreationTimestamp: %v", zkProof.PubSignals[15])
		}
	}

	userData := VotingUserData{
		Nullifier:                 nullifier,
		Citizenship:               new(big.Int).SetBytes([]byte(citizenship)),
		IdentityCreationTimestamp: identityCreationTimestamp,
	}

	var pollResults []PollResult
	if err := json.Unmarshal(pollResultsJSON, &pollResults); err != nil {
		return nil, err
	}

	var vote []*big.Int
	for _, v := range pollResults {
		voteI := makeByteWithBits(v.AnswerIndex)

		vote = append(vote, big.NewInt(int64(voteI)))
	}

	registrationRootRaw, ok := new(big.Int).SetString(zkProof.PubSignals[11], 10)
	if !ok {
		return nil, fmt.Errorf("error setting registrationRoot: %v", zkProof.PubSignals[11])
	}

	currentDate := zkProof.PubSignals[13]
	currentDateBigUInt, ok := new(big.Int).SetString(currentDate, 10)
	if !ok {
		return nil, fmt.Errorf("error setting currentDate: %v", currentDate)
	}

	var registrationRoot [32]byte
	copy(registrationRoot[:], registrationRootRaw.Bytes())

	abi, err := newVotingCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("vote", registrationRoot, currentDateBigUInt, big.NewInt(proposalID), vote, userData, VerifierHelperProofPoints{A: a, B: b, C: c})
}

func (s *CallDataBuilder) BuildFaceRegistryRegisterUser(zkPointsJSON []byte) ([]byte, error) {
	zkProof, proofPoints, err := prepareZKProofForEVMVerification(zkPointsJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare zk proof for evm verification: %v", err)
	}

	userAddress, ok := new(big.Int).SetString(zkProof.PubSignals[2], 10)
	if !ok {
		return nil, fmt.Errorf("error setting userAddress: %v", zkProof.PubSignals[2])
	}

	featureHash, ok := new(big.Int).SetString(zkProof.PubSignals[0], 10)
	if !ok {
		return nil, fmt.Errorf("error setting featureHash: %v", zkProof.PubSignals[0])
	}

	abi, err := newFaceRegistryCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("registerUser", userAddress, featureHash, proofPoints)
}

func (s *CallDataBuilder) BuildFaceRegistryUpdateRule(
	newState string,
	zkPointsJSON []byte,
) ([]byte, error) {
	zkProof, proofPoints, err := prepareZKProofForEVMVerification(zkPointsJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare zk proof for evm verification: %v", err)
	}

	userAddress, ok := new(big.Int).SetString(zkProof.PubSignals[0], 10)
	if !ok {
		return nil, fmt.Errorf("error setting userAddress: %v", zkProof.PubSignals[0])
	}

	newStateInt, ok := new(big.Int).SetString(newState, 10)
	if !ok {
		return nil, fmt.Errorf("error setting newState: %v", newState)
	}

	abi, err := newFaceRegistryCoder()
	if err != nil {
		return nil, err
	}

	return abi.Pack("updateRule", userAddress, newStateInt, proofPoints)
}
