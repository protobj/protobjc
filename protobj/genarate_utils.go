package protobj

import (
	"strconv"
	"strings"
)

func FirstUpper(value string) string {
	firstString, lastString := value[0:1], value[1:]
	return strings.ToUpper(firstString) + lastString
}

func I(format string, params ...interface{}) string {
	for i, v := range params {
		index := "${" + strconv.Itoa(i) + "}"
		var value = ToString(v)
		format = strings.ReplaceAll(format, index, value)
	}
	return format
}

func NI(format string, params ...interface{}) string {
	for i := 0; i < len(params); i += 2 {
		v := params[i+1]
		index := "${" + params[i].(string) + "}"
		var value = ToString(v)
		format = strings.ReplaceAll(format, index, value)
	}

	return format
}
func N(format string, params map[string]interface{}) string {
	for k, v := range params {
		index := "${" + k + "}"
		var value = ToString(v)
		format = strings.ReplaceAll(format, index, value)
	}
	return format
}

func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int:
		return strconv.Itoa(value.(int))
	case []string:
		return "[" + strings.Join(value.([]string), ",") + "]"
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return ""
	}
}
