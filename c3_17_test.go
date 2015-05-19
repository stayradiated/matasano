package matasano

import "testing"

func TestEncryptRandomMessage(t *testing.T) {
	c, iv := EncryptRandomMessage()
	valid := DecryptAndValidatePadding(c, iv)
	if valid == false {
		t.Fatal("Did not decrypt properly")
	}
}

func TestCBCPaddingOracle(t *testing.T) {

	ciphertext, iv := EncryptRandomMessage()
	CBCPaddingOracle(ciphertext, iv)

}
