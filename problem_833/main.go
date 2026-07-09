package main

import (
	"fmt"
	"sort"
)

type testcase struct {
	s       string
	indices []int
	sources []string
	targets []string
}

func main() {
	testcases := []testcase{
		// {"abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}},
		// {"abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}},
		// {"abcd", []int{0, 1}, []string{"ac", "bc"}, []string{"eee", "ffff"}},
		{"vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}},
		{"abcde", []int{2, 2, 3}, []string{"cde", "cdef", "dk"}, []string{"fe", "f", "xyz"}},
		{"abcdef", []int{2, 2}, []string{"cdef", "feg"}, []string{"feg", "abc"}},
	}
	for _, t := range testcases {
		s := findReplaceString(t.s, t.indices, t.sources, t.targets)
		fmt.Println(s)
	}
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	order := make([]int, len(indexes))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(i, j int) bool {
		return indexes[order[i]] < indexes[order[j]]
	})

	prevIndex, prevLength := 0, 0
	result := ""
	for _, replacement := range order {
		index := indexes[replacement]
		source := sources[replacement]
		target := targets[replacement]
		consumed := prevIndex + prevLength
		if index < consumed || index+len(source) > len(S) {
			continue
		}
		if S[index:index+len(source)] == source {
			result += S[consumed:index] + target
			prevIndex, prevLength = index, len(source)
		}
	}
	return result + S[prevIndex+prevLength:]
}
