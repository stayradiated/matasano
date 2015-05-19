package matasano

import "strings"

/*

16. CBC bit flipping
====================

Generate a random AES key.

Combine your padding code and CBC code to write two functions.

The first function should take an arbitrary input string, prepend the
string:
"comment1=cooking%20MCs;userdata="
and append the string:
";comment2=%20like%20a%20pound%20of%20bacon"

The function should quote out the ";" and "=" characters.

The function should then pad out the input to the 16-byte AES block
length and encrypt it under the random AES key.

The second function should decrypt the string and look for the
characters ";admin=true;" (or, equivalently, decrypt, split the string
on ;, convert each resulting string into 2-tuples, and look for the
"admin" tuple. Return true or false based on whether the string exists.

If you've written the first function properly, it should not be
possible to provide user input to it that will generate the string the
second function is looking for.

Instead, modify the ciphertext (without knowledge of the AES key) to
accomplish this.

You're relying on the fact that in CBC mode, a 1-bit error in a
ciphertext block:

* Completely scrambles the block the error occurs in

* Produces the identical 1-bit error (/edit) in the next ciphertext
block.

Before you implement this attack, answer this question: why does CBC
mode have this property?

*/

func CookData(userData string) (ciphertext []byte) {

	prefix := "comment1=cooking%20MCs;userdata="
	suffix := ";comment2=%20like%20a%20pound%20of%20bacon"

	userData = strings.Replace(userData, "=", "%3D", -1)
	userData = strings.Replace(userData, ";", "%3B", -1)

	plaintext := PadBytes([]byte(prefix+userData+suffix), 16)

	key := []byte("diagrammatically")
	iv := []byte("disillusionments")

	ciphertext, _ = EncryptAESInCBCMode(key, iv, plaintext)

	return ciphertext
}

func UncookData(ciphertext []byte) map[string]string {

	key := []byte("diagrammatically")
	iv := []byte("disillusionments")

	plaintext, _ := DecryptAESInCBCMode(key, iv, ciphertext)
	plaintext = UnpadBytes(plaintext, 16)

	m := make(map[string]string)

	for _, pair := range strings.Split(string(plaintext), ";") {
		data := strings.Split(pair, "=")
		if len(data) >= 2 {
			key := data[0]
			value := data[1]
			m[key] = value
		}
	}

	return m
}

func FlipTheBits() []byte {

	// what we want to set
	x := []byte(";admin=true")

	// comment1=cookin% 20MCs;userdata=0 0000000000000000 ;admin=true;comm
	// ----++++----++++ ----++++----++++ ----++++----++++ ----++++----++++

	// give us enough room to work with
	ciphertext := CookData(string(make([]byte, 27)))

	// flip those bits!
	for i, c := range x {
		ciphertext[i+32] ^= c
	}

	return ciphertext
}
