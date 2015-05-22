package score

import "github.com/stayradiated/matasano/xor"

// Bytes tests how well some bytes matches an alphabet
func Bytes(input []byte, alphabet Alphabet) float64 {
	frequency := make(map[byte]float64)

	// count number of repeating bytes
	for _, letter := range input {
		if alphabet.Contains(letter) {
			frequency[letter] += 1
		}
	}

	// convert counts into
	n := float64(len(frequency))
	score := float64(0)

	for letter, count := range frequency {
		optimal := alphabet.Frequency(letter)
		score -= count - (optimal * n)
	}

	return score
}

func SolveSingleByteXor(input []byte) ([]byte, byte, float64) {
	bestScore := float64(0)
	bestKey := make([]byte, 1)

	for i := 0; i < 256; i++ {
		key := []byte{byte(i)}
		decoded := xor.Repeat(key, input)
		score := Bytes(decoded, English)

		if score > bestScore {
			bestScore = score
			bestKey[0] = byte(i)
		}
	}

	return xor.Repeat(bestKey, input), bestKey[0], bestScore
}
