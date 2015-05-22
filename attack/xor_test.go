package attack

import (
	"bytes"
	"testing"

	"github.com/stayradiated/matasano/xor"
)

func TestRepeatingXOR(t *testing.T) {
	key := byte(40)
	message := []byte("Learn from yesterday, live for today, hope for tomorrow. The important thing is not to stop questioning.")

	ciphertext := xor.Repeat([]byte{key}, message)
	output, foundKey, _ := RepeatingXOR(ciphertext)

	if foundKey != key {
		t.Log(foundKey)
		t.Log(key)
		t.Error("TestBytes: found incorrect key")
	}

	if bytes.Equal(output, message) != true {
		t.Log(string(output))
		t.Log(string(message))
		t.Error("TestBytes: did not crack xor cipher")
	}
}
