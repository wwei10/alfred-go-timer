package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	. "wwei10.com/go-timer/utils"
)

func createTimers() {
	var query string
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[create_timers] query=%s", query)

	// Parse query
	timer := *NewTimerFromQuery(query)
	fileName := DIRECTORY + "/" + GetFileNameFromTimer(timer)
	log.Println("fileName", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("create success!")
		file.Close()
	}

	// Sleep
	log.Println("End time: ", timer.EndTime, "now: ", time.Now().Unix())
	duration := time.Unix(timer.EndTime, 0).Sub(time.Now())
	log.Println("Preparing to sleep", duration.Minutes())
	time.Sleep(duration)
	log.Println("Finished sleeping", duration.Minutes())

	// Check whether timer is removed before sending notification
	_, err = os.Stat(fileName)

	// Do nothing if file no longer exists
	if !os.IsNotExist(err) {
		err = os.Remove(fileName)
		if err != nil {
			log.Println("Remove file failed")
		}
		fmt.Println("Timer finished")
	}
}

func main() {
	createTimers()
}
