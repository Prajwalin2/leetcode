package main

import "fmt"

type testcase struct {
	n           int
	k           int
	invocations [][]int
}

func main() {
	testcases := []testcase{
		// {n: 4, k: 1, invocations: [][]int{{1, 2}, {0, 1}, {3, 2}}},
		{5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}}},
		// {3, 2, [][]int{{1, 2}, {0, 1}, {2, 0}}},
		// {2, 0, [][]int{}},
		// {4, 1, [][]int{{1, 2}, {0, 1}, {3, 2}}},
		// {3, 1, [][]int{{1, 2}, {0, 1}}},
	}
	for _, t := range testcases {
		ans := remainingMethods(t.n, t.k, t.invocations)
		fmt.Println(ans)
	}
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	edges := make([][]int, n)
	inDegree := make([]int, n)

	for _, inv := range invocations {
		u, v := inv[0], inv[1]
		edges[u] = append(edges[u], v)
		inDegree[v]++
	}

	queue := []int{k}
	sus := make([]bool, n)
	sus[k] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range edges[u] {
			inDegree[v]--

			if !sus[v] {
				queue = append(queue, v)
				sus[v] = true
			}
		}
	}

	canRemoveAll := true
	rem := []int{}

	for i := 0; i < n; i++ {
		if sus[i] && inDegree[i] > 0 {
			canRemoveAll = false
			break
		} else if !sus[i] {
			rem = append(rem, i)
		}
	}

	if !canRemoveAll {
		allNodes := make([]int, n)
		for i := 0; i < n; i++ {
			allNodes[i] = i
		}
		return allNodes
	}

	return rem
}
