# Contributing to teams-cli

Thanks for helping improve teams-cli. This document explains how to propose changes effectively.

## Code of conduct

By participating, you agree to uphold our [**Code of Conduct**](./CODE_OF_CONDUCT.md).

## What to contribute

- Bug fixes and reliability improvements  
- Documentation and error messages  
- Tests for existing behavior  
- Small, focused features that match the scope of the CLI (Teams chats / messages / auth)

For larger ideas, open an issue first so maintainers can agree on direction before you invest significant time.

## Development setup

1. **Go** — use the version in [`go.mod`](./go.mod) (currently 1.26.x). [setup-go](https://go.dev/dl/) or your package manager is fine.
2. **Clone** your fork and add this repo as `upstream` if you use one:

   ```bash
   git clone https://github.com/abhiudayg/teams-cli.git
   cd teams-cli
   ```

3. **Build and test:**

   ```bash
   go build -o teams-cli .
   go test ./... -count=1
   ```

   Optional, same as CI locally:

   ```bash
   make fmt
   test -z "$(gofmt -l .)"
   golangci-lint run ./...
   go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...
   ```

   Install [golangci-lint](https://golangci-lint.run/welcome/install/) if you don’t have it.

## Project layout (brief)

| Path | Role |
|------|------|
| `main.go` | Entrypoint |
| `cmd/` | Cobra commands and flags |
| `internal/api/` | Graph and chatsvc clients |
| `internal/auth/` | Token validation, MSAL helpers |
| `internal/config/` | Config load/save and API backend selection |
| `docs/` | Logo and other static docs |

## Pull request guidelines

1. **Branch** — use a short descriptive name, e.g. `fix-doctor-hint` or `docs-readme-install`.
2. **Commits** — small logical commits; message should explain *why* when it isn’t obvious.
3. **Before you push:**
   - `gofmt` all Go files (CI fails on `gofmt -l .`).
   - `go test ./...` passes.
   - Avoid committing tokens, real refresh strings, or personal config paths.
4. **Description** — link related issues, note user-visible behavior changes, and mention how you tested (manual and/or tests).

Maintainers will review as soon as they can. You may be asked to add tests or adjust scope to keep PRs reviewable.

## Reporting security issues

See [**SECURITY.md**](./SECURITY.md). Do not file security problems as public issues.

## Questions

Open a [GitHub issue](https://github.com/abhiudayg/teams-cli/issues) with context and what you’re trying to do. Clear questions get faster answers.
