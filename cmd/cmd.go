package cmd

import (
	"aliffatulmf/flus/args"
	fio "aliffatulmf/flus/io"
	"aliffatulmf/flus/logger"
	"aliffatulmf/flus/scan"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func DoScan(av *args.Value) {
	metas, err := scan.All(av.Target)
	if err != nil {
		logger.Errorf("failed to scan files: %v", err)
	}

	for _, meta := range *metas {
		src := meta.Info.Name()
		dst := meta.Destination

		logger.Infof("transferring file: %s \t", src)
		logger.Infof("to: %s \n", dst)

		if err := processFile(av, meta); err != nil {
			logger.Error(err)
			continue
		}
	}
}

func processFile(args *args.Value, meta scan.Metadata) error {
	dir := filepath.Join(args.Target, meta.Destination)
	if _, err := os.Lstat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, fs.ModeDir); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
	}

	if err := fio.Relocate(dir, meta); err != nil {
		return err
	}

	if args.Move {
		if err := os.Remove(meta.Source); err != nil {
			logger.Error("failed to remove source file")
		}
	}

	return nil
}
