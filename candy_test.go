package main

import (
	"fmt"
	"testing"
)

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	alloc := make([]int, len(ratings))
	alloc[0] = 1

	i:=1
	for {
		if i >= len(alloc) {
			break
		}

		if ratings[i] > ratings[i-1] {
			alloc[i] = alloc[i-1] + 1
		} else if ratings[i] < ratings[i-1] {
			if alloc[i-1] - 1 < 1 {
				fixAlloc(alloc, i, ratings)
			} else {
				alloc[i] = 1
			}
		} else { // equal
			alloc[i] = 1
		}

		i++
	}

	sum := 0
	for _, num := range alloc {
		sum += num
	}
	return sum
}

func fixAlloc(alloc []int, i int, ratings []int) {
	alloc[i] = 1
	for {
		if i-1 < 0 || ratings[i-1] <= ratings[i] {
			break
		}

		// ratings[i-1] > ratings[i]
		if alloc[i-1] > alloc[i] {
			break
		}

		alloc[i-1] = alloc[i] + 1
		i--
	}
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1,0,2}))
}
