package main

import (
	"victor-ferrer/solver/domain"
	"victor-ferrer/solver/ui"

	"github.com/rivo/tview"
)

func main() {

	board := domain.NewBoard()

	app := tview.NewApplication()

	// Create menu to select game mode
	list := tview.NewList().
		AddItem("Manual Game", "Play the game manually", 'm', nil).
		AddItem("Automatic Mode", "Watch the solver play", 'a', nil).
		AddItem("Quit", "Exit the application", 'q', nil)

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		var autoMode bool
		switch index {
		case 0: // Manual Game
			autoMode = false
		case 1: // Automatic Mode
			autoMode = true
		case 2: // Quit
			app.Stop()
			return
		}

		// Clear the app and set up the game board
		flex := tview.NewFlex()
		flex.SetDirection(tview.FlexColumn)

		uiBoard := ui.NewUIBoard(&board, app, autoMode)
		flex.AddItem(uiBoard.Table, 57, 4, false)
		flex.AddItem(uiBoard.DebugTxt, 0, 1, false)

		app.SetRoot(flex, true).SetFocus(uiBoard.Table)
	})

	app.SetRoot(list, true).SetFocus(list)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
