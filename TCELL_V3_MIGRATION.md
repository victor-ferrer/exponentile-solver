# TCEll v2 to v3 Migration Analysis

## Current State
- **Current Version**: tcell v2.13.4
- **Import Statement**: `github.com/gdamore/tcell/v2`
- **Integration**: Used indirectly through tview v0.42.0
- **Usage File**: `ui/cli_board.go`

## Breaking Changes in tcell v3 (That Affect This Project)

### 1. **Import Path Change** ✅ MINIMAL IMPACT
- **Change**: Import path changes from `github.com/gdamore/tcell/v2` → `github.com/gdamore/tcell/v3`
- **Current Usage**: Line 9 in `cli_board.go`
- **Action Required**: Update import statement
- **Code Changes**:
  ```go
  // BEFORE:
  "github.com/gdamore/tcell/v2"
  
  // AFTER:
  "github.com/gdamore/tcell/v3"
  ```

### 2. **Key Event Handler Changes** ✅ NO IMPACT FOR THIS PROJECT
- **What Changed**: `EventKey.Rune()` method removed, replaced with `Str()` which returns a string instead of single rune
- **Why No Impact**: This project only uses `tcell.Key` constants (`KeyEscape`, `KeyEnter`), not rune events
- **Usage in Project**: Line 54-60 in `cli_board.go`
  ```go
  uiBoard.Table.SetDoneFunc(func(key tcell.Key) {
      if key == tcell.KeyEscape { ... }
      if key == tcell.KeyEnter { ... }
  })
  ```
- **Status**: No changes needed - key constants work the same

### 3. **Color API Changes** ✅ FULLY COMPATIBLE
- **What Changed**: `Color` type reduced from 64-bit to 32-bit
- **Why No Impact**: tcell v3 maintains backward compatibility for color constants
- **Usage in Project**: Lines 148, 156, 161-187 in `cli_board.go`
  ```go
  cell.SetBackgroundColor(tcell.ColorRed)
  cell.SetBackgroundColor(tcell.ColorGreen)
  // ... etc (all color constants work the same)
  ```
- **Status**: No changes needed - color names and usage identical

### 4. **Terminal Type Support** ⚠️ DEPENDENCY CONSIDERATION
- **What Changed**: Terminfo database removed, legacy terminal support dropped
- **Minimum Requirement**: Windows 10 build 1703+ on Windows
- **Modern Terminals**: All modern terminal emulators work fine (xterm, iTerm2, Terminal.app, ConEmu, Windows Terminal, etc.)
- **Risk Level**: Very low for modern systems
- **Status**: No code changes needed, but requires updated terminal environment

### 5. **Event System Changes** ⚠️ DEPENDENCY CONSIDERATION (via tview)
- **What Changed**: 
  - `PostEvent()`, `PollEvent()`, `ChannelEvents()`, `PostEventWait()` functions removed
  - New direct channel access via `EventQ`
- **Why Low Impact**: This project doesn't directly use these functions - they're handled by tview
- **Risk**: Only matters if tview hasn't updated to tcell v3 yet
- **Status**: Waiting on tview compatibility

## Dependency Chain Analysis

### Current Chain:
```
exponentile-solver
  └─ github.com/gdamore/tcell/v2 (v2.13.4)
  └─ github.com/rivo/tview (v0.42.0)
       └─ github.com/gdamore/tcell/v2 (v2.13.4)
```

### After Migration:
```
exponentile-solver
  └─ github.com/gdamore/tcell/v3 (v3.x.x)
  └─ github.com/rivo/tview (v0.42.0 OR NEWER)
       └─ github.com/gdamore/tcell/v3 (v3.x.x)
```

### Critical: tview Compatibility
- **Status**: Must check if tview v0.42.0+ supports tcell v3
- **Risk**: If tview hasn't updated, this migration will break the application
- **Recommendation**: Verify tview compatibility before proceeding

## Migration Checklist

- [ ] Check tview compatibility with tcell v3
  - Run: `go get -u github.com/rivo/tview@latest` and check if it supports tcell/v3
  - Or search: https://github.com/rivo/tview/issues for tcell v3 discussion
  
- [ ] Update go.mod dependency:
  ```bash
  go get github.com/gdamore/tcell/v3@latest
  ```

- [ ] Update import in ui/cli_board.go (line 9):
  ```go
  "github.com/gdamore/tcell/v3"  // was: "github.com/gdamore/tcell/v2"
  ```

- [ ] Run tests:
  ```bash
  go test ./...
  ```

- [ ] Build and test the application:
  ```bash
  go build
  ./solver.exe
  ```

- [ ] Verify all colors still display correctly
- [ ] Verify keyboard input (Escape, Enter) still works
- [ ] Test on target Windows version (ensure 1703+)

## Code Impact Summary

### Files to Change: 1
- `ui/cli_board.go` - Line 9 only

### Lines of Code Affected: 1
- Import statement change

### API Breaking Changes Affecting This Code: 0
- All used features (tcell.Key constants, tcell.Color constants) are fully compatible

### Risk Level: **LOW**
- No application code changes required
- Only import path and dependency updates needed
- All used APIs remain unchanged

## Recommended Migration Path

1. **Verify tview support first** (BLOCKING)
   - Check if tview has released a version supporting tcell v3
   - If not, wait for tview update OR consider updating tview separately

2. **Update dependencies**:
   ```bash
   go get -u github.com/gdamore/tcell/v3@latest
   go get -u github.com/rivo/tview@latest
   go mod tidy
   ```

3. **Update single import line** in `ui/cli_board.go`

4. **Test thoroughly** (even though changes are minimal, testing is important)

5. **Update CI/CD** if using GitHub Actions to test on Windows 1703+

## Additional Notes

- The project's use of tcell is very minimal and indirect (through tview)
- Most breaking changes in tcell v3 don't affect this codebase
- The main risk is tview compatibility, not tcell direct usage
- Modern terminal support is actually improved in v3
- No functional changes needed in the game logic or UI code
