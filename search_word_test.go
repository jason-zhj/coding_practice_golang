package main


type SNode struct {
	Children map[rune]*SNode
	IsLeaf bool
}

type WordDictionary struct {
	root *SNode
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &SNode{
			Children: map[rune]*SNode{},
		},
	}
}


func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
	for _,c := range word {
		if _,ok := cur.Children[c]; !ok {
			cur.Children[c] = &SNode{
				Children: map[rune]*SNode{},
			}
		}
		cur = cur.Children[c]
	}
	cur.IsLeaf = true
}


func (this *WordDictionary) Search(word string) bool {
	cur := this.root
	for _,c := range word {
		if _,ok := cur.Children[c]; !ok {
			return false
		}
		cur = cur.Children[c]
	}
	return cur.IsLeaf
}

// root - (d) -> x - (a) -> x - (d) -> leaf
//
func search(word string, cur *SNode) bool {
	if len(word) == 0 {
		return cur.IsLeaf
	}

	c := rune(word[0])
	if c == '.' {
		// try matching any
		for _, child := range cur.Children {
			if search(word[1:], child) {
				return true
			}
		}
		return false
	} else {
		// match character c
		if _, ok := cur.Children[c]; !ok {
			return false
		}
		return search(word[1:], cur.Children[c])
	}

	return false
}