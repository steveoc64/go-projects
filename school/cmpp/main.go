package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var filename string

// Filename makes sure that we get the month and year to create a file.
func Filename() {
	// Get current year from time server.
	year := time.Now().Year()

	// Get current month and determine VT or HT term.
	month := time.Now().Month()
	var term string
	switch month {
	case time.January, time.February, time.March, time.April, time.May, time.June:
		term = "VT"
	case time.August, time.September, time.October, time.November, time.December:
		term = "HT"
	default:
		log.Fatalln("You really shouldn't work in July. Please take some time off! :)")
	}

	// Make the file name to write data to.
	filename = term + "-" + strconv.Itoa(year) + ".xml"
}

// ReadDataFromXML reads data from an xml file, couldn't be simpler.
func ReadDataFromXML() Data {

	// Open up the xml file that already exists.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to close it also.
	defer file.Close()

	// Read the data from the opened file.
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the xml data in to our Data struct.
	data := Data{}
	xml.Unmarshal(byteValue, &data)

	return data
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
