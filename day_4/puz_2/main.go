package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

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

func validRange(val string, min int, max int) bool {
	numericValue, err := strconv.Atoi(val)
	if err != nil || !(numericValue >= min && numericValue <= max) {
		return false
	}
	return true
}

func getStringOrDefault(model map[string]interface{}, key string) string {
	val, ok := model[key]
	if ok {
		return val.(string)
	}
	return ""
}

func getIntOrDefault(model map[string]interface{}, key string) int {
	val, ok := model[key]
	if ok {
		conv, err := strconv.Atoi(val.(string))
		if err != nil {
			return 0
		}
		return conv
	}
	return 0
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

	if strings.HasSuffix(val, "cm") && !validRange(strings.TrimSuffix(val, "cm"), 150, 193) {
		return errors.New("Invalid range for cm")
	}

	if strings.HasSuffix(val, "in") && !validRange(strings.TrimSuffix(val, "in"), 59, 76) {
		return errors.New("Invalid range for in")
	}

	return nil
}

func main() {
	validator.SetValidationFunc("height", heightValidator)
	lines, ok := Read("input.txt")
	if !ok {
		panic("stop")
	}

	var (
		modelPattern = regexp.MustCompile(`(?P<key>byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(?P<value>\S+)`)
		valid        = 0
		passports    = make([]map[string]interface{}, 0)
		temp         = make(map[string]interface{})
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
		p := &passport{
			BirthYear:      getIntOrDefault(v, "byr"),
			IssueYear:      getIntOrDefault(v, "iyr"),
			ExpirationYear: getIntOrDefault(v, "eyr"),
			Height:         getStringOrDefault(v, "hgt"),
			HairColor:      getStringOrDefault(v, "hcl"),
			EyeColor:       getStringOrDefault(v, "ecl"),
			PassportID:     getStringOrDefault(v, "pid"),
		}

		err := validator.Validate(p)
		if err != nil {
			continue
		}
		valid++
	}
	fmt.Println(valid)
}
