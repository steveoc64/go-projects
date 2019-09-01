package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

// Create creates a file to download data to. It will save to specified file path if output is true, otherwice it will save to pwd.
func create(outpath string, output bool) (out File, err error) {
	// Save to current working directory if output is false.
	if !output {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		outpath = pwd
	}

	// Create file to download data to.
	out, err = os.Create(outpath)
	if err != nil {
		return err
	}

	// Defer closing of file to save memory.
	defer out.Close()

	return out, nil
}

// Download downloads data from an online source in to created file from create function.
func download(url string, outpath string, output *bool) (err error) {
	// Run function to create the file we should save to.
	out, err := create(outpath, *output)
	if err != nil {
		return err
	}

	// Download actual data.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Close the download body to save memory.
	defer resp.Body.Close()

	// Write downloaded body to created file on file system.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Set up the command line flag.
	var output *bool = flag.Bool("o", false, "Tell the program where to download the file to.")

	// Parse all the flags and arguments.
	flag.Parse()

	// Set up the string avriables for url to fetch and output path.
	var url string = flag.Arg(0)
	var path string = flag.Arg(2)

	// Handle exceptions.
	if url == "" {
		log.Fatalln("Usage: get [url to file]\nOr:    get [url to file] -o [path to save to]")
	} else if *output == true && path == "" {
		log.Fatalln("In order to specify output path, you need to add a path after -o:\nUsage: get [url to file] -o [path to save to]")
	}

	// Download specified url and stop if we get an error.
	err := download(url, path, output)
	if err != nil {
		log.Fatal(err)
	}
}
