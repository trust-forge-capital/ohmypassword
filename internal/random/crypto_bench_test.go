package random

import "testing"

// BenchmarkCryptoRNG_Password mirrors the per-password usage pattern: a fresh
// RNG draws one index per character. This is where buffering reduces syscalls.
func BenchmarkCryptoRNG_Password(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rng := NewCryptoRNG()
		for j := 0; j < 16; j++ {
			if _, err := rng.Intn(70); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkCryptoRNG_Intn measures a single draw on a reused RNG.
func BenchmarkCryptoRNG_Intn(b *testing.B) {
	b.ReportAllocs()
	rng := NewCryptoRNG()
	for i := 0; i < b.N; i++ {
		if _, err := rng.Intn(70); err != nil {
			b.Fatal(err)
		}
	}
}
