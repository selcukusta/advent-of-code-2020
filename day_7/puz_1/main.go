package main

import (
	"fmt"
	"regexp"
	"strings"

	l "github.com/selcukusta/adventfocode/lib"
)

const mainColor = "shiny gold"

var (
	valid   = 0
	pattern = regexp.MustCompile(`^((?:\S+\s+){1}\S+)(.*)`)
	checked = make([]string, 0)
)

func search(dict map[string]string, val string) {
	for k, v := range dict {
		if l.Contains(checked, k) == -1 && strings.Contains(v, val) {
			checked = append(checked, k)
			valid++
			search(dict, k)
		}
	}
}

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

func main() {
	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}

	search(bags(lines), mainColor)
	fmt.Println(valid)
}
