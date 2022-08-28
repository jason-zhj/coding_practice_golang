package main

import (
	"strings"
	"testing"
	"fmt"
)

var m = map[string]bool{
	"a":true,
	"e":true,
	"i":true,
	"o":true,
	"u":true,
}

func process(word string, index int) string {
	word += "ma"
	for i:=0;i<index+1;i++{
		word += "a"
	}
	if !isVowel(word[0]) {
		return word[1:]
	}
	return word
}

func isVowel(c byte) bool {
	lc := strings.ToLower(string(c))
	_,ok := m[lc]
	return ok
}

func TestGoatLatin(t *testing.T) {
	fmt.Println(process("speak",1))
}
