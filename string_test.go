package toolkit

import (
	"strings"
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

func TestAllowCapitalGenerateSlice(t *testing.T) {
	var ans1 = strings.Join([]string{"hello", "world", "zhong", "guo"}, "-")

	if actual := AllowCapitalGenerateSlice("helloWorldZhongGuo"); strings.Join(actual, "-") != ans1 {
		t.Errorf("got %s, expected %s\n", actual, ans1)
	}

	if actual := AllowCapitalGenerateSlice("HelloWorldZhongGuo"); strings.Join(actual, "-") != ans1 {
		t.Errorf("got %s, expected %s\n", actual, ans1)
	}
}

func TestCamelCase2Underline(t *testing.T) {
	ans := "hello_world_user_info"
	if actual := CamelCase2Underline("HelloWorldUserInfo"); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}
