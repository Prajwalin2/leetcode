package main

import (
	"sort"
	"strconv"
)

type item struct {
	value int
	index int
}

func main() {
	for _, test := range [][]int{{10, 3, 8, 9, 4}} {
		findRelativeRanks(test)
	}
}

func findRelativeRanks(score []int) []string {
	items := make([]item, len(score))
	for i, v := range score {
		items[i] = item{value: v, index: i}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})
	arr := make([]string, len(score))
	var val string
	for i, item := range items {
		switch i {
		case 0:
			val = "Gold Medal"
		case 1:
			val = "Silver Medal"
		case 2:
			val = "Bronze Medal"
		default:
			val = strconv.Itoa(i + 1)
		}
		arr[item.index] = val
	}
	return arr
}
