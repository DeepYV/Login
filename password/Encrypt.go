package password

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func RSA_Encrypt(Message []byte, key *rsa.PrivateKey) []byte {
	msgHash := sha256.New()
	_, err := msgHash.Write(Message)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message

	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		fmt.Print("Error encode")
		panic(err)
	}
	return signature
}
