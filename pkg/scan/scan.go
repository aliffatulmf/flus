package scan

import (
	"aliffatulmf/flus/pkg/ext"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileMeta struct {
	Root          string
	Path          string
	Info          fs.FileInfo
	FileDirectory string
}

func Files(root string) ([]FileMeta, error) {
	var fileMetas []FileMeta

	root, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path of %s: %v", root, err)
	}

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %v", path, err)
		}

		// Skip if the path is a directory and not the root directory.
		if d.IsDir() && path != root {
			return filepath.SkipDir
		}

		fileInfo, err := getFileInfo(d)
		if err != nil {
			return err
		}

		if ext.IsSupported(fileInfo.Name()) {
			fileDirectory, err := ext.GetFileDirectory(fileInfo.Name())
			if err != nil {
				return err
			}
			fileMetas = append(fileMetas, FileMeta{
				Root:          root,
				Path:          path,
				Info:          fileInfo,
				FileDirectory: fileDirectory,
			})
		}

		return nil
	})

	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error directory %s does not exist", root)
		}

		return nil, fmt.Errorf("error reading directory %s: %v", root, err)
	}
	return fileMetas[1:], nil
}

func getFileInfo(d fs.DirEntry) (fs.FileInfo, error) {
	fi, err := d.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting file info from %s: %v", d.Name(), err)
	}

	return fi, nil
}
