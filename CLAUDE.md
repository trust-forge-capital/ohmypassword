# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

- `make run` — Run the CLI via `go run ./cmd/ohmypassword`
- `make build` — Build binary to `bin/ohmypassword`
- `make test` — Run all tests with `-race` and coverage
- `go test -race -run TestName ./...` — Run a single test
- `make lint` — Run `golangci-lint` (requires v2.x; config is `.golangci.yml`)
- `make fmt` — Format with `go fmt` and `gofmt -s`
- `make build-all` — Cross-compile for linux/darwin/windows (amd64/arm64)
- `make install` — `go install ./cmd/ohmypassword`

## Architecture

This is a Go CLI built with Cobra. It generates passwords via pluggable strategies and supports multiple output formats and languages.

### Entry Point and Command Routing

- `cmd/ohmypassword/main.go` — Entry point; delegates to `cmd/cli.RootCmd.Execute()`.
- `cmd/cli/root.go` — **Important**: rewrites `os.Args` so bare positional arguments default to the `generate` subcommand. For example, `ohmypassword -l 24` becomes `ohmypassword generate -l 24`. It also parses `--lang`/`-L` from raw `os.Args` before Cobra runs so i18n messages are available during flag initialization.
- `cmd/cli/generate.go` — `generate` command (alias `gen`). Defaults: length 16, charset `all`, strategy `simple`. Passphrase strategy changes the default length to 4 (words) and segmented to 12 (4 segments of 3 chars) when the user does not explicitly set `-l`.
- `cmd/cli/check.go` — `check` command (alias `ck`) for checking password strength.
- **Flags note**: `-v` / `--version` at the root level prints version info (Cobra built-in). `generate --validate` uses short flag `-V` (not `-v`).

### Password Generation Pipeline

1. `generator.GeneratePasswords(opts)` validates options and selects a strategy.
2. `strategy.GetStrategy(name)` returns an implementation of the `Strategy` interface.
3. Strategies live in `internal/strategy/`:
   - `SimpleStrategy` — random characters from selected charset.
   - `PronounceableStrategy` — alternating consonants and vowels.
   - `PassphraseStrategy` — word-based passphrases.
   - `MemorableStrategy` — CVC-pattern memorable passwords.
   - `SegmentedStrategy` — hyphen-delimited 3-character segments.
4. All strategies use `internal/random.CryptoRNG`, which wraps `crypto/rand`. Do not use `math/rand`.
5. `pkg/charset/` provides rune sets and similar-character exclusion.

### Output and Formatting

- `internal/ui/output.go` handles four output formats: `simple`, `json`, `csv`, `table`.
- `quiet` mode suppresses metadata and prints only the password.

### Validation / Strength Analysis

- `internal/validator/validator.go` defines strength levels and display names.
- `internal/validator/strength.go` computes entropy, crack time estimates, scores, and improvement suggestions.
- Strength is calculated from entropy bits and character set size.

### Internationalization

- `internal/i18n/locale.go` manages translations.
- Supported languages: `en`, `zh`, `zh-TW`, `ja`, `ko`, `es`, `fr`.
- Translations are plain maps in `internal/i18n/messages/`.
- `i18n.SetLanguage` is called early in `root.go` by scanning raw `os.Args` for `-L`/`--lang`.

## CI / Release

- `.github/workflows/release.yml` triggers on tags matching `v*`.
- Jobs: `lint` → `test` → `build` (matrix) → `release`.
- The `release` job depends on the others and generates a single GitHub release with all binaries.
- `golangci-lint` v2.x is used both locally and in CI (`golangci/golangci-lint-action@v9` with `version: v2.11`).

## Constraints

- No config files or environment variables — CLI flags only.
- Always use `crypto/rand` (via `internal/random`); never `math/rand`.
- Tests must run with `-race`.
