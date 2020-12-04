package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var (
		pattern = regexp.MustCompile(`(?P<key>byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(?P<value>\S+)`)
		valid   = 0
		temp    = make(map[string]interface{})
	)

	for idx, line := range lines {
		match := pattern.FindAllStringSubmatch(line, -1)
		for i := range match {
			if match[i][1] == "cid" {
				continue
			}
			temp[match[i][1]] = match[i][2]
		}
		if line == "" || idx+1 == len(lines) {
			if len(temp) == 7 {
				valid++
			}
			temp = make(map[string]interface{})
		}
	}
	fmt.Println(valid)
}
