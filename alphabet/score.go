package alphabet

// Bytes tests how well some bytes matches an alphabet
func Bytes(input []byte, alphabet Alphabet) (score float64) {
	frequency := make(map[byte]float64)

	// count number of bytes that are in the alphabet
	for _, letter := range input {
		if _, ok := alphabet[letter]; ok == true {
			frequency[letter] += 1
		}
	}

	l := float64(len(frequency))

	// score based on how well the byte frequency matches the alphabet
	for letter, actual := range frequency {
		optimal := alphabet[letter]
		score -= actual - (optimal * l)
	}

	return score
}
