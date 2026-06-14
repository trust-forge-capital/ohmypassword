package generator

import "testing"

func benchGeneratePasswords(b *testing.B, count int) {
	b.Helper()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		opts := &Options{
			Length:   16,
			Charset:  "all",
			Strategy: "simple",
			Count:    count,
		}
		if _, err := GeneratePasswords(opts); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGeneratePasswords_1(b *testing.B)   { benchGeneratePasswords(b, 1) }
func BenchmarkGeneratePasswords_10(b *testing.B)  { benchGeneratePasswords(b, 10) }
func BenchmarkGeneratePasswords_100(b *testing.B) { benchGeneratePasswords(b, 100) }
