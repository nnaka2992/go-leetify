package leet

import (
	"errors"
	"math/rand"
	"sort"
)

var leetCharCorr = map[string][]string {
	"a":[]string{"4","aye",},
	"b":[]string{"l3","8","l3",},
	"d":[]string{"lo","l7","cl",},
	"e":[]string{"3",},
	"f":[]string{"ph"},
	"g":[]string{"6","9","gee",},
	"i":[]string{"1","l","eye","3y3",},
	"k":[]string{"x","lx",},
	"l":[]string{"1","7","i","el",},
	"m":[]string{"em","aa",},
	"o":[]string{"0","p","oh",},
	"p":[]string{"lo","9","l7"},
	"r":[]string{"l2","l9","2","12","i2",},
	"s":[]string{"5","z","ehs","es",},
	"t":[]string{"7",},
	"u":[]string{"v",},
	"w":[]string{"vv","uu",},
	"x":[]string{"ecks",},
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

func keys(m map[string]string) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {return len(ks[i]) > len(ks[j])})

	return ks
}

var leetWordKey = keys(leetWordCorr)

var keyDoesNotExist = errors.New("Key does not exist")

func LeetChar(char string) (string, error) {
	if val, ok := leetCharCorr[char]; ok {
		return val[rand.Intn(len(val))], nil
	}
	return "", keyDoesNotExist
}

func LeetWord(word string) (string, error) {
	if val, ok := leetWordCorr[word]; ok {
		return val, nil
	}
	return "", keyDoesNotExist
}

func GetLeetWordKeys() ([]string) {
	return leetWordKey
}
