package matasano

import (
	"bytes"
	"testing"

	"github.com/stayradiated/matasano/enc/hex"
	"github.com/stayradiated/matasano/xor"
)

func Test(t *testing.T) {
	a, _ := hex.Decode("1c0111001f010100061a024b53535009181c")
	b, _ := hex.Decode("686974207468652062756c6c277320657965")
	e, _ := hex.Decode("746865206b696420646f6e277420706c6179")

	output := xor.Slices(a, b)

	if bytes.Equal(output, e) != true {
		t.Fatal()
	}
}
