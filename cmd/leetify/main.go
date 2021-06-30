package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"

	"github.com/nnaka2992/go-leetify"
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func applyLeet(word string, random bool) {
	applyLeetWord(word, random)
}

// chack whether leet word is in word or not
func applyLeetWord(word string, random bool) {
	leetWordKeys := leet.LeetWordKeys()
	maxTry := len(leetWordKeys)

	for _, leetWordKey := range leet.LeetWordKeys() {
		wordParts := strings.Split(word, leetWordKey)
		if len(wordParts) > 1 {
			for _, wordPart := range wordParts {
				if len(wordPart) != 0 {
					applyLeetWord(wordPart, random)
				} else if len(wordPart) == 0 {
					if val, err := leet.LeetWord(leetWordKey); err == nil {
						fmt.Printf(val)
					}
				}
			}
		} else {
			maxTry--
		}
	}
	if maxTry == 0 {
		applyLeetChar(word, random)
	}
}

// if the word is not leet word, then convert charcters to leet character
func applyLeetChar(word string, random bool) {
	for _, rune := range word {
		char := string(rune)
		if random {
			if randBool() {
				if val, err := leet.LeetChar(char); err != nil {
					fmt.Printf(char)
				} else {
					fmt.Printf(val)
				}
			} else {
				fmt.Printf(char)
			}
		} else {
			if val, err := leet.LeetChar(char); err != nil {
				fmt.Printf(char)
			} else {
				fmt.Printf(val)
			}
		}
	}
}

func main() {
	var random bool
	flag.BoolVar(&random, "r", false, "Randomly apply leetify")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		applyLeet(string(word), random)
		fmt.Println("")
	}
}
