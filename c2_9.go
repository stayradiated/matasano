package matasano

func PadBytes(bytes []byte, blocksize int) []byte {

	bytesize := len(bytes)
	padsize := blocksize - (bytesize % blocksize)

	padded := make([]byte, bytesize+padsize)
	copy(padded, bytes)

	for i := 0; i < padsize; i++ {
		padded[bytesize+i] = byte(padsize)
	}

	return padded
}

func UnpadBytes(bytes []byte, blocksize int) []byte {

	bytesize := len(bytes)
	padsize := int(bytes[bytesize-1])

	if padsize > blocksize {
		panic("Invalid padding scheme")
	}

	unpadded := make([]byte, bytesize-padsize)
	copy(unpadded, bytes)

	return unpadded
}
