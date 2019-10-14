package main

import (
	"fmt"
	"log"
	"strconv"
)

// PrintLessThan print all names with less than x visits.
func PrintLessThan(lessthan int) {
	data := ReadDataFromXML()
	for i := 0; i < len(data.Person); i++ {
		if data.Person[i].Visits < lessthan {
			fmt.Println(data.Person[i].Name, data.Person[i].Visits)
		}
	}
}

// StringLessThan does almost the same thing as PrintLessThan but makes one big string out of it instead.
func StringLessThan(lessthan int) string {
	var text string
	data := ReadDataFromXML()
	for i := 0; i < len(data.Person); i++ {
		if data.Person[i].Visits < lessthan {
			text += fmt.Sprintf("%s: %v besök\n", data.Person[i].Name, data.Person[i].Visits)
		}
	}

	return text
}

// CheckNumber checks for inputed number from command less.
func CheckNumber(command string) int {

	// Convert string to int.
	lessthan, err := strconv.Atoi(command)
	if err != nil || lessthan < 2 || lessthan > 10 {
		log.Fatalln("Enter a value from two up to ten.")
	}

	return lessthan
}
