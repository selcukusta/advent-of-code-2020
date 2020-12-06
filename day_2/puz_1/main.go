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

	valid := 0
	for _, v := range lines {
		var (
			policy = strings.Split(v, " ")
			rng    = strings.Split(policy[0], "-")
			min, _ = strconv.Atoi(rng[0])
			max, _ = strconv.Atoi(rng[1])
			req    = strings.TrimSuffix(policy[1], ":")
			input  = policy[2]
			count  = strings.Count(input, req)
		)
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
}
