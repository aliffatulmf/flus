package move

import (
	"aliffatulmf/flus/hashutil"
	"aliffatulmf/flus/scan"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Copy(fm *scan.FileMeta, safe bool) error {
	dst := filepath.Join(fm.Root, fm.FileDirectory)

	_, err := os.Lstat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			if err := createDir(dst); err != nil {
				return err
			}
		}
	}

	if err := copy(fm.Path, filepath.Join(dst, fm.Info.Name()), safe); err != nil {
		return err
	}
	return nil
}

func copy(src, dst string, safe bool) error {
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
		return fmt.Errorf("error copying file %s: %v", src, err)
	}

	if !safe {
		// Unsafe mode, skip checksum verification.
		return nil
	}

	// Seek to the beginning of the file.
	// This is needed because io.Copy() will copy the file from the current offset.
	// If we don't seek to the beginning of the file, the checksum will be empty.
	// This is because io.Copy() will copy the file from the end of the file.
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file %s: %v", src, err)
	}

	_, err = w.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file %s: %v", src, err)
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
