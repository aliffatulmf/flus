package io

import (
	"aliffatulmf/flus/scan"
	"os"
	"path/filepath"
)

func Relocate(path string, meta scan.Metadata) error {
	srcStream, err := os.Open(meta.Source)
	if err != nil {
		return err
	}
	defer srcStream.Close()

	filePath := filepath.Join(path, meta.Info.Name())
	dstStream, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dstStream.Close()

	if err := CopyAndVerify(srcStream, dstStream); err != nil {
		return err
	}

	return nil
}
