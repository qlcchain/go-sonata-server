package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const (
	DefaultKey = `-----BEGIN PRIVATE KEY-----
MIHcAgEBBEIBp0kGRR9nSq1lQPZ3WspSjjGtGduopifh5epB5MMziYdHwILlGf0e
ARYLiyZvRJ8LjK5DT2OVasAFG6efpMmHBcygBwYFK4EEACOhgYkDgYYABACfNcNR
IAv1wSoqXI2Ve2hKDEx5hoL037dqC0Hs3GhrGiR0fidivhdTEu4wWCV3Heg0K91x
7fmD6uIlE9K9TWb85wGLsp1HuR+6nXabfNp29f3LaZtrnkW/4sN4+iaVR9LaIMJZ
wwK/9cMKhhtN81oNVp23ake+WwKRI7LETAtJR/jB2w==
-----END PRIVATE KEY-----`
)

func Decode(pemEncoded string) *ecdsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	return privateKey
}

func Encode(privateKey *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return string(pemEncoded)
}

func New() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

func NewBase64() string {
	key, _ := New()
	encode := Encode(key)
	return base64.StdEncoding.EncodeToString([]byte(encode))
}
