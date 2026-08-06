package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 1, 3},   // n=1, m=3 → 3
		{2, 1, 3},   // n=2, m=3 → 6
		{3, 4, 5},   // n=3, m=2 → 2
		{3, 3, 5},   // n=3, m=3 → 10
		{4, 1, 2},   // n=4, m=2 → 2
	}
	for _, t := range testcases {
		fmt.Println(zigZagArrays(t[0], t[1], t[2]))
	}
}

func zigZagArrays(n int, l int, r int) int {
	const MOD = 1_000_000_007
	m := r - l + 1
	if n == 1 {
		return m
	}

	dp := make([][2]int64, m)
	for x := range dp {
		dp[x][0], dp[x][1] = 1, 1
	}

	for i := 1; i < n; i++ {
		ndp := make([][2]int64, m)

		var pre int64
		for y := range m {
			ndp[y][0] = pre
			pre = (pre + dp[y][1]) % MOD
		}

		var suf int64
		for y := m - 1; y >= 0; y-- {
			ndp[y][1] = suf
			suf = (suf + dp[y][0]) % MOD
		}

		dp = ndp
	}

	var ans int64
	for x := range m {
		ans = (ans + dp[x][0] + dp[x][1]) % MOD
	}
	return int(ans)
}
