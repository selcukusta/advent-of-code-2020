package main

import (
	"fmt"
	"regexp"
	"strconv"

	l "github.com/selcukusta/adventfocode/lib"
)

type move struct {
	Action string
	Step   int
}

func parseMove(idx int, container []string) move {
	step, _ := strconv.Atoi(container[2])
	m := move{
		Action: container[1],
		Step:   step,
	}
	return m
}

func actions(lines []string) []move {
	var (
		pattern = regexp.MustCompile(`^(jmp|acc|nop)\s{1}([\+|\-][0-9]{1,})$`)
		val     = make([]move, 0)
	)
	for idx, line := range lines {
		match := pattern.FindAllStringSubmatch(line, -1)
		for i := range match {
			val = append(val, parseMove(idx, match[i]))
		}
	}
	return val
}

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}
	var (
		accumulator = 0
		visited     = make(map[int]int)
		tried       = make(map[int]int)
		reversed    = false
	)

	acts := actions(lines)
	for idx := 0; idx < len(acts); {

		if _, exists := visited[idx]; exists {
			accumulator = 0
			visited = make(map[int]int)
			reversed = false
			idx = 0
		}

		visited[idx] = idx
		current := acts[idx]

		switch current.Action {
		case "acc":
			accumulator += current.Step
			idx++

		case "jmp", "nop":
			if _, exists := tried[idx]; !exists && !reversed {
				tried[idx] = idx
				//jmp -> nop || nop -> jmp
				reversed = true
				idx++
				continue
			}

			if current.Action == "jmp" {
				idx += current.Step
			} else {
				idx++
			}
		}
	}
	fmt.Println(accumulator)
}
