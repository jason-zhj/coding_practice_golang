package main

import (
	"testing"
	"fmt"
)

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	if k == 0 {
		return result
	}

	// if k < 0
	if k < 0 {
		sum := 0
		for i:=0; i< -k; i++ {
			sum += code[i]
		}

		start := -k
		result[start] = sum
		i := 0
		for {
			if i == len(code) - 1 {
				break
			}

			first := i
			next := (start + i) % len(code)
			sum = sum - code[first] + code[next]
			result[(start + i + 1) % len(code)] = sum
			i++
		}
		return result
	}


	// if k > 0
	sum := 0
	for i:=0; i<k; i++ {
		sum += code[i]
	}

	start := len(code) - 1
	result[start] = sum
	i := 0
	for {
		if i == len(code) - 1 {
			break
		}

		first := (k - i - 1 + len(code)) % len(code)
		next := start - i
		sum = sum - code[first] + code[next]
		result[start-i-1] = sum
		i++
	}
	return result
}

func TestDefuseBomb(t *testing.T) {
	fmt.Println(decrypt([]int{2,4,9,3},-2))
}