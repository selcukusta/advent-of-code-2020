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
