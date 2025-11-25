package ui

import (
	"fmt"
	"time"
	"victor-ferrer/solver/domain"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	COLS = 8
	ROWS = 8
)

func NewUIBoard(board domain.Board, app *tview.Application) (*tview.Table, *tview.TextView) {

	debugTxt := tview.NewTextView()
	debugTxt.SetText("Debug:")

	table := tview.NewTable().SetBorders(false).SetSeparator('│')
	renderTileStates(board.GetTileState(), table)

	selectFunc := createSelectFunc(board, app, table, debugTxt)
	doneFunc := createDoneFunc(app, table)
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(doneFunc).SetSelectedFunc(selectFunc)

	return table, debugTxt

}

func createDoneFunc(app *tview.Application, table *tview.Table) func(key tcell.Key) {
	return func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}
}

func createSelectFunc(
	board domain.Board,
	app *tview.Application,
	table *tview.Table,
	debugTxt *tview.TextView,
) func(row int, column int) {
	firstSelectedX := -1
	firstSelectedY := -1
	secondSelectX := -1
	secondSelectY := -1

	return func(row int, column int) {
		if firstSelectedX < 0 {
			firstSelectedX = row
			firstSelectedY = column
		} else {
			secondSelectX = row
			secondSelectY = column

			events := board.MakeMove(domain.CreateTile(firstSelectedX, firstSelectedY), domain.CreateTile(secondSelectX, secondSelectY))

			go func() {
				for _, evt := range events {

					if len(evt.Group.Tiles) > 0 {
						app.QueueUpdateDraw(func() {
							highlightTiles(evt, table)
						})
					}
					time.Sleep(1 * time.Second)

					app.QueueUpdateDraw(func() {
						debugTxt.SetText(fmt.Sprintf("Debug: \n - Event type: %s \n - Score: %d", evt.Type, evt.Score))

						switch evt.Type {
						case domain.EVENT_TYPE_GAME_UPDATED:

							if len(evt.Group.Tiles) > 0 {
								groupedTilesTxt := ""
								for _, tile := range evt.Group.Tiles {
									groupedTilesTxt += fmt.Sprintf("(%d,%d)(%d) ", tile.X, tile.Y, evt.Group.Value)
								}
								debugTxt.SetText(fmt.Sprintf("%s \n - Grouped tiles: %s", debugTxt.GetText(true), groupedTilesTxt))
							}

							renderTileStates(evt.Tiles, table)
						case domain.EVENT_TYPE_GAME_OVER:
							debugTxt.SetText(fmt.Sprintf("%s \n - Game Over", debugTxt.GetText(true)))
							table.SetSelectable(false, false)
						}
					})
				}
			}()

			firstSelectedX = -1
			firstSelectedY = -1

			secondSelectX = -1
			secondSelectY = -1
		}
	}
}

func highlightTiles(evt domain.GameEvent, table *tview.Table) {
	for _, tileState := range evt.Group.Tiles {
		cell := tview.NewTableCell(fmt.Sprintf(" \n %d \n ", evt.Group.Value))
		cell.SetBackgroundColor(tcell.ColorRed)
		table.SetCell(tileState.X, tileState.Y, cell)
	}
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
		return tcell.ColorGold
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
