package hashutil

import (
	"bytes"
	"errors"
	"hash/crc32"
	"io"
)

var (
	ErrHashFile  = errors.New("error hashing file")
	ErrMissMatch = errors.New("hash mismatch")
)

var table = crc32.MakeTable(crc32.Koopman)

func Hash(f io.ReadSeeker) ([]byte, error) {
	h := crc32.New(table)

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	if _, err := io.Copy(h, f); err != nil {
		return nil, ErrHashFile
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func Verify(a, b []byte) error {
	if !bytes.Equal(a, b) {
		return ErrMissMatch
	}
	return nil
}
