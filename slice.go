package toolkit

import "strings"

// 生成指定范围内的有序整型切片
func GenerateSectionIntSliceOfOrderly(min, max int, step int) []int {
	result := make([]int, 0, max)
	for i := min; i <= max; i += step {
		result = append(result, i)
	}
	return result
}

// 生成指定范围内的无序整型切片
func GenerateSectionIntSliceOfDisorderly(min, max int) []int {
	l := max - min
	r := make([]int, 0)
	for i := 0; i < l; i++ {
		r = append(r, Random(min, max))
	}
	return r
}

// 使用指定的定界符连接字符串切片值
func JoinItemOfStringSlice(delimit string, slice ...string) (joined string) {
	return strings.Join(slice, delimit)
}

// 生成斐波那契数
func Fib(n int) (res int) {
	if n <= 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

// 生成指定范围内的斐波那契数列
func FibSliceOfInt(max int) (slice []int) {
	for i := 0; ; i++ {
		var r int
		if r = Fib(i); r >= max {
			return slice
		}
		slice = append(slice, r)
	}
}