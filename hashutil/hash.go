package hashutil

import (
	"errors"
	"hash/crc32"
	"io"
	"sync"
)

var ErrHashFile = errors.New("error hashing file")
var ErrMissMatch = errors.New("hash mismatch")

var table = crc32.MakeTable(crc32.Koopman)

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 64*1024)
	},
}

func Hash(f io.ReadSeeker) ([]byte, error) {
	h := crc32.New(table)
	buf := bufPool.Get().([]byte)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			bufPool.Put(buf)
			return nil, ErrHashFile
		}
		if n == 0 {
			break
		}
		if _, err := h.Write(buf[:n]); err != nil {
			bufPool.Put(buf)
			return nil, ErrHashFile
		}
	}
	bufPool.Put(buf)

	return h.Sum(nil), nil
}

func Verify(a, b []byte) error {
	if crc32.Checksum(a, table) != crc32.Checksum(b, table) {
		return ErrMissMatch
	}
	return nil
}
