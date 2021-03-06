package toolkit

import (
	"strings"
	"unicode"
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

// Underline2SmallCamelCase 下划线转换为小驼峰
func Underline2SmallCamelCase(str string) string {
	s := strings.Split(strings.ToLower(str), "_")
	return ConvertToSmallCamelCase(strings.Join(s, " "))
}

// Underline2BigCamelCase 下划线转换为大驼峰
func Underline2BigCamelCase(str string) string {
	s := strings.Split(strings.ToLower(str), "_")
	return ConvertToBigCamelCase(strings.Join(s, " "))
}

// CamelCase2Underline 使用驼峰转下划线生成标记
func CamelCase2Underline(str string) string {
	return strings.Join(AllowCapitalGenerateSlice(str), Underline)
}

// SmallCamelCase2BigCamelCase 使用小驼峰转大驼峰
func SmallCamelCase2BigCamelCase(str string) string {
	words := AllowCapitalGenerateSlice(str)
	for index, word := range words {
		words[index] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// BigCamelCase2SmallCamelCase 使用大驼峰转小驼峰
func BigCamelCase2SmallCamelCase(str string) string {
	words := AllowCapitalGenerateSlice(str)
	for index, word := range words {
		if index == 0 {
			continue
		}
		words[index] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// AllowCapitalGenerateSlice 将传递参数 str 通过单词首字母大写，进行分割返回全小写单词切片
func AllowCapitalGenerateSlice(str string) (words []string) {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return words
	}

	var letters []string

	for index, letter := range str {
		if unicode.IsUpper(letter) && len(letters) > 0 {
			words = append(words, strings.ToLower(strings.Join(letters, "")))
			letters = make([]string, 0)
		}

		if unicode.IsSpace(letter) {
			continue
		}

		letters = append(letters, string(letter))

		if index == len(str)-1 {
			words = append(words, strings.ToLower(strings.Join(letters, "")))
		}
	}
	return words
}
