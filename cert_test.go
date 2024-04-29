package identity_test

import (
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

	master, err := x508Util.GetMaster(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master: %v", err)
	}

	assert.NotEqual(t, nil, master)

	_, err = x508Util.PublicKeyToPem(master)
	if err != nil {
		t.Errorf("failed to get public key from pem: %v", err)
	}

	_, err = x508Util.CertificateToPem(master)
	if err != nil {
		t.Errorf("failed to get master pem: %v", err)
	}
}
