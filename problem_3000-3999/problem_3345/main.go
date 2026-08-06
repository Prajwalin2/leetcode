package main

import "fmt"

func main() {
	ans := smallestNumber(1, 2)
	fmt.Println(ans)
}

func smallestNumber(n int, t int) int {
	for n%10 != 0 {
		i := n
		mult := 1
		for i != 0 {
			mult *= i % 10
			i = i / 10
		}
		if mult%t == 0 {
			return n
		}
		n++
	}
	return n
}
