# ohmypassword

A secure, high-entropy password generator CLI tool.

## Features

- Cryptographically secure random password generation
- Multiple generation strategies (simple, memorable, passphrase)
- Multi-language support (EN, ZH, JA, KO, ES, FR)
- Multiple output formats (simple, JSON, CSV, table)
- Password strength analysis
- Bulk password generation

## Installation

```bash
# Download pre-built binary from releases
# or build from source
make build
```

## Usage

```bash
# Generate a password with default settings (16 chars)
ohmypassword generate

# Generate password with custom length
ohmypassword generate -l 24

# Generate password with specific charset
ohmypassword generate -c upper,lower,digit

# Generate memorable password
ohmypassword generate -s memorable

# Generate passphrase
ohmypassword generate -s passphrase

# Generate multiple passwords
ohmypassword generate -n 10

# Output in JSON format
ohmypassword generate -o json

# Show password strength
ohmypassword generate -v

# Use Chinese language
ohmypassword generate -L zh

# Check password strength
ohmypassword check "myPassword123"

# Check multiple passwords
ohmypassword check "weak" "StrongP@ssw0rd!" "123456"
```

## Commands

- `generate` - Generate passwords (alias: `gen`)
- `check` - Check password strength
- `version` - Show version information
- `help` - Show help information

## Options

- `-l, --length int` - Password length (8-128, default: 16)
- `-c, --charset string` - Character set (upper/lower/digit/symbol/all, default: all)
- `-s, --strategy string` - Generation strategy (simple/pronounceable/passphrase/memorable, default: simple)
- `-n, --count int` - Number of passwords to generate (1-100, default: 1)
- `-v, --validate` - Show password strength
- `-q, --quiet` - Quiet mode (output password only)
- `-L, --lang string` - Language (en/zh/zh-TW/ja/ko/es/fr, default: en)
- `-o, --output string` - Output format (simple/json/csv/table, default: simple)
- `--version` - Show version
- `-h, --help` - Show help

## License

MIT License - see LICENSE file for details.