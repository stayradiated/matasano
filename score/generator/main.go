package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/stayradiated/matasano/score"
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
	dict := make(map[byte]float64)
	data := make([]byte, 1024)

	for {
		n, err := reader.Read(data)

		if n > 0 {
			for _, b := range data[:n] {
				if _, ok := dict[b]; ok != true {
					dict[b] = 0
				}
				dict[b] += 1
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

	alphabet := make(score.Alphabet, 0)
	for b, v := range dict {
		alphabet = append(alphabet, score.Letter{
			Value:     b,
			Frequency: v / total * 100,
		})
	}

	sort.Sort(alphabet)

	for _, letter := range alphabet {
		fmt.Printf("{%s, %.5f},\n", strconv.QuoteRune(rune(letter.Value)), letter.Frequency)
	}
}
