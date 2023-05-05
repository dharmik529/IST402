package main

import (
	"fmt"
	"strings"
)

// Original Alphabet
var alphabet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// 3 Random Rotors
var rotor1 = []rune{'Y', 'T', 'J', 'H', 'B', 'E', 'L', 'U', 'I', 'W', 'V', 'R', 'C', 'Z', 'Q', 'G', 'X', 'N', 'O', 'K', 'M', 'F', 'P', 'A', 'S', 'D'}
var rotor2 = []rune{'J', 'O', 'X', 'G', 'N', 'W', 'E', 'F', 'V', 'Z', 'Y', 'T', 'S', 'C', 'K', 'L', 'U', 'A', 'Q', 'D', 'P', 'I', 'B', 'R', 'M', 'H'}
var rotor3 = []rune{'Q', 'E', 'R', 'Z', 'U', 'M', 'C', 'T', 'S', 'N', 'B', 'K', 'L', 'I', 'V', 'W', 'G', 'Y', 'O', 'X', 'F', 'A', 'J', 'H', 'D', 'P'}

// Random Alpha Rotor
var reflector = []rune{'C', 'N', 'X', 'L', 'G', 'A', 'P', 'Y', 'M', 'V', 'I', 'O', 'H', 'T', 'F', 'W', 'J', 'K', 'U', 'Z', 'S', 'B', 'D', 'Q', 'E', 'R'}

func encryptPrompt(input string) {
	var (
		Key1, Key2, Key3 int
	)
	fmt.Println("First Key (1-26):")
	fmt.Scanln(&Key1)
	fmt.Println("Second Key (1-26):")
	fmt.Scanln(&Key2)
	fmt.Println("Third Key (1-26):")
	fmt.Scanln(&Key3)

	// Encrypt the Input
	var encrypted = encrypt(input, Key1, Key2, Key3)
	fmt.Println(encrypted)

	decryptPrompt(encrypted)
}

func decryptPrompt(input string) {
	input = strings.ToUpper(input)

	var (
		Key1, Key2, Key3 int
	)
	fmt.Println("First Key (1-26):")
	fmt.Scanln(&Key1)
	fmt.Println("Second Key (1-26):")
	fmt.Scanln(&Key2)
	fmt.Println("Third Key (1-26):")
	fmt.Scanln(&Key3)

	fmt.Println(decrypt(input, Key1, Key2, Key3))
}

// Encryption Function
func encrypt(str string, r1 int, r2 int, r3 int) string {
	// Input validation
	if len(str) < 1 {
		return "Input too short"
	}
	checkRotorPosition("Rotor 1", r1)
	checkRotorPosition("Rotor 2", r2)
	checkRotorPosition("Rotor 3", r3)

	rotate(rotor1, r1)
	rotate(rotor2, r2)
	rotate(rotor3, r3)

	chars := []rune(str)
	for i, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[i] = rune(str[i])
		}
	}

	// Original Input
	fmt.Println("Original:", str)

	// Encrypt the String one Character at a time
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 65 && chars[i] <= 90 {
			// Map the Letter using Rotor 1
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
			// Map the Letter using Rotor 2
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 3
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using the Reflector
			chars[i] = reflector[chars[i]-'A']
			// Map the Letter using Rotor 3
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using Rotor 2
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 1
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
		}
	}

	// Revert the Rotors
	rotate(rotor1, 48-r1)
	rotate(rotor2, 48-r2)
	rotate(rotor3, 48-r3)

	return string(chars)
}

// Decryption Function
func decrypt(str string, r1 int, r2 int, r3 int) string {
	// Input validation
	if len(str) < 1 {
		return "Invalid Input Length"
	}

	// Check that r1, r2, and r3 are integers between 1 and 26
	checkRotorPosition("Rotor 1", r1)
	checkRotorPosition("Rotor 2", r2)
	checkRotorPosition("Rotor 3", r3)

	// Rotate the Rotors 6 positions to account for the Encryption Positions
	rotate(rotor1, r1+6)
	rotate(rotor2, r2+6)
	rotate(rotor3, r3+6)


	chars := []rune(str)
	for i, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[i] = rune(str[i])
		}
	}

	// Print the Original input
	fmt.Println("Encrypted: ", str)

	// Decrypt the String one Character at a time
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] >= 65 && chars[i] <= 90 {
			rotate(rotor1, 26)
			rotorIndex := Indexing(rotor1, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotate(rotor2, 26)
			rotorIndex = Indexing(rotor2, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotate(rotor3, 26)
			rotorIndex = Indexing(rotor3, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotorIndex = Indexing(reflector, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotate(rotor3, 26)
			rotorIndex = Indexing(rotor3, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotate(rotor2, 26)
			rotorIndex = Indexing(rotor2, chars[i])
			chars[i] = alphabet[rotorIndex]
			rotate(rotor1, 26)
			rotorIndex = Indexing(rotor1, chars[i])
			chars[i] = alphabet[rotorIndex]
		}
	}
	return string(chars)
}

func Indexing(haystack []rune, needle rune) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func checkRotorPosition(rotorName string, position int) {
	if position < 1 || position > 26 {
		fmt.Println("Use a positive integer between 1-26.\n", rotorName)
	}
}

func rotate(rotorMap []rune, rotation int) {
	rotation--
	rotated := make([]rune, len(rotorMap))

	for i := range rotorMap {
		rotated[(i+rotation)%len(rotorMap)] = rotorMap[i]
	}

	copy(rotorMap, rotated)
}

func main() {

	var input string
	fmt.Println("Enter a String to Encrypt:")
	fmt.Scan(&input)
	// Convert the Input to uppercase
	input = strings.ToUpper(input)

	encryptPrompt(input)
}
