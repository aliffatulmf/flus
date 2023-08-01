package cmd

import (
	"aliffatulmf/flus/io"
	"aliffatulmf/flus/logger"
	"aliffatulmf/flus/scan"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	target   string
	unsafe   bool
	move     bool
	skip     bool
	buffSize uint
)

func init() {
	fl := flag.NewFlagSet("Flus", flag.ExitOnError)
	fl.StringVar(&target, "target", "", "Target directory to process")
	fl.BoolVar(&move, "move", false, "Use move mode instead of copy")
	fl.UintVar(&buffSize, "buffer", 64*1024, "Buffer size to use when copying files")
	fl.Parse(os.Args[1:])
}

func Command() {
	if target == "" {
		logger.Error("target directory is not specified")
		os.Exit(1)
	}

	target, err := filepath.Abs(target)
	if err != nil {
		logger.Errorf("failed to get absolute path: %v", err)
		os.Exit(1)
	}

	metas, err := scan.All(target)
	if err != nil {
		logger.Errorf("failed to scan files: %v", err)
		os.Exit(1)
	}

	for _, meta := range metas {
		src := meta.Info.Name()
		dst := filepath.Join(strings.ToLower(meta.Destination), meta.Info.Name())
		logger.Infof("transferring file: %s => %s", src, dst)

		if err := processFile(target, meta); err != nil {
			logger.Errorf("failed to process file: %v", err)
			continue
		}
	}
}

func processFile(target string, meta scan.Metadata) error {
	dir := filepath.Join(target, meta.Destination)
	if _, err := os.Lstat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, fs.ModeDir); err != nil {
				return fmt.Errorf("failed to create destination directory: %v", err)
			}
		} else {
			return fmt.Errorf("failed to get destination directory info: %v", err)
		}
	}

	if unsafe {
		logger.Infof("marking file as unsafe: %s", meta.Info.Name())
		if err := meta.Unsafe(); err != nil {
			return fmt.Errorf("failed to mark file as unsafe: %v", err)
		}
	}

	if move {
		logger.Infof("marking file as moved: %s", meta.Info.Name())
		if err := meta.Move(); err != nil {
			return fmt.Errorf("failed to mark file as moved: %v", err)
		}
	}

	if err := io.MoveFile(dir, meta, buffSize); err != nil {
		if err == io.SkipErr {
			logger.Warn("file is skipped")
			return nil
		}
		return fmt.Errorf("failed to transfer file: %v", err)
	}

	return nil
}
