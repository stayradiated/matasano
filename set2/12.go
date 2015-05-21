package matasano

import (
	"bytes"
	"encoding/base64"
)

// note should seed rand before using
func AESBlackBoxWithFixedKey(plaintext []byte) (ciphertext []byte) {

	// unknown to the attacker
	key := []byte("INCOMPREHENSIBLE")

	// unkown to the attacker
	message := []byte("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

	suffix := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(suffix, message)

	padded := PadBytes(append(plaintext, suffix...), 16)

	ciphertext, _ = EncryptAESInECBMode(key, padded)

	return ciphertext
}

func ByteAtATimeDecrypt() []byte {

	bs := 0
	ctextsize := 0
	mtext := make([]byte, 0)

	// figure out blocksize
	for {
		ctext := AESBlackBoxWithFixedKey(mtext)

		if ctextsize > 0 && len(ctext) > ctextsize {
			bs = len(ctext) - ctextsize
			break
		}

		ctextsize = len(ctext)
		mtext = append(mtext, 0)
	}

	// detect ecb
	mtext = make([]byte, bs*4)
	ctext := AESBlackBoxWithFixedKey(mtext)

	if DetectAESinECBMode(ctext) == false {
		panic("black box is not using ECB!")
	}

	plaintext := make([]byte, 0)

	for {
		var more bool
		plaintext, more = byteAtATimeDecryptBlock(plaintext, bs)

		if more == false {
			break
		}
	}

	return plaintext
}

func byteAtATimeDecryptBlock(plaintext []byte, bs int) ([]byte, bool) {
	// chars to test for
	chars := []byte(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'-?,. \n",
	)

	mtext := make([]byte, len(plaintext)+bs)
	copy(mtext[bs-1:], plaintext)
	size := len(mtext)

outer:
	for i := 0; i < bs; i++ {
		empty := make([]byte, bs-i-1)
		xtext := AESBlackBoxWithFixedKey(empty)

		for _, c := range chars {
			mtext[size-1] = c
			ctext := AESBlackBoxWithFixedKey(mtext)

			if bytes.Equal(xtext[0:size], ctext[0:size]) {
				if i < 15 {
					mtext = append(mtext[1:], c)
				}
				continue outer
			}
		}

		return mtext[bs-i-1:], false
	}

	return mtext, true
}
