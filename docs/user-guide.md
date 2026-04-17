# User Guide

## Installation

### From Binary

Download the latest release for your platform:

```bash
# Linux
wget https://github.com/trust-forge-capital/ohmypassword/releases/latest/download/ohmypassword-linux-amd64
chmod +x ohmypassword-linux-amd64
sudo mv ohmypassword-linux-amd64 /usr/local/bin/ohmypassword

# macOS
curl -LO https://github.com/trust-forge-capital/ohmypassword/releases/latest/download/ohmypassword-darwin-arm64
chmod +x ohmypassword-darwin-arm64
sudo mv ohmypassword-darwin-arm64 /usr/local/bin/ohmypassword

# Windows
# Download .exe from releases and add to PATH
```

### From Source

```bash
# Clone repository
git clone https://github.com/trust-forge-capital/ohmypassword.git
cd ohmypassword

# Build
make build

# Or install
make install
```

## Quick Start

### Basic Usage

```bash
# Generate default password (16 characters)
ohmypassword generate

# Short form
ohmypassword gen
```

### Common Options

```bash
# Custom length
ohmypassword generate -l 24
ohmypassword generate --length 32

# Include specific character types
ohmypassword generate -c upper,lower,digit    # No symbols
ohmypassword generate -c lower,digit          # Only lowercase + numbers

# Generate multiple passwords
ohmypassword generate -n 5
ohmypassword generate --count 10

# Show password strength
ohmypassword generate -V
ohmypassword generate --validate

# Exclude similar characters
ohmypassword generate --exclude-similar
```

### Output Formats

```bash
# Simple (default)
ohmypassword generate -o simple

# JSON (for scripting)
ohmypassword generate -o json

# CSV
ohmypassword generate -o csv

# Table
ohmypassword generate -o table
```

### Language

```bash
# Use Chinese
ohmypassword generate -L zh

# Use Japanese
ohmypassword generate -L ja
```

### Quiet Mode

```bash
# Output password only (useful for scripts)
ohmypassword generate -q
ohmypassword generate --quiet
```

## Generation Strategies

### Simple (Default)

Random characters from selected charset:

```bash
ohmypassword generate -s simple
# Output: aB3$kL9@mN2pQ
```

### Pronounceable

Human-readable passwords:

```bash
ohmypassword generate -s pronounceable
# Output: xK9mP2nL5!
```

### Memorable

CVC-pattern memorable passwords that are easier to recall while maintaining good entropy:

```bash
ohmypassword generate -s memorable
# Output: xob-ube-fim2!
```

### Passphrase

Word-based passwords (specify word count with -l):

```bash
ohmypassword generate -s passphrase -l 4
# Output: dragon-forest-thunder-42!
```

### Segmented

Product-key style passwords with 3-character segments separated by hyphens:

```bash
ohmypassword generate -s segmented
# Output: htV-jQ4-A9s-hbY

# Custom length (rounds up to nearest multiple of 3)
ohmypassword generate -s segmented -l 15
# Output: abc-Def-12g-hIj-3Kl
```

## Use Cases

### Generate Strong Password for Website

```bash
# 24 chars, all character types, show strength
ohmypassword generate -l 24 -c all -V
```

### Generate PIN Code

```bash
# 6-digit PIN
ohmypassword generate -l 6 -c digit
```

### Generate Multiple Passwords for New Accounts

```bash
# 5 passwords in JSON format
ohmypassword generate -n 5 -o json
```

### Use in Shell Script

```bash
#!/bin/bash
PASSWORD=$(ohmypassword generate -l 32 -q)
echo "Generated: $PASSWORD"
```

## Password Strength Guide

### Entropy Levels

| Entropy | Rating | Use Case |
|---------|--------|----------|
| < 28 bits | Very Weak | Not recommended |
| 28-35 bits | Weak | Low-security accounts |
| 36-59 bits | Reasonable | General use |
| 60-79 bits | Strong | Important accounts |
| ≥ 80 bits | Very Strong | Critical accounts |

### Recommendations

- **Minimum**: 16 characters for most uses
- **Banking/Finance**: 24+ characters
- **Password Manager Master**: 32+ characters
- **Passphrase**: 4+ words recommended

## Password Strength Check

### Basic Usage

```bash
# Check single password
ohmypassword check "myPassword123"

# Check multiple passwords
ohmypassword check "weak" "StrongP@ssw0rd!" "123456"
```

### Output Formats

```bash
# Simple output (default)
ohmypassword check "password"

# Table output
ohmypassword check -o table "password"

# JSON output
ohmypassword check -o json "password"

# CSV output
ohmypassword check -o csv "password"
```

### Table Output Example

```
+------------------+-------+-------------+-------------+----------------------+
| PASSWORD         | SCORE | ENTROPY     | STRENGTH    | ESTIMATED CRACK TIME |
+------------------+-------+-------------+-------------+----------------------+
| weak             | 1/5   | 37.60 bits  | Weak        | < 1 year             |
+------------------+-------+-------------+-------------+----------------------+
| StrongP@ssw0rd!  | 5/5   | 127.38 bits | Very Strong | millennia+           |
+------------------+-------+-------------+-------------+----------------------+
```

### With Suggestions

```bash
# Check password and get improvement suggestions
ohmypassword check -o table "short"
```

## Troubleshooting

### "Invalid charset" Error

Make sure charset is valid:
```bash
# Valid: all, upper, lower, digit, symbol, or combinations
ohmypassword generate -c all          # OK
ohmypassword generate -c upper,lower  # OK
ohmypassword generate -c invalid      # Error!
```

### "Invalid length" Error

Length must be between 8 and 128:
```bash
ohmypassword generate -l 8   # OK (minimum)
ohmypassword generate -l 128 # OK (maximum)
ohmypassword generate -l 5   # Error!
```

### UTF-8 Display Issues

Ensure your terminal supports UTF-8:
```bash
# Check locale
echo $LANG

# Set UTF-8 if needed
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8
```

## Security Notes

1. **Clipboard Security**: Be careful when copying passwords to clipboard
2. **Screen Recording**: Avoid generating passwords while screen is recorded
3. **Memory**: Passwords exist only in memory, not written to disk
4. **Offline**: Tool works completely offline, no network required

## Commands

- `generate` (alias: `gen`) - Generate passwords
- `check` - Check password strength
- `version` - Show version information
- `help` - Show help

## Getting Help

```bash
# Show help
ohmypassword --help
ohmypassword generate --help
ohmypassword check --help

# Show version
ohmypassword version
```

## Exit Codes

- `0`: Success
- `1`: Error (invalid arguments or generation failure)