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

	signatureHex := "96277aebb727fb5a15755b0f2bd41febeb0dec810a7ba1fc002c3326d57696caa07804bbc0bb54cdc4fa4d0baeee14b2ba99a0ff0e5a35839551cb3fb15ab241e8229124868d1cb8c13d02b27005c4d3eac24d994b75bd54c6b61174d0687d7549f553c24832755802a95a7c92ae842c04bf634d98df4827ed085d773c0465bab92d27a211c41ee5eadd8cd78fa17c49ad7d9827844b185873f7968330b015ce"
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		t.Fatal(err)
	}

	builder := &identity.CallDataBuilder{}

	masterRoot, _ := hex.DecodeString("242f36929b6d99785dec1a9ba087033a5d09e8ed58835f710fe689956e30a801")

	calldata, err := builder.BuildRegisterCalldata(proofJson, signature, pubKeyPem, 1440, masterRoot, false, "registerIdentity_11_256_3_3_576_248_1_1184_5_264")
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

	registerCertificateCalldataResponse, err := calldataBuilder.BuildRegisterCertificateCalldata(mastersPem, slavePem)
	if err != nil {
		t.Errorf("failed to build calldata: %v", err)
	}

	fmt.Printf("calldata: %v\n", hex.EncodeToString(registerCertificateCalldataResponse.Calldata))
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
