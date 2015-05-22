package matasano

import (
	"testing"

	"github.com/stayradiated/matasano/enc/b64"
	"github.com/stayradiated/matasano/enc/hex"
)

func Test(t *testing.T) {

	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	bytes, err := hex.Decode(input)
	if err != nil {
		t.Fatal(err)
	}

	actual := b64.Encode(bytes)

	if actual != expected {
		t.Errorf("\n%s\n%s", expected, actual)
	}
}
