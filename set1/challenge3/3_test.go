package matasano

import (
	"bytes"
	"testing"

	"github.com/stayradiated/matasano/enc/hex"
	"github.com/stayradiated/matasano/score"
)

func Test(t *testing.T) {
	input, _ := hex.Decode("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expected := []byte("Cooking MC's like a pound of bacon")

	output, _, _ := score.SolveSingleByteXor(input)

	if bytes.Equal(output, expected) != true {
		t.Log(string(output))
		t.Log(string(expected))
		t.Error("did not solve cipher")
	}
}
