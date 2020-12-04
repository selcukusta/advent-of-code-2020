package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read is using to read a whole file into memory and return a slice.
func Read(path string) ([]string, bool) {
	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, true
}

func main() {
	lines, ok := Read("input.txt")
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
