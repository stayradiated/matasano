package matasano

func RepeatingKeyXor(key, input []byte) []byte {
	n := len(key)
	output := make([]byte, len(input))
	for i, value := range input {
		output[i] = value ^ key[i%n]
	}
	return output
}
