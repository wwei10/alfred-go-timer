package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	. "wwei10.com/go-timer/alfred"
)

func listTimers() {
	args := os.Args[1:]
	t := time.Now()
	for _, arg := range args {
		t = addDays(t, arg)
	}
	unixtime := t.Unix()
	// Note: Format Jan 2 15:04:05 2006 MST
	sql_time := t.Format("2006-01-02")
	regular := t.Format("Mon Jan 2 15:04 MST 2006")
	short := t.Format("1/2/2006")

	response := MakeResponse()
	response.Items = append(
		response.Items,
		Item{Title: fmt.Sprintf("%d", unixtime), Valid: true, Subtitle: "unixtime", Arg: strconv.FormatInt(unixtime, 10)},
	)
	response.Items = append(
		response.Items,
		Item{Title: sql_time, Valid: true, Subtitle: "sql", Arg: sql_time},
	)
	response.Items = append(
		response.Items,
		Item{Title: regular, Valid: true, Subtitle: "system", Arg: regular},
	)
	response.Items = append(
		response.Items,
		Item{Title: short, Valid: true, Subtitle: "short", Arg: short},
	)
	ret := ToJson(response)
	fmt.Println(ret)
}

// add/substract days from current date.
func addDays(t time.Time, query string) time.Time {
	r := regexp.MustCompile(`([+|-]\d+)d`)
	match := r.FindStringSubmatch(query)
	if match == nil {
		return t
	}
	durationInDays, _ := strconv.ParseInt(match[1], 10, 64)
	duration := time.Duration(durationInDays*24) * time.Hour
	return t.Add(duration)
}

func main() {
	listTimers()
}
