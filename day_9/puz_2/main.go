package main

import (
	"fmt"

	l "github.com/selcukusta/adventfocode/lib"
)

func main() {
	lines, ok := l.ReadAsInt("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		sum     = 0
		visited = make([]int, 0)
		found   = false
		// goal    = 127 //puz_1 sample result
		goal = 217430975 //puz_1 actual result
	)

	for i := range lines {
		for idx := i + 1; idx < len(lines); idx++ {
			next := lines[idx]
			sum += next
			visited = append(visited, next)
			if sum == goal {
				found = true
				break
			} else if sum > goal {
				sum = 0
				visited = make([]int, 0)
				break
			}
		}
		if found {
			break
		}

	}
	min, max, err := l.GetIntMinAndMaxValuesFromSlice(visited)
	if err != nil {
		panic("stop")
	}

	fmt.Println(min + max)
}
