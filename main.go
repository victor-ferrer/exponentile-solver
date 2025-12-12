package main

import (
	"victor-ferrer/solver/domain"
	"victor-ferrer/solver/ui"

	"github.com/rivo/tview"
)

func main() {

	board := domain.NewBoard()

	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexColumn)

	uiBoard := ui.NewUIBoard(&board, app, false)
	flex.AddItem(uiBoard.Table, 57, 4, false)
	flex.AddItem(uiBoard.DebugTxt, 0, 1, false)

	app.SetRoot(flex, true).SetFocus(uiBoard.Table)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
