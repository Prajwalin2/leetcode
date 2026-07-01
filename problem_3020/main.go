package main

import (
	"maps"
	"slices"
)

func main() {
	testcases := [][]int{{2, 4, 2, 4, 2}}
	for _, test := range testcases {
		maximumLength(test)
	}
}

func maximumLength(nums []int) int {
	m := map[int]int{}
	for _, i := range nums {
		m[i] += 1
	}
	max := m[1]
	delete(m, 1)
	keys := slices.Sorted(maps.Keys(m))
	for _, val := range keys {
		cur := Max(val, m)
		if max < cur {
			max = cur
		}
	}
	return max
}

func Max(current int, m map[int]int) int {
	val := m[current]
	delete(m, current)
	switch val {
	case 0:
		return -1
	case 1:
		return 1
	default:
		return 2 + Max(current*current, m)
	}
}
