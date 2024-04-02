package identity

import (
	"encoding/hex"
	"fmt"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy: %w", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}

	return mnemonic, nil
}

func SeedFromMnemonic(mnemonic string) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, "")
	return seed, nil
}

func NewSecretKeyFromSeed(seed []byte) (string, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", fmt.Errorf("failed to generate master key: %w", err)
	}

	var secretKey babyjub.PrivateKey
	copy(secretKey[:], masterKey.Key[:32])

	return hex.EncodeToString(secretKey[:]), nil
}
