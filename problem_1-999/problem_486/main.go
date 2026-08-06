package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 5, 2}, {1, 5, 233, 7},
	}
	for _, t := range testcases {
		ans := predictTheWinner(t)
		fmt.Println(ans)
	}
}

func predictTheWinner(nums []int) bool {
	return pickA(0, 0, nums)
}

func pickA(a, b int, nums []int) bool {
	if len(nums) == 0 {
		return a >= b
	}
	return pickB(a+nums[0], b, nums[1:]) || pickB(a+nums[len(nums)-1], b, nums[:len(nums)-1])
}

func pickB(a, b int, nums []int) bool {
	if len(nums) == 0 {
		return a >= b
	}
	return pickA(a, b+nums[0], nums[1:]) && pickA(a, b+nums[len(nums)-1], nums[:len(nums)-1])
}
