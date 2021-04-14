package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	"encoding/json"
)

func main() {
	privStr := "MHcCAQEEICfrz3CsrsscS9h04p4Tt7JYuUmMvb0a/bLAE99lj8y5oAoGCCqGSM49AwEHoUQDQg" +
		"AEaMDIHXKFZyLgNzintGwRYoXBo6hQ7vyEpudHeB8iHQr9n1fffwkJP3nIBm5TD9XEUjk72rAZEbylkSJOekTW0A=="
	privB, _ := base64.StdEncoding.DecodeString(privStr)
	priv, _ := x509.ParseECPrivateKey(privB)

	fmt.Println("pub key:")
	fmt.Println(hex.EncodeToString(elliptic.MarshalCompressed(priv.Curve, priv.X, priv.Y)))

	hash := bytes.Repeat([]byte{'1'}, 32)
	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash)
	sig := make([]byte, 0, 64)
	sig = append(sig, r.Bytes()...)
	sig = append(sig, s.Bytes()...)
	fmt.Println("signature:")
	fmt.Println(hex.EncodeToString(sig))

	fmt.Println("hash:")
	fmt.Println(hex.EncodeToString(hash))

	f, _ := os.Create("sig")
	defer f.Close()
	f.Write(sig)
}
