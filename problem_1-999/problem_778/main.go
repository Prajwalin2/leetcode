package main

import "container/heap"

func main() {
	swimInWater([][]int{{7, 23, 21, 9, 5}, {3, 20, 8, 18, 15}, {14, 13, 1, 0, 22}, {2, 10, 24, 17, 12}, {6, 16, 19, 4, 11}})
}

type pos struct {
	val      int
	position []int
}

var traversal = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type PosHeap []pos

func (h PosHeap) Len() int           { return len(h) }
func (h PosHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h PosHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PosHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pos))
}

func (h *PosHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func swimInWater(grid [][]int) int {
	h := &PosHeap{}
	heap.Init(h)
	heap.Push(h, pos{grid[0][0], []int{0, 0}})
	mx := grid[0][0]
	grid[0][0] = -1
	for {
		p := heap.Pop(h).(pos)
		mx = max(mx, p.val)
		if p.position[0] == len(grid)-1 && p.position[1] == len(grid[0])-1 {
			break
		}
		for _, t := range traversal {
			x := t[0] + p.position[0]
			y := t[1] + p.position[1]
			if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] != -1 {
				heap.Push(h, pos{grid[x][y], []int{x, y}})
				grid[x][y] = -1
			}
		}
	}
	return mx
}
