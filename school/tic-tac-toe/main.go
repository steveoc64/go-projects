package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Define all our global variables. We use these in order to not redeclare them every time we start the game.
var (
	// Variable for defining if a button is pressed or not. Redefined on every press of the button.
	clicked = [9]bool{}

	// Variables for which parts the players own. Index 0 is the first tile and index 8 is the ninth tile.
	player1 = [9]bool{}
	player2 = [9]bool{}

	// Variable to handle if buttons are already pressed. Does not change on button press.
	pressed = [9]bool{}

	// Index for defining if it's player one or player two's turn to play. Use 8bit variable to save on memory allocation.
	index uint8
)

func main() {
	InitGUI()
}

// InitGUI starts up the whole interface for out program.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Set the application icon for our program.
	app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Tic-Tac-Toe")

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
		// Make sure that all buttons are reset before we start.
		button1.SetIcon(nil)
		button2.SetIcon(nil)
		button3.SetIcon(nil)
		button4.SetIcon(nil)
		button5.SetIcon(nil)
		button6.SetIcon(nil)
		button7.SetIcon(nil)
		button8.SetIcon(nil)
		button9.SetIcon(nil)

		// Clear the index to make sure that we start from scratch.
		index = 0

		// Clear the board information for each player, each button click and all buttons that have already been pressed.
		player1 = [9]bool{false, false, false, false, false, false, false, false, false}
		player2 = [9]bool{false, false, false, false, false, false, false, false, false}
		clicked = [9]bool{false, false, false, false, false, false, false, false, false}
		pressed = [9]bool{false, false, false, false, false, false, false, false, false}

		// The main for loop where our game plays from. We want to always loop until we manually break it.
		for {

			// Handles all our button presses during the play time.
			switch {
			case clicked[0] && !pressed[0]:
				index = PressHandler(button1, 0, index)
			case clicked[1] && !pressed[1]:
				index = PressHandler(button2, 1, index)
			case clicked[2] && !pressed[2]:
				index = PressHandler(button3, 2, index)
			case clicked[3] && !pressed[3]:
				index = PressHandler(button4, 3, index)
			case clicked[4] && !pressed[4]:
				index = PressHandler(button5, 4, index)
			case clicked[5] && !pressed[5]:
				index = PressHandler(button6, 5, index)
			case clicked[6] && !pressed[6]:
				index = PressHandler(button7, 6, index)
			case clicked[7] && !pressed[7]:
				index = PressHandler(button8, 7, index)
			case clicked[8] && !pressed[8]:
				index = PressHandler(button9, 8, index)
			default:
				fmt.Print("") // Just run empty print command so we don't stall the gui when nothing happends in the loop.
			}

			// Check if index is bigger or equal to 5, because it's the earliest time we can win. If index is 9, we have a tie and nobody won.
			if index == 9 {
				// It is a tie if the game hasn't ended before index reaches 9.
				message := dialog.NewInformation("It is a tie!", "Nobody has won. Please try better next time.", window)
				message.Show()
				break
			} else if index >= 5 {
				if CheckWon(player1) {
					// Show a dialogue informing the first player that he won!
					message := dialog.NewInformation("Player 1 has won!", "Congratulations to player 1 for winning.", window)
					message.Show() // Need to have a return value so we wait for the function to complete fully before continuing.

					break
				} else if CheckWon(player2) {
					// Show a dialogue informing the second player that he won!
					message := dialog.NewInformation("Player 2 has won!", "Congratulations to player 2 for winning.", window)
					message.Show()
					break
				}
			}
		}
	}))

	// Append each new row as a new container with grid layout and three buttons.
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button1, button2, button3))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button4, button5, button6))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button7, button8, button9))

	// Add our vertical box to be viewed.
	window.SetContent(vbox)

	// Set a sane default for the window size.
	window.Resize(fyne.NewSize(400, 100))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}

// PressHandler Handles the press of a button and updates button icons and player arrays accordingly.
func PressHandler(button *widget.Button, num, player uint8) uint8 {
	// Check if player one or two presses the button and handle accordingly.
	if player%2 == 0 {
		button.SetIcon(circle)
		player1[num] = true
	} else {
		button.SetIcon(cross)
		player2[num] = true
	}

	// Set the button as pressed.
	pressed[num] = true

	// Need to have a return value so we wait for the function to complete fully before continuing. Thus we bump player index.
	return player + 1
}

// CheckWon checks all possible combinations for winning.
func CheckWon(player [9]bool) bool {

	// Switch statement with all possible combinations for winning.
	switch {
	case player[0] && player[1] && player[2]:
		return true
	case player[3] && player[4] && player[5]:
		return true
	case player[6] && player[7] && player[8]:
		return true
	case player[0] && player[3] && player[6]:
		return true
	case player[1] && player[4] && player[7]:
		return true
	case player[2] && player[5] && player[8]:
		return true
	case player[0] && player[4] && player[8]:
		return true
	case player[2] && player[4] && player[6]:
		return true
	}

	return false
}
