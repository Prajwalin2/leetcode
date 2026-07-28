package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, t := range []int{31, 22, 124} {
		ans := maxProduct(t)
		fmt.Println(ans)
	}
}

func maxProduct(n int) int {
	arr := [2]int{}
	s := strconv.Itoa(n)
	for _, i := range s {
		val := int(i - '0')
		if val > arr[0] {
			arr[1] = arr[0]
			arr[0] = val
		} else if val > arr[1] {
			arr[1] = val
		}
	}
	return arr[0] * arr[1]
}
