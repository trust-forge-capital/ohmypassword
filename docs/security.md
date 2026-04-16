# Security Design

## Overview

ohmypassword is designed with security as the primary concern. This document outlines the security architecture and measures.

## Cryptographic Requirements

### Random Number Generation

- **Source**: Go's `crypto/rand` package
- **Quality**: Cryptographically secure, suitable for security-sensitive applications
- **Implementation**: OS-provided entropy source (e.g., `/dev/urandom` on Unix, CryptGenRandom on Windows)

```go
// Using crypto/rand
func (r *CryptoRNG) Intn(n int) (int, error) {
    // Uses crypto/rand.Reader internally
    // Never falls back to math/rand
}
```

### Character Selection

- Each character is selected independently using fresh random bytes
- No bias or predictable patterns
- Full entropy from the random source

## Entropy Analysis

### Entropy Calculation

Entropy (in bits) = length × log2(charset_size)

| Charset | Size | 8 chars | 16 chars | 32 chars |
|---------|------|---------|----------|----------|
| lower (26) | 26 | 37.6 | 75.2 | 150.4 |
| lower+digit (36) | 36 | 41.4 | 82.7 | 165.5 |
| all (94) | 94 | 51.9 | 103.8 | 207.6 |

### Entropy Thresholds

| Level | Min Entropy | Description |
|-------|-------------|-------------|
| Very Weak | < 28 bits | Easily cracked |
| Weak | 28-35 bits | Vulnerable to attack |
| Reasonable | 36-59 bits | Acceptable for low-risk |
| Strong | 60-79 bits | Good for most uses |
| Very Strong | ≥ 80 bits | Excellent security |

## Threat Model

### Addressed Threats

1. **Brute Force Attacks**
   - Mitigation: High entropy passwords (≥60 bits recommended)
   - 80+ bits provides protection against modern GPU clusters

2. **Dictionary Attacks**
   - Mitigation: Random character selection, no dictionary words (except passphrase mode)
   - Passphrase mode uses large word list with separator

3. **Rainbow Table Attacks**
   - Mitigation: Random salt through unique random selection per generation
   - No deterministic password generation

4. **Social Engineering**
   - Mitigation: No personal information in passwords
   - Completely random generation

### Not Addressed

- Physical security of the machine
- Keyloggers or screen capture
- Memory scraping attacks
- Clipboard security (user responsibility)

## Memory Security

- Passwords exist only in memory during generation
- No logging of generated passwords
- No persistence to disk
- Garbage collection handles memory cleanup

## Best Practices

### Password Length

- **Minimum**: 8 characters
- **Recommended**: 16+ characters
- **High Security**: 24+ characters
- **Maximum**: 128 characters

### Character Sets

- Use `all` charset for maximum entropy
- Avoid reducing character set without good reason
- Consider password manager storage when using complex passwords

### Generation Strategy

- **simple**: Maximum entropy, hard to remember
- **pronounceable**: Balance of security and memorability
- **passphrase**: Best memorability, acceptable security with 4+ words

## Compliance

Generated passwords meet requirements for:
- PCI DSS (payment cards)
- NIST SP 800-63B (US government)
- OWASP recommendations
- Most corporate security policies

## Security Auditing

- Code uses only standard library crypto packages
- No custom cryptographic implementations
- Entropy calculations are transparent and verifiable