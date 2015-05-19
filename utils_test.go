package matasano

import (
	"encoding/base64"
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

	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(bytes, data)

	return bytes, nil
}
