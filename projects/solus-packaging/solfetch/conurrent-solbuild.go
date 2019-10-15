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
func fetch(command string, result chan []byte) {
	fetch := exec.Command("sh", "-c", command)
	output, err := fetch.CombinedOutput()
	if err != nil && command != "" {
		panic(err)
	}

	result <- output
}

func main() {
	// Parse three first arguments in to repo slice.
	flag.Parse()
	repo := [3]string{flag.Arg(0), flag.Arg(1), flag.Arg(2)}

	// Don't proceed if they alread are fetched.
	if exists(repo[0]) || exists(repo[1]) || exists(repo[2]) {
		log.Fatalln("Don't fetch repos that are already fetched!")
	}

  // Handle empty arguments and tell user how to use program.
  if repo[0] == "" {
    log.Fatalln("Usage: solfetch [repository name] [optional] [optional]")
  }

	// Start up the three channels for communication.
	var first, second, third chan []byte

  if repo[0] != "" {
    first = make(chan []byte)
  }
  if repo[1] != "" {
    second = make(chan []byte)
  }
  if repo[2] != "" {
    third = make(chan []byte)
  }

	// Start up cocurrent taskt for fetching up to three repos at once.
  if repo[0] != "" {
		go fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[0]), first)
  }
  if repo[1] != "" {
		go fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[1]), second)
  }
  if repo[2] != "" {
		go fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[2]), third)
  }

	// Fetch the outputs from each fetch run.
  var one, two, three []byte

  if repo[0] != "" && repo[1] == "" && repo[2] == "" {
    one = <-first
  } else if repo[0] != "" && repo[1] != "" && repo[2] == "" {
    one, two = <-first, <-second
  } else if repo[0] != "" && repo[1] != "" && repo[2] != "" {
    one, two, three = <-first, <-second, <-third
  }

	// Print output to terminal:
	fmt.Printf("%s%s%s", one, two, three)
}
