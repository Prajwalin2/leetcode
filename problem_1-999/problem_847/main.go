package main

func main() {
	shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}})
}

type State struct {
	mask, node, dist int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	allVisited := (1 << n) - 1
	queue := []State{}
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		queue = append(queue, State{1 << i, i, 0})
		visited[(1<<i)*16+i] = true
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.mask == allVisited {
			return cur.dist
		}

		for _, neighbor := range graph[cur.node] {
			newMask := cur.mask | (1 << neighbor)
			hashValue := newMask*16 + neighbor

			if !visited[hashValue] {
				visited[hashValue] = true
				queue = append(queue, State{newMask, neighbor, cur.dist + 1})
			}
		}
	}

	return -1
}
