package leet

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var charTable = map[string][]string {
	"a":[]string{"4",},
	"b":[]string{"8",},
	"e":[]string{"3",},
	"g":[]string{"6","9",},
	"i":[]string{"1","l",},
	"k":[]string{"x","lx",},
	"l":[]string{"1","7","i",},
	"o":[]string{"0","p",},
	"p":[]string{"lo","9",},
	"r":[]string{"2",},
	"s":[]string{"5","z",},
	"t":[]string{"7",},
	"u":[]string{"v",},
	"y":[]string{"j","7",},
	"z":[]string{"2","s",},
}

var wordTable = map[string]string {
	"one":"1",
	"won":"1",
	"two":"2",
	"too":"2",
	"to":"2",
	"three":"3",
	"four":"4",
	"for":"4",
	"fore":"4",
	"five":"5",
	"six":"6",
	"seven":"7",
	"eight":"8",
	"ate":"8",
	"nine":"9",
	"be":"b",
	"bee":"b",
	"tea":"t",
	"see":"c",
	"are":"r",
	"you":"u",
	"why":"y",
	"owe":"o",
	"love":"luv",
	"through":"thru",
	"night":"nite",
	"easy":"ez",
}

func initWordKeys(m map[string]string) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {return ks[i] < ks[j]})
	sort.Slice(ks, func(i, j int) bool {return len(ks[i]) > len(ks[j])})


	return ks
}

func initCharKeys(m map[string][]string) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {return ks[i] < ks[j]})

	return ks
}

var charKeys = initCharKeys(charTable)
var wordKeys = initWordKeys(wordTable)

func LeetChar(char string) (string, bool) {
	val, ok := charTable[char]
	if !ok {
		return "", ok
	}
	return val[rand.Intn(len(val))], ok
}

func LeetWord(word string) (string, bool) {
	word, ok := wordTable[word]
	return word, ok
}

func LeetCharKeys() []string {
	return charKeys
}

func LeetWordKeys() []string {
	return wordKeys
}

func randBool() bool {
	return rand.Intn(5) == 0
}

func ApplyLeet(str string) string {
	newStr := ""
	for ; 0 < len(str); {
		if substr, val, ok := applyWord(str); ok {
			str = substr
			newStr += val
			continue
		}

		substr, val, _ := applyChar(str)
		str = substr
		newStr += val

	}
	return newStr
}

func ApplyWord(str string) string {
	newStr := ""
	for ; 0 < len(str); {
		if substr, val, ok := applyWord(str); ok {
			str = substr
			newStr += val
			continue
		}
		newStr += str[:1]
		str = str[1:]
	}
	return newStr
}

func ApplyChar(str string) string {
	newStr := ""
	for ; 0 < len(str); {
		substr, val, _ := applyChar(str)
		str = substr
		newStr += val
	}
	return newStr
}

func applyWord (str string) (string, string, bool){
	for i := 0; i < len(wordKeys); i++ {
		if strings.HasPrefix(str, wordKeys[i]) {
			if word, ok := LeetWord(wordKeys[i]); ok {
				return str[len(wordKeys[i]):], word, true
			}
		}
	}
	return str, "",  false
}

func applyChar(str string) (string, string, bool) {
	char, ok := LeetChar(str[:1])
	if ok && randBool() {
		return str[1:], char, true
	} else {
		return str[1:], str[:1], false
	}
}

func PrintChatTable(){
	for key, value := range charTable {
		fmt.Printf("%s\t: %s\n", key, value)
	}
}


func PrintWordTable(){
	for key, value := range wordTable {
		fmt.Printf("%s\t: %s\n", key, value)
	}
}
