package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	. "wwei10.com/go-timer/alfred"
	. "wwei10.com/go-timer/utils"
)

func listTimers() {
	flag.Parse()
	var query string
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)

	// Fetch all active timers
	files, err := ioutil.ReadDir(DIRECTORY)
	if err != nil {
		log.Fatal(err)
	}

	response := MakeResponse()
	for _, file := range files {
		timer := *NewTimerFromFileName(file.Name())
		response.Items = append(
			response.Items,
			Item{
				Title: fmt.Sprintf("Timer %s %d minutes left", timer.Name, GetRemainingMinutes(timer)),
				Valid: false,
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
