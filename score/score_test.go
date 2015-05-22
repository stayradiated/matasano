package score

import (
	"bytes"
	"testing"

	"github.com/stayradiated/matasano/xor"
)

func TestBytes(t *testing.T) {
	key := []byte{40}
	message := []byte("Learn from yesterday, live for today, hope for tomorrow. The important thing is not to stop questioning.")

	encrypted := xor.Repeat(key, message)
	output, bestKey, _ := SolveSingleByteXor(encrypted)

	if bestKey != key[0] {
		t.Log(bestKey)
		t.Log(key[0])
		t.Error("TestBytes: found incorrect key")
	}

	if bytes.Equal(output, message) != true {
		t.Log(string(output))
		t.Log(string(message))
		t.Error("TestBytes: did not crack xor cipher")
	}
}
