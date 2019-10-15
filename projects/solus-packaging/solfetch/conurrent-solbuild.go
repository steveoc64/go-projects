package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// exists returns whether the given file or directory exists
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
func fetch(command string, result chan string) {
	fetch := exec.Command("sh", "-c", command)
	output, err := fetch.CombinedOutput()
	if err != nil && command != "" {
		panic(err)
	}

	result <- string(output)
}

func main() {
	// Parse three first arguments in to repo slice.
	flag.Parse()
	repo := flag.Args()

	// Don't proceed if they alread are fetched.
	for i := range repo {
		if exists(repo[i]) {
			log.Fatalln("Don't fetch repos that are already fetched!")
		}
	}

	// Handle empty arguments and tell user how to use program.
	if repo[0] == "" {
		log.Fatalln("Usage: solfetch [repository name] [optional] [optional]")
	}

	// Start up the three channels for communication.
	var chanel []chan string
	for range repo {
		chanel = append(chanel, make(chan string))
	}

	// Start up cocurrent taskt for fetching up to three repos at once.
	for i := range repo {
		go fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[i]), chanel[i])
	}

	// Fetch the outputs from each fetch run.
	var outputs []string
	for i := range repo {
		outputs = append(outputs, <-chanel[i])
	}

	// Print output to terminal:
	for i := range outputs {
		fmt.Print(outputs[i])
	}
}
