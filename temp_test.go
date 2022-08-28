package main

import (
	"fmt"
	"testing"
)

type Node struct {
	Val int
	Next *Node
}

/*
reverse the singly linked list every k nodes
1 -> 2 -> 3 ...
prev cur next
cur.Next = prev

prev = cur
cur = next
next = secondNext

1,2,3,4,5,6,7,8
s
    cur
3,2,1

count = 2
3,2,1,
k = 3

4,5,6,7,8
cur
start
6,5,4
cur

7,8
    cur
count = 1
 */

func reverseK(head * Node, k int) *Node {
	if head == nil {
		return nil
	}

	// assume there are at least k nodes
	cur := head
	start := head
	count := 0
	for {
		cur = cur.Next
		count++
		if cur == nil { // there are < k elements in the list
			return start
		}

		if count == k-1 {
			next := reverseK(cur.Next, k)
			cur.Next = nil
			reverse(start)
			start.Next = next
			break
		}
	}

	return cur
}

func reverse(head *Node) *Node{
	// special cases where length of list if < 3
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// two elements
	if head.Next.Next == nil {
		tail := head.Next
		tail.Next = head
		head.Next = nil
		return tail
	}

	var prev *Node
	var cur *Node
	var next *Node

	// intialisation
	prev = head
	cur = prev.Next
	next = cur.Next

	head.Next = nil

	for {
		cur.Next = prev

		if next == nil {
			// cur is the new head
			return cur
		}

		prev = cur
		cur = next
		next = next.Next
	}
}

func printLinkedList(head *Node) {
	if head == nil {
		fmt.Println("empty")
	}
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

type testcase struct {
	nums []int
	k int
}

/*
number of elements < k
case: nil, 2
case: 1  , 2
case: 1 -> 2, 3
number of elements is a multiple of (or equal) k
case: 1  , 1
case: 1 -> 2, 2
case: 1 -> 2 -> 3, 3
case: 1 -> 2 -> 3 -> 4, 2
number of elements > k, but not multiple of k
case: 1 -> 2 -> 3 -> 4 -> 5, 2
 */



func TestTemp(t *testing.T) {
	cases := []*testcase{
		{
			nums:[]int{},
			k:2,
		},
		{
			nums:[]int{1},
			k:2,
		},
		{
			nums:[]int{1,2},
			k:3,
		},
		{
			nums:[]int{1},
			k:1,
		},
		{
			nums:[]int{1,2},
			k:2,
		},
		{
			nums:[]int{1,2,3},
			k:3,
		},
		{
			nums:[]int{1,2,3,4},
			k:2,
		},
		{
			nums:[]int{1,2,3,4,5},
			k:2,
		},
		{
			nums:[]int{1,2,3,4,5},
			k:1,
		},
	}

	for _,c := range cases {
		testScenario(c.nums,c.k)
	}
}

func testScenario(nums []int, k int) {
	var head *Node
	if len(nums) > 0 {
		head = createList(nums)
	}
	newHead := reverseK(head, k)
	fmt.Printf("[Case] nums = %v, k = %v, result is:\n",nums, k)
	printLinkedList(newHead)
}

func createList(nums []int) *Node{
	head := &Node{
		Val:nums[0],
	}
	cur := head
	for i:=1; i<len(nums); i++ {
		cur.Next = &Node{
			Val:nums[i],
		}
		cur = cur.Next
	}
	cur.Next = nil
	return head
}



