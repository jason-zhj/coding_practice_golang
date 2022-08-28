package main

import "testing"

type Node struct {
	Val rune
	Children map[rune]*Node
	IsLeaf bool
}

func wordBreak(s string, wordDict []string) bool {
	root := buildTree(wordDict)
	cur := root
	for _, c := range s {
		if cur == root {
			// start looking from root again
			if child, ok := cur.Children[c]; !ok {
				return false
			} else {
				cur = child
			}
		} else {
			// we are in the middle of the tree
			if cur.IsLeaf { // reaches the end
				cur = root
			}
			if child, ok := cur.Children[c]; !ok {
				return false
			} else {
				cur = child
			}
		}
	}

	return cur.IsLeaf
}

func buildTree(wordDict []string) *Node {
	root := &Node{
		Children: map[rune]*Node{},
	}
	for _, word := range wordDict {
		cur := root
		for i, c := range word {
			if _, ok := cur.Children[c]; ok {
				cur = cur.Children[c]
			} else {
				newNode := &Node{
					Children: map[rune]*Node{},
					Val: c,
				}
				cur.Children[c] = newNode
				cur = newNode
			}

			if i == len(word)-1 {
				cur.IsLeaf = true
			}
		}
	}
	return root
}

func TestWordBreak(t *testing.T) {
	wordBreak("aaaaaaa",[]string{"aaaa","aaa"})
}
