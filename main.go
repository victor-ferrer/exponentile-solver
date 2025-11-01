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

	table := ui.NewUIBoard(board, app)
	flex.AddItem(table, 0, 7, false)

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
