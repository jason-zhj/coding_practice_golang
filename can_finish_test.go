package main

import (
	"testing"
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	pre := make([][]int, numCourses)
	for i := range pre {
		pre[i] = []int{}
	}

	for _, p := range prerequisites {
		pre[p[0]] = append(pre[p[0]], p[1])
	}

	visited := make([]int, numCourses)
	marker := 1
	for i:=0; i<numCourses; i++ {
		if visited[i] == 0 { // not visited
			if dfs(i, pre, visited, marker) {
				return false
			}
			marker++
		}
	}

	return true
}

// dfs returns true if cycle is found
// dfs assumes the start hasn't been visited
func dfs(start int, graph [][]int, visited []int, marker int) bool {
	visited[start] = marker
	pres := graph[start]

	for _,pre := range pres {
		if visited[pre] == marker {
			return true
		}
		if visited[pre] == 0 {
			if dfs(pre, graph, visited, marker) {
				return true
			}
		}
	}

	visited[start] = 0
	return false
}

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(3,[][]int{
		{1,0},
		{0,1},
		{1,2},
	}))
}
