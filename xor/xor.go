package xor

// Slices xor's two slices together
func Slices(a, b []byte) (dst []byte) {
	if len(a) != len(b) {
		panic("slices must be the same length")
	}

	dst = make([]byte, len(a))
	for i, x := range a {
		y := b[i]
		dst[i] = x ^ y
	}
	return dst
}

// Repeat xor's a slice against another slice that is constantly repeated
func Repeat(b []byte, src []byte) (dst []byte) {
	dst = make([]byte, len(src))
	for i, y := range src {
		x := b[i%len(b)]
		dst[i] = x ^ y
	}
	return dst
}
