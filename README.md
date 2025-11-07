# Exponentile Solver

Golang engine that solves the game Exponentile: https://www.bellika.dk/exponentile

The game is played on an 8x8 board by swapping contiguous tiles. When three or more tiles with the same value line up, they combine into a single tile bearing the next power of two (2, 4, 8, 16, ...).

The game becomes harder as more numbers appear on the board, making it increasingly difficult to line up 3 or more tiles with the same value.

## Game Mechanics

### Group Detection

The engine scans for matching tiles both horizontally and vertically:
- Groups must contain **3 or more** contiguous tiles with the same value
- Both horizontal and vertical runs are detected and combined if they form a cross/plus shape
- The center tile is counted only once when both runs exist

**Example:**
```
Row 7: [16, 16, 16, 16, 8, 2, 2, 2]
- Positions (7,0) through (7,3): 4-tile group of 16s
- Positions (7,5) through (7,7): 3-tile group of 2s
```

### Group Merging

When a valid group is formed after a swap:
1. The **middle tile** (at index `len(group)/2`) is upgraded to the next power of 2
2. All other tiles in the group are dropped and replaced with random tiles (values: 2, 4, 8, 16, or 32)
3. **Score** is calculated as: `currentValue × groupSize`

**Example:**
- Group of 4 tiles with value 16
- Middle tile at position 2 becomes 32
- Other 3 tiles are dropped and replaced
- Score: 16 × 4 = 64 points

## Project Status

### Completed Features ✅

**Board Operations:**
- ✅ Swap tiles
- ✅ Drop tiles that match
- ✅ Get groups of tiles (supports 3+ tiles, horizontal and vertical runs)
- ✅ Calculate scores of removed tiles

**Infrastructure:**
- ✅ Tests with GitHub Actions
- ⏳ Linters on CI pipeline (planned)

**UI:**
- ✅ Basic CLI interface using TVIEW
- ⏳ Map basic board operations to UI (in progress)

### Future Plans

**Implement solving strategies:**
- Top-bottom, bottom-top, random, etc.
- Benchmark: score after N moves, maximum score before game end, execution time
- GitHub Actions to post benchmark results

## How to Run

**Build and run:**
```bash
go build
.\solver.exe
```

**Or run directly:**
```bash
go run main.go
```

**Gameplay:**
1. Once the board appears, hit Enter
2. Select two tiles to swap them
3. Valid swaps that create groups of 3+ matching tiles will merge them

## How to Test

```bash
go test ./...
```

## Current UI

![Current Board status](./docs/ui_board.PNG)

Black and White color schema for now.

## Technology Stack

- **Language:** Go 1.22
- **CLI UI:** [TVIEW](https://github.com/rivo/tview) for rendering the terminal UI
- **Matrix Operations:** [GoNum](https://github.com/gonum/matrix) matrix package

## Architecture

This project follows Domain-Driven Design principles:

- **Domain Layer** (`domain/` package): Contains core game logic and board operations
- **Event-Based Communication**: Domain communicates via events, not direct calls
- **UI Layer** (`ui/` package): CLI interface implementation (other UIs may be added later)
- **Separation of Concerns**: Domain logic is independent of UI and infrastructure

See [AGENTS.md](AGENTS.md) for detailed development guidelines and game mechanics documentation.
