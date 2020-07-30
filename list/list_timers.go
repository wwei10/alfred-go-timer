package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	. "wwei10.com/go-timer/alfred"
	. "wwei10.com/go-timer/utils"
)

func listTimers() {
	listExpired := flag.Bool("expired", false, "true if you want to list expired timers")
	flag.Parse()
	var query string
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)

	// Fetch all active timers
	directory := DIRECTORY
	if *listExpired {
		directory = EXPIRED_DIRECTORY
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	response := MakeResponse()
	for _, file := range files {
		timer := *NewTimerFromFileName(file.Name())
		response.Items = append(
			response.Items,
			Item{
				Title: fmt.Sprintf(
					"Name: %s, duration: %d minutes, %d minutes left",
					timer.Name,
					GetDuration(timer),
					GetRemainingMinutes(timer)),
				Valid: true,
				Arg:   strconv.FormatInt(timer.EndTime, 10),
			})
		log.Println("timer", timer)
	}
	log.Println("before tojson", response)
	ret := ToJson(response)
	log.Println("json response", ret)
	fmt.Println(ret)
}

func main() {
	listTimers()
}
