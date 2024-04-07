# Exponentile Solver

Golang engine that solves the game Exponentile: https://www.bellika.dk/exponentile

# Ideas

## Mimic the current game engine (ongoing):
-  Board operations: Swap tiles, drop tiles that match, finch machting tiles, etc.
-  Add tests and some Github Actions executing them.
-  Add the most basic form of UI to play the game: A CLI. Just for laughs.

## Implement some strategies that solve the game:
- Top-bottom, bottom-top, random, etc.
- Benchmark the score after a given number of moves, maximum score before game end, execution time, etc.
- Have some Github action that postes the results of such benchmarks.

# Current status
- Basic Board model Poc implemented: Swap operation + test
- Basic UI: Board layout + swap operation


![Current Board status](./docs/ui_board.PNG)





# Stack
- Golang 1.22
- [TVIEW](https://github.com/rivo/tview) for rendering the CLI UI.