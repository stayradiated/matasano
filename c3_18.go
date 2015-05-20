package matasano

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
)

/*

18. Implement CTR, the stream cipher mode
=========================================

The string:

L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==

... decrypts to something approximating English in CTR mode, which is an AES
block cipher mode that turns AES into a stream cipher, with the following
parameters:

key=YELLOW SUBMARINE
nonce=0
format=64 bit unsigned little endian nonce,
 64 bit little endian block count (byte count / 16) CTR mode is very simple.

Instead of encrypting the plaintext, CTR mode encrypts a running counter,
producing a 16 byte block of keystream, which is XOR'd against the plaintext.

For instance, for the first 16 bytes of a message with these parameters:

keystream = AES("YELLOW SUBMARINE",
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

... for the next 16 bytes:

keystream = AES("YELLOW SUBMARINE",
	"\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00")

... and then:

keystream = AES("YELLOW SUBMARINE",
	"\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00")

CTR mode does not require padding; when you run out of plaintext, you just stop
XOR'ing keystream and stop generating keystream.

Decryption is identical to encryption. Generate the same keystream, XOR, and
recover the plaintext.

Decrypt the string at the top of this function, then use your CTR function to
encrypt and decrypt other things.

This is the only block cipher mode that matters in good code.  Most modern
cryptography relies on CTR mode to adapt block ciphers into stream ciphers,
because most of what we want to encrypt is better described as a stream than as
a sequence of blocks. Daniel Bernstein once quipped to Phil Rogaway that good
cryptosystems don't need the "decrypt" transforms. Constructions like CTR are
what he was talking about.

*/

type Nonce interface {
	Next() []byte
}

type CounterNonce struct {
	counter int64
}

func (c *CounterNonce) Next() []byte {
	buf := bytes.NewBuffer(make([]byte, 8))
	binary.Write(buf, binary.LittleEndian, c.counter)
	c.counter++
	return buf.Bytes()
}

type KeyStream struct {
	cipher cipher.Block
	nonce  Nonce
}

func NewAESKeyStream(key []byte, nonce Nonce) (*KeyStream, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	k := &KeyStream{
		cipher: cipher,
		nonce:  nonce,
	}

	return k, nil
}

func (k *KeyStream) Read(p []byte) (n int, err error) {
	blockSize := k.cipher.BlockSize()

	if len(p) < blockSize {
		return 0, nil
	}

	k.cipher.Encrypt(p, k.nonce.Next())

	return blockSize, nil
}

func XORWithAESInCTRMode(key, input []byte, nonce Nonce) (output []byte) {

	keystream, err := NewAESKeyStream(key, nonce)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(keystream)

	buf := make([]byte, 1)
	output = make([]byte, len(input))

	for i, c := range input {
		reader.Read(buf)
		output[i] = c ^ buf[0]
	}

	return output
}
