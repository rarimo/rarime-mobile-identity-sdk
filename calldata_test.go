package identity_test

import (
	"encoding/hex"
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

	calldata, err := builder.BuildRegisterCalldata(proofJson, signature, pubKeyPem, []byte{})
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

	calldataBuilder := &identity.CallDataBuilder{}

	cosmosAddr := "core-api.node1.mainnet-beta.rarimo.com:443"

	_, err = calldataBuilder.BuildRegisterCertificateCalldata(cosmosAddr, slavePem, "rarimo-temp", "icaopkd-list.ldif")
	if err != nil {
		t.Errorf("failed to build calldata: %v", err)
	}
}
