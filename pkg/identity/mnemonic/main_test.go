package identity

import "testing"

func TestMnemonic(t *testing.T) {
	// Test NewMnemonic
	mnemonic, err := NewMnemonic()
	if err != nil {
		t.Errorf("NewMnemonic() failed: %v", err)
	}
	if mnemonic == "" {
		t.Errorf("NewMnemonic() failed: mnemonic is empty")
	}

	// Test SeedFromMnemonic
	seed, err := SeedFromMnemonic(mnemonic)
	if err != nil {
		t.Errorf("SeedFromMnemonic() failed: %v", err)
	}
	if len(seed) == 0 {
		t.Errorf("SeedFromMnemonic() failed: seed is empty")
	}

	// Test NewSecretKeyFromSeed
	secretKey, err := NewSecretKeyFromSeed(seed)
	if err != nil {
		t.Errorf("NewSecretKeyFromSeed() failed: %v", err)
	}

	if secretKey == "" {
		t.Errorf("NewSecretKeyFromSeed() failed: secretKey is empty")
	}
}
