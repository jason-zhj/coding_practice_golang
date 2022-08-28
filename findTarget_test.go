package main

import (
	"testing"
	"fmt"
)


func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

// cur is not nil
func dfs(cur *TreeNode, root *TreeNode, k int) bool {
	if cur.Val != k - cur.Val && search(root, k - cur.Val) {
		return true
	}
	if cur.Left != nil && dfs(cur.Left, root, k) {
		return true
	}
	if cur.Right != nil && dfs(cur.Right, root, k) {
		return true
	}
	return false
}

func search(root *TreeNode, target int) bool {
	if root.Val == target {
		return true
	}
	if root.Val < target {
		if root.Right != nil && search(root.Right, target) {
			return true
		}
	} else { // root.Val > target
		if root.Left != nil && search(root.Left, target) {
			return true
		}
	}
	return false
}

func TestFindTarget(t *testing.T) {
	root := &TreeNode{
		Val:2,
		Left:&TreeNode{
			Val:1,
		},
		Right:&TreeNode{
			Val:3,
		},
	}
	fmt.Println(findTarget(root,4))
}
