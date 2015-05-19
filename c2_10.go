package matasano

import (
	"crypto/aes"
	"errors"
)

func EncryptAESInCBCMode(key, iv, plaintext []byte) (ciphertext []byte, err error) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := cipher.BlockSize()

	if len(plaintext)%bs != 0 {
		return nil, errors.New("Plaintext length must be a multiple of blocksize")
	}

	ciphertext = make([]byte, len(plaintext))
	count := len(plaintext) / bs

	lastblock := iv

	for i := 0; i < count; i++ {

		nextblock, err := FixedXor(plaintext[i*bs:(i+1)*bs], lastblock)
		if err != nil {
			return nil, err
		}

		cipher.Encrypt(
			ciphertext[i*bs:(i+1)*bs],
			nextblock,
		)

		lastblock = ciphertext[i*bs : (i+1)*bs]

	}

	return ciphertext, nil
}

func DecryptAESInCBCMode(key, iv, ciphertext []byte) (plaintext []byte, err error) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := cipher.BlockSize()

	if len(ciphertext)%bs != 0 {
		return nil, errors.New("Ciphertext length must be a multiple of blocksize")
	}

	plaintext = make([]byte, len(ciphertext))
	count := len(plaintext) / bs

	// used as a scratchpad during decrypting
	block := make([]byte, bs)

	for i := count - 1; i >= 0; i-- {

		var nextblock []byte
		if i == 0 {
			nextblock = iv
		} else {
			nextblock = ciphertext[(i-1)*bs : i*bs]
		}

		cipher.Decrypt(
			block,
			ciphertext[i*bs:(i+1)*bs],
		)

		block, err = FixedXor(block, nextblock)
		if err != nil {
			return nil, err
		}

		copy(
			plaintext[i*bs:(i+1)*bs],
			block,
		)

	}

	return plaintext, nil
}
