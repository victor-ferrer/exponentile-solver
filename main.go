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
	flex.SetDirection(tview.FlexColumn)

	table, debugTxt := ui.NewUIBoard(&board, app)
	flex.AddItem(table, 57, 4, false)
	flex.AddItem(debugTxt, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
