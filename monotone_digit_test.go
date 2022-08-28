package main

import (
	"testing"
	"fmt"
)

func monotoneIncreasingDigits(n int) int {
	digits := []int{}
	for {
		if n == 0 {
			break
		}

		digits = append(digits, n % 10)
		n /= 10
	}

	if len(digits) == 1 {
		return n
	}

	first := digits[len(digits)-1]
	val, ok := search(digits, len(digits)-2, first, first, false)
	if !ok {
		return first * tens(len(digits)-1) - 1
	}
	return val
}

func tens(numDigits int) int {
	prod := 1
	for i:=0; i<numDigits; i++ {
		prod *= 10
	}
	return prod
}

func search(digits []int, i int, prefix int, lastDigit int, alreadySmaller bool) (int, bool) {
	if alreadySmaller {
		return prefix * tens(i+1) + tens(i+1) - 1, true
	}

	if i < 0 {
		return prefix, true
	}

	// not already smaller
	if digits[i] < lastDigit {
		return 0, false
	}

	for d:=digits[i]; d >= lastDigit; d-- {
		smaller := d < digits[i]
		if val, ok := search(digits, i-1, prefix*10+d, d, smaller); ok {
			return val, true
		}
	}

	return 0, false
}

func TestDigits(t *testing.T) {
	fmt.Println(monotoneIncreasingDigits(10))
}
