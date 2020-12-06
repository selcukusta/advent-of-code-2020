package lib

import (
	"bufio"
	"os"
	"strconv"
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

// ReadAsInt is using to read a whole file into memory and return a int slice.
func ReadAsInt(path string) ([]int, bool) {
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
