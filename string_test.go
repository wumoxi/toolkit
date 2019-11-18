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

func TestUnderline2SmallCamelCase(t *testing.T) {
	var ans = "helloWorldZhongGuo"
	if actual := Underline2SmallCamelCase("hello world zhong guo"); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}

func TestUnderline2BigCamelCase(t *testing.T) {
	var ans = "HelloWorldZhongGuo"
	if actual := Underline2BigCamelCase("hello world zhong guo"); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}