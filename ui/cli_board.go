package ui

import (
	"fmt"
	"time"
	"victor-ferrer/solver/domain"
	bruteforce "victor-ferrer/solver/solvers/brute_force"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	COLS = 8
	ROWS = 8
)

type UIBoard struct {
	Table    *tview.Table
	DebugTxt *tview.TextView
	app      *tview.Application
}

func NewUIBoard(board domain.Board, app *tview.Application, auto bool) UIBoard {

	debugTxt := tview.NewTextView()
	debugTxt.SetText("Debug:")

	table := tview.NewTable().SetBorders(false).SetSeparator('│')

	uiBoard := UIBoard{
		Table:    table,
		DebugTxt: debugTxt,
		app:      app,
	}

	renderTileStates(board.GetTileState(), table)

	if auto {
		uiBoard.createSolver(board)
	} else {
		uiBoard.createSelectFunc(board)
		table.Select(0, 0).SetFixed(1, 1)
	}

	uiBoard.createDoneFunc()

	return uiBoard

}

// Controls when the app ends
func (uiBoard *UIBoard) createDoneFunc() {
	uiBoard.Table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			uiBoard.app.Stop()
		}
		if key == tcell.KeyEnter {
			uiBoard.Table.SetSelectable(true, true)
		}
	})
}

// Creates a solver for the board and feeds its events to the ui table
func (uiBoard *UIBoard) createSolver(board domain.Board) {
	eventsChan := bruteforce.Solve(board)
	go func() {
		for event := range eventsChan {

			if len(event.Group.Tiles) > 0 {
				uiBoard.app.QueueUpdateDraw(func() {
					highlightTiles(event, uiBoard.Table)
				})
			}
			time.Sleep(1 * time.Second)
			uiBoard.app.QueueUpdateDraw(func() {
				uiBoard.DebugTxt.SetText(fmt.Sprintf("Debug: \n - Event type: %s \n - Score: %d", event.Type, event.Score))
				renderTileStates(event.Tiles, uiBoard.Table)

			})
			time.Sleep(1 * time.Second)
		}
	}()
}

func (uiBoard *UIBoard) createSelectFunc(board domain.Board) {
	firstSelectedX := -1
	firstSelectedY := -1
	secondSelectX := -1
	secondSelectY := -1

	selectFunc := func(row int, column int) {
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
						uiBoard.app.QueueUpdateDraw(func() {
							highlightTiles(evt, uiBoard.Table)
						})
					}
					time.Sleep(1 * time.Second)

					uiBoard.app.QueueUpdateDraw(func() {
						uiBoard.DebugTxt.SetText(fmt.Sprintf("Debug: \n - Event type: %s \n - Score: %d", evt.Type, evt.Score))

						switch evt.Type {
						case domain.EVENT_TYPE_GAME_UPDATED:

							if len(evt.Group.Tiles) > 0 {
								groupedTilesTxt := ""
								for _, tile := range evt.Group.Tiles {
									groupedTilesTxt += fmt.Sprintf("(%d,%d)(%d) ", tile.X, tile.Y, evt.Group.Value)
								}
								uiBoard.DebugTxt.SetText(fmt.Sprintf("%s \n - Grouped tiles: %s", uiBoard.DebugTxt.GetText(true), groupedTilesTxt))
							}

							renderTileStates(evt.Tiles, uiBoard.Table)
						case domain.EVENT_TYPE_GAME_OVER:
							uiBoard.DebugTxt.SetText(fmt.Sprintf("%s \n - Game Over", uiBoard.DebugTxt.GetText(true)))
							uiBoard.Table.SetSelectable(false, false)
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
	uiBoard.Table.SetSelectedFunc(selectFunc)
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
