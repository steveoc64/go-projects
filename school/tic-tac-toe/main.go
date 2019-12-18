package main

import (
	"fyne.io/fyne/app"
)

// Define all our global variables. We use these in order to not redeclare them every time we start the game.
var ()

func main() {
	// Initialize our new fyne interface application.
	a := app.New()

	// Set the application icon for our program.
	a.SetIcon(icon)

	game := NewGame(a)
	game.ShowAndRun()
}
