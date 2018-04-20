// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//var myMap map[int]int
	myMap := make(map[int]int)

	for k, v := range nums {
		myMap[v] = k
	}
	for k, v := range nums {
		index := target - v
		if _, ok := myMap[index]; ok {
			if myMap[index] != k {
				return []int{k, myMap[index]}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{3, 3}
	fmt.Println(twoSum(a, 6))
}
