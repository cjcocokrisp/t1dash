package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	mathrand "math/rand/v2"
	"net"
	"os"
	"time"

	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/pkg/util"

	log "github.com/sirupsen/logrus"
)

// InitTLS checks to see if TLS keys and certs are there if they are not it generates them
// It returns a bool representing if everything was successful that can be used to determine
// if to include TLS in the server or not. If any of the files are missing, all will be generated.
func InitTLS() bool {
	type CheckParams struct {
		path          string
		checkIfExists bool
	}

	checks := []CheckParams{
		{path: config.AppCfg.TLSCAKeyPath, checkIfExists: false},
		{path: config.AppCfg.TLSCACertPath, checkIfExists: true},
		{path: config.AppCfg.TLSKeyPath, checkIfExists: false},
		{path: config.AppCfg.TLSCertPath, checkIfExists: true},
	}

	for _, e := range checks {
		create, err := checkIfExists(e.path, e.checkIfExists)
		if err != nil {
			return false
		}

		if create {
			log.Infof("%s does not exist or has expired, generating new tls certs and keys", e.path)
			err = setupTLS()
			if err != nil {
				return false
			}
			break
		}
	}

	return true
}

// checkIfExists checks to see if the file exists and returns if it needs to be created or not
func checkIfExists(path string, checkExp bool) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			util.LogError(fmt.Sprintf("Error while checking to see if %s exists", path), "tls init", err)
			return false, err
		}
		return true, nil
	} else if checkExp {
		expired, err := checkCertExp(path)
		if err != nil {
			util.LogError(fmt.Sprintf("Error while checking if %s has expired", path), "tls init", err)
			return false, err
		}
		if expired {
			log.Infof("Cert %s is expired, going to regenerate", path)
		}
		return expired, nil
	}
	return false, nil
}

// checkCertExp checks the certs expiration and returns if it is valid or not. If invalid
// generates a new one
func checkCertExp(path string) (bool, error) {
	certPem, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	block, _ := pem.Decode(certPem)
	if block == nil || block.Type != "CERTIFICATE" {
		return false, fmt.Errorf("Invalid contents of cert %s", path)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, err
	}

	return time.Now().After(cert.NotAfter), nil
}

// setupTLS creates the server cert and key for TLS and saves it to file specified
// path is specified by environment variable.
func setupTLS() error {
	// Generate ca related cert and key
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(mathrand.Int64()),
		Subject: pkix.Name{
			Organization:  []string{"T1Dash"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{""},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		util.LogError("Error creating ca key", "tls gen", err)
		return err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caKey.PublicKey, caKey)
	if err != nil {
		util.LogError("Error creating ca bytes", "tls gen", err)
		return err
	}

	caPem := new(bytes.Buffer)
	pem.Encode(caPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	err = os.WriteFile(config.AppCfg.TLSCACertPath, caPem.Bytes(), 0600)
	if err != nil {
		util.LogError("Error writing ca cert", "tls gen", err)
		return err
	}

	caKeyPem := new(bytes.Buffer)
	pem.Encode(caKeyPem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caKey),
	})

	err = os.WriteFile(config.AppCfg.TLSCAKeyPath, caKeyPem.Bytes(), 0600)
	if err != nil {
		util.LogError("Error writing ca key", "tls gen", err)
		return err
	}

	// Generate server cert and key
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(mathrand.Int64()),
		Subject: pkix.Name{
			Organization:  []string{"T1Dash"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{""},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		IPAddresses:           []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback}, // Need to alter off of localhost at some point (maybe do lookup)
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	certKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		util.LogError("error creating key for server", "tls gen", err)
		return err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &certKey.PublicKey, certKey)
	if err != nil {
		util.LogError("error creating bytes for server", "tls gen", err)
		return err
	}

	certPem := new(bytes.Buffer)
	pem.Encode(certPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	err = os.WriteFile(config.AppCfg.TLSCertPath, certPem.Bytes(), 0600)
	if err != nil {
		util.LogError("error writing server cert", "tls gen", err)
		return err
	}

	certKeyPem := new(bytes.Buffer)
	pem.Encode(certKeyPem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certKey),
	})

	err = os.WriteFile(config.AppCfg.TLSKeyPath, certKeyPem.Bytes(), 0600)
	if err != nil {
		util.LogError("error writing server key", "tls gen", err)
		return err
	}

	log.Info("Successfully generated TLS certs and keys for ca and server")

	return nil
}
