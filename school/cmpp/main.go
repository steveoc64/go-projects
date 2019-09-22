package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

var filename string

// Filename makes sure that we get the month and year to create a file.
func Filename() {
	// Determine year and month from local time server.
	year := time.Now().Year()
	month := time.Now().Month()

	// Handle if it's the spring term or the autumn term.
	var term string
	switch month {
	case time.January, time.February, time.March, time.April, time.May, time.June:
		term = "VT"
	case time.August, time.September, time.October, time.November, time.December:
		term = "HT"
	default:
		log.Fatalln("You really shouldn't work in July. Please take some time off! :)")
	}

	// Put together the whole filename from teh data we defined earlier.
	filename = term + "-" + strconv.Itoa(year) + ".xml"
}

func main() {
	// Parse input from terminal.
	flag.Parse()
	command := flag.Arg(0)
	fileToParse := flag.Arg(1)

	// Check the filename for month and year.
	Filename()

	// Handle different commands for the program.
	if command == "import" {
		Importer(fileToParse)
	} else if command == "less" {
		PrintLessThan(CheckNumber(fileToParse))
	} else {
		fmt.Println("Usage:\n	Importing a PDF:\n		cmpp import [file.pdf]\n\n	Show users with < x visits:\n		cmpp less [1 < value < 11]")
		return
	}

}
