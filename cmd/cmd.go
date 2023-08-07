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

func DoScan(args *args.Args) {
	metas, err := scan.All(args.Target)
	if err != nil {
		logger.Errorf("failed to scan files: %v", err)
	}

	for _, meta := range *metas {
		src := meta.Info.Name()
		dst := meta.Destination

		logger.Infof("transferring file: %s \t", src)
		logger.Infof("to: %s \n", dst)

		if err := processFile(args, meta); err != nil {
			logger.Error(err)
			continue
		}
	}
}

func processFile(args *args.Args, meta scan.Metadata) error {
	dir := filepath.Join(args.Target, meta.Destination)
	if _, err := os.Lstat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, fs.ModeDir); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
	}

	if err := fio.Relocate(dir, meta, args.Buffer); err != nil {
		return err
	}

	if args.Move {
		if err := os.Remove(meta.Source); err != nil {
			logger.Error("failed to remove source file")
		}
	}

	return nil
}
