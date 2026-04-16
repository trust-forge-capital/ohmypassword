# AGENTS.md

## Dev Commands

```bash
export GOROOT=/opt/homebrew/Cellar/go/1.26.2/libexec  # Required on this machine

make run        # Run CLI (go run ./cmd/ohmypassword)
make build      # Build to bin/ohmypassword
make test       # Test with -race flag
make lint       # golangci-lint v2.x (auto-installed by action in CI)
make build-all  # Cross-platform builds (linux/darwin/windows, amd64/arm64)
```

## Project Structure

- `cmd/ohmypassword/` - Main entry point (main.go)
- `cmd/cli/` - Cobra commands (root.go, generate.go)
- `internal/generator/` - Core password generation
- `internal/strategy/` - Strategies: simple, pronounceable, passphrase
- `internal/random/` - Crypto random (crypto/rand)
- `internal/validator/` - Password strength analysis
- `internal/i18n/` - Translations (en, zh, zh-TW, ja, ko, es, fr)
- `internal/ui/` - Output formatting (simple/json/csv/table)
- `pkg/charset/` - Character set utilities

## CI/Release

- Release triggers on tag push (`v*`)
- Workflow: `.github/workflows/release.yml` (lint → test → build → upload assets)
- Release notes auto-generated; edit manually to avoid duplicates
- **Important**: Use golangci-lint v2.x (both local and CI)
  - Action: `golangci/golangci-lint-action@v9` with `version: v2.11`
  - Config: `.golangci.yml` uses v2 format (`version: "2"`, `linters.settings`)

## Notes

- No config files or env vars - CLI flags only
- Always use `crypto/rand` (not math/rand)
- Tests require `-race` flag
- Lint config: `.golangci.yml` (v2 format with `version: "2"`, use `linters.settings` not `linters-settings`)
- Build target: `./cmd/ohmypassword` (not `./cmd/cli`)