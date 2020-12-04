package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

// Read is using to read a whole file into memory and return a slice.
func Read(path string) ([]int, bool) {
	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lines = append(lines, val)

		}
	}
	return lines, true
}

func main() {
	lines, ok := Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		cs       = combin.Combinations(len(lines), 3)
		multiply = 0
	)
	for _, c := range cs {
		if lines[c[0]]+lines[c[1]]+lines[c[2]] == 2020 {
			multiply = lines[c[0]] * lines[c[1]] * +lines[c[2]]
			break
		}
	}
	fmt.Println(multiply)
}
