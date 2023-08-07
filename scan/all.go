package scan

import (
	"aliffatulmf/flus/file"
	"aliffatulmf/flus/logger"
	"fmt"
	"io/fs"
	"path/filepath"
)

func All(root string) (*[]Metadata, error) {
	var metas []Metadata

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != root {
			return fs.SkipDir
		}

		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

		if !file.IsSupported(info.Name()) {
			// Skip unsupported files.
			return nil
		}

		dir, err := file.FileToDir(info.Name())
		if err != nil {
			return err
		}

		metas = append(metas, Metadata{
			Source:      path,
			Info:        info,
			Destination: dir,
		})

		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return &metas, nil
}
