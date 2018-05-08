// 数字的补数
// 给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。
// 注意:
// 给定的整数保证在32位带符号整数的范围内。
// 你可以假定二进制数不包含前导零位。

package main

import (
	"fmt"
)

func findComplement(num int) int {
	res := 0
	base := 1
	for num > 0 {
		if num%2 == 0 {
			res += base
		}

		num = num >> 1
		if num == 0 {
			break
		}
		base = base << 1
	}
	return res
}

func main() {
	fmt.Println(findComplement(5))
}
