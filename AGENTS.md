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

## Current Development Focus

1. **Board Operations** (ongoing):
   - ✅ Swap tiles
   - ✅ Drop tiles that match
   - ⏳ Get groups of tiles
   - ⏳ Calculate scores of removed tiles

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
