package matasano

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stayradiated/matasano/attack"
	"github.com/stayradiated/matasano/enc/hex"
)

func Test(t *testing.T) {
	file, err := os.Open("./4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	lines := bytes.Split(data, []byte{'\n'})

	for i, line := range lines {
		lines[i], _ = hex.Decode(string(line))
	}

	plaintext := DetectSingleByteXor(lines)
	expected := []byte("Now that the party is jumping\n")

	if bytes.Equal(plaintext, expected) != true {
		t.Log(string(plaintext))
		t.Log(string(expected))
		t.Error("did not find the correct ciphertext")
	}
}

func DetectSingleByteXor(input [][]byte) (plaintext []byte) {
	plaintext = []byte{}
	score := float64(0)

	for _, line := range input {
		if p, _, s := attack.RepeatingByteXOR(line); s > score {
			score = s
			plaintext = p
		}
	}

	return plaintext
}
