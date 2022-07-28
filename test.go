package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/veraison/go-cose"
)

func main() {
	// create a signer
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	signer, _ := cose.NewSigner(cose.AlgorithmES512, privateKey)

	// create message header
	headers := cose.Headers{
		Protected: cose.ProtectedHeader{
			cose.HeaderLabelAlgorithm: cose.AlgorithmES512,
		},
	}

	// sign and marshal message
	sig, _ := cose.Sign1(rand.Reader, signer, headers, []byte("hello world"), nil)
	fmt.Println(sig)
}
