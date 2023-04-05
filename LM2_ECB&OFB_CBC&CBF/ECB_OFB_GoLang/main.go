package main

import (
    "bytes"
    "crypto/cipher"
    "crypto/des"
    "fmt"
)

//These are two helper functions that add padding to plaintext 
//and remove padding from ciphertext, respectively. Padding is 
//used to ensure that the plaintext is a multiple of the block size, 
//which is 8 bytes for DES.

func pad(plaintext []byte, blockSize int) []byte {
    padding := blockSize - (len(plaintext) % blockSize)
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(plaintext, padtext...)
}

func unpad(padded []byte) []byte {
    padding := int(padded[len(padded)-1])
    return padded[:len(padded)-padding]
}

func main() {
    // ECB example
    key := []byte("01234567") // 8-byte key
    plaintext := []byte("Hello world!")
    fmt.Println("Plaintext:", string(plaintext))

    plaintext = pad(plaintext, des.BlockSize)
    block, err := des.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, len(plaintext))
    for i := 0; i < len(plaintext); i += block.BlockSize() {
        block.Encrypt(ciphertext[i:], plaintext[i:])
    }
    fmt.Printf("ECB encrypted: %x\n", ciphertext)

    decryptedtext := make([]byte, len(ciphertext))
    for i := 0; i < len(ciphertext); i += block.BlockSize() {
        block.Decrypt(decryptedtext[i:], ciphertext[i:])
    }
    decryptedtext = unpad(decryptedtext)
    fmt.Println("ECB decrypted:", string(decryptedtext))

    // OFB example
    iv := []byte("76543210") // 8-byte initialization vector
    ofb := cipher.NewOFB(block, iv)

    plaintext = []byte("Hello world!")
    fmt.Println("Plaintext:", string(plaintext))

    plaintext = pad(plaintext, des.BlockSize)
    ciphertext = make([]byte, len(plaintext))
    ofb.XORKeyStream(ciphertext, plaintext)
    fmt.Printf("OFB encrypted: %x\n", ciphertext)

    decryptedtext = make([]byte, len(ciphertext))
    ofb = cipher.NewOFB(block, iv)
    ofb.XORKeyStream(decryptedtext, ciphertext)
    decryptedtext = unpad(decryptedtext)
    fmt.Println("OFB decrypted:", string(decryptedtext))
}

