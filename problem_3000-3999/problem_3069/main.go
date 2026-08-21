package main

func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	one, sec := 0, 0
	for i := 2; i < len(nums); i++ {
		num := nums[i]
		if arr1[one] > arr2[sec] {
			arr1 = append(arr1, num)
			one++
		} else {
			arr2 = append(arr2, num)
			sec++
		}
	}
	return append(arr1, arr2...)
}
