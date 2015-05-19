package matasano

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDecryptAESInECBMode(t *testing.T) {
	file, err := os.Open("./c1_7.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(ciphertext, data)

	plaintext, err := DecryptAESInECBMode([]byte("YELLOW SUBMARINE"), ciphertext)
	result := string(plaintext)

	expectedPrefix := "I'm back and I'm ringin' the bell"

	if strings.HasPrefix(result, expectedPrefix) == false {
		t.Log(result[0:len(expectedPrefix)])
		t.Log(expectedPrefix)
		t.Error("Corrupted prefix")
	}
}

func TestEncryptAESInECBMode(t *testing.T) {
	file, err := os.Open("./c1_7.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("YELLOW SUBMARINE")

	expectedCiphertext := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(expectedCiphertext, data)

	plaintext, err := DecryptAESInECBMode(key, expectedCiphertext)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext, err := EncryptAESInECBMode(key, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(ciphertext, expectedCiphertext) == false {
		t.Fatal("Did not encrypt correctly")
	}
}
