//反转字符串
//请编写一个函数，其功能是将输入的字符串反转过来。
package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	ss := []byte(s)
	for from, to := 0, len(ss)-1; from < to; from, to = from+1, to-1 {
		ss[from], ss[to] = ss[to], ss[from]
	}
	return string(ss)
}

func reverseString2(s string) string {
	result := make([]string, len(s))
	for i, one := range s {
		result[len(s)-i-1] = string(one)
	}
	return strings.Join(result, "")
}

func main() {
	str := "abcdef"
	fmt.Println(reverseString(str))
}
