# AGENTS.md

## Dev Commands

```bash
export GOROOT=/opt/homebrew/Cellar/go/1.26.2/libexec  # Required on this machine

make run        # Run CLI (go run ./cmd/ohmypassword)
make build      # Build to bin/ohmypassword
make test       # Test with -race flag
make lint       # golangci-lint (use v1.64.8: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /tmp v1.64.8)
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
- golangci-lint v1.64.8 in CI, v2.x locally - set `verify: false` in workflow

## Notes

- No config files or env vars - CLI flags only
- Always use `crypto/rand` (not math/rand)
- Tests require `-race` flag
- Lint config: `.golangci.yml` (v1 format, no `version` field)
- Build target: `./cmd/ohmypassword` (not `./cmd/cli`)