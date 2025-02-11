package identity_test

import (
	"encoding/hex"
	"os"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
	"github.com/stretchr/testify/assert"
)

func TestCert(t *testing.T) {
	mastersPem, err := identity.LoadMasterCertificatesPem("rarimo-temp", "icaopkd-list.ldif")
	if err != nil {
		t.Errorf("failed to load master certificates pem: %v", err)
	}

	slavePem, err := os.ReadFile("assets/slave.pem")
	if err != nil {
		t.Errorf("failed to read slave pem: %v", err)
	}

	x508Util := identity.X509Util{}

	_, master, err := x508Util.GetMaster(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master: %v", err)
	}

	t.Logf("master signature: %v", hex.EncodeToString(master.Signature))

	masterPem, err := x508Util.CertificateToPem(master)
	if err != nil {
		t.Errorf("failed to convert certificate to pem: %v", err)
	}

	t.Logf("master pem: %v", string(masterPem))

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

	masterCertificateIndex, err := x508Util.GetSlaveCertificateIndex(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master certificate index: %v", err)
	}

	t.Log("master certificate index len : ", len(masterCertificateIndex))

	t.Logf("master certificate index: %v", hex.EncodeToString(masterCertificateIndex))
}

func TestCertBrail(t *testing.T) {
	slavePem := []byte("-----BEGIN CERTIFICATE-----\nxxx\n-----END CERTIFICATE-----\n")
	mastersPem := []byte("-----BEGIN CERTIFICATE-----\nxxx\n-----END CERTIFICATE-----\n")

	x508Util := identity.X509Util{}

	masterCertificateIndex, err := x508Util.GetSlaveCertificateIndex(slavePem, mastersPem)
	if err != nil {
		t.Errorf("failed to get master certificate index: %v", err)
	}

	t.Log("master certificate index len : ", len(masterCertificateIndex))

	t.Logf("master certificate index: %v", hex.EncodeToString(masterCertificateIndex))
}
