package main

import "fmt"

func main() {
	for _, test := range []int{1, 2, 4} {
		ans := winnerSquareGame(test)
		fmt.Println(ans)
	}
}

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)

	dp[0] = false

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp[i-j*j] {
				dp[i] = true

				break
			}
		}
	}

	return dp[n]
}
