package xor

import (
	"bytes"
	"testing"

	"github.com/stayradiated/matasano/enc/hex"
)

func TestSlices(t *testing.T) {

	tests := map[string][]string{
		"001122334455": []string{"000000000000", "001122334455"},
		"aaaaeeeeaaaa": []string{"aabbccddeeff", "001122334455"},
	}

	for e, test := range tests {

		a, _ := hex.Decode(test[0])
		b, _ := hex.Decode(test[1])
		expected, _ := hex.Decode(e)

		actual := Slices(a, b)

		if bytes.Equal(actual, expected) != true {
			t.Logf("%x", actual)
			t.Logf("%x", expected)
			t.Fatal("XOR: Failed")
		}
	}

}

func TestRepeat(t *testing.T) {

	tests := map[string][]string{
		"0a0a":                     []string{"0a", "0000"},
		"3f0011":                   []string{"ff", "c0ffee"},
		"abcdefabcdefabcdefabcdef": []string{"abcdef", "000000000000000000000000"},
	}

	for e, test := range tests {

		b, _ := hex.Decode(test[0])
		src, _ := hex.Decode(test[1])
		expected, _ := hex.Decode(e)

		actual := Repeat(b, src)

		if bytes.Equal(actual, expected) != true {
			t.Logf("%x", actual)
			t.Logf("%x", expected)
			t.Fatal("XOR: Failed")
		}
	}

}
