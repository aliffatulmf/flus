package io

import (
	"aliffatulmf/flus/logger"
	"aliffatulmf/flus/scan"
	"errors"
	"os"
	"path/filepath"
)

var SkipErr = errors.New("skip copy")

func MoveFile(path string, meta scan.Metadata, bufSize uint) error {
	if meta.IsSkipped() {
		return SkipErr
	}

	srcStream, err := os.Open(meta.Source)
	if err != nil {
		return err
	}

	filePath := filepath.Join(path, meta.Info.Name())
	dstStream, err := os.Create(filePath)
	if err != nil {
		return err
	}

	if err := SafeCopy(srcStream, dstStream, bufSize); err != nil {
		return err
	}

	if err := srcStream.Close(); err != nil {
		logger.Error("failed to close source file")
		return err
	}

	if err := dstStream.Close(); err != nil {
		logger.Error("failed to close destination file")
		return err
	}

	if meta.IsMoved() {
		logger.Info("removing source file")
		if err := os.Remove(meta.Source); err != nil {
			logger.Error("failed to remove source file")
		}
	}

	return nil
}
