package main

import (
	"testing"
	"container/heap"
	"fmt"
)

type Item struct {
	Value int
	Index []int
}

type ItemHeap struct {
	items []*Item
}

func (i *ItemHeap) Len() int {
	return len(i.items)
}

func (h *ItemHeap) Less(i int,j int) bool {
	return h.items[i].Value < h.items[j].Value
}


func (h *ItemHeap) Swap(i,j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *ItemHeap) Push(x interface{}) {
	h.items = append(h.items, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
	last := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	return last
}

func TestHeap(t *testing.T) {
	itemHeap := &ItemHeap{}
	heap.Push(itemHeap,&Item{
		Value:10,
	})
	heap.Push(itemHeap,&Item{
		Value:1,
	})
	heap.Push(itemHeap,&Item{
		Value:5,
	})

	poped := heap.Pop(itemHeap)
	fmt.Println(poped.(*Item).Value)
	poped = heap.Pop(itemHeap)
	fmt.Println(poped.(*Item).Value)
	poped = heap.Pop(itemHeap)
	fmt.Println(poped.(*Item).Value)
}