package main

import "math"

type testcase struct {
	edges  [][]int
	online []bool
	k      int64
}

func main() {
	testcases := []testcase{
		// {[][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}}, []bool{true, true, true, true}, 10},
		// {[][]int{{0, 2, 90}, {1, 2, 36}, {0, 1, 61}}, []bool{true, true, true}, 129},
		{[][]int{{0, 1, 0}, {0, 2, 7}, {1, 3, 9}, {0, 4, 7}, {2, 4, 9}, {3, 4, 2}, {0, 3, 5}, {2, 3, 3}, {1, 4, 6}, {1, 2, 0}}, []bool{true, true, true, true, true}, 5},
	}

	for _, t := range testcases {
		findMaxPathScore(t.edges, t.online, t.k)
	}
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	dp := map[int]int{}
	grid := make([][]int, len(online))
	for i := range grid {
		grid[i] = make([]int, len(online))
		for j := range grid[i] {
			grid[i][j] = -1
		}
	}
	for _, edge := range edges {
		if !online[edge[1]] {
			continue
		}
		grid[edge[0]][edge[1]] = edge[2]
	}
	maximum := -1
	q := queue{{0, 0, math.MaxInt64}}
	for len(q) != 0 {
		ele := q.Pop()
		pos := int(ele[0])
		val := ele[1]
		m := ele[2]
		if pos == len(online)-1 {
			if val <= k {
				maximum = max(int(m), maximum)
			}
			continue
		}
		for j := 0; j < len(online); j++ {
			if grid[pos][j] != -1 {
				curr := int64(grid[pos][j])
				total := val + int64(grid[pos][j])
				if total <= k {
					temp := -1
					if _, ok := dp[j]; ok {
						temp = dp[j]
					}
					mini := min(m, curr)
					if mini >= int64(temp) {
						dp[j] = int(mini)
						q.Push([]int64{int64(j), total, min(m, curr)})
					}
				}
			}
		}
	}
	return int(maximum)
}

type queue [][]int64

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i][1] < q[j][1]
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(ele interface{}) {
	e := ele.([]int64)

	*q = append(*q, e)
}

func (q *queue) Pop() []int64 {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
