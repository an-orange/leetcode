//判断路线成圈
// 初始位置 (0, 0) 处有一个机器人。给出它的一系列动作，判断这个机器人的移动路线是否形成一个圆圈，
// 换言之就是判断它是否会移回到原来的位置。

// 移动顺序由一个字符串表示。每一个动作都是由一个字符来表示的。
// 机器人有效的动作有 R（右），L（左），U（上）和 D（下）。
// 输出应为 true 或 false，表示机器人移动路线是否成圈。

package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") &&
		strings.Count(moves, "L") == strings.Count(moves, "R")
}

func judgeCircle2(moves string) bool {
	coordX, coordY := 0, 0
	for _, v := range []byte(moves) {
		switch v {
		case 'R':
			coordX++
		case 'L':
			coordX--
		case 'U':
			coordY++
		case 'D':
			coordY--
		}
	}
	if coordX == 0 && coordY == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(judgeCircle("UDUD"))
}
