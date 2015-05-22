package alphabet

import "testing"

func TestBytes(t *testing.T) {
	english := []byte("Learn from yesterday, live for today, hope for tomorrow.")
	gibberish := []byte("!@#$%^&*()1234567890-=+_:{}[]|<>,.?/")

	if Bytes(english, English) < Bytes(gibberish, English) {
		t.Error("Gibberish got a higher score than english")
	}
}
