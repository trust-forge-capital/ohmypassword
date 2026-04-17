# ohmypassword

## 1.3.1 (2026-04-17)

### Fixed
- Charset validation now accepts all valid combinations (e.g., `lower,digit`) by delegating to `pkg/charset.IsValidCharset`
- Removed duplicate `-` from `CharsetSymbol`; corrected unique symbol count to 31 and `all` charset to 93
- Fixed spinner race condition by moving `spinnerIndex` into the `Spinner` struct
- Fixed integration test that mistakenly used `-v` instead of `-V` for validate
- Added comprehensive unit tests for `internal/ui/output.go` (81% coverage)
- Updated `docs/api-spec.md` with missing `check` command and `custom` strategy
- Updated `docs/architecture.md` with `check` command flow, `custom` strategy, and `pkg/charset` details

## 1.3.0 (2026-04-17)

### Added
- `segmented` strategy: hyphen-delimited 3-character segments
- Auto-validation for `json`/`csv`/`table` outputs (no `-V` required)
- `CLAUDE.md` for Claude Code guidance

### Fixed
- RNG rejection sampling to eliminate modulo bias
- Root argument rewriting for bare flags defaulting to `generate`
- Locale panic on unsupported languages

## 1.2.0 (2026-04-17)

### Added
- `memorable` strategy: CVC-pattern memorable passwords
- `-q` / `--quiet` mode for script-friendly password-only output
- `--exclude-similar` flag to remove visually ambiguous characters
- Bulk generation via `-n` / `--count`

## 1.1.0 (2026-04-17)

### Added
- `check` command (alias `ck`): password strength analysis with entropy, crack time, and suggestions
- `-v` / `--version` short flag at root level
- Multi-language support: EN, ZH, ZH-TW, JA, KO, ES, FR
- Multiple output formats: JSON, CSV, table (structured formats always include strength data)

## 1.0.0 (2024-XX-XX)

### Added
- Initial release
- Cryptographically secure password generation
- Multiple generation strategies (`simple`, `pronounceable`, `passphrase`)
- Multi-language support
- Multiple output formats
