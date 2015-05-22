package matasano

import (
	"bytes"
	"testing"
)

func TestEncryptRandomMessage(t *testing.T) {
	c, iv := EncryptBase64EncodedMessageWithAESInCBCMode("MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=")
	valid := DecryptAndValidatePadding(c, iv)
	if valid == false {
		t.Fatal("Did not decrypt properly")
	}
}

func TestCBCPaddingOracle(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	messages := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}

	for _, m := range messages {
		ciphertext, iv := EncryptBase64EncodedMessageWithAESInCBCMode(m)
		actual := CBCPaddingOracle(ciphertext, iv)
		expected := ReadBase64EncodedBytes([]byte(m))

		if bytes.Equal(expected, actual) == false {
			t.Log(string(actual))
			t.Log(string(expected))
			t.Log(actual)
			t.Log(expected)
			t.Error("Did not decrypt correctly")
		}
	}

}
