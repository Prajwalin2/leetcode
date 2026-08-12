package main

import (
	"fmt"
	"sort"
)

func main() {
	testcases := [][]int{
		{2, 6, 4},
		{3, 6, 2, 8},
	}
	for _, test := range testcases {
		ans := gcdSum(test)
		fmt.Println(ans)
	}
}

func gcdSum(nums []int) int64 {
	gcd := make([]int, len(nums))
	mx := nums[0]
	for i, num := range nums {
		mx = max(mx, num)
		gcd[i] = GCD(mx, num)
	}
	sort.Ints(gcd)
	low, high := 0, len(nums)-1
	var sum int64 = 0
	for low < high {
		sum += int64(GCD(gcd[low], gcd[high]))
		low++
		high--
	}
	return sum
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
