package identity

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/rarimo/certificate-transparency-go/x509"
)

// PubKeyANS1Prefix is the prefix for the public key in ANS1 format
var PubKeyANS1Prefix []byte = []byte{0x02, 0x82, 0x02, 0x01, 0x00}

// X509Util used to simplify work with x509 certificates
type X509Util struct{}

// GetMaster takes a slave certificate and returns its master
func (x *X509Util) GetMaster(slavePem []byte, mastersPem []byte) (*x509.Certificate, *x509.Certificate, error) {
	slaveCert, err := x.ParseCertificate(slavePem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse slave: %v", err)
	}

	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(mastersPem)

	foundCerts, err := slaveCert.Verify(x509.VerifyOptions{
		Roots: roots,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("invalid certificate: %w", err)
	}

	if len(foundCerts) == 0 {
		return nil, nil, fmt.Errorf("invalid certificate: no valid certificate found")
	}

	return slaveCert, foundCerts[0][1], nil
}

// PublicKeyToPem takes an x509 certificate and returns its public key in PEM format
func (x *X509Util) PublicKeyToPem(cert *x509.Certificate) ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

// CertificateToPem takes an x509.Certificate and returns it in PEM format
func (x *X509Util) CertificateToPem(cert *x509.Certificate) ([]byte, error) {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}

	certBytes := pem.EncodeToMemory(block)

	return certBytes, nil
}

// ParseCertificate parses a PEM certificate
func (x *X509Util) ParseCertificate(pemFile []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemFile)
	if block == nil {
		return nil, fmt.Errorf("invalid certificate: invalid PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse cert: %v", err)
	}

	return cert, nil
}

// BuildPartialRegistrationCircuitInputs returns the inputs for the registration circuit
func (x *X509Util) BuildPartialRegistrationCircuitInputs(slavePem []byte, mastersPem []byte) (*PassportCertificateInputs, error) {
	slaveCert, masterCert, err := x.GetMaster(slavePem, mastersPem)
	if err != nil {
		return nil, fmt.Errorf("failed to get master: %v", err)
	}

	slaveSignedAttributes := slaveCert.RawTBSCertificate
	slaveSignature := slaveCert.Signature

	var masterModulus []byte
	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		masterModulus = pub.N.Bytes()
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	inputs := &PassportCertificateInputs{
		SlaveSignedAttributes: ByteArrayToBits(slaveSignedAttributes),
		SlaveSignature:        SmartChunking(new(big.Int).SetBytes(slaveSignature), 64),
		MasterModulus:         SmartChunking(new(big.Int).SetBytes(masterModulus), 64),
	}

	return inputs, nil
}

// FindKeyPositionInSignedAttributes finds the position of the key in the signed attributes
func (x *X509Util) FindKeyPositionInSignedAttributes(cert *x509.Certificate) (*big.Int, error) {
	signedAttributes := cert.RawTBSCertificate

	var publicKey []byte
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		publicKey = pub.N.Bytes()
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	index := bytes.Index(signedAttributes, publicKey)
	if index == -1 {
		return nil, errors.New("subarray not found in array")
	}

	return new(big.Int).SetInt64(int64(index)), nil
}

// FindExpirationPositionInSignedAttributes finds the position of the expiration date in the signed attributes
func (x *X509Util) FindExpirationPositionInSignedAttributes(cert *x509.Certificate) (*big.Int, error) {
	expiration := cert.NotAfter

	expirationUTCTime := expiration.UTC().Format("060102150405Z")

	signedAttributes := cert.RawTBSCertificate

	index := bytes.Index(signedAttributes, []byte(expirationUTCTime))
	if index == -1 {
		return nil, errors.New("subarray not found in array")
	}

	return new(big.Int).SetInt64(int64(index)), nil
}

// GetSlaveCertificateIndex returns the index of the master certificate of the slave certificate
func (x *X509Util) GetSlaveCertificateIndex(slavePem []byte, mastersPem []byte) ([]byte, error) {
	slaveCert, _, err := x.GetMaster(slavePem, mastersPem)
	if err != nil {
		return nil, fmt.Errorf("failed to get master: %v", err)
	}

	var masterCertificateIndex *big.Int
	switch pub := slaveCert.PublicKey.(type) {
	case *rsa.PublicKey:
		pubKeyRaw := pub.N.Bytes()

		masterCertificateIndex, err = HashPacked(pubKeyRaw)
		if err != nil {
			return nil, fmt.Errorf("failed to hash key: %v", err)
		}
	case *ecdsa.PublicKey:
		pubKeyRaw := pub.X.Bytes()
		pubKeyRaw = append(pubKeyRaw, pub.Y.Bytes()...)

		masterCertificateIndex, err = Hash512(pubKeyRaw)
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	return masterCertificateIndex.Bytes(), nil
}

// GetRSASize returns the size of the RSA key
func (x *X509Util) GetRSASize(pubKeyPem []byte) (int, error) {
	rsaPubKeyN, _, err := pubKeyPemToRaw(pubKeyPem)
	if err != nil {
		return 0, fmt.Errorf("error parsing public key: %v", err)
	}

	return len(rsaPubKeyN) * 8, nil
}
