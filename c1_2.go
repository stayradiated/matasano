package matasano

import "errors"

func FixedXor(a, b []byte) ([]byte, error) {

	n := len(a)

	if n != len(b) {
		return nil, errors.New("both slices must have same length")
	}

	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = a[i] ^ b[i]
	}

	return out, nil
}
