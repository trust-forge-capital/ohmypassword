package random

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
)

type CryptoRNG struct {
	reader io.Reader
}

func NewCryptoRNG() *CryptoRNG {
	return &CryptoRNG{
		reader: rand.Reader,
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
	var b [4]byte
	_, err := io.ReadFull(r.reader, b[:])
	if err != nil {
		return 0, err
	}

	v := binary.BigEndian.Uint32(b[:])
	v = v % uint32(n)
	return int(v), nil
}

func (r *CryptoRNG) intnLarge(n int) (int, error) {
	var b [8]byte
	_, err := io.ReadFull(r.reader, b[:])
	if err != nil {
		return 0, err
	}

	v := binary.BigEndian.Uint64(b[:])
	v = v % uint64(n)
	return int(v), nil
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