package toolkit

import "regexp"

// GetWordsFormString 获取 str 中的单词
func GetWordsFormString(str string) (words []string) {
	compile := regexp.MustCompile(`\w+`)
	for _, word := range compile.FindAll([]byte(str), -1) {
		words = append(words, string(word))
	}
	return words
}
