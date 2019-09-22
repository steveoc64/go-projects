package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

// Data has the xml data for the initial data tag.
type Data struct {
	XMLName xml.Name `xml:"data"`
	Person  []Person `xml:"person"`
}

// Person keeps track of the data per person inside the data tag.
type Person struct {
	Name   string `xml:"name"`
	Visits int    `xml:"visits"`
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
