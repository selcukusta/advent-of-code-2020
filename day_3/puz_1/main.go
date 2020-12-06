package main

import (
	"fmt"

	l "github.com/selcukusta/adventfocode/lib"
)

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}
	var (
		right    = 3
		bottom   = 1
		trees    = 0
		rightIdx = right
	)
	for row := bottom; row < len(lines); row += bottom {
		current := lines[row]
		if rightIdx >= len(current) {
			rightIdx -= len(current)
		}

		var block = l.Substring(current, rightIdx, 1)
		rightIdx += right
		if block == "#" {
			trees++
		}
	}
	fmt.Println(trees)
}
