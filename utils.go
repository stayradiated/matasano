package matasano

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadBase64EncodedFile(fp string) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return ReadBase64EncodedBytes(data), nil
}

func ReadBase64EncodedBytes(input []byte) []byte {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	base64.StdEncoding.Decode(bytes, input)
	return bytes
}

func Block(index, blocksize int, bytes []byte) []byte {
	return bytes[index*blocksize : (index+1)*blocksize]
}

func HighlightBytes(bytes []byte, indexes ...int) {
	dark := "\x1b[31m"
	reset := "\x1b[0m"
	ptr := 0
	for _, i := range indexes {
		fmt.Printf("%s% x %s% x ",
			dark, bytes[ptr:i], reset, bytes[i:i+1])
		ptr = i + 1
	}
	if ptr < len(bytes) {
		fmt.Printf("%s% x%s", dark, bytes[ptr:], reset)
	}
	print("\n")
}
