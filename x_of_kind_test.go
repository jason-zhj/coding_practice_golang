package main

import (
	"sort"
	"testing"
	"fmt"
)

func hasGroupsSizeX(deck []int) bool {
	sort.Slice(deck, func(i,j int) bool {
		return deck[i] < deck[j]
	})

	x := 0
	lastIndex := 0

	for i:=1; i<len(deck); i++ {
		if x == 0 { // x not determined
			if deck[i] != deck[i-1] {
				if i < 2 {
					return false
				}
				x = i
				lastIndex = i
			}
		} else { // x already determined
			if deck[i] != deck[i-1] {
				if (i - lastIndex) % x != 0 {
					return false
				}
				lastIndex = i
			}
		}
	}

	if x == 0 {
		return len(deck) >= 2
	}

	return (len(deck) - lastIndex) % x == 0
}

func TestXOfKind(t *testing.T) {
	fmt.Println(hasGroupsSizeX([]int{1,2,3,4,4,3,2,1}))
}
