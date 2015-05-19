package matasano

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDetectSingleByteXor(t *testing.T) {
	file, err := os.Open("./c1_4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	lines := bytes.Split(data, []byte{10})

	for i, line := range lines {
		rawBytes := make([]byte, len(line)/2)
		hex.Decode(rawBytes, line)
		lines[i] = rawBytes
	}

	bestLine := DetectSingleByteXor(lines)
	expected := "Now that the party is jumping\n"

	if string(bestLine) != expected {
		t.Errorf("did not find correct solution: %v", bestLine)
	}
}
