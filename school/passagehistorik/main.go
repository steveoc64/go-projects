package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
)

// ParseNames takes a string of data from the pdf document medley sends out and parses the names from it.
func ParseNames(content string) (names []string) {

	// Arrays with the void and new line that we peplace with.
	var removal = [2]string{"\n", ""}
	var replace1 = [7]string{"Person:", "TidKortnummerVärdekortResultatLäsareMeddelandeNytt besök", "Dumtumintervall:", " totalt:", "Passagehistorik per person220Antal,", "Curt Nicolingymnasiet AB", "Curt Nicolingymnasiet AB (elever)"}

	// Loop through the things we should remove instead of having an abbomination of removals.
	for i := 0; i < len(replace1); i++ {
		if i == 0 || i == 1 {
			content = strings.ReplaceAll(content, replace1[i], removal[0])
		} else {
			content = strings.ReplaceAll(content, replace1[i], removal[1])
		}
	}

	// Split the string in to an array for every new line.
	textArr := strings.Split(content, removal[0])

	// Loop through the text array and append every line with uneaven index to names array.
	for i := 0; i < len(textArr); i++ {
		if i%2 != 0 {
			names = append(names, textArr[i])
		}
	}

	// Array housing a character list for characters to strip out from all the names.
	var replace2 = [15]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", " - ", "Passagehistorik per personAntal,", "Passagehistorik per personAntal,    (elever)", "    (elever)"}

	// Interate through every name and remove left over chracters.
	for i := 0; i < len(names); i++ {
		for f := 0; f < len(replace2); f++ {
			names[i] = strings.ReplaceAll(names[i], replace2[f], "")
		}
	}

	return
}

// Importer is  just a handler function to clean up the code in main().
func Importer(fileToParse string) {
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

	fmt.Println("Antal elever på medley under veckan:", len(names))

	CheckForData(names)
}

// CheckForData checks that we have the data file, if not, we create it.
func CheckForData(names []string) {
	// Get current year from time server.
	year := time.Now().Year()

	// Get current month and determine VT or HT term.
	month := time.Now().Month()
	var term string
	switch string(month) {
	case "January", "February", "March", "April", "May", "June":
		term = "VT"
	case "August", "September", "October", "November", "December":
		term = "HT"
	default:
		log.Fatalln("You really shouldn't work in July. Please take some time off!")
	}

	// Make the file name to write data to.
	filename := term + "-" + string(year) + ".xml"

	// Checking if we have a file with the set name for the term.
	if _, err := os.Stat(filename); err == nil {
		// It exsists, we should call a function to update the data there.
	} else if os.IsNotExist(err) {
		// File isn't there, we should create it.
		CreateFile(filename)

		// Now call function to update the data from the names array.
	} else {
		// Test some schrodinger stuff where the file may or may not exist.
		panic(err)
	}
}

// UpdateXMLFile updates the xml file with imported data.
func UpdateXMLFile() {

}

// CreateFile uses os.Create to make a file.
func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return file
}

func main() {
	flag.Parse()
	command := flag.Arg(0)
	fileToParse := flag.Arg(1)

	if command == "import" {
		/*names := */ Importer(fileToParse)
	} else {
		log.Fatalln("Usage:\n			Importing a PDF:\n						cmpp import [file.pdf]")
	}

	// TODO:
	// - Add function to update xml data.
	// - Add option to display users with less than x visits.
	// - Make a user interface.
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
