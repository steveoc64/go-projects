package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	flag.Parse()

	var repo string = flag.Arg(0)
	if repo == "" {
		log.Fatalln("Usage: solfetch [repository name]")
	}

	path := "https://dev.getsol.us/source/" + repo + ".git"
	command := "git clone " + path

	cmd := exec.Command("/bin/bash", "-c", command)
	if err := cmd.Run(); err != nil {
		log.Fatal(err, "\nSpecified repository doesn't exist or your internet connection is down")
	}
}
