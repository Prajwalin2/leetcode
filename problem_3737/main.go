package main

import "fmt"

type testcase struct {
	nums   []int
	target int
}

func main() {
	testcases := []testcase{
		{
			[]int{1, 2, 2, 3}, 2,
		},
	}
	for _, t := range testcases {
		countMajoritySubarrays(t.nums, t.target)
	}
}

func countMajoritySubarrays(nums []int, target int) int {
	total := 0
	for i := range nums {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == target {
				count++
			}
			if count > (j+1-i)/2 {
				fmt.Println(i, j, "  	", nums[i:j], "	", count)
				total++
			}
		}
	}
	return total
}
