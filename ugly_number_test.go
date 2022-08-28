package main

import (
	"testing"
	"fmt"
)

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	arr := make([]int, n)
	arr[0] = 1
	i2 := 0
	i3 := 0
	i5 := 0

	for i:=1; i<n; i++ {
		next := min(arr[i2]*2, arr[i3] * 3, arr[i5] * i5)
		if next == arr[i2]*2 {
			i2++
		}
		if next == arr[i3] * 3 {
			i3++
		}
		if next == arr[i5] * 5 {
			i5++
		}

		arr[i] = next
	}

	return arr[n-1]
}

func min(nums ...int) int {
	result := nums[0]
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}

func TestUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(10))
}
