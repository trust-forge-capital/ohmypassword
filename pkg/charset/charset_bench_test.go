package charset

import "testing"

func BenchmarkGetCharsetRunes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetCharsetRunes("upper,lower,digit,symbol")
	}
}

func BenchmarkGetCharsetSize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetCharsetSize("upper,lower,digit,symbol")
	}
}
