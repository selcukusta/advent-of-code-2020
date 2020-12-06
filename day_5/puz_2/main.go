package main

import (
	"fmt"
	"strconv"
	"strings"

	l "github.com/selcukusta/adventfocode/lib"
)

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		seats    = make([]int64, 0)
		replacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	)

	//algorithm from https://github.com/ogun/advent-of-code-2020/blob/main/main.py#L182
	for _, line := range lines {
		seat, err := strconv.ParseInt(replacer.Replace(line), 2, 0)
		if err != nil {
			panic("stop")
		}
		seats = append(seats, seat)
	}
	min, max, err := l.GetMinAndMaxValuesFromSlice(seats)
	if err != nil {
		panic("stop")
	}
	for i := min; i <= max; i++ {
		if l.Contains(seats, i) == -1 {
			fmt.Println(i)
			break
		}
	}
}
