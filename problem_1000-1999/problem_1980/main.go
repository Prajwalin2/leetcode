package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for _, test := range [][]string{{"00", "01"}} {
		findDifferentBinaryString(test)
	}
}

func findDifferentBinaryString(nums []string) string {
	m := map[int64]struct{}{}
	for _, n := range nums {
		i, err := strconv.ParseInt(n, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		m[i] = struct{}{}
	}
	for i := int64(0); i < math.MaxInt64; i++ {
		if _, ok := m[i]; !ok {
			s := strconv.FormatInt(i, 2)
			for len(s) != len(nums) {
				s = "0" + s
			}
			return s
		}
	}
	return ""
}
