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
		right  = []int{1, 3, 5, 7, 1}
		bottom = []int{1, 1, 1, 1, 2}
		mTree  = 1
	)

	for i := 0; i < len(right); i++ {
		var (
			rightIdx = right[i]
			trees    = 0
		)
		for row := bottom[i]; row < len(lines); row += bottom[i] {
			current := lines[row]
			if rightIdx >= len(current) {
				rightIdx -= len(current)
			}
			var block = l.Substring(current, rightIdx, 1)
			rightIdx += right[i]
			if block == "#" {
				trees++
			}
		}
		mTree = mTree * trees
	}

	fmt.Println(mTree)
}
