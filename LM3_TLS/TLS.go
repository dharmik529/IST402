package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"io"
	"net"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Connect to the server
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Get user input 
	fmt.Println("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	//Encrypt user input
	ciphertext, err := encrypt(input, conn)
	if err != nil {
		log.Fatalf("Failed to encrypt: %v", err)
	}
	fmt.Printf("Ciphered Text: %v\n", ciphertext)


	//Decrypt user input
	plaintext, err := decrypt(ciphertext, conn)
	if err != nil {
		log.Fatalf("Failed to decrypt: %v", err)
	}
	fmt.Printf("Plain Text: %v\n", plaintext)
}

func encrypt(plaintext string, conn net.Conn) ([]byte, error) {
	//Generate a random encryption key and write it to the server
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate ket: %v", err)
	}
	_, err = conn.Write(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to write key: %v", err)
	}

	//Create an encryption cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to create cipher: %v", err)
	}

	//Encrypt the plaintext using cipher
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("Failed to generate IV: %v", err)
    }
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return ciphertext, nil;
}

func decrypt(ciphertext []byte, conn net.Conn) (string, error) {
	//Read the encryption key from the server
	key := make([]byte, 32)
	_, err := io.ReadFull(conn, key)
	if err != nil {
		return "", fmt.Errorf("Failed to create cipher: %v", err)
	}

	// Create an encryption cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("Failed to create cipher: %v", err)
	}

	//Decrypt the ciphertext using the cipher
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return string(plaintext), nil
} 
