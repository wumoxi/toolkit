package toolkit

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

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

// GetFieldNameByTagValue is to get an array of field names from the label value
// The parameter `source` must be a structure，
// The parameter `tagValue` specifies a structure tagValue.
func GetFieldNameByTagValue(source interface{}, tagValue string) (fields []string, err error) {
	stags, err := GetAllStructTags(source)
	if err != nil {
		return nil, err
	}

	for field, tags := range stags {
		for _, tag := range tags {
			for _, tv := range tag {
				if strings.Contains(tv, ",") {
					mtv := strings.Split(tv, ",")
					for _, stv := range mtv {
						if stv == tagValue {
							fields = append(fields, field)
						}
					}
				} else {
					if tv == tagValue {
						fields = append(fields, field)
					}
				}
			}
		}
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i] < fields[j]
	})
	return fields, nil
}
