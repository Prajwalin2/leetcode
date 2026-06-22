package main

import "fmt"

func main() {
	testcases := [][]int{{5, 2, 6, 1},
		{-1},
		{-1, -1},
		{0, -1, -1},
	}
	for _, t := range testcases {
		fmt.Println(countSmaller(t))
	}
}
func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}

	var mergeSort func(left, right int)

	mergeSort = func(left, right int) {
		if right-left <= 1 {
			return
		}

		mid := left + (right-left)/2

		mergeSort(left, mid)
		mergeSort(mid, right)

		temp := make([]int, 0, right-left)

		i := left
		j := mid
		rightSmallerCount := 0

		for i < mid && j < right {
			if nums[indexes[j]] < nums[indexes[i]] {
				rightSmallerCount++
				temp = append(temp, indexes[j])
				j++
			} else {
				res[indexes[i]] += rightSmallerCount
				temp = append(temp, indexes[i])
				i++
			}
		}

		for i < mid {
			res[indexes[i]] += rightSmallerCount
			temp = append(temp, indexes[i])
			i++
		}

		for j < right {
			temp = append(temp, indexes[j])
			j++
		}

		for k := 0; k < len(temp); k++ {
			indexes[left+k] = temp[k]
		}
	}

	mergeSort(0, n)

	return res
}
