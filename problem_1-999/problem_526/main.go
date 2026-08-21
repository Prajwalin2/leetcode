package main

import "fmt"

func main() {
	for i := range 10 {
		ans := countArrangement(i + 1)
		fmt.Println(i, ":", ans)
	}
	// countArrangement(10)
	// fmt.Println(ans)
}

func countArrangement(n int) int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = i + 1
	}

	return recursive(arr, 1)
}

func recursive(arr []int, pos int) int {
	if len(arr) == 0 {
		return 1
	}
	var count int
	for i, ele := range arr {
		if ele%pos == 0 || pos%ele == 0 {
			next := make([]int, 0, len(arr)-1)
			next = append(next, arr[:i]...)
			next = append(next, arr[i+1:]...)
			count += recursive(next, pos+1)
		}
	}
	return count
}
