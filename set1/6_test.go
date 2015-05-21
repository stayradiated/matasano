package matasano

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")
	distance, _ := HammingDistance(a, b)
	expected := 37
	if distance != expected {
		fmt.Println(distance)
		t.Errorf("not expected distance")
	}
}

func TestTransposeBytes(t *testing.T) {
	input := []byte{1, 2, 3, 1, 2, 3, 1, 2, 3}
	output := TransposeBytes(input, 3)

	if len(output) != 3 {
		t.Errorf("output is not correct length")
	}

	joined := bytes.Join(output, []byte{})
	expected := []byte{1, 1, 1, 2, 2, 2, 3, 3, 3}

	if bytes.Compare(joined, expected) != 0 {
		t.Errorf("did not transpose properly")
	}
}

func TestFindKeysize(t *testing.T) {
	key := []byte("neato")
	message := []byte("This is not gibberish, it just looks like it")

	ciphertext := RepeatingKeyXor(key, message)
	keysize := FindKeysize(ciphertext)

	if keysize != len(key) {
		fmt.Println(keysize)
		t.Errorf("did not find correct keysize")
	}
}

func TestSolveRepeatingKeyXor(t *testing.T) {
	file, err := os.Open("./6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(ciphertext, data)

	plaintext, key := SolveRepeatingKeyXor(ciphertext)

	expectedKey := "Terminator X: Bring the noise"
	expectedPlaintext := "I'm back and I'm ringin' the bell"

	if string(key) != expectedKey {
		fmt.Printf(string(key))
		t.Errorf("did not get correct key")
	}

	if strings.HasPrefix(string(plaintext), expectedPlaintext) == false {
		t.Errorf("did not get correct plaintext")
	}

}
