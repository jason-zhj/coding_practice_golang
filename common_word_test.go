package main

import (
	"strings"
	"testing"
)

func mostCommonWord(paragraph string, banned []string) string {
	banMap := map[string]bool{}
	for _,b := range banned {
		banMap[b] = true
	}

	words := strings.Split(strings.ToLower(paragraph), " ")
	counts := map[string]int{}
	maxWord := ""
	for _, word := range words {
		if _,ok := banMap[word]; ok {
			continue
		}
		// not banned

		if _,ok := counts[word]; !ok {
			counts[word] = 1
			if maxWord == "" {
				maxWord = word
			}
		} else {
			counts[word]++
			if counts[word] > counts[maxWord] {
				maxWord = word
			}
		}
	}

	return maxWord
}

func TestCommonWord(t *testing.T) {
	mostCommonWord("")
}
