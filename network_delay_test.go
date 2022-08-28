package main

import (
	"testing"
	"fmt"
)

func TestMap(t *testing.T) {
	m := map[int][]int{}
	m[1] = []int{1,2}
	fmt.Println(m[1][0])
}