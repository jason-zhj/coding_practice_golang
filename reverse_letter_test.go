package main

import (
	"strings"
	"testing"
	"fmt"
)

func reverseOnlyLetters(S string) string {
	result := make([]string, len(S))
	start := -1
	for i, c := range S {
		if !isascii(c) {
			result[i] = string(c)
		} else {
			if start < 0 {
				start = i
			}
		}
	}

	for i:=len(S)-1; i>=0; i-- {
		if isascii(rune(S[i])) {
			result[start] = string(S[i])
			start++
			// keep finding the next available position
			for {
				if start > len(result)-1 || result[start] == "" {
					break
				}
				start++
			}
		}
	}

	return strings.Join(result, "")
}

func isascii(c rune) bool {
	return (int(c) >= 65 && int(c) <= 90) || (int(c) >= 97 && int(c) <= 122)
}

func TestReverseOnlyLetters(t *testing.T) {
	reversed := reverseOnlyLetters("a-bC-dEf-ghIj")
	fmt.Println(reversed)
}