package main

import "strconv"

type testcase struct {
	s       string
	queries [][]int
}


func main() {
	testcases := []testcase{
		// {"10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}},
		{"9876543210", [][]int{{0, 9}}},
	}
	for _, t := range testcases {
		sumAndMultiply(t.s, t.queries)
	}
}
var mod int64 = 1000000007
func sumAndMultiply(s string, queries [][]int) []int {
	var total []int
	for _, query := range queries {
		no, err := strconv.Atoi(s[query[0] : query[1]+1])
		if err != nil {
			continue
		}
		total = append(total, sumMultiply(int64(no)))
	}
	return total
}

func sumMultiply(n int64) int {
	var multiplier, sum, no int64 = 1, 0, 0
	for n != 0 {
		remainder := n % 10
		if remainder != 0 {
			no += remainder * multiplier
			sum += remainder
			multiplier *= 10
		}
		n /= 10
	}
	if sum == n {
		return 0
	}
	var total int64 = (sum * no) % mod
	return int(total)
}
