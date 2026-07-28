package main

import "sort"

func main() {
	maximumProduct([]int{1, 2, 3})
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	p1 := nums[n-1] * nums[n-2] * nums[n-3]
	p2 := nums[0] * nums[1] * nums[n-1]

	if p1 > p2 {
		return p1
	}

	return p2
}
