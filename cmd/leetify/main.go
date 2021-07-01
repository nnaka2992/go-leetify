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
	random := flag.Bool("r", false, "Randomly apply leetify")
	show_input := flag.Bool("s", false, "Show input")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

    s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := string(s.Text())
		new_word := leet.ApplyLeet(word, *random)
		if *show_input {
			fmt.Printf("%v: %v\n", word, new_word)
		} else {
			fmt.Println(new_word)
		}
	}

}
