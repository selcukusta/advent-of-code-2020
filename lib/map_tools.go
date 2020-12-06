package lib

import "strconv"

//GetStringOrDefault is using to get map value as string or "" (empty string).
func GetStringOrDefault(model map[string]interface{}, key string) string {
	val, ok := model[key]
	if ok {
		return val.(string)
	}
	return ""
}

//GetIntOrDefault is using to get map value as int or 0 (zero).
func GetIntOrDefault(model map[string]interface{}, key string) int {
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
