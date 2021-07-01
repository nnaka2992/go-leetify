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
	var random bool
	flag.BoolVar(&random, "r", false, "Randomly apply leetify")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

    s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		new_word := leet.ApplyLeet(string(word), random)
		fmt.Println(new_word)
	}

}
