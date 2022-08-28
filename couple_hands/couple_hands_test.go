package couple_hands

import (
	"testing"
	"fmt"
)

func minSwapsCouples(row []int) int {
	m := map[int]int{} // id -> position
	for i, id := range row {
		m[id] = i
	}

	return solve(row, 0, m)
}

func solve(row []int, start int, m map[int]int) int {
	if start >= len(row) {
		return 0
	}

	var look int
	if row[start] % 2 == 1 {
		look = row[start] - 1
	} else {
		look = row[start] + 1
	}

	if row[start]+1 == look {
		return solve(row, start+2, m)
	}

	lookPos := m[look]
	toReplace := row[start+1]
	row[start+1], row[lookPos] = row[lookPos], row[start+1]
	m[toReplace] = lookPos
	m[look] = start+1

	return solve(row, start+2, m) + 1
}

func TestMinSwapCouples(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{3,2,0,1}))
}
