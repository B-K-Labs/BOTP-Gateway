package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKey(privkey *rsa.PrivateKey) string {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privkeyBytes}
	privateKeyStr := string(pem.EncodeToMemory(privateKeyPEM))
	return privateKeyStr
}

func ExportRsaPublicKey(pubkey *rsa.PublicKey) string {
	pubkeyBytes, _ := x509.MarshalPKIXPublicKey(pubkey)
	publicKeyPEM := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubkeyBytes}
	publicKeyStr := string(pem.EncodeToMemory(publicKeyPEM))
	return publicKeyStr
}

func ParseRsaPublicKey(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not an RSA public key")
	}

	return rsaPublicKey, nil
}

func VerifyRSASignature(publicKeyPem, signatureBase64 string, message []byte) bool {
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Println("Error decoding signature:", err)
		return false
	}
	publicKey, err := ParseRsaPublicKey(publicKeyPem)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		return false
	}
	opts := &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}

	hash := sha256.Sum256(message)
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hash[:], signature, opts)

	fmt.Println("err", err)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return false
	}

	return true
}
