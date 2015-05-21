package matasano

import (
	"testing"

	"github.com/stayradiated/matasano/enc/b64"
	"github.com/stayradiated/matasano/enc/hex"
)

/*

1. Convert hex to base64 and back.
==================================

The string:

  49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

should produce:

  SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

Now use this code everywhere for the rest of the exercises. Here's a
simple rule of thumb:

  Always operate on raw bytes, never on encoded strings. Only use hex
  and base64 for pretty-printing.

*/

func TestConvertHexToBase64(t *testing.T) {

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
