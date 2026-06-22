package main

import "fmt"

func main() {

	testcases := [][]int{{-5, 1, 5, 0, -7}, {-4, -3, -2, -1, 4, 3, 2}}
	for _, c := range testcases {
		fmt.Println(largestAltitude(c))
	}
}
func largestAltitude(gain []int) int {
	current := 0
	m := 0
	for _, i := range gain {
		current = current + i
		if current > m {
			m = current
		}
	}
	return m
}
