package toolkit

import (
	"testing"
)

func TestConvertToSmallCamelCase(t *testing.T) {
	var ans = "helloWorldZhongGuo"
	if actual := ConvertToSmallCamelCase("hello world zhong guo"); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}

func TestConvertToBigCamelCase(t *testing.T) {
	var ans = "HelloWorldZhongGuo"
	if actual := ConvertToBigCamelCase("hello world zhong guo"); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}