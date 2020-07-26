package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	. "wwei10.com/go-timer/alfred"
)

func listTimers() {
	flag.Parse()
	var query string
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)
	now := time.Now()
	unixtime := now.Unix()
	// Note: Format Jan 2 15:04:05 2006 MST
	sql_time := now.Format("2006-01-02")
	regular := now.Format("Mon Jan 2 15:04 MST 2006")
	short := now.Format("1/2/2006")

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

func main() {
	listTimers()
}
