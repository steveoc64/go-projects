package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// InitGui starts up the whole interface for out program.
func InitGui() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Create the window for our user interface.
	window := app.NewWindow("CNG Medley PDF Parser")

	// Make a vertical box to house all our widgets in.
	box := widget.NewVBox()

	// Fyne doesn't support file chooser yet so we need to enter the filename manually. Should not be a problem for files in the same folder though.
	inputedFile := widget.NewEntry()

	// Add inputedFile to the box.
	box.Append(inputedFile)

	// Create out number input entry.
	inputedNumber := widget.NewEntry()

	// Create the label where we should be printing our information to.
	dataLabel := widget.NewLabel("")
	dataLabel.Resize(fyne.NewSize(400, 600))

	// Just a little spacer as a quick workaround getting a gap.
	spacer := layout.NewSpacer()

	// Create the import button for our file.
	importPDF := widget.NewButton("Importera data från pdf ovan", func() {
		// Load the imported data from the inputed pdf.
		visitors := Importer(inputedFile.Text)

		// Update the label to show that we have inported stuff.
		dataLabel.SetText("Antal elever under den veckan: " + strconv.Itoa(visitors))
	})

	// Add three more widgets to the box.
	box.Append(importPDF)
	box.Append(spacer)
	box.Append(inputedNumber)

	// Create the button for showing visitors in the gui.
	displayLessThan := widget.NewButton("Visa antal elever med färre besök än valt nummer ovan", func() {
		// Run the name printer for strings and add the data to the label.
		lessthan := CheckNumber(inputedNumber.Text)

		if lessthan < 2 || lessthan > 10 {
			box.Append(widget.NewLabel("Vänligen välj ett nummer mellan (eller lika med) 2 och 10."))
			window.Resize(fyne.NewSize(400, 200))
		} else {
			// Resize the window to something that's useable.
			window.Resize(fyne.NewSize(400, 550))

			// Loop through names and add a new label for each name.
			data := ReadDataFromXML()
			for i := 0; i < len(data.Person); i++ {
				if data.Person[i].Visits < lessthan {
					box.Append(widget.NewLabel(fmt.Sprintf("%s: %v besök", data.Person[i].Name, data.Person[i].Visits)))
				}
			}

			// Resize the window a little bit to avoid overlapping data.
			window.Resize(fyne.NewSize(400, 600))
		}
	})

	// Add the last button to the box.
	box.Append(displayLessThan)

	// Make the whole box scrollable.
	scroll := widget.NewScrollContainer(box)

	// Set our content in the application.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), scroll))

	// Set a default size for the window.
	window.Resize(fyne.NewSize(400, 170))

	// Show all our widgets and initialize out main gui loop.
	window.ShowAndRun()
}
