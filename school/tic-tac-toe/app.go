package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Game is our TicTacToe app
type Game struct {
	app fyne.App

	// the tiles and other UI elements
	buttons     Buttons
	container   *fyne.Container
	startButton *widget.Button

	// Variables for which parts the players own. Index 0 is the first tile and index 8 is the ninth tile.
	player1 []bool
	player2 []bool

	// Variable to handle if buttons are already pressed.
	pressed []bool

	// Tell if we are in game or not.
	inGame bool

	// Index for defining if it's player one or player two's turn to play. Use 8bit variable to save on memory allocation.
	index uint8

	w fyne.Window
}

// NewGame returns a new TTT Game
func NewGame(app fyne.App) *Game {
	g := &Game{
		app:     app,
		w:       app.NewWindow("Tic-Tac-Toe"),
		player1: make([]bool, 9),
		player2: make([]bool, 9),
		pressed: make([]bool, 9),
	}
	g.buttons = NewButtons(g)
	// Create our start button for the whole game.
	g.startButton = widget.NewButton("Click to ReStart", func() {
		g.Start()
	})

	g.container = fyne.NewContainerWithLayout(layout.NewGridLayout(3), layout.NewSpacer(), g.startButton, layout.NewSpacer())
	// Add all the buttons in to a three column grid layout inside a container.
	for _, v := range g.Objects() {
		g.container.AddObject(v)
	}

	// Set the conatiner as what is being displayed.
	g.w.SetContent(g.container)
	// Set a sane default for the window size.
	g.w.Resize(fyne.NewSize(400, 250))
	return g
}

// ShowAndRun launches the UI
func (g *Game) ShowAndRun() {
	// Show all of our set content and run the gui.
	g.Start()
	g.w.ShowAndRun()
}

// Objects gets the tiles
func (g *Game) Objects() []fyne.CanvasObject {
	return g.buttons.Objects()
}

// Start a new game
func (g *Game) Start() {
	// Make sure that all buttons are reset to remove icons before we start.
	g.buttons.Clear()
	g.inGame = true
	g.index = 0
}

// CheckWon returns true if the player has won
func (g *Game) CheckWon(data []bool) bool {
	// Switch statement with all possible combinations for winning.
	switch {
	case data[0] && data[1] && data[2]:
		return true
	case data[3] && data[4] && data[5]:
		return true
	case data[6] && data[7] && data[8]:
		return true
	case data[0] && data[3] && data[6]:
		return true
	case data[1] && data[4] && data[7]:
		return true
	case data[2] && data[5] && data[8]:
		return true
	case data[0] && data[4] && data[8]:
		return true
	case data[2] && data[4] && data[6]:
		return true
	}
	return false
}

// Check the state of the game
func (a *Game) Check() {
	// If index < 4, cant have a result yet
	if a.index < 4 {
		return
	}
	if a.index == 8 {
		// It is a tie if the game hasn't ended before index reaches 8 and no one wins on the ninth placement.
		a.End(dialog.NewInformation("It is a tie!", "Nobody has won. Please try better next time.", a.w))
		return
	}

	if a.CheckWon(a.player1) {
		a.End(dialog.NewInformation("Player 1 has won!", "Congratulations to player 1 for winning.", a.w))
		return
	}

	if a.CheckWon(a.player2) {
		a.End(dialog.NewInformation("Player 2 has won!", "Congratulations to player 2 for winning.", a.w))
		return
	}
}

// End the game
func (g *Game) End(d dialog.Dialog) {
	// Clean up after our game finishes and do it on an other goroutine to speed it up.
	g.inGame = false
	for i := 0; i < 9; i++ {
		g.player1[i] = false
		g.player2[i] = false
		g.pressed[i] = false
	}
	d.Show()
	g.Start()
}

// Press performs a button press in the game state
func (g *Game) Press(i uint8) {
	if !g.inGame {
		return
	}
	// Check if player one or two presses the button and handle it accordingly.
	if g.index%2 == 0 {
		g.buttons[i].SetIcon(circle)
		g.player1[i] = true
	} else {
		g.buttons[i].SetIcon(cross)
		g.player2[i] = true
	}

	// Set the button as pressed to not make it pressable again.
	g.pressed[i] = true
	g.index++
	g.Check()
}
