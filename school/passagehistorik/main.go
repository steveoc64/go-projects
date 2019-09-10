package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

// ParseNames takes a string of data from the pdf document medley sends out and parses the names from it.
func ParseNames(content string) (names []string) {
	// Abommination of cleanups for the content string.
	content = strings.ReplaceAll(content, "Curt Nicolingymnasiet AB (elever)", "")
	content = strings.ReplaceAll(content, "Curt Nicolingymnasiet AB", "")
	content = strings.ReplaceAll(content, "Passagehistorik per person220Antal,", "")
	content = strings.ReplaceAll(content, "Person:", "\n")
	content = strings.ReplaceAll(content, " totalt:", "")
	content = strings.ReplaceAll(content, "Dumtumintervall:", "")
	content = strings.ReplaceAll(content, "TidKortnummerVärdekortResultatLäsareMeddelandeNytt besök", "\n")

	// Split the string in to an array for every new line.
	textArr := strings.Split(content, "\n")

	// Loop through the text array and append every line with uneaven index to names array.
	for i := 0; i < len(textArr); i++ {
		if i%2 != 0 {
			names = append(names, textArr[i])
		}
	}

	// Unalocate mamory from first array
	textArr = nil

	// Array housing a character list for characters to strip out from all the names.
	var replaces = [12]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", " - "}

	// Interate through every name and remove left over chracters.
	for i := 0; i < len(names); i++ {
		for f := 0; f < len(replaces); f++ {
			names[i] = strings.ReplaceAll(names[i], replaces[f], "")
		}
		fmt.Println(names[i])
	}

	return
}

func main() {
	flag.Parse()
	fileToParse := flag.Arg(0)

	// Read the pdf file using ReadPDF function.
	content, err := ReadPDF(fileToParse)
	if err != nil {
		panic(err)
	}

	// Run the PDF data through the name parser.
	names := ParseNames(content)

	// Interate through every name and print it to terminal.
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

// ReadPDF reads the content of the whole pdf file and prints it as text.
func ReadPDF(path string) (string, error) {

	// Open the pdf file using the pdf library.
	file, result, err := pdf.Open(path)
	if err != nil {
		return "", err
	}

	// remember to close file.
	defer file.Close()

	// Create new buffer.
	var buf bytes.Buffer

	// Make sure to get plain text from pdf.
	buffer, err := result.GetPlainText()
	if err != nil {
		return "", err
	}

	// Read the buffer that we took plaintext from.
	buf.ReadFrom(buffer)

	return buf.String(), nil
}
