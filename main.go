package main

import (
	"fmt"
	"victor-ferrer/solver/solver"
	"victor-ferrer/solver/ui"

	"github.com/rivo/tview"
)

func main() {

	// Model
	board := solver.NewBoard()

	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	titleView := tview.NewTextView()
	debugView := tview.NewTextView()

	flex.AddItem(titleView, 0, 1, false)
	table := ui.NewUIBoard(board, app, debugView)
	flex.AddItem(table, 0, 7, false)

	flex.AddItem(debugView, 0, 3, false)

	fmt.Fprintf(titleView, "Exponentile solver")

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
