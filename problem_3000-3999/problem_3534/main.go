package main

import "sort"

func main() {
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	type node struct {
		value int
		index int
	}

	nodes := make([]node, n)
	for i, num := range nums {
		nodes[i] = node{num, i}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].value < nodes[j].value
	})

	position := make([]int, n)
	for i, node := range nodes {
		position[node.index] = i
	}

	next := make([]int, n)
	right := 0
	for left := 0; left < n; left++ {
		if right < left {
			right = left
		}
		for right+1 < n && nodes[right+1].value-nodes[left].value <= maxDiff {
			right++
		}
		next[left] = right
	}

	log := 1
	for (1 << log) <= n {
		log++
	}

	jump := make([][]int, log)
	jump[0] = next
	for i := 1; i < log; i++ {
		jump[i] = make([]int, n)
		for j := 0; j < n; j++ {
			jump[i][j] = jump[i-1][jump[i-1][j]]
		}
	}

	answer := make([]int, len(queries))
	for i, query := range queries {
		start, end := position[query[0]], position[query[1]]
		if start > end {
			start, end = end, start
		}
		if start == end {
			continue
		}
		if jump[log-1][start] < end {
			answer[i] = -1
			continue
		}

		steps := 0
		current := start
		for j := log - 1; j >= 0; j-- {
			if jump[j][current] < end {
				current = jump[j][current]
				steps += 1 << j
			}
		}
		answer[i] = steps + 1
	}
	return answer
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
