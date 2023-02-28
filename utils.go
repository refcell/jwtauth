package main

import (
	"encoding/hex"
	"fmt"
)

var ErrEmptyString = fmt.Errorf("empty string")
var ErrMissingPrefix = fmt.Errorf("missing 0x prefix")

// Decode decodes a hex string with 0x prefix.
func Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrEmptyString
	}
	hexStr := input
	if !Has0xPrefix(input) {
		hexStr = "0x" + input
	}
	b, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		return nil, err
	}
	return b, err
}

// Encode encodes a byte slice to a hex string with 0x prefix.
func Encode(input []byte) string {
	return "0x" + hex.EncodeToString(input)
}

// Has0xPrefix returns true if the input string has a 0x prefix.
func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
