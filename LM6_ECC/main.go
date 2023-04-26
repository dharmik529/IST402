package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Generate a random curve for elliptic-curve cryptography
	curve := elliptic.P256()
	// Generate a new private key using the chosen curve
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Extract the public key from the private key
	publicKey := &privateKey.PublicKey

	// Get input from user
	fmt.Print("Please enter the message you wish to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// Remove newline character from input
	input = strings.TrimSpace(input)

	// Convert input to a byte slice
	message := []byte(input)

	// Generate a random nonce for use with AES-GCM encryption
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// Generate a shared secret between the private and public keys
	sharedSecretX, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	sharedSecret := sharedSecretX.Bytes()

	// Encrypt the message using AES-GCM
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		panic(err)
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		panic(err)
	}

	ciphertext := aesGcm.Seal(nil, nonce, message, nil)

	// Convert the ciphertext to a hexadecimal string for display
	hexCiphertext := hex.EncodeToString(ciphertext)

	// Display the encrypted message to the user
	fmt.Printf("Encrypted message: %s\n", hexCiphertext)

	// Decrypt the ciphertext using the same nonce and shared secret
	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// Display the decrypted message to the user
	fmt.Printf("Decrypted message: %s\n", plaintext)
}
