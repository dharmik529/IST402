package main

import (
	"fmt"
	"strings"
)

func bruteForce(cipherText string) {
	for i := 1; i <= 25; i++ { // we try all possible shift values from 1 to 25
		plainText := ""
		for _, ch := range cipherText {
			if ch >= 'a' && ch <= 'z' {
				plainText += string('a' + ((ch-'a'+rune(i))%26))
			} else if ch >= 'A' && ch <= 'Z' {
				plainText += string('A' + ((ch-'A'+rune(i))%26))
			} else {
				plainText += string(ch)
			}
		}
		fmt.Printf("Shift %d: %s\n", i, plainText)
	}
}

func main() {
	cipherText := "Frph ryhu khuh Zdwvrq"
	bruteForce(strings.ToLower(cipherText)) // we convert the ciphertext to lowercase
}
