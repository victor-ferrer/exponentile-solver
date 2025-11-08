package ui

import (
	"fmt"
	"victor-ferrer/solver/domain"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	COLS = 8
	ROWS = 8
)

func NewUIBoard(board domain.Board, app *tview.Application, debugTxt *tview.TextView) *tview.Table {

	table := tview.NewTable().SetBorders(false).SetSeparator('│')

	firstSelectedX := -1
	firstSelectedY := -1

	secondSelectX := -1
	secondSelectY := -1

	for r := range ROWS {
		for c := range COLS {
			color := tcell.ColorWhite
			value := board.Get(r, c)
			table.SetCell(r, c,

				tview.NewTableCell(fmt.Sprintf(" \n %d \n ", value)).
					SetTextColor(color).
					SetBackgroundColor(getTileColor(value)).
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

			events := board.MakeMove(domain.CreateTile(firstSelectedX, firstSelectedY), domain.CreateTile(secondSelectX, secondSelectY))

			for _, evt := range events {
				debugTxt.SetText(fmt.Sprintf("Debug: \n - Event type: %s \n - Score: %d", evt.Type, evt.Score))

				switch evt.Type {
				case domain.EVENT_TYPE_GAME_UPDATED:

					if len(evt.GroupedTiles) > 0 {
						groupedTilesTxt := ""
						for _, tile := range evt.GroupedTiles {
							groupedTilesTxt += fmt.Sprintf("(%d,%d)(%d) ", tile.X, tile.Y, board.Get(tile.X, tile.Y))
						}
						debugTxt.SetText(fmt.Sprintf("%s \n - Grouped tiles: %s", debugTxt.GetText(true), groupedTilesTxt))
					}

					renderTileStates(evt.Tiles, table)
				case domain.EVENT_TYPE_GAME_OVER:
					debugTxt.SetText(fmt.Sprintf("%s \n - Game Over", debugTxt.GetText(true)))
					table.SetSelectable(false, false)
				}
			}

			// Clear selection
			firstSelectedX = -1
			firstSelectedY = -1

			secondSelectX = -1
			secondSelectY = -1
		}

	})

	return table

}

func renderTileStates(tiles []domain.TileState, table *tview.Table) {
	for _, tileState := range tiles {
		cell := tview.NewTableCell(fmt.Sprintf(" \n %d \n ", tileState.Value))
		cell.SetBackgroundColor(getTileColor(tileState.Value))
		table.SetCell(tileState.Position.X, tileState.Position.Y, cell)
	}
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
	case 256:
		return tcell.ColorTurquoise
	case 512:
		return tcell.ColorTurquoise
	case 1024:
		return tcell.ColorMaroon
	case 2048:
		return tcell.ColorDarkCyan
	default:
		return tcell.ColorGray
	}
}
