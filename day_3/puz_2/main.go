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
		right  = []int{1, 3, 5, 7, 1}
		bottom = []int{1, 1, 1, 1, 2}
		mTree  = 1
	)

	for i := 0; i < len(right); i++ {
		var (
			rightIdx = right[i]
			trees    = 0
		)
		for row := bottom[i]; row < len(lines); row += bottom[i] {
			current := lines[row]
			if rightIdx >= len(current) {
				rightIdx -= len(current)
			}
			var block = substr(current, rightIdx, 1)
			rightIdx += right[i]
			if block == "#" {
				trees++
			}
		}
		mTree = mTree * trees
	}

	fmt.Println(mTree)
}
