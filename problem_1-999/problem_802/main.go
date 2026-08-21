package main

import "fmt"

func main() {
	testcases := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	states := eventualSafeNodes(testcases)
	fmt.Println(states)
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	state := make([]int, n) // 0: unvisited, 1: visiting, 2: safe
	var dfs func(int) bool

	dfs = func(node int) bool {
		if state[node] > 0 {
			return state[node] == 2
		}
		state[node] = 1 // Mark as visiting
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		state[node] = 2 // Mark as safe
		return true
	}

	safeNodes := []int{}
	for i := range n {
		if dfs(i) {
			safeNodes = append(safeNodes, i)
		}
	}

	return safeNodes
}
