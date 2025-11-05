package ui

import (
	"fmt"
	"victor-ferrer/solver/domain"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewUIBoard(board domain.Board, app *tview.Application, text *tview.TextView) *tview.Table {

	table := tview.NewTable().SetBorders(true)

	firstSelectedX := -1
	firstSelectedY := -1

	secondSelectX := -1
	secondSelectY := -1

	cols, rows := 8, 8
	for r := range rows {
		for c := range cols {
			color := tcell.ColorWhite
			value := board.Get(r, c)
			table.SetCell(r, c,

				tview.NewTableCell(fmt.Sprintf("%d", value)).
					SetTextColor(color).
					SetBackgroundColor(tcell.ColorBlack). //getTileColor(value)).
					SetAlign(tview.AlignCenter))

		}
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {

		if firstSelectedX < 0 {
			firstSelectedX = row
			firstSelectedY = column
		} else {
			secondSelectX = row
			secondSelectY = column

			evt := board.MakeMove(domain.CreateTile(firstSelectedX, firstSelectedY), domain.CreateTile(secondSelectX, secondSelectY))

			text.SetText(fmt.Sprintf("Debug: Event type is %s, score=%d", evt.Type, evt.Score))

			if evt.Type == domain.EVENT_TYPE_GAME_UPDATED {

			}

			// Swap tiles in the GUI
			swapTiles(table, firstSelectedX, firstSelectedY, secondSelectX, secondSelectY)

			// Clear selection
			firstSelectedX = -1
			firstSelectedY = -1

			secondSelectX = -1
			secondSelectY = -1
		}

	})

	return table

}

func swapTiles(table *tview.Table, firstSelectedX, firstSelectedY, secondSelectX, secondSelectY int) {

	src := table.GetCell(firstSelectedX, firstSelectedY)
	tgt := table.GetCell(secondSelectX, secondSelectY)

	auxValue := src.Text
	auxColor := src.BackgroundColor

	src.SetText(tgt.Text)
	src.SetBackgroundColor(tgt.BackgroundColor)

	tgt.SetText(auxValue)
	tgt.SetBackgroundColor(auxColor)

}

func getTileColor(value int) tcell.Color {
	switch value {
	case 2:
		return tcell.ColorGreen
	case 4:
		return tcell.ColorBlue
	case 8:
		return tcell.ColorYellow
	case 16:
		return tcell.ColorOrange
	case 32:
		return tcell.ColorPink
	case 64:
		return tcell.ColorRed
	case 128:
		return tcell.ColorPurple
	case 518:
		return tcell.ColorTurquoise
	case 1024:
		return tcell.ColorMaroon
	case 2048:
		return tcell.ColorGray
	default:
		return tcell.ColorGray
	}
}
