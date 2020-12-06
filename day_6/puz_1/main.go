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
		total   = 0
		answers = make(map[string]string)
	)

	for _, line := range lines {
		if line == "" {
			total += len(answers)
			answers = make(map[string]string)
		} else {
			for _, q := range line {
				answers[string(q)] = `👌`
			}
		}
	}
	total += len(answers)
	fmt.Println(total)
}
