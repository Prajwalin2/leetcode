package main

func main() {
	testcases := []int{10203004}
	for _, t := range testcases {
		sumAndMultiply(t)
	}
}

func sumAndMultiply(n int) int64 {
	var multiplier, sum, no int64 = 1, 0, 0
	for n != 0 {
		var remainder int64 = int64(n % 10)
		if remainder != 0 {
			no += remainder * multiplier
			sum += remainder
			multiplier *= 10
		}
		n /= 10
	}
	return sum * no
}
