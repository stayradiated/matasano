package matasano

import "bytes"

func DetectAESinECBMode(ciphertext []byte) bool {

	bs := 16                   // blocksize
	bc := len(ciphertext) / bs // blockcount

	// go through each block
	for i := 0; i < bc; i++ {
		block1 := ciphertext[i*bs : (i+1)*bs]

		// go through each block after block1
		for j := i + 1; j < bc; j++ {

			block2 := ciphertext[j*bs : (j+1)*bs]

			// check if block1 and block2 are equal
			if bytes.Equal(block1, block2) {
				return true
			}

		}

	}

	return false
}
