package main

import (
	"sort"
	"math"
	"testing"
	"fmt"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	results := make([][]int, R*C)
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++ {
			results[i*R+j] = []int{i,j}
		}
	}

	sort.Slice(results, func(i,j int) bool {
		return mandist(results[i][0],results[i][1],r0,c0) < mandist(results[j][0],results[j][1],r0,c0)
	})

	return results
}

func mandist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1 - x2))) + int(math.Abs(float64(y1- y2)))
}

func TestManDist(t *testing.T) {
	fmt.Println(allCellsDistOrder(2,3,1,2))
}