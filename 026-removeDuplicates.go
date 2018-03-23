package main

import "fmt"

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}

	j := 0

	for i := 1; i < len; i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j++
		}

	}
	fmt.Println(nums)

	return j + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 3, 5, 6, 8, 8, 9}))
}
