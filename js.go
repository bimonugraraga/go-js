package gojs

import (
	"errors"
	"reflect"
)

func Ternary(condition bool, result1 interface{}, result2 interface{}) interface{} {
	if condition {
		return result1
	}
	return result2
}

func IsFalsy(params interface{}) bool {
	var err error
	errNew := errors.New("Err1")
	result := false
	switch params {
	case "", nil, 0:
		result = true
	default:
		if reflect.TypeOf(params) == reflect.TypeOf(errNew) || reflect.TypeOf(params) == reflect.TypeOf(err) {
			return true
		}
		result = false
	}
	return result
}