package matasano

import "crypto/aes"

func DecryptAESInECBMode(key, ciphertext []byte) (plaintext []byte, err error) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := cipher.BlockSize()
	count := len(ciphertext) / bs

	plaintext = make([]byte, len(ciphertext))

	for i := 0; i < count; i++ {
		cipher.Decrypt(
			plaintext[i*bs:(i+1)*bs],
			ciphertext[i*bs:(i+1)*bs],
		)
	}

	return plaintext, nil
}

func EncryptAESInECBMode(key, plaintext []byte) (ciphertext []byte, err error) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := cipher.BlockSize()
	count := len(plaintext) / bs

	ciphertext = make([]byte, len(plaintext))

	for i := 0; i < count; i++ {
		cipher.Encrypt(
			ciphertext[i*bs:(i+1)*bs],
			plaintext[i*bs:(i+1)*bs],
		)
	}

	return ciphertext, nil
}
