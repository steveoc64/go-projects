package main

import (
	//"fyne.io/fyne"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	//"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	InitGUI()
}

// InitGUI starts up the whole interface for out program.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Set the application icon for our program.
	//app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Hangman")

	// Make a new vertical box where we can plae our stuff.
	vbox := widget.NewVBox()

	start := widget.NewButton("Press to start", func() {

	})

	display := widget.NewIcon(hangman)

	vbox.Append(start)

	// Add our vertical box to be viewed.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), vbox, display))

	// Set a sane default for the window size and set it to not be user resizable.
	window.Resize(fyne.NewSize(400, 500))
	window.SetFixedSize(true)

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
