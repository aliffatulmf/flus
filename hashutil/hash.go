package hashutil

import (
	"errors"
	"hash/crc32"
	"io"
)

var ErrHashFile = errors.New("error hashing file")
var ErrMissMatch = errors.New("hash mismatch")

var table = crc32.MakeTable(crc32.Koopman)

func Hash(f io.Reader) ([]byte, error) {
	h := crc32.New(table)

	if _, err := io.Copy(h, f); err != nil {
		return nil, ErrHashFile
	}
	return h.Sum(nil), nil
}

func Verify(a, b []byte) error {
	if crc32.Checksum(a, table) != crc32.Checksum(b, table) {
		return ErrMissMatch
	}
	return nil
}
