package main

import (
	"fmt"
	"strings"
)

func main() {
	testcases := []string{"z", "babab", "daccad"}
	for _, test := range testcases {
		s := smallestPalindrome(test)
		fmt.Println(s)
	}
}

func smallestPalindrome(s string) string {
	pre, middle, post := strings.Builder{}, strings.Builder{}, strings.Builder{}

	ele := [26]int{}
	for _, r := range s {
		ele[r-'a']++
	}
	for i := range ele {
		for ele[i] != 0 {
			if ele[i] == 1 {
				middle.WriteRune(rune('a' + i))
				ele[i]--
			} else {
				pre.WriteRune(rune('a' + i))
				post.WriteRune(rune('a' + i))
				ele[i] -= 2
			}
		}
	}
	return pre.String() + middle.String() + reverse(post.String())
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
