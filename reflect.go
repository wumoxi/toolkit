package toolkit

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName 获取函数名称
func GetFuncName(fn interface{}) (string, error) {
	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), nil
	}
	return "", errors.New("fn parameter is not func type")
}

// GetAllStructTag gets the struct field name and tag mappings，the parameter `source` must be a structure
func GetAllStructTags(source interface{}) (map[string][]map[string]string, error) {
	if reflect.TypeOf(source).Kind() != reflect.Struct {
		return nil, fmt.Errorf("invalid parameter source")
	}
	r := make(map[string][]map[string]string)
	t := reflect.TypeOf(source)
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		ft := t.Field(i).Tag
		if ft != "" {
			if _, exists := r[fn]; !exists {
				parts := strings.Fields(string(ft))
				var tags = make([]map[string]string, 0)
				for _, part := range parts {
					single := strings.Split(part, ":")
					if single == nil || len(single) != 2 {
						continue
					}
					tn := single[0]
					tv := strings.Trim(single[1], "\"")
					if tv == "-" {
						continue
					}
					tags = append(tags, map[string]string{tn: tv})
				}
				r[fn] = tags
			}
		}
	}
	return r, nil
}
