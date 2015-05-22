package matasano

import (
	"bytes"

	"github.com/stayradiated/matasano/xor"
)

/*

3. Single-character XOR Cipher

The hex encoded string:

	1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

... has been XOR'd against a single character. Find the key, decrypt
the message.

Write code to do this for you. How? Devise some method for "scoring" a
piece of English plaintext. (Character frequency is a good metric.)
Evaluate each output and choose the one with the best score.

Tune your algorithm until this works.

*/

type Letter struct {
	Value     byte
	Frequency float64
}

type Alphabet []Letter

func (a Alphabet) Contains(value byte) bool {
	for _, letter := range a {
		if letter.Value == value {
			return true
		}
	}
	return false
}

func (a Alphabet) Frequency(value byte) float64 {
	for _, letter := range a {
		if letter.Value == value {
			return letter.Frequency
		}
	}
	return 0
}

var English = Alphabet{
	{'\n', 15},
	{' ', 15},
	{'E', 12.02},
	{'T', 9.10},
	{'A', 8.12},
	{'O', 7.68},
	{'I', 7.31},
	{'N', 6.95},
	{'S', 6.28},
	{'R', 6.02},
	{'H', 5.92},
	{'D', 4.32},
	{'L', 3.98},
	{'U', 2.88},
	{'C', 2.71},
	{'M', 2.61},
	{'F', 2.30},
	{'Y', 2.11},
	{'W', 2.09},
	{'G', 2.03},
	{'P', 1.82},
	{'B', 1.49},
	{'V', 1.11},
	{'K', 0.69},
	{'X', 0.17},
	{'Q', 0.11},
	{'J', 0.10},
	{'Z', 0.07},
}

// ScoreEnglish tests how well some bytes matches english
func ScoreEnglish(input []byte) float64 {
	input = bytes.ToUpper(input)
	frequency := make(map[byte]float64)

	// count number of repeating bytes
	for _, letter := range input {
		if English.Contains(letter) {
			frequency[letter] += 1
		}
	}

	// convert counts into
	n := float64(len(frequency))
	score := float64(0)

	for letter, count := range frequency {
		optimal := English.Frequency(letter)
		score -= count - (optimal * n)
	}

	return score
}

func SolveSingleByteXor(input []byte) ([]byte, byte, float64) {
	bestScore := float64(0)
	bestKey := byte(0)

	for i := 0; i < 256; i++ {
		decoded := xor.Repeat(byte(i), input)
		score := ScoreEnglish(decoded)

		if score > bestScore {
			bestScore = score
			bestKey = byte(i)
		}
	}

	return xor.Repeat(bestKey, input), bestKey, bestScore
}
