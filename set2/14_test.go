package matasano

import "testing"

func TestByteAtATimeDecryptWithPrefix(t *testing.T) {

	result := string(ByteAtATimeDecryptWithPrefix())
	expected := "I better have a complete analysis."

	if result != expected {
		t.Log(result)
		t.Log(expected)
		t.Log([]byte(result))
		t.Log([]byte(expected))
		t.Fatal("Did not decrypt correctly")
	}

}
