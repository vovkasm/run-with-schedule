package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/robfig/cron"
)

var (
	timeSpec string
)

func main() {
	flag.StringVar(&timeSpec, "spec", "", "Timespec in cron format")

	flag.Parse()
	job := flag.Args()
	if len(job) == 0 {
		log.Println("Job is empty")
		os.Exit(1)
	}

	c := cron.New()
	err := c.AddFunc(timeSpec, func() {
		log.Printf("JOB %#v\n", job)
		cmd := exec.Command(job[0], job[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
	})
	if err != nil {
		log.Printf("Incorrect spec: '%s'\n", timeSpec)
		os.Exit(1)
	}

	c.Run()
}
