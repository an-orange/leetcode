//汉明距离
// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
// 给出两个整数 x 和 y，计算它们之间的汉明距离。
// 注意：0 ≤ x, y < 231(31次方).
package main

import "fmt"

func hammingDistance(x int, y int) int {
	val := x ^ y
	fmt.Println(val)
	dist := 0
	for val != 0 {
		dist++
		val &= (val - 1)
	}
	return dist
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
