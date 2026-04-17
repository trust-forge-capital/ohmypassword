# Algorithm Documentation

## Random Number Generation

### Cryptographically Secure RNG

ohmypassword uses Go's `crypto/rand` package for all random number generation.

```go
type RNG interface {
    Intn(n int) (int, error)
    Uint64() (uint64, error)
    Bytes(n int) ([]byte, error)
}
```

### Implementation Details

- Uses OS-provided entropy source (`/dev/urandom` on Unix, `CryptGenRandom` on Windows)
- Never falls back to `math/rand` 
- Thread-safe
- No predictable patterns

### Intn Algorithm

Uses rejection sampling to eliminate modulo bias.

For small n (≤ 2^31-1):
1. Compute rejection limit: `2^32 - (2^32 % n)`
2. Read 4 random bytes and convert to uint32
3. If value ≥ limit, reject and repeat
4. Return `value % n`

For large n (> 2^31-1):
1. Compute rejection limit: `2^64 - (2^64 % n)`
2. Read 8 random bytes and convert to uint64
3. If value ≥ limit, reject and repeat
4. Return `value % n`

## Entropy Calculation

### Formula

```
Entropy (bits) = length × log2(charset_size)
```

### Examples

| Length | Charset Size | Entropy (bits) |
|--------|--------------|----------------|
| 8 | 26 | 37.6 |
| 8 | 36 | 41.4 |
| 8 | 62 | 47.6 |
| 8 | 93 | 52.3 |
| 16 | 93 | 104.6 |
| 32 | 93 | 209.3 |

### Implementation

```go
func CalculateEntropy(password string, charsetSize int) float64 {
    if charsetSize <= 0 || len(password) == 0 {
        return 0
    }
    return float64(len(password)) * math.Log2(float64(charsetSize))
}
```

## Password Strength Detection

### Entropy Levels

| Level | Min Entropy | Interpretation |
|-------|-------------|----------------|
| Very Weak | < 28 bits | Easily cracked |
| Weak | 28-35 bits | Vulnerable |
| Reasonable | 36-59 bits | Acceptable |
| Strong | 60-79 bits | Good |
| Very Strong | ≥ 80 bits | Excellent |

### Crack Time Estimation

Based on attack scenarios:

| Scenario | Attempts/Second |
|----------|-----------------|
| Online (throttled) | 100 |
| Online (unthrottled) | 1,000 |
| Offline (slow hash) | 1,000,000 |
| Offline (fast hash) | 10,000,000,000 |
| Offline (GPU) | 1,000,000,000,000 |

Crack time = 2^entropy / attempts_per_second

## Generation Strategies

### Simple Strategy

1. Build character pool from charset
2. For each position, select random character
3. Return concatenated result

### Pronounceable Strategy

1. Alternates between consonants and vowels
2. Ensures readable pattern: C-V-C-V-C-V...
3. Ends with digit and symbol for complexity
4. Example: `xK9mP2nL5`

### Passphrase Strategy

1. Select random words from word list
2. Join with separator (-)
3. Append digit and symbol
4. Example: `dragon-forest-thunder-42!`

Word list contains 777 common English words for variety.

### Memorable Strategy

1. Generates CVC (consonant-vowel-consonant) patterns
2. Creates human-memorable syllables
3. Joins multiple syllables with separator (-)
4. Appends digit and symbol for complexity
5. Example: `xob-ube-fim2!`

### Segmented Strategy

1. Generates random characters from selected charset
2. Splits result into 3-character segments
3. Joins segments with hyphen separator
4. Length rounds up to nearest multiple of 3
5. Example: `htV-jQ4-A9s-hbY`

## Character Sets

### Predefined Sets

| Name | Characters | Size |
|------|------------|------|
| upper | A-Z | 26 |
| lower | a-z | 26 |
| digit | 0-9 | 10 |
| symbol | !@#$%^&*()_+-=[]{}|;:,.<>?/~`"' | 31 |
| all | All above combined | 93 |

### Similar Character Exclusion

Optional feature to avoid ambiguous characters:
- 0 ↔ O
- 1 ↔ l ↔ I
- | ↔ I