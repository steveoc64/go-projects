package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var pwd, url, path string

func init() {
  // Get the present working directory.
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Define our variable for the url.
	path = *flag.String("o", pwd, "Specify output path and filename to save file.")

	// Parse all the flags and arguments.
	flag.Parse()

	// The very first argument will be our url string.
	url = flag.Arg(0)
  if url == "" {
    panic("No url was specified.")
  }
}

func main() {
	// Start downloading of our specified url.
  download(url, path)
}

// Download fetches data from an online source in to a created file.
func download(url string, outpath string) {
	// Run function to create the file we should save to.
	out := create(outpath, outpath)

	// Make sure that we close the file when we are done. Should save us some memory.
	defer out.Close()

	// Download actual data from the website.
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// Close the download body to save some memory.
	defer resp.Body.Close()

	// Write downloaded body to the file we created on the file system.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

// Create creates a file to download data to. It will save to specified file path if output is true, otherwice it will save to pwd.
func create(outpath string, path string) (out *os.File) {

	// Add the filename in the url if we don't specify a path, or if the path doesn't contain a filename.
	if path == pwd || match("*/", path){
		// Join the present working directory with the filename at the end of the url.
		outpath = filepath.Join(path, filepath.Base(url))
	}

	// Create the file where we should save our downloaded data.
	out, err := os.Create(outpath)
	if err != nil {
		panic(err)
	}

	return out
}

// Match is filepath.Match() without the error.
func match(pattern, path string) bool {
   truth, _ := filepath.Match(pattern, path)
   return truth
}
