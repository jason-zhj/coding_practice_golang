package main

import (
	"sort"
	"testing"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) == 1 {
		return []int{1}
	}

	indices := make([]int, len(arr))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i,j int) bool {
		return arr[i] <= arr[j]
	})

	rank := 1
	ranks := make([]int, len(arr))
	ranks[indices[0]] = rank
	for i:=1; i<len(indices); i++ {
		if arr[indices[i]] > arr[indices[i-1]] {
			rank++
		}
		ranks[indices[i]] = rank
	}

	return ranks
}

func TestRankTransform(t *testing.T) {
	arrayRankTransform([]int{40,10,20,30})
}
