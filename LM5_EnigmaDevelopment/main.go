package main

// Rotor represents a rotating wheel in the Enigma machine
type Rotor struct {
    wiring          [26]int
    turnoverIndexes []int
    offset          int
}

// RotorSet represents a set of rotors in the Enigma machine
type RotorSet struct {
    rotors []*Rotor
}

// InputRotor represents the rotor that receives input from the plugboard
type InputRotor struct {
    wiring [26]int
    offset int
}

// Reflector represents the fixed reflector in the Enigma machine
type Reflector struct {
    wiring [26]int
}

// Plugboard represents the user-configurable board that swaps pairs of letters
type Plugboard struct {
    swaps map[rune]rune
}

// EnigmaMachine represents the complete Enigma machine
type EnigmaMachine struct {
    plugboard *Plugboard
    rotorSet  *RotorSet
    reflector *Reflector
    input     *InputRotor
}
