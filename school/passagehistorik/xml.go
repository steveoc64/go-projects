package main

import "encoding/xml"

type Data struct {
	XMLName   xml.Name   `xml:"data"`
	Person    []Person   `xml:"person"`
}

type Person struct {
	XMLName    xml.Name   `xml:"person"`
	Name       string     `xml:"name"`
	Visits     int        `xml:"visits"`
}
