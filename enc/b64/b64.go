package b64

import "encoding/base64"

func Decode(s string) (dst []byte, err error) {
	return base64.StdEncoding.DecodeString(s)
}

func Encode(src []byte) (s string) {
	return base64.StdEncoding.EncodeToString(src)
}
