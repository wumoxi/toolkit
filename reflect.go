package toolkit

import (
	"errors"
	"reflect"
	"runtime"
)

// GetFuncName 获取函数名称
func GetFuncName(fn interface{}) (string, error) {
	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), nil
	}
	return "", errors.New("fn parameter is not func type")
}
