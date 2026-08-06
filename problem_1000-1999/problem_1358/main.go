package main

import "fmt"

func main() {
	testcases := []string{"abcabc", "aaacb"}
	for _, test := range testcases {
		fmt.Println(numberOfSubstrings(test))
	}
}

func numberOfSubstrings(s string) int {
	count := 0
	left := 0
	right := 0
	a, b, c := 0, 0, 0
	for left <= right && right < len(s) {
		switch s[right] {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}
		for a > 0 && b > 0 && c > 0 {
			count += len(s) - right
			switch s[left] {
			case 'a':
				a--
			case 'b':
				b--
			case 'c':
				c--
			}
			left++
		}
		right++
	}
	return count
}
