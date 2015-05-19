package matasano

import (
	"math/rand"
	"testing"
)

func TestSimpleAESBlackBoxEncrypt(t *testing.T) {

	data := make([]byte, 128)

	rand.Seed(2)
	ciphertext := SimpleAESBlackBoxEncrypt(data)

	if DetectAESinECBMode(ciphertext) == false {
		t.Fatal("Incorrect")
	}

	rand.Seed(4)
	ciphertext = SimpleAESBlackBoxEncrypt(data)

	if DetectAESinECBMode(ciphertext) == true {
		t.Fatal("Incorrect")
	}

}
