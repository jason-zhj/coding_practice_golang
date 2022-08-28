package main

import (
	"strings"
	"testing"
	"fmt"
)

func reverseVowels(s string) string {
	vowels := []string{}
	for i:=len(s)-1; i>=0; i-- {
		if vowel(s[i]) {
			vowels = append(vowels, string(s[i]))
		}
	}

	v := 0
	chs := strings.Split(s, "")
	for i, ch := range chs {
		if vowel(byte(ch[0])) {
			chs[i] = vowels[v]
			v++
		}
	}

	return strings.Join(chs, "")
}

func vowel(c byte) bool {
	return c == 'a' || c == 'o' || c == 'i' || c == 'u' || c == 'e'
}

func TestReverseVowels(t *testing.T) {
	fmt.Println(reverseVowels("hello"))
}