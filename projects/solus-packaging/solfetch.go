package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// The fetch command runs the git clone command and prints the output to terminal.
func fetch(command string) {
	compile := exec.Command("sh", "-c", command)
	output, err := compile.CombinedOutput()
	if err != nil {
		log.Fatalf("%s", output)
	} else {
		fmt.Printf("%s", output)
	}
}

func main() {
	// Let flag parse the command line arguments and then input the first argument in to repo variable.
	flag.Parse()
	var repo string = flag.Arg(0)

	// Handle empty arguments by printing usage
	if repo == "" {
		log.Fatalln("Usage: solfetch [repository name]")
	}

	// Stitch together the full path to the repository.
	path := "git clone https://dev.getsol.us/source/" + repo + ".git"

	// Run the git command to download the whole repository.
	fetch(path)
}
