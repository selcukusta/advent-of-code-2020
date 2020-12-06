package main

import (
	"fmt"
	"regexp"

	l "github.com/selcukusta/adventfocode/lib"
)

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		pattern = regexp.MustCompile(`(?P<key>byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(?P<value>\S+)`)
		valid   = 0
		temp    = make(map[string]interface{})
	)

	onValidated := func() {
		if len(temp) == 7 {
			valid++
		}
	}

	for _, line := range lines {
		match := pattern.FindAllStringSubmatch(line, -1)
		if line == "" {
			onValidated()
			temp = make(map[string]interface{})
		}
		for i := range match {
			if match[i][1] == "cid" {
				continue
			}
			temp[match[i][1]] = match[i][2]
		}
	}
	onValidated()
	fmt.Println(valid)
}
