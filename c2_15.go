package matasano

/*

PKCS#7 padding validation
=========================

Write a function that takes a plaintext, determines if it has valid PKCS#7
padding, and strips the padding off.

The string:

"ICE ICE BABY\x04\x04\x04\x04"

... has valid padding, and produces the result "ICE ICE BABY".

The string:

"ICE ICE BABY\x05\x05\x05\x05"

... does not have valid padding, nor does:

"ICE ICE BABY\x01\x02\x03\x04"

If you are writing in a language with
exceptions, like Python or Ruby, make your function throw an exception on
bad padding.

Crypto nerds know where we're going with this. Bear with us.

*/

// ValidatePKCS7 checks the padding on input is correct. It does not alter the
// input in any way.
func ValidatePKCS7(input []byte, blockSize int) bool {
	l := len(input)

	// input length must be a multiple of the blocksize.
	if l%blockSize != 0 {
		return false
	}

	lastByte := input[l-1]

	// last byte must be between 1 and blocksize (inclusive).
	if lastByte > byte(blockSize) || lastByte <= 0 {
		return false
	}

	// check that each of the padding bytes are correct
	for i := 1; byte(i) <= lastByte; i++ {
		if input[l-i] != lastByte {
			return false
		}
	}

	return true
}
