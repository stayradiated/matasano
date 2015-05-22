package attack

import (
	"github.com/stayradiated/matasano/alphabet"
	"github.com/stayradiated/matasano/xor"
)

func RepeatingXOR(ciphertext []byte) (plaintext []byte, key byte, score float64) {
	for i := 0; i < 256; i++ {
		k := byte(i)
		plaintext = xor.Repeat([]byte{k}, ciphertext)

		if s := alphabet.Bytes(plaintext, alphabet.English); s > score {
			score = s
			key = k
		}
	}

	plaintext = xor.Repeat([]byte{key}, ciphertext)
	return plaintext, key, score
}
