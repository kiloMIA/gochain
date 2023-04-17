package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func GeneratePublicKey() (string, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}
	pubKeyBytes := key.PublicKey.N.Bytes()
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKeyBytes)
	return pubKeyBase64, nil
}
