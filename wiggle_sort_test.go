package main

func wiggleSort(nums []int)  {
	m := map[int]int{}
	for _, num := range nums {
		if _,ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}

	result := make([]int, len(nums))
	for num := range m {
		selected := []int{num}
		m[num]--
		if searchWiggle(&selected, m, len(nums), &result) {
			break
		}
		m[num]++
	}

	for i := range nums {
		nums[i] = result[i]
	}
}

// selected has at least one element
func searchWiggle(selected *[]int, m map[int]int, totalLen int, result *[]int) bool {
	if len(*selected) == totalLen {
		for i:= range *result {
			(*result)[i] = (*selected)[i]
		}
		return true
	}

	requireBigger := len(*selected) % 2 == 1
	lastAdd := (*selected)[len(*selected)-1]

	for val, count := range m {
		if count == 0 {
			continue
		}

		if (requireBigger && val > lastAdd) || (!requireBigger && val < lastAdd) {
			// dfs
			*selected = append(*selected, val)
			m[val]--
			if searchWiggle(selected, m, totalLen, result) {
				return true
			}
			*selected = (*selected)[:len(*selected)-1]
			m[val]++
		}
	}

	return false
}
