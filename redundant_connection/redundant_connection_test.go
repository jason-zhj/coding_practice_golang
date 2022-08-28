package redundant_connection

import (
	"testing"
	"fmt"
)

func findRedundantConnection(edges [][]int) []int {
	graph := map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _,ok := graph[a]; !ok {
			graph[a] = []int{b}
		} else {
			graph[a] = append(graph[a], b)
		}

		if _,ok := graph[b]; !ok {
			graph[b] = []int{a}
		} else {
			graph[b] = append(graph[b], a)
		}
	}

	path := []int{}
	visited := map[int]bool{}
	record := map[int]bool{}
	dfs(1, graph, &path, visited, record)

	for i:=len(edges)-1; i>=0; i-- {
		a, b := edges[i][0], edges[i][1]
		if record[a] && record[b] {
			return edges[i]
		}
	}

	return edges[0]
}

func dfs(node int, graph map[int][]int, path *[]int, visited map[int]bool, record map[int]bool) bool {
	if visited[node] {
		for i:=len(*path)-1; i>=0; i-- {
			record[(*path)[i]] = true
			if (*path)[i] == node {
				break
			}
		}
		return true
	}

	visited[node] = true
	parent := -1
	if len(*path) > 0 {
		parent = (*path)[len(*path)-1]
	}
	*path = append(*path, node)

	for _, child := range graph[node] {
		if child == parent {
			continue
		}
		if dfs(child, graph, path, visited, record) {
			return true
		}
	}

	visited[node] = false
	*path = (*path)[:len(*path)-1]
	return false
}

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{
		{
			1,2,
		},
		{
			1,3,
		},
		{
			2,3,
		},
	}
	fmt.Println(findRedundantConnection(edges))
}
