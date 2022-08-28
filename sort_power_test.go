package main

import (
	"testing"
	"sort"
	"fmt"
)

func getKth(lo int, hi int, k int) int {
	arr := make([]int, hi - lo + 1)
	for i := range arr {
		arr[i] = getPower(lo + i)
	}
	sort.Slice(arr, func(i,j int) bool {
		return arr[i] <= arr[j]
	})
	return arr[k-1]
}

func getPower(x int) int {
	count := 0
	for {
		if x == 1 {
			break
		}

		if x % 2 == 0 {
			x /= 2
		} else {
			x = 3 *x + 1
		}
		count++
	}

	return count
}

func TestSortPower(t *testing.T) {
	fmt.Println(getPower(570))
}
