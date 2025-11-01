package main

import (
	"victor-ferrer/solver/domain"
	"victor-ferrer/solver/ui"

	"github.com/rivo/tview"
)

func main() {

	// Model
	board := domain.NewBoard()

	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	text := tview.NewTextView()

	text.SetText("Debug:")
	flex.AddItem(text, 0, 1, false)

	table := ui.NewUIBoard(board, app, text)
	flex.AddItem(table, 0, 8, false)

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
