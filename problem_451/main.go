package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	for _, t := range []string{"tree", "Aabb"} {
		ans := frequencySort(t)
		fmt.Println(ans)
	}
}

func frequencySort(s string) string {
	m := map[rune]int{}
	for _, a := range s {
		m[a]++
	}
	chars := make([]rune, 0, len(m))
	for ch := range m {
		chars = append(chars, ch)
	}

	slices.SortFunc(chars, func(a rune, b rune) int {
		return cmp.Compare(m[b], m[a])
	})
	sb := strings.Builder{}
	for _, c := range chars {
		for range m[c] {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
