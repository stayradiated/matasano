package matasano

func DetectSingleByteXor(input [][]byte) []byte {
	bestScore := float64(0)
	bestOutput := []byte{}

	for _, line := range input {
		output, _, score := SolveSingleByteXor(line)

		if score > bestScore {
			bestScore = score
			bestOutput = output
		}
	}

	return bestOutput
}
