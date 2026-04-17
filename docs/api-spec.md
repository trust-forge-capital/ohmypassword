# API Specification

## Command Structure

```
ohmypassword [global options] <command> [subcommand] [options]
```

## Global Options

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--lang` | `-L` | string | `en` | Language (en/zh/zh-TW/ja/ko/es/fr) |
| `--output` | `-o` | string | `simple` | Output format (simple/json/csv/table) |
| `--version` | | bool | false | Show version |
| `--help` | `-h` | bool | false | Show help |

## Commands

### generate (gen)

Generate a new password.

```
ohmypassword generate [options]
ohmypassword gen [options]
```

#### Options

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--length` | `-l` | int | 16 | Password length (8-128) |
| `--charset` | `-c` | string | `all` | Character set |
| `--strategy` | `-s` | string | `simple` | Generation strategy |
| `--count` | `-n` | int | 1 | Number of passwords |
| `--validate` | `-V` | bool | false | Show password strength |
| `--quiet` | `-q` | bool | false | Quiet mode |
| `--exclude-similar` | | bool | false | Exclude similar characters (0, O, 1, l, I, \|) |

#### Character Set Options

- `upper` - Uppercase letters (A-Z)
- `lower` - Lowercase letters (a-z)
- `digit` - Numbers (0-9)
- `symbol` - Symbols (!@#$%^&*...)
- `all` - All character types (default)

Combined example: `upper,lower,digit`

#### Strategy Options

- `simple` - Random character selection (default)
- `pronounceable` - Human-readable passwords (e.g., xK9mP2)
- `passphrase` - Word-based passwords (e.g., dragon-forest-thunder)
- `memorable` - CVC-pattern memorable passwords
- `segmented` - Hyphen-delimited 3-character segments (e.g., htV-jQ4-A9s-hbY)

### version

Show version information.

```
ohmypassword version
```

### help

Show help information.

```
ohmypassword help
ohmypassword generate --help
```

## Output Formats

### simple (default)

```
aB3$kL9@mN2pQ
```

With validation (`-V`):
```
Password: aB3$kL9@mN2pQ
  Entropy: 95.27 bits
  Strength: Strong
  Crack Time: centuries
```

### json

Default:
```json
[
  {
    "password": "aB3$kL9@mN2pQ"
  }
]
```

With validation (`-v`):
```json
[
  {
    "password": "aB3$kL9@mN2pQ",
    "entropy": 95.27,
    "strength": {
      "level": "Strong",
      "crack_time": "centuries",
      "score": 5
    }
  }
]
```

### csv

Default:
```
password
aB3$kL9@mN2pQ
```

With validation (`-v`):
```
password,entropy,strength,crack_time
aB3$kL9@mN2pQ,95.27,Strong,centuries
```

### table

Default:
```
aB3$kL9@mN2pQ
```

With validation (`-v`):
```
PASSWORD                     ENTROPY         STRENGTH     CRACK_TIME
-----------------------------------------------------------------------------
aB3$kL9@mN2pQ               95.27 bits       Strong       centuries
```

## Quiet Mode

In quiet mode (`-q`), only the password is output (one per line for multiple passwords), same as default mode without validation.

```
$ ohmypassword generate -q
aB3$kL9@mN2pQ

$ ohmypassword generate -n 3 -q
xK9mP2nL5
qR8wT3yZ1
bJ7vN4cX6
```

## Exit Codes

- `0` - Success
- `1` - Error (invalid arguments, generation failure)

## Examples

```bash
# Generate 16-character password
ohmypassword generate

# Generate 24-character password
ohmypassword generate -l 24

# Generate password with only lowercase and numbers
ohmypassword generate -c lower,digit

# Generate 5 passwords at once
ohmypassword generate -n 5

# Generate passphrase
ohmypassword generate -s passphrase -l 4

# Generate segmented password
ohmypassword generate -s segmented

# Show password strength
ohmypassword generate -V

# Output as JSON
ohmypassword generate -o json

# Use Chinese language
ohmypassword generate -L zh

# All options combined
ohmypassword generate -l 32 -c all -s simple -n 10 -v -o json
```