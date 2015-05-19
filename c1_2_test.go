package matasano

import (
	"encoding/hex"
	"testing"
)

func TestFixedXor(t *testing.T) {
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	output, _ := FixedXor(a, b)
	actual := hex.EncodeToString(output)
	expected := "746865206b696420646f6e277420706c6179"

	if actual != expected {
		t.Errorf("did not produce expected output: %v", actual)
	}
}
