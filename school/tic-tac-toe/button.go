package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// Button object for each tile
type Button struct {
	g   *Game
	btn *widget.Button
	id  uint8
}

// NewButton returns a new Button
func NewButton(g *Game, id uint8) *Button {
	return &Button{
		id: id,
		btn: widget.NewButton("", func() {
			g.Press(id)
		}),
	}
}

type Buttons []*Button

// NewButtons returns a fresh new slice of buttons
func NewButtons(g *Game) Buttons {
	data := make(Buttons, 0, 9)
	var i uint8
	for i = 0; i < 9; i++ {
		data = append(data, NewButton(g, i))
	}
	return data
}

// SetIcon sets the tile
func (b *Button) SetIcon(i fyne.Resource) {
	b.btn.SetIcon(i)
}

// Clear the tile
func (b Buttons) Clear() {
	for _, v := range b {
		v.SetIcon(nil)
	}
}

// Objects gets the buttons
func (b Buttons) Objects() []fyne.CanvasObject {
	data := make([]fyne.CanvasObject, 0, 9)
	for _, v := range b {
		data = append(data, v.btn)
	}
	return data
}
