package toolkit

import (
	"testing"
)

func TestGetFuncName(t *testing.T) {
	ans := "github.com/wumoxi/toolkit.CamelCase2Underline"

	actual, err := GetFuncName(CamelCase2Underline)
	if err != nil && actual == "" {
		t.Error(err)
	}

	if actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}
