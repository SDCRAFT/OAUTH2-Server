package globals

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func generateRSAKeys() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		},
	)

	return publicKeyPEM, privateKeyPEM
}

func init() {
	pub, pri := generateRSAKeys()
	RSAPublicKey = string(pub)
	RSAPrivateKey = pri
}

var (
	Hash       string = ""
	Version    string = ""
	Buildstamp string = ""
)

var (
	RSAPublicKey  string
	RSAPrivateKey []byte
)
