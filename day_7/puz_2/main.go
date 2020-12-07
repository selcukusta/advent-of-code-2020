package main

import (
	"fmt"
	"regexp"
	"strconv"

	l "github.com/selcukusta/adventfocode/lib"
)

const mainColor = "shiny gold"

var (
	valid         = 0
	pattern       = regexp.MustCompile(`^((?:\S+\s+){1}\S+)(.*)`)
	numberPattern = regexp.MustCompile(`(\d+)\s*(\w+\s{1}\w+)`)
)

func bags(lines []string) map[string]string {
	val := make(map[string]string)
	for _, line := range lines {
		match := pattern.FindAllStringSubmatch(line, -1)
		for i := range match {
			var (
				key   = match[i][1]
				value = match[i][2]
			)
			val[key] = value
		}
	}
	return val
}

func search(dict map[string]string, key string) int {
	val := 0
	match := numberPattern.FindAllStringSubmatch(dict[key], -1)
	for i := range match {
		var (
			count = match[i][1]
			color = match[i][2]
		)
		n, _ := strconv.Atoi(count)
		val += n + (n * search(dict, color))
	}
	return val
}

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}

	values := bags(lines)
	for k := range values {
		if k == mainColor {
			valid += search(values, mainColor)
		}
	}

	fmt.Println(valid)
}
