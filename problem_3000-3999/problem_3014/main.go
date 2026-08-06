package main

func main() {
	for _, t := range []string{"abc", "abcdefg", "abcdefghijkl"} {
		minimumPushes(t)
	}
}

func minimumPushes(word string) int {
	l := len(word)
	count := 0
	iter := 1
	for l > 0 {
		if l > 8 {
			count += iter * 8
		} else {
			count += iter * l
		}
		l -= 8
		iter++
	}
	return count
}
