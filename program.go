package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/deanishe/awgo"
)

type Timer struct {
	name      string
	startTime int64
	endTime   int64
}

func newTimer(name string, startTime int64, endTime int64) *Timer {
	t := Timer{name: name, startTime: startTime, endTime: endTime}
	return &t
}

var (
	// Stores timer name => timer object
	m map[string]Timer

	// Command-line flags
	query string

	// Workflow
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
	m = make(map[string]Timer)
	m["test"] = *newTimer("hello", time.Now().Unix(), time.Now().Unix()+3600)
}

func run() {
	wf.Args()
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)
	for k, v := range m {
		log.Printf("key[%s] value[%s]\n", k, v)
		wf.NewItem(fmt.Sprintf("key[%s] value[%s]\n", k, v))
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
