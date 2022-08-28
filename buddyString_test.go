package main

import (
	"testing"
	"fmt"
)

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	diffs := [][]byte{}
	m := map[byte]bool{}
	hasDuplicate := false
	for i:=0; i<len(A); i++ {
		if A[i] != B[i] {
			diffs = append(diffs, []byte{A[i],B[i]})
		} else {
			if _,ok := m[A[i]]; !ok {
				m[A[i]] = true
			} else {
				hasDuplicate = true
			}
		}
	}

	if len(diffs) == 0 {
		return hasDuplicate
	}
	if len(diffs) != 2 {
		return false
	}
	return diffs[0][0] == diffs[1][1] && diffs[0][1] == diffs[1][0]
}

func TestBuddyString(t *testing.T) {
	fmt.Println(buddyStrings("ab","ab"))
}
