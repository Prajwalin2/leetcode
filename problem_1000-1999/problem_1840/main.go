package main

import (
	"fmt"
	"sort"
)

func main() {
	testcases := [][][]int{
		{{2, 1}, {4, 1}},
		{},
		{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
	}
	testcasesN := []int{5, 6, 10}

	for i, r := range testcases {
		fmt.Println(maxBuilding(testcasesN[i], r))
	}

	fmt.Println(maxBuilding(10, [][]int{
		{8, 5}, {9, 0}, {6, 2}, {4, 0}, {3, 2},
		{10, 0}, {5, 3}, {7, 3}, {2, 4},
	}))
}

func maxBuilding(n int, restrictions [][]int) int {
	// Add building 1 with height 0.
	restrictions = append(restrictions, []int{1, 0})

	// Add building n with maximum possible unrestricted height n-1.
	restrictions = append(restrictions, []int{n, n - 1})

	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	// Left to right pass:
	// Each building can only be at most previousHeight + distance.
	for i := 1; i < len(restrictions); i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]
		allowed := restrictions[i-1][1] + dist

		if restrictions[i][1] > allowed {
			restrictions[i][1] = allowed
		}
	}

	// Right to left pass:
	// Same idea, but constraints also apply from the right.
	for i := len(restrictions) - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]
		allowed := restrictions[i+1][1] + dist

		if restrictions[i][1] > allowed {
			restrictions[i][1] = allowed
		}
	}

	ans := 0

	for i := 1; i < len(restrictions); i++ {
		id1 := restrictions[i-1][0]
		h1 := restrictions[i-1][1]

		id2 := restrictions[i][0]
		h2 := restrictions[i][1]

		dist := id2 - id1
		diff := abs(h1 - h2)

		peak := max(h1, h2) + (dist-diff)/2

		if peak > ans {
			ans = peak
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
