package main

import "testing"

func moveZeroes(nums []int)  {
	s := -1

	for i, num := range nums {
		if s < 0 {
			if num == 0 {
				s = i
			}
			continue
		}
		if num != 0 { // swap nums[i] with nums[s]
			nums[i], nums[s] = nums[s], nums[i]
			s++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	moveZeroes([]int{0,1,0,3,12})
}
