package main

func main() {
	stoneGameIII([]int{1, 2, 3, 4})
}

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {

		dp[i] = -(1 << 30)

		sum := 0

		for j := i; j < n && j < i+3; j++ {

			sum += stoneValue[j]

			if sum-dp[j+1] > dp[i] {
				dp[i] = sum - dp[j+1]
			}
		}
	}

	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}
