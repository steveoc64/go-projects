package main

import (
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

func setupEntry() *gtk.Entry {
	// Create the entry.
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	return entry
}

func initGui() {
	// Initialize gtk without arguments.
	gtk.Init(nil)

	// Run setup of window.
	win := setupWindow("Cng Medley PDF Parser")

	win.Add(windowWidgets())

	// Show everything and start the main interface loop.
	win.ShowAll()
	gtk.Main()
}

func windowWidgets() *gtk.Widget {

	// Make a new scrolled window for our grid.
	scrolled, _ := gtk.ScrolledWindowNew(nil, nil)

	// Create a grid to house all out little widgets.
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	// Set the orientation of our spiffy new grid.
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Create a new button with the label show.
	show, err := gtk.ButtonNewWithLabel("show")
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Make a new entry for inputing text.
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	// Create a new label for displaying text.
	label, err := gtk.LabelNew("Names with less than x visits:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Connect the show button to printing out names to teh label.
	show.Connect("clicked", func() {
		text, _ := entry.GetText()
		names := StringLessThan(CheckNumber(text))
		label.SetText(names)
	})

	// Add the entry, button and labels to out grid.
	grid.Add(entry)
	grid.Add(show)
	grid.Add(label)

	// Tell the entry that it can't expand and that the label can.
	entry.SetHExpand(false)
	label.SetHExpand(true)

	// Attatch the grid next to the entry.
	grid.AttachNextTo(show, entry, gtk.POS_RIGHT, 1, 1)

	// Add out grid to the scrolled window.
	scrolled.Add(grid)

	return &scrolled.Container.Widget
}
