package main

import (
	"strconv"
	"testing"
	"fmt"
)

func calculate(s string) int {
	nums, ops := parse(s)

	if len(ops) == 0 {
		return nums[0]
	}

	processedNums := []int{}
	processedOps := []string{}
	// first pass
	i := 0 // point to num
	j := 0 // point to op
	for {
		if ops[j] == "+" || ops[j] == "-" {
			processedNums = append(processedNums, nums[i])
			processedOps = append(processedOps, ops[j])
			i++
			j++
		} else { // * or /
			// keep calculating
			fst := nums[i]
			i++
			snd := nums[i]
			for {
				fst = calc(fst, snd, ops[j])
				i++
				j++
				if j >= len(ops) || ops[j] == "+" || ops[j] == "-" {
					break
				}
				snd = nums[i]
			}
			i--
			nums[i] = fst
		}

		if j >= len(ops) {
			break
		}
	}

	if i < len(nums) {
		processedNums = append(processedNums, nums[len(nums)-1])
	}

	if len(processedOps) == 0 {
		return processedNums[0]
	}

	for i, op := range processedOps {
		processedNums[i+1] = calc(processedNums[i], processedNums[i+1], op)
	}

	return processedNums[len(processedNums)-1]
}

func calc(a,b int, op string) int {
	if op == "*" {
		return a * b
	} else if op == "/" {
		return a / b
	} else if op == "+" {
		return a + b
	}
	return a - b
}

func parse(s string) ([]int, []string) {
	num := ""
	nums := []int{}
	ops := []string{}

	for _,c := range s {
		if c >= '0' && c <= '9' {
			num = num + string(c)
		} else {
			if len(num) > 0 {
				val, _ := strconv.Atoi(num)
				nums = append(nums, val)
				num = ""
			}
			if isOp(byte(c)) {
				ops = append(ops, string(c))
			}
		}
	}

	if len(num) > 0 {
		val, _ := strconv.Atoi(num)
		nums = append(nums, val)
	}

	return nums, ops
}

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c =='/'
}


func TestCalculate(t *testing.T) {
	fmt.Println(calculate("1*2+3*4"))
	fmt.Println(calculate("1+1"))
	fmt.Println(calculate("2*3*4"))
}