package main

import "fmt"

func main() {
	testcases := [][]string{{"ab", ".*"}}
	for _, v := range testcases {
		fmt.Println(isMatch(v[0], v[1]))
	}
}

func isMatch(s string, p string) bool {
	memo := make([][]int, len(s)+1)
	for i := range memo {
		memo[i] = make([]int, len(p)+1)
	}

	var dp func(i, j int) bool

	dp = func(i, j int) bool {
		if memo[i][j] != 0 {
			return memo[i][j] == 1
		}

		var temp bool

		if j == len(p) {
			temp = i == len(s)
		} else {
			firstMatch := i < len(s) &&
				(p[j] == s[i] || p[j] == '.')

			if j+1 < len(p) && p[j+1] == '*' {
				temp =
					// use zero occurrence of p[j]
					dp(i, j+2) ||
						// use one occurrence of p[j]
						(firstMatch && dp(i+1, j))
			} else {
				temp = firstMatch && dp(i+1, j+1)
			}
		}

		if temp {
			memo[i][j] = 1
		} else {
			memo[i][j] = -1
		}

		return temp
	}

	return dp(0, 0)
}
