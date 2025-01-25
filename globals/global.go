package globals

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/sirupsen/logrus"
	"sdcraft.fun/oauth2/models"
)

var (
	Hash       string = ""
	Version    string = ""
	Buildstamp string = ""
)

var (
	RSAPublicKey  string
	RSAPrivateKey []byte
	Config        models.Config = models.NewConfig()
)

func generateRSAKeys() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyPKCS1 := x509.MarshalPKCS1PrivateKey(privateKey)

	if err != nil {
		logrus.Fatalf("Failed to generate private key: %v", err)
	}

	publicKeyD, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

	if err != nil {
		logrus.Fatalf("Failed to generate private key: %v", err)
	}

	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyD,
		},
	)

	return publicKeyPEM, privateKeyPKCS1
}

func init() {
	pub, pri := generateRSAKeys()
	RSAPublicKey = string(pub)
	RSAPrivateKey = pri
}
