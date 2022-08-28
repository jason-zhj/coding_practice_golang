package main

import "testing"

func closestDivisors(num int) []int {
	start := 1
	end := num + 2

	var mid int
	var sqr int
	for {
		if end <= start {
			break
		}

		mid = (start + end) / 2
		sqr = mid * mid
		// end condition

		if sqr > num + 2 {
			end = mid - 1
		} else if sqr < num + 1 {
			start = mid + 1
		} else if sqr == num + 1 || sqr == num + 2 {
			return []int{mid, mid}
		}
	}

	for {
		if num + 1 % start == 0 {
			return []int{start, (num+1) / start}
		}
		if num + 2 % start == 0 {
			return []int{start, (num+2) / start}
		}

		start--
	}

	return nil
}

func TestClosestDivisor(t *testing.T) {
	closestDivisors(123)
}
