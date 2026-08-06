package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	for _, t := range []string{"abcde", "xyzxyzxyzxyz", "aabbccddeeffgghhiiiiii"} {
		total := minimumPushes(t)
		fmt.Println(total)
	}
}

func minimumPushes(word string) int {
	arr := [26]int{}
	for _, c := range word {
		arr[c-'a']++
	}
	slices.SortFunc(arr[:], func(a, b int) int {
		return cmp.Compare(b, a) // Reverse order
	})
	iter, counter, total := 1, 8, 0
	for _, i := range arr {
		if i == 0 {
			break
		}
		total += i * iter
		counter--
		if counter == 0 {
			counter = 8
			iter++
		}
	}
	return total
}
