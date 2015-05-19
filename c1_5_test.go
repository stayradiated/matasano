package matasano

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRepeatingKeyXor(t *testing.T) {
	key := "ICE"
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	output := hex.EncodeToString(RepeatingKeyXor([]byte(key), []byte(input)))
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	if output != expected {
		fmt.Println(output)
		fmt.Println(expected)
		t.Errorf("does not match expected output")
	}
}
