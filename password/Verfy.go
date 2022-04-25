package password

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func Verify(message []byte, signature []byte, pk rsa.PublicKey) bool {

	msgHash := sha256.New()
	_, err := msgHash.Write(message)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err = rsa.VerifyPSS(&pk, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return false
	}
	// If we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	return true
}
