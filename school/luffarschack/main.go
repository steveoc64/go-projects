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

// InitGUI starts up the whole interface for out program.
func InitGUI() {
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

			// Ugly switch statement to set the icons for our buttons.
			if clicked[0] && !pressed[0] {
				pressed[0] = true
				if index%2 == 0 {
					button1.SetIcon(circle)
					player1[0] = true
				} else {
					button1.SetIcon(cross)
					player2[0] = true
				}
				index++
			} else if clicked[1] && !pressed[1] {
				pressed[1] = true
				if index%2 == 0 {
					button2.SetIcon(circle)
					player1[1] = true
				} else {
					button2.SetIcon(cross)
					player2[1] = true
				}
				index++
			} else if clicked[2] && !pressed[2] {
				pressed[2] = true
				if index%2 == 0 {
					button3.SetIcon(circle)
					player1[2] = true
				} else {
					button3.SetIcon(cross)
					player2[2] = true
				}
				index++
			} else if clicked[3] && !pressed[3] {
				pressed[3] = true
				if index%2 == 0 {
					button4.SetIcon(circle)
					player1[3] = true
				} else {
					button4.SetIcon(cross)
					player2[3] = true
				}
				index++
			} else if clicked[4] && !pressed[4] {
				pressed[4] = true
				pressed[4] = true
				if index%2 == 0 {
					button5.SetIcon(circle)
					player1[4] = true
				} else {
					button5.SetIcon(cross)
					player2[4] = true
				}
				index++
			} else if clicked[5] && !pressed[5] {
				pressed[5] = true
				if index%2 == 0 {
					button6.SetIcon(circle)
					player1[5] = true
				} else {
					button6.SetIcon(cross)
					player2[5] = true
				}
				index++
			} else if clicked[6] && !pressed[6] {
				pressed[6] = true
				if index%2 == 0 {
					button7.SetIcon(circle)
					player1[6] = true
				} else {
					button7.SetIcon(cross)
					player2[6] = true
				}
				index++
			} else if clicked[7] && !pressed[7] {
				pressed[7] = true
				if index%2 == 0 {
					button8.SetIcon(circle)
					player1[7] = true
				} else {
					button8.SetIcon(cross)
					player2[7] = true
				}
				index++
			} else if clicked[8] && !pressed[8] {
				pressed[8] = true
				if index%2 == 0 {
					button9.SetIcon(circle)
					player1[8] = true
				} else {
					button9.SetIcon(cross)
					player2[8] = true
				}
				index++
			} else {
				// Just run the print command so we don't stall the gui when nothing happends in the loop. Avoid printing anything, it's just useless memory usage.
				fmt.Print("")
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
					message.Show()
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

	// Append each new row as a new container with gird layout and three buttons.
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

func main() {
	InitGUI()
}

// CheckWon checks all possible combinations for winning.
func CheckWon(player [9]bool) bool {

	// Switch statement with all possible combinations for winning.
	switch {
	case player[0] && player[1] && player[2]:
		return true
	case player[0] && player[3] && player[6]:
		return true
	case player[0] && player[3] && player[5]:
		return true
	case player[6] && player[7] && player[8]:
		return true
	case player[2] && player[5] && player[8]:
		return true
	case player[0] && player[4] && player[8]:
		return true
	case player[2] && player[4] && player[6]:
		return true
	case player[1] && player[4] && player[7]:
		return true
	case player[3] && player[4] && player[5]:
		return true
	}

	return false
}
