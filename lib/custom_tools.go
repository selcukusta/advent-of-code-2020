package lib

import "strconv"

//ValidRange is using to check provided value is in valid range.
func ValidRange(val string, min int, max int) bool {
	numericValue, err := strconv.Atoi(val)
	if err != nil || !(numericValue >= min && numericValue <= max) {
		return false
	}
	return true
}
