package scan

import (
	"aliffatulmf/flus/ext"
	"fmt"
	"io/fs"
	"path/filepath"
)

func All(root string) ([]Metadata, error) {
	var metas []Metadata

	root, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	metaCh := make(chan Metadata)
	errCh := make(chan error)

	walkFunc := func(path string, d fs.DirEntry, err error) error {
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

		fext := ext.NewFile(info.Name())
		if !fext.IsSupported() {
			// Skip unsupported files.
			return nil
		}

		dir, err := fext.GetDirName()
		if err != nil {
			return err
		}

		metaCh <- Metadata{
			Source:      path,
			Info:        info,
			Destination: dir,
		}
		return nil
	}

	go func() {
		defer close(metaCh)

		if err := filepath.WalkDir(root, walkFunc); err != nil {
			errCh <- fmt.Errorf("failed to walk directory: %w", err)
		}
	}()

	for meta := range metaCh {
		metas = append(metas, meta)
	}

	select {
	case err := <-errCh:
		return nil, err
	default:
		return metas, nil
	}
}
