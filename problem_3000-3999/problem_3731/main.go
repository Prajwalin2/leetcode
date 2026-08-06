package main

import (
	"fmt"
	"sort"
)

func main() {
	for _, test := range [][]int{{1, 4, 2, 5}} {
		ans := findMissingElements(test)
		fmt.Println(ans)
	}
}

func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	mn, mx := nums[0], nums[len(nums)-1]
	counter, pos := 0, 0
	l := 1 + mx - (mn + len(nums))
	ret := make([]int, l)
	for mn < mx {
		for mn != nums[counter] {
			ret[pos] = mn
			pos++
			mn++
		}
		counter++
		mn++
	}
	return ret
}
