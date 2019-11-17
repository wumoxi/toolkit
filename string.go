package toolkit

import (
	"strings"
)

// ConvertToSmallCamelCase 转换为小驼峰命令法
func ConvertToSmallCamelCase(str string) string {
	if len(str) <= 0 {
		return ""
	}

	words := GetWordsFormString(str)

	if len(words) <= 1 {
		return strings.Join(words, "")
	}

	for index, w := range words {
		if index == 0 {
			continue
		}
		words[index] = strings.Title(w)
	}

	return strings.Join(words, "")
}

// ConvertToBigCamelCase 转换为大驼峰命令法
func ConvertToBigCamelCase(str string) string {
	if len(str) <= 0 {
		return ""
	}

	words := GetWordsFormString(str)

	if len(words) <= 1 {
		return strings.Join(words, "")
	}

	for index, w := range words {
		words[index] = strings.Title(w)
	}

	return strings.Join(words, "")
}
