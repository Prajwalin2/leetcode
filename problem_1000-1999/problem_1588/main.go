package main

import "fmt"

func main() {
	testcases := [][]int{{1, 4, 2, 3, 5}}
	for _, t := range testcases {
		fmt.Println(sumOddLengthSubarrays(t))
	}
}

func sumOddLengthSubarrays(arr []int) int {
	l := len(arr)
	total := 0
	for i, val := range arr {
		left := (l - i) / 2
		right := (i) / 2
		total += val * (left + right)
	}
	return total
}
