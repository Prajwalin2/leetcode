package main

import "fmt"

func main() {
	testcases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{30, 60, 90},
		// [1,1,4,2,1,1,0,0]
	}
	for _, test := range testcases {
		ans := dailyTemperatures(test)
		fmt.Println(ans)
	}
}

func dailyTemperatures(temps []int) []int {
	results := make([]int, len(temps))
	stack := make([]int, 0)
	for i, temp := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}
		stack = append(stack, i)
	}

	return results
}
