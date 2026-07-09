package main

type testcase struct {
	n       int
	nums    []int
	maxDiff int
	queries [][]int
}

// n = 4, nums = [2,5,6,8], maxDiff = 2, queries = [[0,1],[0,2],[1,3],[2,3]]
func main() {
	testcases := []testcase{
		// {n: 2, nums: []int{1, 3}, maxDiff: 1, queries: [][]int{{0, 0}, {0, 1}}},
		{4, []int{2, 5, 6, 8}, 2, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
	}
	for _, t := range testcases {
		pathExistenceQueries(t.n, t.nums, t.maxDiff, t.queries)
	}
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= maxDiff {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		if q[1]-q[0] == dp[q[1]]-dp[q[0]] {
			ret[i] = true
		}
	}
	return ret
}
