package identity

import (
	"crypto/rsa"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/rarimo/certificate-transparency-go/x509"
)

// X509Util used to simplify work with x509 certificates
type X509Util struct{}

// GetMaster takes a slave certificate and returns its master
func (x *X509Util) GetMaster(slavePem []byte, mastersPem []byte) (*x509.Certificate, *x509.Certificate, error) {
	slaveCert, err := parseCertificate(slavePem)
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

func parseCertificate(pemFile []byte) (*x509.Certificate, error) {
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

	fmt.Println("slaveSignedAttributes: ", hex.EncodeToString(slaveSignedAttributes))

	var masterModulus []byte
	switch pub := masterCert.PublicKey.(type) {
	case *rsa.PublicKey:
		masterModulus = pub.N.Bytes()
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}

	inputs := &PassportCertificateInputs{
		SlaveSignedAttributes: ByteArrayToBits(slaveSignedAttributes),
		SlaveSignature:        SmartChunking(new(big.Int).SetBytes(slaveSignature)),
		MasterModulus:         SmartChunking(new(big.Int).SetBytes(masterModulus)),
	}

	return inputs, nil
}
