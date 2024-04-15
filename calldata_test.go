package identity_test

import (
	"os"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
)

func TestBuildRegisterCalldata(t *testing.T) {
	proofJson, err := os.ReadFile("assets/test_register_proof.json")
	if err != nil {
		t.Fatal(err)
	}

	calldata, err := identity.BuildRegisterCalldata(string(proofJson))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(calldata)
}
