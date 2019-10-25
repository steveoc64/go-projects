package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// InitGUI starts up the whole interface for out program.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Create the window for our user interface.
	window := app.NewWindow("Luffarschack")

	// Make a new vertical box where we can stack our horizontal grids.
	vbox := widget.NewVBox()

	// Define all variables we need for the buttons.
	var (
		clicked = [9]bool{}

		button1 = widget.NewButton("", func() {
			clicked[0] = true
		})
		button2 = widget.NewButton("", func() {
			clicked[1] = true
		})
		button3 = widget.NewButton("", func() {
			clicked[2] = true
		})
		button4 = widget.NewButton("", func() {
			clicked[3] = true
		})
		button5 = widget.NewButton("", func() {
			clicked[4] = true
		})
		button6 = widget.NewButton("", func() {
			clicked[5] = true
		})
		button7 = widget.NewButton("", func() {
			clicked[6] = true
		})
		button8 = widget.NewButton("", func() {
			clicked[7] = true
		})
		button9 = widget.NewButton("", func() {
			clicked[8] = true
			fmt.Println(true)
		})
	)

	// Append each new row as a new container with gird layout and three buttons.
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button1, button2, button3))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button4, button5, button6))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button7, button8, button9))

	// Add our vertical box to be viewed.
	window.SetContent(vbox)

	// Set the size to something small but good looking and usable.
	window.Resize(fyne.NewSize(200, 100))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}

func main() {
	InitGUI()
}
