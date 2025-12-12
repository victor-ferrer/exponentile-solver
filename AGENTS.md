# AGENTS.md

This file provides guidance for AI coding agents working on the Exponentile Solver project.

## Project Overview

Golang engine that solves the Exponentile board game (https://www.bellika.dk/exponentile).

The game is played on an 8x8 board by swapping contiguous tiles. When three or more tiles with the same value line up, they combine into a single tile bearing the next power of two (2, 4, 8, 16, ...).

## Architecture Principles

### Domain-Driven Design

- **Domain Layer** (`domain/` package): Contains the core game logic and board operations. This is the domain model and should remain pure and isolated.
- **Event-Based Communication**: The domain communicates with the outside world via events, not direct calls.
- **UI Layer** (`ui/` package): Currently implements a CLI interface using TVIEW. Other UIs (event sourcing, web, etc.) may be added later.
- **Separation of Concerns**: Keep domain logic independent of UI and infrastructure concerns.

### Package Structure

- `domain/`: Domain layer - board game logic, tile operations, game rules
- `ui/`: User interface implementations (currently CLI only)
- `main.go`: Application entry point

## Technology Stack

- **Language**: Go 1.24
- **CLI UI**: [TVIEW](https://github.com/rivo/tview) for rendering the terminal UI
- **Matrix Operations**: [GoNum](https://github.com/gonum/matrix) matrix package

## Common Commands

### Build and Run
```bash
go build
.\solver.exe
```

### Testing
```bash
go test ./...
```

### Linting (planned)
```bash
# Linters will be added to CI pipeline
```

## Code Conventions

- Follow standard Go conventions and idiomatic patterns
- Keep domain logic in `domain/` package pure and testable
- UI implementations should consume domain events rather than directly manipulating domain objects
- Write tests for new domain logic in `*_test.go` files
- **Update README.md** along with code changes to keep documentation in sync

## Game Mechanics

### Group Detection (`findGroup`)

The `findGroup(x, y int)` function identifies matching tiles that can be combined:

**Algorithm:**
1. Scans horizontally (left/right) from position (x,y) to find contiguous tiles with the same value
2. Scans vertically (up/down) from position (x,y) to find contiguous tiles with the same value
3. Returns all tiles from runs of 3 or more in both directions
4. If both horizontal and vertical runs exist (forming a cross/plus shape), returns tiles from both runs
5. The center tile at (x,y) is only included once when both runs exist

**Example:**
- Row 7: [16, 16, 16, 16, 8, 2, 2, 2]
- `findGroup(7, 0)` returns 4 tiles: [(7,0), (7,1), (7,2), (7,3)]
- `findGroup(7, 7)` returns 3 tiles: [(7,5), (7,6), (7,7)]

### Group Merging (`MakeMove`)

When a valid group is found after a swap:

1. **Upgrade moved tile**: The tile that was moved (t2 position, or t1 if t2 is not in the group) is upgraded based on group size. The replacement value is calculated as: `value * 2^(group_size - 2)`. For 3 tiles, this is `value * 2`, for 4 tiles `value * 4`, for 5 tiles `value * 8`, etc.
2. **Drop other tiles**: All other tiles in the group are dropped and replaced with random tiles (values 2, 4, 8, 16, or 32)
3. **Calculate and increment score**: Score increments by the sum of all tile values in the group (`value * group_size`)
4. **Cascade detection**: After dropping tiles, scan the board for new groups formed by the dropped tiles. Repeat steps 1-3 until no more groups form.

**Example with 4-tile group:**
- Group: [(7,0), (7,1), (7,2), (7,3)] all with value 16
- Move tile from (7,3) to (7,2): creates a group
- After merge: position (7,2) (the moved tile) becomes 64 (16 * 2^(4-2) = 16 * 4), others are dropped and replaced
- Score increment: 16 + 16 + 16 + 16 = 64 points (total score is cumulative across all moves)
- If dropped tiles form new groups, cascade continues with additional score increments

### Cascade Detection

After each group merge, the board is scanned for cascading groups:

**Algorithm:**
1. Scan from bottom-to-top (x: width-1 → 0), left-to-right (y: 0 → width-1)
2. For each position, detect if a group exists
3. If found: process the group, create a `GameEvent`, and restart scan
4. Repeat until no more groups are found

**Event Output:**
- Each group (initial + cascades) generates a separate `GAME_UPDATED` event
- All events from a single `MakeMove` have the same sequence number
- Scores are cumulative across events within a move

## Current Development Focus

1. **Board Operations** (completed):
- ✅ Swap tiles
- ✅ Drop tiles that match
- ✅ Get groups of tiles (supports 3+ tiles, horizontal and vertical runs)
- ✅ Calculate scores of removed tiles
- ✅ Reevaluate the board once a group is gone (detect cascading matches)
- ✅ Bigger groups return higher new tiles   
- ✅ Cascade detection and processing

2. **UI Development** (completed):
   - ✅ Basic CLI interface using TVIEW
   - ✅ Map basic board operations to UI
   - ✅ Game mode menu (Manual vs Automatic)
   - ✅ Manual mode (player input and tile selection)
   - ✅ Automatic mode (solver visualization)

3. **Infrastructure**:
   - ✅ Tests with GitHub Actions
   - ⏳ Add linters to CI pipeline

4. **Future**: Implement solving strategies (top-bottom, bottom-top, random, etc.) with benchmarking

## Notes for AI Agents

- When adding new UI implementations, ensure they follow the event-based communication pattern with the domain
- Domain logic should never depend on UI or infrastructure code
- Keep the `domain/` package framework-agnostic and testable
- Consider event sourcing as a potential future UI/persistence layer
