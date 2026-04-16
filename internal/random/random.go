package random

type RNG interface {
	Intn(n int) (int, error)
	Uint64() (uint64, error)
	Bytes(n int) ([]byte, error)
}