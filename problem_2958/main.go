package main

import "fmt"

type testcase struct {
	nums []int
	k    int
}

func main() {
	testcases := []testcase{
		{[]int{1, 2, 3, 1, 2, 3, 1, 2}, 2},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 1},
		{[]int{5, 5, 5, 5, 5, 5, 5}, 4},
	}
	for _, test := range testcases {
		ans := maxSubarrayLength(test.nums, test.k)
		fmt.Println(ans)
	}
}

func maxSubarrayLength(nums []int, k int) int {
	m := map[int]int{}
	mx, l := 0, 0

	for r, num := range nums {
		m[num]++
		for m[num] > k {
			m[nums[l]]--
			l++
		}
		mx = max(mx, 1+r-l)
	}
	return mx
}
