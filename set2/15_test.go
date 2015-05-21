package matasano

import "testing"

func TestValidatePKCS7(t *testing.T) {

	tests := map[string]bool{
		"ICE ICE BABY\x04\x04\x04\x04": true,
		"ICE ICE BABY...\x01":          true,
		"ICE ICE BABY\x05\x05\x05\x05": false,
		"ICE ICE BABY\x01\x02\x03\x04": false,
		"ICE ICE BABY...\x00":          false,
		"ICE ICE BABY....":             false,

		"ICE ICE ICE BABY\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10":    true,
		"ICE ICE ICE BAB\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11": false,
	}

	for s, e := range tests {
		a := ValidatePKCS7([]byte(s), 16)
		if a != e {
			t.Errorf("Failed validation: %s. Expected: %t. Actual: %t", s, e, a)
		}
	}

}
