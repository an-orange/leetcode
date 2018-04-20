//宝石与石头
//给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。
//S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
// J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。
package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	MapJ := make(map[rune]int)
	for k, v := range J {
		MapJ[v] = k
	}
	var num int
	for _, v := range S {
		if _, ok := MapJ[v]; ok {
			num++
		}
	}
	return num
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
