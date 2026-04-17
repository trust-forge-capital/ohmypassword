# ohmypassword

## 1.1.0 (2026-04-17)

### Added
- `memorable` strategy: CVC-pattern memorable passwords
- `segmented` strategy: hyphen-delimited 3-character segments
- `check` command (alias `ck`): password strength analysis with entropy, crack time, and suggestions
- `-v` / `--version` short flag at root level
- `-q` / `--quiet` mode for script-friendly password-only output
- `--exclude-similar` flag to remove visually ambiguous characters
- Multi-language support: EN, ZH, ZH-TW, JA, KO, ES, FR
- Multiple output formats: JSON, CSV, table (structured formats always include strength data)
- Bulk generation via `-n` / `--count`
- Auto-validation for `json`/`csv`/`table` outputs (no `-V` required)
- `CLAUDE.md` for Claude Code guidance

### Fixed
- RNG rejection sampling to eliminate modulo bias
- Root argument rewriting for bare flags defaulting to `generate`
- Locale panic on unsupported languages
- Duplicate `-` in symbol charset (corrected unique symbol count to 31)
- Charset validation now accepts all valid combinations (e.g., `lower,digit`)
- Spinner race condition on concurrent `Start()` calls

## 1.0.0 (2024-XX-XX)

### Added
- Initial release
- Cryptographically secure password generation
- Multiple generation strategies (`simple`, `pronounceable`, `passphrase`)
- Multi-language support
- Multiple output formats
