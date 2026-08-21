package main

import "fmt"

func main() {
	for _, test := range []string{"bcbbbcba", "aaaa", "aabbcc"} {
		ans := maximumLengthSubstring(test)
		fmt.Println(ans)
	}
}

func maximumLengthSubstring(s string) int {
	arr := make([]int, 26)
	l := 0
	mx := 0
	for i, ch := range s {
		arr[ch-'a']++
		for arr[ch-'a'] > 2 {
			arr[s[l]-'a']--
			l++
		}
		mx = max(mx, i-l+1)
	}
	return mx
}
