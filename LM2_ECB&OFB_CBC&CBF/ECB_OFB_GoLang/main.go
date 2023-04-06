package main

import "fmt"

var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var codebook2 = [4][2]int{{0b00, 0b010}, {0b01, 0b00}, {0b10, 0b01}, {0b11, 0b11}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}

func codebookLookup(xor int) int {
	for i := 0; i < len(codebook); i++ {
		if codebook[i][0] == xor {
			return codebook[i][1]
		}
	}
	return 0
}

func codebookReverseLookup(value int) int {
	for i := 0; i < len(codebook); i++ {
		if codebook[i][1] == value {
			return codebook[i][0]
		}
	}
	return 0
}

func main() {
	fmt.Println("Codebook page 1:")
	for i, m := range message {
		cipheredValue := codebookLookup(m)
		fmt.Printf("Block %d ciphered: %b -> %b\n", i+1, m, cipheredValue)
	}
	fmt.Println("Decrypting...")
	for i, c := range codebook2 {
		decryptedValue := codebookReverseLookup(c[1])
		fmt.Printf("Block %d decrypted: %b -> %b\n", i+1, c[0], decryptedValue)
	}
}

