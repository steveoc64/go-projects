package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// The fetch command runs the git clone command and prints the output to terminal.
func fetch(command string, finished chan bool) {
	fetch := exec.Command("sh", "-c", fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", command))
	output, err := fetch.CombinedOutput()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", output)
	}

	finished <- true
}

func main() {
	// Grab an array of all our arguments.
	flag.Parse()
	repo := flag.Args()

	// Don't proceed if any any of the repos are already fetched.
	for i := range repo {
		if exists(repo[i]) {
			log.Fatalln("Don't fetch repos that are already fetched!")
		}
	}

	// Handle empty arguments and tell user how to use program.
	if repo[0] == "" {
		log.Fatalln("Usage: solfetch [repository name] [optional] [optional]")
	}

	// Spin up all of our communication channels.
	var channel []chan bool
	for range repo {
		channel = append(channel, make(chan bool))
	}

	// Start up cocurrent tasks for fetching all repos at once.
	for i := range repo {
		go fetch(repo[i], channel[i])
	}

	// Loop through and grab the output of each boolean channel without storing it.
	for i := range repo {
		<-channel[i]
	}
}
