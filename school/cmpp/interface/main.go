package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// Create and initialize out window.
func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	// Set title and connect it to destroy event.
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Set the default size.
	win.SetDefaultSize(800, 600)
	return win
}

func setupButton(label string, onClick func()) *gtk.Button {
	// Create the button.
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Connect the button to the on click function.
	btn.Connect("clicked", onClick)

	return btn
}

func main() {
	// Initialize gtk without arguments.
	gtk.Init(nil)

	// Run setup of window.
	win := setupWindow("Cng Medley PDF Parser")

	fixed, err := gtk.FixedNew()
	if err != nil {
		log.Fatal("Unable to create GtkFixed:", err)
	}

	btn := setupButton("Import", func() {
		fmt.Println("Thanks for importing a pdf!")
	})

	fixed.Put(btn, 100, 200)

	win.Add(fixed)

	// Show everything and start the main interface loop.
	win.ShowAll()
	gtk.Main()
}
