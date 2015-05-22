package matasano

import (
	"bytes"
	crand "crypto/rand"
)

/*

17. The CBC padding oracle
==========================

This is the best-known attack on modern block-cipher cryptography.

Combine your padding code and your CBC code to write two functions.

The first function should select at random one of the following 10 strings:

MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=
MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=
MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==
MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==
MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl
MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==
MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==
MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=
MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=
MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93

...generate a random AES key (which it should save for all future encryptions),
pad the string out to the 16-byte AES block size and CBC-encrypt it under that
key, providing the caller the ciphertext and IV.

The second function should consume the ciphertext produced by the first
function, decrypt it, check its padding, and return true or false depending on
whether the padding is valid.

- What you're doing here. -----------------------------------------------------
This pair of functions approximates AES-CBC encryption
as its deployed serverside in web applications; the second function models the
server's consumption of an encrypted session token, as if it was a cookie.
-------------------------------------------------------------------------------

It turns out that it's possible to decrypt the ciphertexts provided by the
first function.

The decryption here depends on a side-channel leak by the decryption function.
The leak is the error message that the padding is valid or not.

You can find 100 web pages on how this attack works, so I won't re-explain it.
What I'll say is this:

The fundamental insight behind this attack is that the byte 01h is valid
padding, and occur in 1/256 trials of "randomized" plaintexts produced by
decrypting a tampered ciphertext.

02h in isolation is not valid padding.

02h 02h is valid padding, but is much less likely to occur randomly than 01h.

03h 03h 03h is even less likely.

So you can assume that if you corrupt a decryption AND it had valid padding,
you know what that padding byte is.

It is easy to get tripped up on the fact that CBC plaintexts are "padded".
Padding oracles have nothing to do with the actual padding on a CBC plaintext.
It's an attack that targets a specific bit of code that handles decryption. You
can mount a padding oracle on any CBC block, whether it's padded or not.
*/

func EncryptBase64EncodedMessageWithAESInCBCMode(message string) (ciphertext, iv []byte) {
	key := []byte("irresponsibility")

	iv = make([]byte, 16)
	crand.Read(iv)

	plaintext := ReadBase64EncodedBytes([]byte(message))
	plaintext = PadBytes(plaintext, 16)
	ciphertext, _ = EncryptAESInCBCMode(key, iv, plaintext)

	return ciphertext, iv
}

func DecryptAndValidatePadding(ciphertext, iv []byte) bool {
	key := []byte("irresponsibility")
	plaintext, err := DecryptAESInCBCMode(key, iv, ciphertext)
	if err != nil {
		return false
	}
	return ValidatePKCS7(plaintext, 16)
}

func CBCPaddingOracle(ciphertext, iv []byte) (plaintext []byte) {

	search := func(ciphertext []byte, iv []byte) []byte {
		plaintext := make([]byte, 16)
		xor := bytes.Repeat([]byte{255}, 32)
		copy(xor[16:], ciphertext[0:16])

		for i := 15; i >= 0; i-- {
			for c := byte(0); c < 255; c++ {

				// guess char
				xor[i] = iv[i] ^ c

				// load plaintext
				for j := i + 1; j < 16; j++ {
					xor[j] = iv[j] ^ plaintext[j] ^ byte(16-i)
				}

				valid := DecryptAndValidatePadding(xor, iv)
				if valid {
					plaintext[i] = c ^ byte(16-i)
					break
				}
			}
		}

		return plaintext
	}

	bc := len(ciphertext) / 16
	plaintext = make([]byte, 0)

	for i := 0; i < bc; i++ {

		block := Block(i, 16, ciphertext)
		var before []byte

		if i == 0 {
			before = iv
		} else {
			before = Block(i-1, 16, ciphertext)
		}

		plaintext = append(plaintext, search(block, before)...)
	}

	return UnpadBytes(plaintext, 16)
}
