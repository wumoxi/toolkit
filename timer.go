package toolkit

import "time"

// 记录运行时间
var run time.Time

// 开始记录运行时间
func FunctionStart() {
	run = time.Now()
}

// 获取记录的运行时间
func FunctionEnd() (duration time.Duration) {
	return time.Since(run)
}
