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
		cs       = combin.Combinations(len(lines), 2)
		multiply = 0
	)
	for _, c := range cs {
		if lines[c[0]]+lines[c[1]] == 2020 {
			multiply = lines[c[0]] * lines[c[1]]
			break
		}
	}
	fmt.Println(multiply)
}
