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
		total              = 0
		personCountByGroup = 0
		answers            = make(map[string]int)
	)

	onCalculated := func(answers map[string]int, personCount int) {
		for _, y := range answers {
			if y == personCount {
				total++
			}
		}
	}

	reset := func() {
		personCountByGroup = 0
		answers = make(map[string]int)
	}

	for _, line := range lines {
		if line == "" {
			onCalculated(answers, personCountByGroup)
			reset()
		} else {
			personCountByGroup++
			for _, q := range line {
				_, exists := answers[string(q)]
				if exists {
					answers[string(q)]++
				} else {
					answers[string(q)] = 1
				}
			}
		}
	}
	onCalculated(answers, personCountByGroup)
	fmt.Println(total)
}
