package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func Encrypt_data(rsa_key string, data []byte) []byte {

	publicKeyBlock, _ := pem.Decode([]byte(rsa_key))
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	encrypted_data, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), data)
	if err != nil {
		panic(err)
	}

	return encrypted_data
}
