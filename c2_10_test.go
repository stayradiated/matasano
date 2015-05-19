package matasano

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecryptAESInCBCMode(t *testing.T) {
	ciphertext, err := ReadBase64EncodedFile("./c2_10.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("YELLOW SUBMARINE")
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	plaintext, err := DecryptAESInCBCMode(key, iv, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	expectedPrefix := "I'm back and I'm ringin' the bell"
	if strings.HasPrefix(string(plaintext), expectedPrefix) == false {
		t.Fatal("Invalid decryption")
	}

}

func TestEncryptAESInCBCMode(t *testing.T) {
	expectedCiphertext, err := ReadBase64EncodedFile("./c2_10.txt")
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("YELLOW SUBMARINE")
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	plaintext, err := DecryptAESInCBCMode(key, iv, expectedCiphertext)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext, err := EncryptAESInCBCMode(key, iv, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(ciphertext, expectedCiphertext) == false {
		t.Fatal("Encryption failed")
	}
}
