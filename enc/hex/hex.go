package hex

import "encoding/hex"

func Decode(s string) (dst []byte, err error) {
	return hex.DecodeString(s)
}

func Encode(src []byte) (s string) {
	return hex.EncodeToString(src)
}
