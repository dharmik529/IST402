package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "os"

    "golang.org/x/crypto/chacha20poly1305"
)

func main() {
    // Get user input
    fmt.Print("Enter a message to encrypt: ")
    message, err := readInput()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Generate a random 256-bit key
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Generate a random nonce
    nonce := make([]byte, chacha20poly1305.NonceSize)
    if _, err := rand.Read(nonce); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Encrypt the message using ChaCha20-Poly1305
    ciphertext := encrypt(message, key, nonce)

    // Print the encrypted message and key
    fmt.Printf("Encrypted message: %s\n", hex.EncodeToString(ciphertext))
    fmt.Printf("Key: %s\n", hex.EncodeToString(key))
    fmt.Printf("Nonce: %s\n", hex.EncodeToString(nonce))

    // Decrypt the message using ChaCha20-Poly1305
    plaintext := decrypt(ciphertext, key, nonce)

    // Print the decrypted message
    fmt.Printf("Decrypted message: %s\n", plaintext)
}

func readInput() ([]byte, error) {
    input := make([]byte, 1024)
    n, err := os.Stdin.Read(input)
    if err != nil {
        return nil, err
    }
    return input[:n-1], nil
}

func encrypt(plaintext, key, nonce []byte) []byte {
    // Create a new ChaCha20-Poly1305 cipher
    block, err := chacha20poly1305.New(key)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Encrypt the plaintext
    ciphertext := make([]byte, 0, len(plaintext)+block.Overhead())
    ciphertext = block.Seal(ciphertext, nonce, plaintext, nil)

    return ciphertext
}

func decrypt(ciphertext, key, nonce []byte) []byte {
    // Create a new ChaCha20-Poly1305 cipher
    block, err := chacha20poly1305.New(key)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Decrypt the ciphertext
    plaintext, err := block.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return plaintext
}

