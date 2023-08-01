package io

import (
	"aliffatulmf/flus/hashutil"
	"io"
)

func SafeCopy(src io.ReadSeeker, dst io.ReadWriteSeeker, bSize uint) error {
	srcHash, err := hashutil.Hash(src)
	if err != nil {
		return err
	}

	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return err
	}

	buff := make([]byte, bSize)
	if _, err := io.CopyBuffer(dst, src, buff); err != nil {
		return err
	}

	if _, err := dst.Seek(0, io.SeekStart); err != nil {
		return err
	}

	dstHash, err := hashutil.Hash(dst)
	if err != nil {
		return err

	}

	if err := hashutil.Verify(srcHash, dstHash); err != nil {
		return err
	}

	return nil
}
