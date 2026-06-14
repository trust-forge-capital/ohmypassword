package random

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
)

type CryptoRNG struct {
	reader io.Reader
}

// NewCryptoRNG returns a CryptoRNG backed by crypto/rand. The reader is wrapped
// in a buffer so that generating a password issues a single read from the OS
// CSPRNG instead of one read per character; the random bytes served are still
// produced by crypto/rand, so the cryptographic guarantees are unchanged.
func NewCryptoRNG() *CryptoRNG {
	return &CryptoRNG{
		reader: bufio.NewReaderSize(rand.Reader, 256),
	}
}

func (r *CryptoRNG) Intn(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid argument to Intn")
	}

	if n <= 1<<31-1 {
		return r.intn31(n)
	}

	return r.intnLarge(n)
}

func (r *CryptoRNG) intn31(n int) (int, error) {
	nn := uint32(n)
	max := uint64(1) << 32
	limit := max - (max % uint64(nn))

	var b [4]byte
	for {
		_, err := io.ReadFull(r.reader, b[:])
		if err != nil {
			return 0, err
		}

		v := uint64(binary.BigEndian.Uint32(b[:]))
		if v >= limit {
			continue
		}
		return int(v % uint64(nn)), nil
	}
}

func (r *CryptoRNG) intnLarge(n int) (int, error) {
	nn := uint64(n)
	limit := ^uint64(0) - (^uint64(0) % nn)

	var b [8]byte
	for {
		_, err := io.ReadFull(r.reader, b[:])
		if err != nil {
			return 0, err
		}

		v := binary.BigEndian.Uint64(b[:])
		if v >= limit {
			continue
		}
		return int(v % nn), nil
	}
}

func (r *CryptoRNG) Uint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r.reader, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[:]), nil
}

func (r *CryptoRNG) Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(r.reader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
