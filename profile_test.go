package identity_test

import (
	"encoding/hex"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
)

func TestProfile(t *testing.T) {
	secretKey := identity.NewBJJSecretKey()

	t.Logf("secretKey: %v\n", hex.EncodeToString(secretKey))

	profile := &identity.Profile{}

	profile, err := profile.NewProfile(secretKey)
	if err != nil {
		t.Errorf("NewProfile() failed: %v", err)
	}

	challenge, err := profile.GetRegistrationChallenge()
	if err != nil {
		t.Errorf("GetRegistrationChallenge() failed: %v", err)
	}

	t.Log(hex.EncodeToString(challenge))
}
