package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	aw "github.com/deanishe/awgo"
	. "wwei10.com/go-timer/utils"
)

var (
	// Use timers directory to store all active timers
	timersDirectory = "/tmp/timers"

	// Command-line flags
	query string

	// Workflow
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	wf.Args()
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)

	// Fetch all active timers
	files, err := ioutil.ReadDir(timersDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Printf(file.Name())
		timer := *NewTimerFromString(file.Name())
		wf.NewItem(fmt.Sprintf("Timer %s: %d minutes left", timer.Name, timer.EndTime))
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
