package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"math/rand"
	"time"

	"github.com/nnaka2992/go-leetify"
)

func main() {
	showInput := flag.Bool("s", false, "Show input")
	showCharList := flag.Bool("list-characters", false, "List characters' conversion table")
	showWordList := flag.Bool("list-words", false, "List words' conversion table")
	applyChar := flag.Bool("c", true, "Change characters")
	applyWord := flag.Bool("w", true, "Change words")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	if *showCharList || *showWordList {
		if *showCharList {
			leet.PrintChatTable()
		}
		if *showWordList {
			leet.PrintWordTable()
		}
		os.Exit(0)
	}

	if !(*applyChar || *applyWord) {
		fmt.Fprintf(os.Stderr, "One of c or w must be true")
		os.Exit(1)
	}

    s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := string(s.Text())
		newWord := ""
		if *applyChar != *applyWord {
			if *applyChar {
				newWord = leet.ApplyChar(word)
			} else if *applyWord {
				newWord = leet.ApplyWord(word)
			}
		} else {
			newWord = leet.ApplyLeet(word)
		}

		if *showInput {
			fmt.Printf("%v: %v\n", word, newWord)
		} else {
			fmt.Println(newWord)
		}
	}

}
