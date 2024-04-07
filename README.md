# Exponentile Solver

Golang engine that solves the game Exponentile: https://www.bellika.dk/exponentile

The game is played in a 8x8 board by swapping contiguos tiles. When three or more are lined up, they are combined into a single one bearing the next power of two (2,4,8,16...).
The game is harder as time passes as more numbers are present in the board and is harded to line 3 of them with the same valu.


# Ideas

## Mimic the current game engine (ongoing):
-  Board operations:
 - [ ] Swap tiles.
 - [ ] Drop tiles that match.
 - [ ] Calculate scores of the removed tiles.
-  Project infrastructure:
 - [ ] Add tests and some Github Actions executing them.
-  A UI to play around: 
 -  The most basic form of UI to play the game: A CLI. Just for laughs.

## Implement some strategies that solve the game (future):
- Top-bottom, bottom-top, random, etc.
- Benchmark the score after a given number of moves, maximum score before game end, execution time, etc.
- Have some Github action that postes the results of such benchmarks.

# Current status
- Basic Board model Poc implemented: Swap operation + test.
- Added basic CI that run the tests.
- Basic UI: Board layout + swap operation:

![Current Board status](./docs/ui_board.PNG)

Run it with:
- `go build`
- Execute `solver`.
- Once the board appears, hit enter and the you can select two tiles in order to swap them. 
- That is for now.


As you can see there are some issues with the colouring, maybe I will have to use a black and white schema.



# Stack
- Golang 1.22
- [TVIEW](https://github.com/rivo/tview) for rendering the CLI UI.