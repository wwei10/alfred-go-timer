package utils_test

import (
	"testing"

	. "wwei10.com/go-timer/utils"
)

func TestParseFileName(t *testing.T) {
	s := "name:hello, world time:1595152620"
	timer := *NewTimerFromString(s)
	if timer.EndTime != 1595152620 || timer.Name != "hello, world" {
		t.Errorf("Parse Fail")
	}
}
