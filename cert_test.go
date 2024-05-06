package identity_test

import (
	"encoding/hex"
	"os"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
	"github.com/stretchr/testify/assert"
)

func TestCert(t *testing.T) {
	slavePem, err := os.ReadFile("assets/slave.pem")
	if err != nil {
		t.Errorf("failed to read slave pem: %v", err)
	}

	mastersPem, err := os.ReadFile("assets/masters.pem")
	if err != nil {
		t.Errorf("failed to read masters pem: %v", err)
	}

	x508Util := identity.X509Util{}

	_, master, err := x508Util.GetMaster(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master: %v", err)
	}

	assert.NotEqual(t, nil, master)

	keyPositionInSignedAttributes, err := x508Util.FindKeyPositionInSignedAttributes(master)
	if err != nil {
		t.Errorf("failed to find key position in signed attributes: %v", err)
	}

	t.Logf("key position in signed attributes: %v", keyPositionInSignedAttributes)

	expirationPositionInSignedAttributes, err := x508Util.FindExpirationPositionInSignedAttributes(master)
	if err != nil {
		t.Errorf("failed to find expiration position in signed attributes: %v", err)
	}

	t.Logf("expiration position in signed attributes: %v", expirationPositionInSignedAttributes)

	masterCertificateIndex, err := x508Util.GetMasterCertificateIndex(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master certificate index: %v", err)
	}

	t.Logf("master certificate index: %v", hex.EncodeToString(masterCertificateIndex))
}
