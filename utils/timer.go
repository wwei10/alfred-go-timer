package utils

import (
	"regexp"
	"strconv"
)

type Timer struct {
	Name    string
	EndTime int64
}

func NewTimer(name string, endTime int64) *Timer {
	t := Timer{Name: name, EndTime: endTime}
	return &t
}

func NewTimerFromString(s string) *Timer {
	r := regexp.MustCompile(`name:(?P<name>.*) time:(?P<time>\d+)`)
	match := r.FindStringSubmatch(s)
	if len(match) < 3 {
		return nil
	}
	endTime, _ := strconv.ParseInt(match[2], 10, 64)

	return NewTimer(match[1], endTime)
}
