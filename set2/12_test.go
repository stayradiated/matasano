package matasano

import (
	"bytes"
	"fmt"
	"testing"
)

func TestByteAtATimeDecrypt(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	plaintext := ByteAtATimeDecrypt()

	expected := []byte(`Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by

`)

	if bytes.Equal(plaintext, expected) == false {
		fmt.Println(plaintext)
		fmt.Println(expected)
		t.Fatal("Bad decrypt")
	}

}
