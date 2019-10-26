package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Global variable for defining if a button is pressed or not.
var clicked = [9]bool{}

// InitGUI starts up the whole interface for out program.
func InitGUI(channel chan bool) {
	// Initialize our new fyne interface application.
	app := app.New()

	// Set the application icon for our program.
	app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Luffarschack")

	// Make a new vertical box where we can stack our horizontal grids.
	vbox := widget.NewVBox()

	// Define all variables we need for the buttons.
	var (
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
		})
	)

	// Append a start button to our vertical box.
	vbox.Append(widget.NewButton("Click to start", func() {
		// Make sure to clear all earlier button presses before proceeding.
		for i := range clicked {
			clicked[i] = false
		}

		// Also remove all our icons before proceeding.
		button1.SetIcon(nil)
		button2.SetIcon(nil)
		button3.SetIcon(nil)
		button4.SetIcon(nil)
		button5.SetIcon(nil)
		button6.SetIcon(nil)
		button7.SetIcon(nil)
		button8.SetIcon(nil)
		button9.SetIcon(nil)

		for index := 0; index < 9; index++ {

			// Sleep in order to wait for inputs from the user.
			time.Sleep(time.Duration(5 * time.Second))

			// Ugly switch statement to set the icons for our buttons.
			switch {
			case clicked[0]:
				if index%2 == 0 {
					button1.SetIcon(circle)
				} else {
					button1.SetIcon(cross)
				}
			case clicked[1]:
				if index%2 == 0 {
					button2.SetIcon(circle)
				} else {
					button2.SetIcon(cross)
				}
			case clicked[2]:
				if index%2 == 0 {
					button3.SetIcon(circle)
				} else {
					button3.SetIcon(cross)
				}
			case clicked[3]:
				if index%2 == 0 {
					button4.SetIcon(circle)
				} else {
					button4.SetIcon(cross)
				}
			case clicked[4]:
				if index%2 == 0 {
					button5.SetIcon(circle)
				} else {
					button5.SetIcon(cross)
				}
			case clicked[5]:
				if index%2 == 0 {
					button6.SetIcon(circle)
				} else {
					button6.SetIcon(cross)
				}
			case clicked[6]:
				if index%2 == 0 {
					button7.SetIcon(circle)
				} else {
					button7.SetIcon(cross)
				}
			case clicked[7]:
				if index%2 == 0 {
					button8.SetIcon(circle)
				} else {
					button8.SetIcon(cross)
				}
			case clicked[8]:
				if index%2 == 0 {
					button9.SetIcon(circle)
				} else {
					button9.SetIcon(cross)
				}
			}

		}
	}))

	// Append each new row as a new container with gird layout and three buttons.
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button1, button2, button3))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button4, button5, button6))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button7, button8, button9))

	// Add our vertical box to be viewed.
	window.SetContent(vbox)

	// Set the size to something small but good looking and usable.
	window.Resize(fyne.NewSize(350, 100))

	// Show all of our set content and run the gui.
	window.ShowAndRun()

	channel <- true
}

func main() {
	channel := make(chan bool)

	go InitGUI(channel)

	<-channel
}
