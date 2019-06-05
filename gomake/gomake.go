package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// compileRelease64 compiles a binary without debug symbols for a smaller size.
func compileRelease64() {
	fmt.Println("Compiling binaries without debug symbols...")
	cmd := exec.Command("/bin/bash", "-c", `go build -ldflags="-s -w"`)
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// crossCompileRelease64 compiles the binary for linux, macOS and windows and all without debug symbols.
func crossCompile64() {
	fmt.Println("Compiling 64-bit binaries (no debug) for Linux, MacOS and Windows...")

	// cmdOne has the command to build for linux, cmdOne.Env has the environment variables and then we run with cmdOne.Run.
	cmdOne := exec.Command("/bin/bash", "-c", `go build -ldflags="-s -w" -o $(pwd)/main-linux-x64`)
	cmdOne.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if err := cmdOne.Run(); err != nil {
		log.Fatal(err)
	}

	// cmdTwo has the command to build for macOS, cmdTwo.Env has the environment variables and then we run with cmdTwo.Run.
	cmdTwo := exec.Command("/bin/bash", "-c", `go build -ldflags="-s -w" -o $(pwd)/main-darwin-x64`)
	cmdTwo.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
	if err := cmdTwo.Run(); err != nil {
		log.Fatal(err)
	}

	// cmdThree has the command to build for windows, cmdThree.Env has the environment variables and then we run with cmdThree.Run.
	cmdThree := exec.Command("/bin/bash", "-c", `go build -ldflags="-s -w" -o $(pwd)/main-windows-x64`)
	cmdThree.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	if err := cmdThree.Run(); err != nil {
		log.Fatal(err)
	}
}

// Just a standard compilation using go build.
func compile() {
	fmt.Println("Running standard compile...")
	cmd := exec.Command("go", "build")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Check that user uses Linux to run the software.
	if runtime.GOOS != "linux" {
		log.Fatalf("%s is not a supported operatingsystem, use Linux instead!\n", runtime.GOOS)
	}

	// Set up command line flags with documentation and then parse them.
	var Release *bool = flag.Bool("release", false, "strip debug data for a smaller binary file.")
	var CrossCompile *bool = flag.Bool("cross-compile", false, "compiles the package for Amd64 on Linux, MacOS and Windows without debug symbols.")
	flag.Parse()

	// Handle values for the command line flags passed by users.
	if *Release {
		compileRelease64()
	} else if *CrossCompile {
		crossCompile64()
	} else {
		compile()
	}
}
