package leet

import (
	"math/rand"
	"sort"
	"strings"
)

var leetCharCorr = map[string][]string {
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

var leetWordCorr = map[string]string {
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

var charKeys = initCharKeys(leetCharCorr)
var wordKeys = initWordKeys(leetWordCorr)

func LeetChar(char string) (string, bool) {
	val, ok := leetCharCorr[char]
	if !ok {
		return "", ok
	}
	return val[rand.Intn(len(val))], ok
}

func LeetWord(word string) (string, bool) {
	word, ok := leetWordCorr[word]
	return word, ok
}

func LeetCharKeys() []string {
	return charKeys
}

func LeetWordKeys() []string {
	return wordKeys
}

func randBool() bool {
	return rand.Intn(2) == 0
}

func ApplyLeet(word string) string{
	new_str := applyLeet(word)
	return new_str
}

func applyLeet(str string) string {
	applyWord := func (str string, new_str *string) (string, bool){
		for i := 0; i < len(wordKeys); i++ {
			if strings.HasPrefix(str, wordKeys[i]) {
				if word, ok := LeetWord(wordKeys[i]); ok {
					*new_str += word
				}
				return str[len(wordKeys[i]):], true
			}
		}
		return str, false
	}

	new_str := ""
	for ; 0 < len(str); {
		if substr, ok := applyWord(str, &new_str); ok {
			str = substr
			continue
		}

		if char, ok := LeetChar(str[:1]); ok{
			new_str += char
		} else {
			new_str += str[:1]
		}
		str = str[1:]
	}
	return new_str
}


