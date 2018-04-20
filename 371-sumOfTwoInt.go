//两整数之和
//不使用运算符 + 和-，计算两整数a 、b之和。

package main

import "fmt"

func getSum(a int, b int) int {
	if 0 == b {
		return a //当没有进位时
	}
	sum := a ^ b
	carry := (a & b) << 1 //进位
	return getSum(sum, carry)
}

func main() {

	fmt.Println(getSum(3, 5))
}
