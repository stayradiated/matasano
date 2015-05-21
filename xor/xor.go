package xor

// Slices xor's two slices together
func Slices(a, b []byte) (dst []byte) {
	// use smallest slice length
	l := len(a)
	if l > len(b) {
		l = len(b)
	}
	dst = make([]byte, l)
	for i := 0; i < l; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}

// Repeat xor's a single byte against all the bytes in a slice
func Repeat(x byte, src []byte) (dst []byte) {
	dst = make([]byte, len(src))
	for i, b := range src {
		dst[i] = x ^ b
	}
	return dst
}
