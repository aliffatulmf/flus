package io

import (
	"aliffatulmf/flus/hashutil"
	"io"
)

func CopyAndVerify(src io.ReadSeeker, dst io.ReadWriteSeeker, buffSize uint) error {
	srcHash, err := hashutil.Hash(src)
	if err != nil {
		return err
	}

	buff := make([]byte, buffSize)
	if _, err := io.CopyBuffer(dst, src, buff); err != nil {
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
