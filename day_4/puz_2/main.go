package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func validRange(val string, min int, max int) bool {
	numericValue, err := strconv.Atoi(val)
	if err != nil || !(numericValue >= min && numericValue <= max) {
		return false
	}
	return true
}

func validRegex(val string, pattern *regexp.Regexp) bool {
	return pattern.Match([]byte(val))
}

func main() {
	lines, ok := Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		modelPattern        = regexp.MustCompile(`(?P<key>byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(?P<value>\S+)`)
		hexPattern          = regexp.MustCompile(`^#(?:[0-9a-f]{3}){2}$`)
		paddedNumberPattern = regexp.MustCompile(`^\d{9}$`)
		valid               = 0
		passports           = make([]map[string]interface{}, 0)
		temp                = make(map[string]interface{})
	)

	for idx, line := range lines {
		match := modelPattern.FindAllStringSubmatch(line, -1)
		for i := range match {
			if match[i][1] == "cid" {
				continue
			}
			temp[match[i][1]] = match[i][2]
		}
		if line == "" || idx+1 == len(lines) {
			passports = append(passports, temp)
			temp = make(map[string]interface{})
		}
	}

	for _, v := range passports {
		if len(v) != 7 {
			continue
		}

		val, ok := v["byr"]
		if !ok || !validRange(val.(string), 1920, 2002) {
			continue
		}

		val, ok = v["iyr"]
		if !ok || !validRange(val.(string), 2010, 2020) {
			continue
		}

		val, ok = v["eyr"]
		if !ok || !validRange(val.(string), 2020, 2030) {
			continue
		}

		val, ok = v["hgt"]
		if ok {
			literalValue, ok := val.(string)
			if !ok || !(strings.HasSuffix(literalValue, "cm") || strings.HasSuffix(literalValue, "in")) {
				continue
			}

			if strings.HasSuffix(literalValue, "cm") && !validRange(strings.TrimSuffix(literalValue, "cm"), 150, 193) {
				continue
			}

			if strings.HasSuffix(literalValue, "in") && !validRange(strings.TrimSuffix(literalValue, "in"), 59, 76) {
				continue
			}
		}

		val, ok = v["hcl"]
		if ok {
			literalValue, ok := val.(string)
			if !ok || !validRegex(literalValue, hexPattern) {
				continue
			}
		}

		val, ok = v["ecl"]
		if ok {
			literalValue, ok := val.(string)
			if !ok || !contains(strings.Split("amb blu brn gry grn hzl oth", " "), literalValue) {
				continue
			}
		}

		val, ok = v["pid"]
		if ok {
			literalValue, ok := val.(string)
			if !ok || !validRegex(literalValue, paddedNumberPattern) {
				continue
			}
		}
		valid++
	}

	fmt.Println(valid)
}
