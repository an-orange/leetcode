// 给定一个数组 nums, 编写一个函数将所有 0 移动到它的末尾，同时保持非零元素的相对顺序。
// 例如， 定义 nums = [0, 1, 0, 3, 12]，调用函数之后， nums 应为 [1, 3, 12, 0, 0]。
package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	s := 0
	for k, v := range nums {
		if v != 0 {
			if k != s {
				nums[s], nums[k] = nums[k], nums[s]
			}
			s++
		}
	}
}

func main() {
	data := []int{0, 1, 0, 3, 12}
	moveZeroes(data)
	fmt.Println(data)
}
