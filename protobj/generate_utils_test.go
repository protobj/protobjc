package protobj

import (
	"fmt"
	"testing"
)

func Test_I(t *testing.T) {
	t.Log("start test_format", t.Name())
	format := "${0} ${1} ${2} ${0} ${1}"
	params := []interface{}{"name", int32(2), "class"}
	fmt.Println(I(format, params))
}

func Test_NI(t *testing.T) {
	t.Log("start test_format", t.Name())
	format := "${name} ${age} ${class}"
	params := []interface{}{"name", "0index", "age", int32(2), "class", "3.2"}
	fmt.Println(NI(format, params))
}

func Test_N(t *testing.T) {
	t.Log("start test_format", t.Name())
	format := "${name} ${age} ${class}"
	params := map[string]interface{}{"name": "0index", "age": int32(2), "class": "3.2"}
	fmt.Println(N(format, params))
}
