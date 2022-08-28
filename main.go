package main

import (
	"container/heap"
	"fmt"
)

type HeapItem struct {
	Weight int
	Val string
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type HeapArr []*HeapItem

func (h HeapArr) Less(i,j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h HeapArr) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HeapArr) Len() int {
	return len(h)
}

func (h *HeapArr) Push(item interface{}) {
	*h = append(*h, item.(*HeapItem))
}

func (h *HeapArr) Pop() interface{} {
	old := *h
	toPop := old[len(old)-1]
	*h = old[:len(old)-1]
	return toPop
}

type MyHeap struct {
	Items HeapArr
}

func (m *MyHeap) Push(item *HeapItem) {
	heap.Push(&m.Items, item)
}

func (m *MyHeap) PopMin() *HeapItem {
	poped := heap.Pop(&m.Items)
	return poped.(*HeapItem)
}

func main() {
	myheap := MyHeap{}
	myheap.Push(&HeapItem{
		Weight:5,
		Val:"5",
	})
	myheap.Push(&HeapItem{
		Weight:1,
		Val:"1",
	})
	myheap.Push(&HeapItem{
		Weight:10,
		Val:"10",
	})
	myheap.Push(&HeapItem{
		Weight:3,
		Val:"3",
	})
	fmt.Println(myheap.PopMin().Val)
	fmt.Println(myheap.PopMin().Val)
	fmt.Println(myheap.PopMin().Val)
	fmt.Println(myheap.PopMin().Val)
}