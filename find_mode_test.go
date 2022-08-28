package main

import (
	"testing"
	"fmt"
)

type BSTTreeNode struct {
	Val int
	Left *BSTTreeNode
	Right *BSTTreeNode
}

func findMode(root *BSTTreeNode) []int {
	maxCount := 0
	maxVals := []int{}
	if root == nil {
		return []int{}
	}
	rootElemCount := countOcc(root.Left, root.Val, &maxCount, &maxVals) + countOcc(root.Right, root.Val, &maxCount, &maxVals) + 1

	if rootElemCount > maxCount {
		maxCount = rootElemCount
		maxVals = []int{root.Val}
	} else if rootElemCount == maxCount {
		maxVals = append(maxVals, root.Val)
	}

	return maxVals
}

// returns the count of parent in this subtree
func countOcc(root *BSTTreeNode, target int, maxCount *int, maxVals *[]int) int {
	if root == nil {
		return 0
	}

	var leftCount int
	var rightCount int
	if root.Val >= target {
		leftCount = countOcc(root.Left, target, maxCount, maxVals)
	}
	if root.Val <= target {
		rightCount = countOcc(root.Right, target, maxCount, maxVals)
	}

	if root.Val == target {
		return leftCount + rightCount + 1
	}

	// i'm not target, count myself
	myCount := countOcc(root.Left, root.Val, maxCount, maxVals) + countOcc(root.Right, root.Val, maxCount, maxVals) + 1
	if myCount > *maxCount {
		*maxVals = []int{root.Val}
		*maxCount = myCount
	} else if myCount == *maxCount {
		*maxVals = append(*maxVals, root.Val)
	}

	return leftCount + rightCount
}

func TestFindMode(t *testing.T) {
	mode := findMode(&BSTTreeNode{
		Val:1,
		Right:&BSTTreeNode{
			Val:2,
			Left:&BSTTreeNode{
				Val:2,
			},
		},
	})
	fmt.Println(mode)
}
