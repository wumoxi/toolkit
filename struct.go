package toolkit

import (
	"sort"
	"strings"
)

// GetFieldsNameByTag is to get an array of field names from the label value
// The parameter `source` must be a structure，
// The parameter `tagValue` specifies a structure tagValue.
func GetFieldsNameByTag(source interface{}, tagValue string) (fields []string, err error) {
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
