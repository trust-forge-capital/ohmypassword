# AGENTS.md

## Dev Commands

```bash
make run        # Run CLI locally
make build      # Build binary to bin/ohmypassword
make test       # Run tests with race detector
make lint       # Run golangci-lint
make fmt        # Format code
make build-all  # Build for all platforms
```

## Project Structure

- `cmd/cli/` - CLI entry point (main.go, root.go, generate.go)
- `internal/generator/` - Core password generation
- `internal/strategy/` - Generation strategies (simple, pronounceable, passphrase)
- `internal/random/` - Crypto random (crypto/rand)
- `internal/validator/` - Password strength analysis
- `internal/i18n/` - Translations (en, zh, zh-TW, ja, ko, es, fr)
- `internal/ui/` - Output formatting (simple, json, csv, table)
- `pkg/charset/` - Exported character set library

## Notes

- No config files or env vars - CLI flags only
- Always use `crypto/rand` (not math/rand)
- Run tests with `-race` flag (enabled in `make test`)
- Lint config in `.golangci.yml`