# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Desktop app (Wails v2, Go + React) that synchronizes toggl track time entries to Redmine, avoiding duplicates and
blocking sync for running or long-closed tickets. UI and docs are in French.

## Commands

Dev environment is provided by Nix + direnv (`flake.nix` / `.envrc`) — Go, the Wails CLI, Node.js, and the
GTK/WebKit system deps are loaded automatically via `direnv allow`.

```sh
wails dev            # dev mode, hot reload of frontend/, devtools at http://localhost:34115
wails build           # production build, binary output in build/bin/
go vet ./...          # from repo root
gofmt -l .            # check formatting
```

There is no Go test suite in this repo currently. Frontend has no lint/test scripts configured either (only
`dev`/`build`/`preview`, run from `frontend/` — normally invoked automatically by Wails, not directly).

The Go build tag `webkit2_41` (required to link against webkitgtk 4.1) is fixed via `build:tags` in `wails.json` —
never pass `-tags` manually.

## Architecture

**Binding boundary**: `main.go` embeds `frontend/dist` + the app icon and starts the Wails runtime, binding the
`App` struct (root `main` package, `app.go`) to the frontend. Every exported method on `App` becomes callable from
JS via generated bindings in `frontend/wailsjs/go/main/App.js` (generated — never hand-edit). This is the only
communication channel between Go and React; there is no REST/GraphQL layer.

**Package layout** (`internal/`):
- `app/` — resolves the OS-conventional config directory (`os.UserConfigDir()`) and migrates it from the legacy
  `~/.toggl-redmine` location on first run after upgrade.
- `cfg/` — config load/save via viper, backed by a JSON file in the app dir; config is stored in the Wails
  `context.Context` (`cfg.ContextWithConfig` / `cfg.ConfigFromContext`) rather than passed as a parameter.
- `api/` — HTTP clients for both external APIs. `api.go` has a shared `call()` helper and a `Service` interface
  (`Prepare`/`CheckUser`); `redmine.go` and `toggltrack.go` implement the per-service specifics (auth header vs.
  basic auth, endpoints, payload shapes). toggl track HTTP 402 responses are surfaced as a typed
  `TogglQuotaExceededError` carrying the reset delay.
- `redmine/` — plain JSON-mapped types for Redmine API responses (`TimeEntry`, `Issue`).
- `toggltrack/` — the core business logic:
  - `issue.go`: regex-parses a toggl entry's description into a Redmine issue number + comment
    (`^(?:[\p{L}\p{N}_]* )?#?([0-9]+)(?: - )?(.*)$`); entries that don't match are flagged invalid and excluded
    from sync.
  - `toggltrack.go`: `ProcessTasks` groups raw toggl entries by day+description (SHA1 hash key) to merge
    duplicates, converts durations to decimal hours rounded to the nearest quarter, sorts by date descending, then
    cross-references against existing Redmine time entries (`mutateWithRedmineEntries`) to prevent re-sending
    already-synced hours. `MarkClosedTooLong` separately disables sync for tasks whose issue closed more than 15
    days ago (checked via `FindIssuesClosedBefore`). A task's final syncability is `IsSyncable()` — valid parse,
    user opted in, not running, not closed-too-long, non-zero duration.
- `logger/` — file-based Wails logger writing to `logs.log` in the app dir.

**Sync flow** (`app.go`): `LoadTasks` fetches toggl entries + Redmine time entries + running-task status
concurrently by intent (sequential calls, shared timeout context), runs them through
`toggltrack.ProcessTasks`/`MarkClosedTooLong`, and returns the merged `AskedTasks` view model to the frontend for
the user to review/select. `SynchronizeTasks` then POSTs only the entries passing `IsSyncable()` to Redmine, one
at a time with an 800ms per-call timeout. Go-side errors are surfaced to the frontend uniformly via
`App.logError(f)`, which both emits a `goError` Wails event and logs through the runtime logger — there's no
per-call error return path to JS.

**Frontend** (`frontend/src/`): React 18 + Vite, `HashRouter` (`AppRouter.jsx`) with three routes — `/` (Home),
`/synchroniser` (Synchronise), `/configurer` (Configure) — wrapped in a shared `Layout`. Pages call the generated
Wails bindings directly (e.g. `CanSynchronize()` from `wailsjs/go/main/App.js`) as if they were local async
functions.

## Release

CI (`.github/workflows/main.yaml`) builds on every pushed tag via `The-Egg-Corp/wails-build-action`, producing
`linux/amd64` and `darwin/universal` binaries, published as a GitHub release. The action auto-detects the Ubuntu
version to install the matching `libwebkit2gtk` dev package and add the `webkit2_41` build tag on 24.04 — this is
redundant with, but harmless alongside, the `build:tags` fixed in `wails.json`.
