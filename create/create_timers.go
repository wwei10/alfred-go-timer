package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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
	expiredTimerFileName := EXPIRED_DIRECTORY + "/" + GetFileNameFromTimer(timer)
	log.Println("fileName", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("create success!")
		file.Close()
	}

	// Integrate with BlockSite app
	cmd := exec.Command(
		"open",
		"-a",
		"Safari",
		"blocksite://add?"+strconv.FormatInt(int64(timer.Duration/60), 10))
	cmd.Run()

	// Sleep
	duration := time.Unix(timer.EndTime, 0).Sub(time.Now())
	log.Println("End time: ", timer.EndTime, "now: ", time.Now().Unix())
	log.Println("Preparing to sleep", duration.Minutes())
	time.Sleep(duration)
	log.Println("Finished sleeping", duration.Minutes())

	// Check whether timer is removed before sending notification
	_, err = os.Stat(fileName)

	// Do nothing if file no longer exists
	if !os.IsNotExist(err) {
		err = os.Rename(fileName, expiredTimerFileName)
		if err != nil {
			log.Fatal("Failure in renaming timer file")
		}
		fmt.Printf("Timer %s finished", timer.Name)
	}
}

func main() {
	createTimers()
}
