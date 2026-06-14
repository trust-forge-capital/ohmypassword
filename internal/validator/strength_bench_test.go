package validator

import "testing"

func BenchmarkCalculateStrength(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = CalculateStrength("aB3$kL9@mN2pQ", "all")
	}
}

func BenchmarkGenerateSuggestions(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = generateSuggestions("aB3$kL9@mN2pQ", 90)
	}
}
