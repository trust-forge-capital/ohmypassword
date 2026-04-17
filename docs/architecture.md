# Architecture Design

## Overview

ohmypassword is a CLI password generator that creates cryptographically secure passwords. The project follows Go best practices with clear separation of concerns.

## Architecture Layers

```
┌─────────────────────────────────────────────┐
│              cmd/cli                        │
│         (Entry Point & Commands)            │
└─────────────────────────────────────────────┘
                    │
┌─────────────────────────────────────────────┐
│              internal/                      │
├──────────────┬──────────────┬───────────────┤
│  generator   │    i18n      │      ui       │
│  (Core)      │  (Language)  │   (Output)    │
├──────────────┼──────────────┼───────────────┤
│   strategy   │   validator  │    random     │
│  (Algorithms)│  (Security)  │  (Crypto)     │
└──────────────┴──────────────┴───────────────┘
                    │
┌─────────────────────────────────────────────┐
│              pkg/                           │
│           (Exported Library)                │
└─────────────────────────────────────────────┘
```

## Module Design

### cmd/cli
- **main.go**: Application entry point
- **root.go**: Root command definition with global flags; rewrites bare positional args to default to `generate`
- **generate.go**: Password generation command
- **check.go**: Password strength check command
- **version.go**: Version command

### internal/generator
Core password generation logic:
- **password.go**: Main generator orchestration
- **options.go**: Configuration options validation
- **entropy.go**: Entropy calculation
- **types.go**: Type definitions
- **errors.go**: Shared error definitions

### internal/random
Cryptographically secure random number generation:
- **random.go**: RNG interface
- **crypto.go**: crypto/rand implementation

### internal/strategy
Generation strategies:
- **strategy.go**: Strategy interface and factory
- **simple.go**: Simple random characters
- **pronounceable.go**: Human-readable passwords
- **passphrase.go**: Word-based passwords
- **memorable.go**: CVC-pattern memorable passwords
- **segmented.go**: Hyphen-delimited segment passwords
- **custom.go**: Rule-based generation for programmatic use

### internal/validator
Password strength analysis:
- **validator.go**: Validator interface
- **strength.go**: Strength calculation

### internal/i18n
Internationalization:
- **locale.go**: Language management
- **messages/**: Translation files (en, zh, zh-TW, ja, ko, es, fr)

### internal/ui
Output formatting:
- **output.go**: Main output functions
- **spinner.go**: Progress display

### pkg/charset
Public character set library:
- **charset.go**: Charset constants, rune builders, and validation
- **sets.go**: Predefined charset sets

## Data Flow

```
User Input → CLI Flags → Options → Generator → Strategy → Random → Password
                                          ↓
                                    Entropy Calc
                                          ↓
Password → Validator → Output Formatter → Terminal

Check Input → CLI Flags → Validator → Strength Report → Output Formatter → Terminal
```

## Security Model

1. **Random Source**: Uses `crypto/rand` for cryptographically secure randomness
2. **No Storage**: Passwords never stored, only in memory
3. **No Network**: Fully offline operation
4. **Clear Memory**: No sensitive data logging

## Configuration Priority

1. CLI flags (highest priority)
2. Default values (lowest priority)

No configuration files or environment variables are used.