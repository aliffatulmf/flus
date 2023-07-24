package pkg

import (
	"aliffatulmf/flus/hashutil"
	"aliffatulmf/flus/scan"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SafeCopy(fm *scan.FileMeta) error {
	dst := filepath.Join(fm.Root, fm.FileDirectory)

	_, err := os.Lstat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			if err := createDir(dst); err != nil {
				return err
			}
		}
	}

	if err := safeCopy(fm.Path, filepath.Join(dst, fm.Info.Name())); err != nil {
		return err
	}
	return nil
}

func safeCopy(src, dst string) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", src, err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", dst, err)
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	h, err := hashutil.Hash(r)
	if err != nil {
		return err
	}

	match, err := hashutil.Verify(w, h)
	if err != nil {
		return err
	}

	if !match {
		return fmt.Errorf("error copying file %s: checksum mismatch", src)
	}

	return nil
}

func createDir(dir string) error {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %v", dir, err)
	}

	return nil
}
