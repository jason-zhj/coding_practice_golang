package main

import (
	"testing"
	"fmt"
)
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{}
	}

	m := map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _,ok := m[a]; !ok {
			m[a] = []int{b}
		} else {
			m[a] = append(m[a], b)
		}

		if _,ok := m[b]; !ok {
			m[b] = []int{a}
		} else {
			m[b] = append(m[b], a)
		}
	}

	// find all leaf nodes
	leaves := []int{}
	for node, neighbors := range m {
		if len(neighbors) == 1 {
			leaves = append(leaves, node)
		}
	}

	// keep trimming until len(leaves) <= 2
	remain := n
	for {
		if remain <= 2 {
			break
		}

		newLeaves := []int{}
		for _, leaf := range leaves {
			neighbor := m[leaf][0]
			remove(m, neighbor, leaf)
			if len(m[neighbor]) == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
		}

		remain -= len(leaves)
		leaves = newLeaves
	}

	return leaves
}

func remove(m map[int][]int, key int, val int) {
	result := []int{}
	for _, item := range m[key] {
		if item != val {
			result = append(result, item)
		}
	}
	m[key] = result
}

func TestFindMinHeightTrees(t *testing.T) {
	fmt.Println(findMinHeightTrees(2,[][]int{
		{1,0},
	}))
}
