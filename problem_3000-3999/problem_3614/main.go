package main

import "fmt"

func main() {
	testcases := []string{"a#b%*", "cd%#*#", "z*#"}
	testK := []int64{1, 3, 0}
	for i, testcase := range testcases {
		fmt.Println(processStr(testcase, testK[i]))
	}
}

func processStr(s string, k int64) byte {
	var l int64 = 0

	// Step 1
	for _, c := range s {
		if c == '*' {
			if l > 0 {
				l--
			}
		} else if c == '#' {
			l *= 2
		} else if c != '%' {
			l++
		}
	}

	if k >= l {
		return '.'
	}

	// Step 2
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]

		if c == '*' {
			l++
		} else if c == '#' {
			half := l / 2
			if k >= half {
				k -= half
			}
			l = half
		} else if c == '%' {
			k = l - 1 - k
		} else {
			if k == l-1 {
				return c
			}
			l--
		}
	}

	return '.'
}
