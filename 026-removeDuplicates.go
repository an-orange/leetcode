// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
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
