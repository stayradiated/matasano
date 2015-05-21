package matasano

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestCounterNonce(t *testing.T) {
	nonce := &CounterNonce{}

	tests := [][]byte{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0},
	}

	for _, expected := range tests {
		n := nonce.Next()
		if bytes.Equal(n, expected) == false {
			t.Log(n)
			t.Log(expected)
			t.Error("Incorrect nonce")
		}
	}
}

func TestAESKeyStream(t *testing.T) {

	ks, err := NewAESKeyStream([]byte("YELLOW SUBMARINE"), &CounterNonce{})
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 16)

	tests := [][]byte{
		[]byte{118, 209, 203, 75, 175, 162, 70, 226, 227, 175, 3, 93, 108, 19, 195, 114},
		[]byte{210, 236, 108, 220, 152, 109, 18, 222, 207, 218, 31, 147, 175, 238, 115, 24},
		[]byte{45, 160, 142, 203, 17, 123, 55, 75, 195, 218, 183, 38, 178, 252, 132, 205},
		[]byte{193, 128, 171, 53, 73, 250, 110, 85, 209, 76, 102, 103, 201, 111, 165, 176},
	}

	for _, expected := range tests {
		n, err := ks.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		if n != 16 {
			t.Log(n)
			t.Fatal("Incorrect byte length")
		}

		if bytes.Equal(buf, expected) == false {
			t.Logf("%#v", buf)
			t.Log(expected)
			t.Error("Incorrect ciphertext")
		}
	}
}

func TestXORWithAESInCTRMode(t *testing.T) {

	ciphertext, _ := base64.StdEncoding.DecodeString("L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")

	key := []byte("YELLOW SUBMARINE")

	plaintext := XORWithAESInCTRMode(key, ciphertext, &CounterNonce{})

	expected := []byte("Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby ")

	if bytes.Equal(plaintext, expected) == false {
		t.Log(string(plaintext))
		t.Log(string(expected))
		HighlightBytes(ciphertext)
		HighlightBytes(plaintext)
		HighlightBytes(expected)
		t.Error("Decode failed")
	}

}
