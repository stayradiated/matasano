package matasano

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stayradiated/matasano/xor"
)

func TestSingleByteXor(t *testing.T) {
	key := byte(5)
	input := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := xor.Repeat(key, input)
	expected := []byte{5, 4, 7, 6, 1, 0, 3, 2, 13, 12}
	if bytes.Compare(output, expected) != 0 {
		t.Errorf("output does not equal expected: %v", output)
	}
}

func TestSolveSingleByteXor(t *testing.T) {
	expected := "It was a dark and stormy night"
	input := xor.Repeat(40, []byte(expected))
	output, _, _ := SolveSingleByteXor(input)
	if string(output) != expected {
		t.Errorf("did not solve cipher: %v", output)
	}

	expected = "Cooking MC's like a pound of bacon"
	input, _ = hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	output, _, _ = SolveSingleByteXor(input)
	if string(output) != expected {
		t.Errorf("did not solve cipher: %v", output)
	}
}
