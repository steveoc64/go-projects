package main

import "encoding/xml"

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
