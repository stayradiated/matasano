package matasano

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestDetectAESinECBMode(t *testing.T) {

	file, err := os.Open("./c1_8.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	lines := bytes.Split(data, []byte{10})

	matches := 0

	for _, line := range lines {

		ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(line)))
		base64.StdEncoding.Decode(ciphertext, line)

		isECB := DetectAESinECBMode(ciphertext)

		if isECB {
			matches += 1
		}

	}

	if matches != 1 {
		t.Fatal("Incorrect number of matches found")
	}

}
