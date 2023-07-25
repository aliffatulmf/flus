package move

import (
	"aliffatulmf/flus/hashutil"
	"aliffatulmf/flus/scan"
	"bytes"
	"encoding/hex"
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

	dst = filepath.Join(dst, fm.Info.Name())
	if err := copy(fm.Path, dst, safe); err != nil {
		// Remove destination file if error copying file.
		if err := os.Remove(dst); err != nil {
			return err
		}
		return err
	}
	return nil
}

func copy(src, dst string, safe bool) error {
	fmt.Println("#\tSource\t\t=>", src)
	fmt.Println("#\tDestination\t=>", dst)
	fmt.Println("#\tSafe\t\t=>", safe)

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

	// More efficient way to copy file. See https://stackoverflow.com/a/21067803/1057058
	// Instead of reading the whole file into memory, we read it in chunks.
	// os.ReadFile and os.WriteFile will read the whole file into memory.
	// This is not efficient for large files.
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

	// Safe way to verify copied file instead using os.SameFile().
	rh, err := hashutil.Hash(r)
	if err != nil {
		return err
	}
	fmt.Println("#\tH\t\t=>", hex.EncodeToString(rh))

	wh, err := hashutil.Hash(w)
	if err != nil {
		return err
	}

	if !bytes.Equal(rh, wh) {
		return fmt.Errorf("error copying file %s: checksum mismatch", src)
	}

	fmt.Println("#\tV\t\t=>", hex.EncodeToString(wh), "=> OK")

	return nil
}

func createDir(dir string) error {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %v", dir, err)
	}

	return nil
}
