package main

// LM6 - ECC in GoLang
// IST 402

import (
	"bufio"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Generate a random curve
	curve := elliptic.P256()

	// Generate a private key
	privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Generate a public key
	// publicKey := elliptic.Marshal(curve, x, y)

	// Get user input
	fmt.Print("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Remove newline character from input
	input = strings.TrimSuffix(input, "\n")

	// Convert input to bytes
	message := []byte(input)

	// Generate a random number k
	k, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		panic(err)
	}

	// Calculate the ephemeral public key R = k * G
	rx, ry := curve.ScalarBaseMult(k.Bytes())
	R := elliptic.Marshal(curve, rx, ry)

	// Calculate the shared secret S = (x,y) * k
	Sx, Sy := curve.ScalarMult(x, y, k.Bytes())
	S := elliptic.Marshal(curve, Sx, Sy)

	// Calculate the hash of the shared secret as the symmetric key K
	hash := Hash(S)
	K := hash[:16] // Use the first 16 bytes as the key for simplicity

	// Encrypt the message using XOR with the symmetric key K
	ciphertext := make([]byte, len(message))
	for i := range message {
		ciphertext[i] = message[i] ^ K[i%len(K)]
	}

	// Convert ciphertext and ephemeral public key to hexadecimal strings
	hexCiphertext := hex.EncodeToString(ciphertext)
	hexR := hex.EncodeToString(R)

	// Print the encrypted message and ephemeral public key
	fmt.Printf("Encrypted message: %s\n", hexCiphertext)
	fmt.Printf("Ephemeral public key: %s\n", hexR)

	// Decode the ephemeral public key and calculate the shared secret
	R, _ = hex.DecodeString(hexR)
	rx, ry = elliptic.Unmarshal(curve, R)
	Sx, Sy = curve.ScalarMult(rx, ry, privateKey)

	// Calculate the hash of the shared secret as the symmetric key K
	S = elliptic.Marshal(curve, Sx, Sy)
	hash = Hash(S)
	K = hash[:16] // Use the first 16 bytes as the key for simplicity

	// Decrypt the ciphertext using XOR with the symmetric key K
	ciphertext, _ = hex.DecodeString(hexCiphertext)
	plaintext := make([]byte, len(ciphertext))
	for i := range ciphertext {
		plaintext[i] = ciphertext[i] ^ K[i%len(K)]
	}

	// Print the decrypted plaintext
	fmt.Printf("Decrypted message: %s\n", plaintext)
}

// Hash calculates the SHA-256 hash of a byte slice
func Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

