package matasano

import (
	"bytes"
	"testing"
)

func TestPadBytes(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4}
	expected := []byte{0, 1, 2, 3, 4, 3, 3, 3}

	padded := PadBytes(data, 8)

	if bytes.Equal(padded, expected) == false {
		t.Fatal("Could not pad bytes")
	}
}

func TestUnpadBytes(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 3, 3, 3}
	expected := []byte{0, 1, 2, 3, 4}

	unpadded := UnpadBytes(data, 8)

	if bytes.Equal(unpadded, expected) == false {
		t.Fatal("Could not unpad bytes")
	}
}
