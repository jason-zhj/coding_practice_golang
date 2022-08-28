package main

import (
	"testing"
	"fmt"
)

func findKthPositive(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	cur := arr[0] + 1
	count := 0
	i := 1

	for {
		if arr[i] > cur { // cur is missing
			count++
			if count == k {
				return cur
			}
			cur++
		} else { // arr[i] == cur
			i++
			cur++
		}
	}

	return 0
}

func TestKthPositive(t *testing.T) {
	fmt.Println(findKthPositive([]int{2,3,4,7,11},5))
}
