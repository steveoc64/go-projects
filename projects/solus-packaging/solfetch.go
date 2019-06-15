package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	// Let flag parse the command line arguments and then input the first argument in to repo variable
	flag.Parse()
	var repo string = flag.Arg(0)
	
	// Handle empty arguments by printing usage
	if repo == "" {
		log.Fatalln("Usage: solfetch [repository name]")
	}

	path := "git clone https://dev.getsol.us/source/" + repo + ".git"

	// Run the git command to download the whole repository
	cmd := exec.Command("/bin/bash", "-c", path)
	if err := cmd.Run(); err != nil {
		log.Fatal(err, "\nSpecified repository doesn't exist or your internet connection is down")
	}
}
