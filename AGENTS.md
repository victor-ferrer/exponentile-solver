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

- **Language**: Go 1.22
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

1. **Upgrade middle tile**: The tile at index `len(group)/2` is upgraded to the next power of 2
2. **Drop other tiles**: All other tiles in the group are dropped and replaced with random tiles (values 2, 4, 8, 16, or 32)
3. **Calculate score**: Score = `currentValue * len(group)` (e.g., 4 tiles of value 16 = 64 points)

**Example with 4-tile group:**
- Group: [(7,0), (7,1), (7,2), (7,3)] all with value 16
- Middle index: 4/2 = 2
- After merge: position (7,2) becomes 32, others are dropped and replaced
- Score: 16 * 4 = 64 points

## Current Development Focus

1. **Board Operations** (ongoing):
   - ✅ Swap tiles
   - ✅ Drop tiles that match
   - ✅ Get groups of tiles (supports 3+ tiles, horizontal and vertical runs)
   - ✅ Calculate scores of removed tiles

2. **UI Development** (ongoing):
   - ✅ Basic CLI interface
   - ⏳ Map basic board operations to UI

3. **Infrastructure**:
   - ✅ Tests with GitHub Actions
   - ⏳ Add linters to CI pipeline

4. **Future**: Implement solving strategies (top-bottom, bottom-top, random, etc.) with benchmarking

## Notes for AI Agents

- When adding new UI implementations, ensure they follow the event-based communication pattern with the domain
- Domain logic should never depend on UI or infrastructure code
- Keep the `domain/` package framework-agnostic and testable
- Consider event sourcing as a potential future UI/persistence layer
