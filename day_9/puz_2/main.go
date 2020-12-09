package main

import (
	"fmt"

	l "github.com/selcukusta/adventfocode/lib"
	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	lines, ok := l.ReadAsInt("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		result = 0
		limit  = 25
	)

	for i := limit; i < len(lines); i++ {
		var (
			previous = lines[i-limit : i]
			c        = combin.Combinations(len(previous), 2)
			sum      = make([]int, 0)
		)

		for _, x := range c {
			sum = append(sum, previous[x[0]]+previous[x[1]])
		}

		current := lines[i]
		if l.Contains(sum, current) == -1 {
			result = current
			break
		}
	}

	fmt.Println(result)
}
