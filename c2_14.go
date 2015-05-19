package matasano

import "bytes"

// note should seed rand before using
func AESBlackBoxWithFixedKeyAndPrefix(plaintext []byte) (ciphertext []byte) {
	key := []byte("Quintessentially")
	prefix := []byte("teas up")
	message := []byte("I better have a complete analysis.")

	padded := PadBytes(append(append(prefix, plaintext...), message...), 16)

	ciphertext, _ = EncryptAESInECBMode(key, padded)
	return ciphertext
}

func ByteAtATimeDecryptWithPrefix() []byte {

	// assume blocksize
	bs := 16

	// encrypt 64 0's
	mtext := make([]byte, bs*4)
	ctext := AESBlackBoxWithFixedKeyAndPrefix(mtext)
	bc := len(ctext) / bs
	fcb := 0 // first clean block

	// look for two consecutive identical blocks
	var lastblock []byte
	for i := 0; i < bc; i++ {
		block := ctext[i*bs : (i+1)*bs]
		if bytes.Equal(block, lastblock) {
			fcb = (i - 1) * bs
			break
		}
		lastblock = block
	}

	// figure out the exact amount of bytes needed to pad the prefix
	mtext = make([]byte, bs)

	prefixsize := 0

outer:
	for {
		ctext = AESBlackBoxWithFixedKeyAndPrefix(mtext)
		bc := len(ctext) / bs

		for i := 0; i < bc; i++ {
			block := ctext[i*bs : (i+1)*bs]

			if bytes.Equal(block, lastblock) {
				prefixsize = fcb - len(mtext) + bs
				break outer
			}
		}
		mtext = append(mtext, 0)
	}

	plaintext := make([]byte, 0)

	for {
		var more bool
		plaintext, more = byteAtATimeDecryptBlockWithPrefix(plaintext, prefixsize, bs)

		if more == false {
			break
		}
	}

	return plaintext
}

func byteAtATimeDecryptBlockWithPrefix(plaintext []byte, prefixsize, bs int) ([]byte, bool) {

	// chars to test for
	chars := []byte(
		"BCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'-?,. \nA",
	)

	padsize := bs - (prefixsize % bs)
	ignore := (prefixsize/bs + 1) * bs

	padding := bytes.Repeat([]byte{170}, padsize)
	mtext := make([]byte, len(plaintext)+bs+padsize)

	copy(mtext[0:padsize], padding)
	copy(mtext[padsize+bs-1:], plaintext)

	mtextsize := len(mtext)

outer:
	for i := 0; i < bs; i++ {
		empty := make([]byte, padsize+bs-i-1)
		xtext := AESBlackBoxWithFixedKeyAndPrefix(empty)

		for _, c := range chars {
			mtext[mtextsize-1] = c
			ctext := AESBlackBoxWithFixedKeyAndPrefix(mtext)

			if bytes.Equal(xtext[ignore:mtextsize+ignore-padsize], ctext[ignore:mtextsize+ignore-padsize]) {
				if i < 15 {
					mtext = append(mtext[1:], c)
				}
				continue outer
			}
		}

		return mtext[padsize+bs-i-1 : len(mtext)-1], false
	}

	return mtext[padsize:], true
}
