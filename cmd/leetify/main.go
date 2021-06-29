package main

import (
	"bufio"
	"flag"
	"os"
	"math/rand"
	"time"

	"github.com/nnaka2992/go-leetify"
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func applyLeet(word string) string {
	val, err := leet.leetWord(word)
	if err != nil {
		str  := ""
		for _, char := range word {
			if val, err = leet.leetChar(char); err != nil {
				str += val
			}
			str += string(char)
		}
		return str
	}
	return val
}

func main() {
	var random bool
	flag.BoolVar(&random, "r", false, "Randomly apply leetify")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if !random && !randBool() {
			continue
		}
		applyLeet(string(word))
	}
}
