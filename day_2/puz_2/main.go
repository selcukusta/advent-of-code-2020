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
			policy        = strings.Split(v, " ")
			rng           = strings.Split(policy[0], "-")
			firstIndex, _ = strconv.Atoi(rng[0])
			lastIndex, _  = strconv.Atoi(rng[1])
			req           = strings.TrimSuffix(policy[1], ":")
			input         = policy[2]
		)
		idx1 := string(input[firstIndex-1])
		idx2 := string(input[lastIndex-1])

		if idx1 == req && idx2 == req {
			continue
		}

		if idx1 != req && idx2 != req {
			continue
		}
		valid++
	}
	fmt.Println(valid)
}
