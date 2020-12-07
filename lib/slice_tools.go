package lib

import "errors"

//Contains looks the item exits in the given slice and return the specific index. If it's not found, it returns -1.
func Contains(slice []int64, val int64) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//ContainsStr looks the item exits in the given slice and return the specific index. If it's not found, it returns -1.
func ContainsStr(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//GetMinAndMaxValuesFromSlice gets the minimum and maximum values from the specified slice.
func GetMinAndMaxValuesFromSlice(array []int64) (int64, int64, error) {
	if len(array) == 0 {
		return -1, -1, errors.New("Array is invalid")
	}
	var (
		max int64 = array[0]
		min int64 = array[0]
	)
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max, nil
}
