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
	_, week := t.ISOWeek()
	start, end := getWeekRange(t)
	start_date_string := start.Format("Jan 2")
	end_date_string := end.Format("Jan 2")
	week_range_string := fmt.Sprintf("%s - %s", start_date_string, end_date_string)

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
		Item{Title: short, Valid: true, Subtitle: "short", Arg: short},
	)
	response.Items = append(
		response.Items,
		Item{
			Title: fmt.Sprintf("Week %d, %s", week, week_range_string),
			Valid: true, Subtitle: "week",
			Arg: week_range_string,
		},
	)
	response.Items = append(
		response.Items,
		Item{Title: regular, Valid: true, Subtitle: "system", Arg: short},
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

// Calculate [beginning of week, end of week]
// Start from Monday
func getWeekRange(t time.Time) (time.Time, time.Time) {
	weekday := t.Weekday()
	var start time.Time
	if weekday == time.Sunday {
		start = t.AddDate(0, 0, -6)
	} else {
		start = t.AddDate(0, 0, -int(weekday)+1)
	}
	return start, start.AddDate(0, 0, 6)
}

func main() {
	listTimers()
}
