package random

import (
	"testing"
)

func TestCryptoRNG_Intn(t *testing.T) {
	rng := NewCryptoRNG()

	tests := []struct {
		name string
		n    int
	}{
		{"small number", 10},
		{"medium number", 100},
		{"large number", 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result, err := rng.Intn(tt.n)
				if err != nil {
					t.Errorf("Intn() error = %v", err)
					return
				}
				if result < 0 || result >= tt.n {
					t.Errorf("Intn() result = %v, want 0 <= result < %v", result, tt.n)
				}
			}
		})
	}
}

func TestCryptoRNG_Uint64(t *testing.T) {
	rng := NewCryptoRNG()

	for i := 0; i < 100; i++ {
		result, err := rng.Uint64()
		if err != nil {
			t.Errorf("Uint64() error = %v", err)
			return
		}
		if result == 0 {
			t.Log("Uint64() returned 0 (unlikely but possible)")
		}
	}
}

func TestCryptoRNG_Bytes(t *testing.T) {
	rng := NewCryptoRNG()

	sizes := []int{16, 32, 64, 128}
	for _, size := range sizes {
		t.Run("", func(t *testing.T) {
			result, err := rng.Bytes(size)
			if err != nil {
				t.Errorf("Bytes() error = %v", err)
				return
			}
			if len(result) != size {
				t.Errorf("Bytes() length = %v, want %v", len(result), size)
			}
			allZero := true
			for _, b := range result {
				if b != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				t.Error("Bytes() returned all zeros (extremely unlikely)")
			}
		})
	}
}