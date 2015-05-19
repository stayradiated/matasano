package matasano

import (
	"errors"
	"strconv"
	"strings"
)

func HammingDistance(a, b []byte) (int, error) {
	n := len(a)
	if n != len(b) {
		return 0, errors.New("Both slices must have the same length")
	}

	distance := 0
	for i, _ := range a {
		// convert to binary and count number of 1's
		bits := strconv.FormatInt(int64(a[i]^b[i]), 2)
		distance += strings.Count(bits, "1")
	}

	return distance, nil
}

func TransposeBytes(input []byte, blocksize int) [][]byte {
	output := make([][]byte, blocksize)
	for i, _ := range output {
		output[i] = make([]byte, 0)
	}

	for i, value := range input {
		block := i % blocksize
		output[block] = append(output[block], value)
	}

	return output
}

func FindKeysize(input []byte) (bestKeysize int) {
	length := len(input)
	bestScore := float64(length)

	min := 3
	max := 40

	for keysize := min; keysize < max; keysize++ {

		if keysize*6 >= length {
			break
		}

		sliceA := input[keysize*2 : keysize*3]
		sliceB := input[keysize*3 : keysize*4]

		sliceC := input[keysize*4 : keysize*5]
		sliceD := input[keysize*5 : keysize*6]

		distAB, _ := HammingDistance(sliceA, sliceB)
		distCD, _ := HammingDistance(sliceC, sliceD)

		avgAB := float64(distAB) / float64(keysize)
		avgCD := float64(distCD) / float64(keysize)

		score := (avgAB + avgCD) / 2.0

		if score < bestScore {
			bestScore = score
			bestKeysize = keysize
		}
	}

	return bestKeysize
}

func SolveRepeatingKeyXor(input []byte) (plaintext, key []byte) {
	keysize := FindKeysize(input)
	blocks := TransposeBytes(input, keysize)

	key = make([]byte, 0)

	for _, block := range blocks {
		_, letter, _ := SolveSingleByteXor(block)
		key = append(key, letter)
	}

	plaintext = RepeatingKeyXor(key, input)

	return plaintext, key
}
