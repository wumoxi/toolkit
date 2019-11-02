package toolkit

import (
	"fmt"
	"testing"
)

func TestFunctionStart(t *testing.T) {
	FunctionStart()
	sum(100000000)
	duration := FunctionEnd()
	fmt.Println(duration)
}

func sum(number int64) (sum int64) {
	var i int64
	for i = 0; i < number; i++ {
		sum += i
	}
	return sum
}
