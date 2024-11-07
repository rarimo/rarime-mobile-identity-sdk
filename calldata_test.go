package identity_test

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
)

func TestBuildRegisterCalldata(t *testing.T) {
	proofJson, err := os.ReadFile("assets/test_register_proof.json")
	if err != nil {
		t.Fatal(err)
	}

	pubKeyPem, err := os.ReadFile("assets/test_register_pubkey.pem")
	if err != nil {
		t.Fatal(err)
	}

	signatureHex := "2fb614007c94084d8fe609a448116eb29ebf04fc3c19cdc326e611ab1c65b0e169a952133c5f1e3e71a83206cce4341dcc51634abc59d438c6dd5e6e9bf0eaf3da4c84cc835b393242b433e9ec5d17593df3145bf7f590bc96a82c06a04dc18781186efc1d64d45bed33b298f7ff4115446804568fe7e9ca5d513cb346ab9ec8"
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		t.Fatal(err)
	}

	builder := &identity.CallDataBuilder{}

	calldata, err := builder.BuildRegisterCalldata(proofJson, signature, pubKeyPem, 2668, []byte{}, false, "placeholder")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hex.EncodeToString(calldata))
}

func TestBuildRegisterCertificate(t *testing.T) {
	slavePem, err := os.ReadFile("assets/slave.pem")
	if err != nil {
		t.Errorf("failed to read slave pem: %v", err)
	}

	mastersPem, err := os.ReadFile("assets/masters.pem")
	if err != nil {
		t.Errorf("failed to read masters pem: %v", err)
	}

	calldataBuilder := &identity.CallDataBuilder{}

	calldata, err := calldataBuilder.BuildRegisterCertificateCalldata(mastersPem, slavePem)
	if err != nil {
		t.Errorf("failed to build calldata: %v", err)
	}

	fmt.Printf("calldata: %v\n", hex.EncodeToString(calldata))
}

func TestRevokeCalldata(t *testing.T) {
	aaPubKey, _ := os.ReadFile("assets/pubKey.pem")

	identityKey, _ := hex.DecodeString("0x039f80b68aaef6e8a73cf16ad9c242bb98471926d76a7e618736adadc746615e")
	signature, _ := hex.DecodeString("0x4a3fde41f253aeb52c2e11d453c952e035742b515d36fe5c5b038c8683b4e1653dc981f6aa8d41b6576b208709f65d72a62898095d1cff5bdae765de7b70901c")
	ecSizeInBits := 2_704

	builder := &identity.CallDataBuilder{}

	calldata, err := builder.BuildRevoceCalldata(identityKey, signature, aaPubKey, ecSizeInBits)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("revoke calldata: %v", hex.EncodeToString(calldata))
}
