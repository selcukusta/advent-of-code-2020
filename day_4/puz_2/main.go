package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	l "github.com/selcukusta/adventfocode/lib"
	"gopkg.in/validator.v2"
)

type passport struct {
	BirthYear      int    `validate:"min=1920,max=2002"`
	IssueYear      int    `validate:"min=2010,max=2020"`
	ExpirationYear int    `validate:"min=2020,max=2030"`
	Height         string `validate:"nonzero,height"`
	HairColor      string `validate:"nonzero,regexp=^#(?:[0-9a-f]{3}){2}$"`
	EyeColor       string `validate:"nonzero,regexp=^(amb|blu|brn|gry|grn|hzl|oth){1}$"`
	PassportID     string `validate:"nonzero,regexp=^[0-9]{9}$"`
}

func heightValidator(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return validator.ErrUnsupported
	}

	val := st.String()
	if !strings.HasSuffix(val, "cm") && !strings.HasSuffix(val, "in") {
		return errors.New("Undefined unit")
	}

	if strings.HasSuffix(val, "cm") && !l.ValidRange(strings.TrimSuffix(val, "cm"), 150, 193) {
		return errors.New("Invalid range for cm")
	}

	if strings.HasSuffix(val, "in") && !l.ValidRange(strings.TrimSuffix(val, "in"), 59, 76) {
		return errors.New("Invalid range for in")
	}

	return nil
}

func main() {
	err := validator.SetValidationFunc("height", heightValidator)
	if err != nil {
		panic("stop")
	}

	lines, ok := l.Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		modelPattern = regexp.MustCompile(`(?P<key>byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(?P<value>\S+)`)
		valid        = 0
		passports    = make([]map[string]interface{}, 0)
		temp         = make(map[string]interface{})
	)

	for _, line := range lines {
		match := modelPattern.FindAllStringSubmatch(line, -1)
		if line == "" {
			passports = append(passports, temp)
			temp = make(map[string]interface{})
		}
		for i := range match {
			if match[i][1] == "cid" {
				continue
			}
			temp[match[i][1]] = match[i][2]
		}
	}
	passports = append(passports, temp)

	for _, v := range passports {
		if validator.Validate(&passport{
			BirthYear:      l.GetIntOrDefault(v, "byr"),
			IssueYear:      l.GetIntOrDefault(v, "iyr"),
			ExpirationYear: l.GetIntOrDefault(v, "eyr"),
			Height:         l.GetStringOrDefault(v, "hgt"),
			HairColor:      l.GetStringOrDefault(v, "hcl"),
			EyeColor:       l.GetStringOrDefault(v, "ecl"),
			PassportID:     l.GetStringOrDefault(v, "pid"),
		}) == nil {
			valid++
		}
	}
	fmt.Println(valid)
}
