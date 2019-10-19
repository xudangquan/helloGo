package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {

}

//message = timestamp + postData
func GenerateSignature(message string) (string, error) {
	privKey, err := loadPrivateKey("<key path>", "signKey")
	if err != nil {
		return "", err
	}

	hashed := sha256.Sum256([]byte(message))

	s, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	signature := "digest-alg=RSA-SHA; key-id=KEY:RSA:rsf.org; data=" + base64.StdEncoding.EncodeToString(s)

	return signature, nil
}

func loadPrivateKey(path, apiKey string) (*rsa.PrivateKey, error) {
	var data []byte
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parsePrivateKey(data)
}

func parsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey *rsa.PrivateKey
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return rawkey, nil
}
