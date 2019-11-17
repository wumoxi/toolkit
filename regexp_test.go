package toolkit

import (
	"strings"
	"testing"
)

func TestGetWordsFormString(t *testing.T) {
	var ans = "hello world wu mo xi"
	words := GetWordsFormString(" hello      world  wu     mo        xi ")
	if actual := strings.Join(words, " "); actual != ans {
		t.Errorf("got %s, expected %s\n", actual, ans)
	}
}
