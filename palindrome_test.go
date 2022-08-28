package main

import (
	"strings"
	"testing"
	"fmt"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for {
		if i>=j {
			break
		}

		for {
			if i<len(s) && !isAlphaNumeric(s[i]) {
				i++
			} else {
				break
			}
		}

		for {
			if j >= 0 && !isAlphaNumeric(s[j]) {
				j--
			} else {
				break
			}
		}

		if i>=j {
			break
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c>='0' && c<='9') || (c>='a' && c<='z') || (c>='A' && c<='Z')
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("race a car"))
}
