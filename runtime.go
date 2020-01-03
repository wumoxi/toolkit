package toolkit

import (
	"fmt"
	"runtime"
)

// 字节大小
type ByteSize float64

// 字节大小常量枚举
const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB

	KBIdent = "KB"
	MBIdent = "MB"
	GBIdent = "GB"
	TBIdent = "TB"
	PBIdent = "PB"
	EBIdent = "EB"
	ZBIdent = "ZB"
	YBIdent = "YB"
)

// ByteSizeMap字节大小及单位映射表
type ByteSizeInfo string

// String 单位大小描述信息
func (s *ByteSizeInfo) String(number ByteSize, ident string) string {
	return fmt.Sprintf("%.2f%s", number, ident)
}

// MemAllocated获取当前程序占用内存大小
func MemAllocated() (info string) {
	var i ByteSizeInfo
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	allocated := ByteSize(m.Alloc / 1024)

	switch {
	case allocated >= 0 && allocated <= KB:
		info = i.String(allocated, KBIdent)
	case allocated > KB && allocated <= MB:
		info = i.String(allocated/KB, MBIdent)
	case allocated > MB && allocated <= GB:
		info = i.String(allocated/MB, GBIdent)
	default:
		info = "The current program occupies more than GB of memory"
	}

	return info
}
