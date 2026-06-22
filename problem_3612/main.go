package main

import (
	"fmt"
)

func main() {
	testcases := []string{"a#b%*", "z*#", "*%"}
	for _, testcase := range testcases {
		fmt.Println(processStr(testcase))
	}
}

func processStr(s string) string {
	str := ""
	for _, char := range s {
		switch char {
		case '#':
			str += str
		case '%':
			str = Reverse(str)
		case '*':
			if len(str) != 0 {
				str = str[:len(str)-1]
			}
		default:
			str += string(char)
		}
	}
	return str
}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
