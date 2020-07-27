package utils

import (
	"math"
	"regexp"
	"strconv"
	"time"
)

type Timer struct {
	Name    string
	EndTime int64
}

const DIRECTORY = "timers"
const EXPIRED_DIRECTORY = "expired"

func NewTimer(name string, endTime int64) *Timer {
	t := Timer{Name: name, EndTime: endTime}
	return &t
}

func GetRemainingMinutes(timer Timer) int64 {
	remaining := time.Unix(timer.EndTime, 0).Sub(time.Now()).Minutes()
	return int64(math.Abs(remaining))
}

func GetFileNameFromTimer(timer Timer) string {
	return "name:" + timer.Name + " time:" + strconv.FormatInt(timer.EndTime, 10)
}

func NewTimerFromFileName(s string) *Timer {
	r := regexp.MustCompile(`name:(?P<name>.*) time:(?P<time>\d+)`)
	match := r.FindStringSubmatch(s)
	if len(match) < 3 {
		return nil
	}
	endTime, _ := strconv.ParseInt(match[2], 10, 64)
	return NewTimer(match[1], endTime)
}

func NewTimerFromQuery(s string) *Timer {
	r := regexp.MustCompile(`(?P<time>\d+) (?P<name>.*)`)
	match := r.FindStringSubmatch(s)
	if len(match) < 3 {
		return nil
	}
	duration, _ := strconv.ParseInt(match[1], 10, 64)
	duration *= 60
	return NewTimer(match[2], duration+time.Now().Unix())
}
