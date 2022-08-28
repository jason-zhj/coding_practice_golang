package main

import "strings"

type CharCount struct {
	Char string
	Locations []int
}

func smallestSubsequence(s string) string {
	chars := strings.Split(s, "")
	m := map[string][]int{}
	base := make([]string, len(s))

	for i, c := range chars {
		if _,ok :=m[c]; !ok {
			m[c] = []int{i}
		} else {
			m[c] = append(m[c], i)
		}
	}

	for i, c := range chars {
		if len(m[c]) == 1 {
			base[i] = c
		}
	}

	dupList := []*CharCount{}
	for s, locations := range m {
		if len(locations) > 1 {
			dupList = append(dupList, &CharCount{
				Char: s,
				Locations: locations,
			})
		}
	}

	minSeq := ""
	findMinSeq(base, dupList, 0, &minSeq)

	return minSeq
}


func findMinSeq(base []string, dupList []*CharCount, index int, minSeq *string) {
	if index >= len(dupList) {
		seq := strings.Join(base, "")
		if len(*minSeq) == 0 || seq < *minSeq {
			*minSeq = seq
		}
		return
	}

	c := dupList[index].Char
	for _, i := range dupList[index].Locations {
		base[i] = c
		findMinSeq(base, dupList, index+1, minSeq)
		base[i] = ""
	}
}

