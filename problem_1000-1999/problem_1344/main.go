package main

import (
	"fmt"
	"math"
)

func main() {
	test := [][]int{{12, 30}, {3, 30}, {3, 15}}
	//test := [][]int{{1, 57}}
	for _, t := range test {
		fmt.Println(angleClock(t[0], t[1]))
	}
}
func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	if minutes == 60 {
		minutes = 0
	}
	h := float64(hour)
	m := float64(minutes)
	ha := (h * 30) + (m / 2)
	ma := m * 6
	total := math.Abs(ha - ma)
	if total > 180 {
		total = 360 - total
	}
	return total
}
