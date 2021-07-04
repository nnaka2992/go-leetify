package leet

import (
	"testing"
)

func TestLeetCharKeys(t *testing.T) {
	var charKeys = LeetCharKeys()
	if charKeys[0] != "a" {
		t.Errorf("First value of charKeys should be 'a' but got '%s'", charKeys[0])
	}
}

func TestLeetWordKeys(t *testing.T) {
	var wordKeys = LeetWordKeys()
	if wordKeys[0] != "through" {
		t.Errorf("First value of wordKeys should be 'through' but got %s", wordKeys[0])
	}
}

func TestLeetChar(t *testing.T) {
	val, ok := LeetChar("g")
	if !ok {
		t.Error("LeetChar with argument 'g' should be ok")
	}
	if val != "6" && val != "9" {
		t.Errorf("LeetChar with argument 'g' should return '6' or '9' but got '%s'", val)
	}
	if _, ok = LeetChar("c"); ok {
		t.Error("LeetChar with argument 'c' should not be ok")
	}
}

func TestLeetWord(t *testing.T) {
	val, ok := LeetWord("through")
	if !ok {
		t.Errorf("LeetWord with argument 'throught' shold be ok: %t", ok)
	}
	if val != "thru" {
		t.Errorf("LeetWord with argument 'throught' shold return 'thru' but got '%s'", val)
	}
	if _, ok = LeetWord("value"); ok {
		t.Error("LeetWord with argument 'value' shold not be ok")
	}
}

func TestApplyLeet(t *testing.T) {
	if val := ApplyLeet("g"); val != "6" && val != "9" && val != "g" {
		t.Errorf("ApplyLeet with argument 'g' should return '6','9' or 'g' but got '%s'", val)
	}
	if val := ApplyLeet("c"); val != "c" {
		t.Errorf("LeetChar with argument 'c' should return 'c' but got '%s'", val)
	}
	if val := ApplyLeet("night"); val != "nite" {
		t.Errorf("ApplyLeet with argument 'night' should return 'nite' but got '%s'", val)
	}
	if val := ApplyLeet("eight night hate"); val != "8 nite h8" {
		t.Errorf("ApplyLeet with argument 'eight night hate' should return '8 nite h8' but got '%s'", val)
	}
}
