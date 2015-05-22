package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var path string

	flag.StringVar(&path, "p", "", "path to a text file which will be used to generate the alphabet")
	flag.Parse()

	if path == "" {
		panic("Must specify path")
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	total := float64(0)
	alphabet := make(map[byte]float64)
	data := make([]byte, 1024)

	for {
		n, err := reader.Read(data)

		if n > 0 {
			for _, b := range data[:n] {
				alphabet[b] += 1
				total += 1
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			break
		}
	}

	for key, value := range alphabet {
		alphabet[key] = value / total * 100
	}

	for key, value := range alphabet {
		fmt.Printf("%s: %.5f,\n", strconv.QuoteRune(rune(key)), value)
	}
}
