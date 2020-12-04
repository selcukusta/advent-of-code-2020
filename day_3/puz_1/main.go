package main

import (
	"bufio"
	"fmt"
	"os"
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

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func main() {
	lines, ok := Read("input.txt")
	if !ok {
		panic("stop")
	}
	var (
		right    = 3
		bottom   = 1
		trees    = 0
		rightIdx = right
	)
	for row := bottom; row < len(lines); row += bottom {
		current := lines[row]
		if rightIdx >= len(current) {
			rightIdx -= len(current)
		}

		var block = substr(current, rightIdx, 1)
		rightIdx += right
		if block == "#" {
			trees++
		}
	}
	fmt.Println(trees)
}
