package matasano

import (
	crand "crypto/rand"
	"math/rand"
)

// note should seed rand before using
func SimpleAESBlackBoxEncrypt(plaintext []byte) (ciphertext []byte) {

	key := make([]byte, 16)
	crand.Read(key)

	prefixsize := rand.Intn(5) + 5
	prefix := make([]byte, prefixsize)
	crand.Read(prefix)

	suffixsize := rand.Intn(5) + 5
	suffix := make([]byte, suffixsize)
	crand.Read(suffix)

	padded := PadBytes(append(append(prefix, plaintext...), suffix...), 16)

	iv := make([]byte, 16)
	crand.Read(iv)

	if rand.Intn(2) == 1 {
		ciphertext, _ = EncryptAESInCBCMode(key, iv, padded)
	} else {
		ciphertext, _ = EncryptAESInECBMode(key, padded)
	}

	return ciphertext

}
