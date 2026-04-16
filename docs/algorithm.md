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

For small n (≤ 2^31-1):
1. Read 4 random bytes
2. Convert to uint32
3. Apply modulo: `value % n`

For large n (> 2^31-1):
1. Read 8 random bytes  
2. Convert to uint64
3. Apply modulo: `value % n`

This approach avoids modulo bias.

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
| 8 | 94 | 51.9 |
| 16 | 94 | 103.8 |
| 32 | 94 | 207.6 |

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

Word list contains 50 common English words for variety.

## Character Sets

### Predefined Sets

| Name | Characters | Size |
|------|------------|------|
| upper | A-Z | 26 |
| lower | a-z | 26 |
| digit | 0-9 | 10 |
| symbol | !@#$%^&*()_+-=[]{}|;:,.<>? | 32 |
| all | All above combined | 94 |

### Similar Character Exclusion

Optional feature to avoid ambiguous characters:
- 0 ↔ O
- 1 ↔ l ↔ I
- | ↔ I