package move

import (
	"aliffatulmf/flus/hashutil"
	"aliffatulmf/flus/scan"
	"aliffatulmf/flus/util"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copy(src, dst string, fn func(r io.ReadSeekCloser, w io.ReadWriteSeeker) error) error {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer w.Close()

	// More efficient way to copy file. See https://stackoverflow.com/a/21067803/1057058
	// Instead of reading the whole file into memory, we read it in chunks.
	// os.ReadFile and os.WriteFile will read the whole file into memory.
	// This is not efficient for large files.
	_, err = io.Copy(w, r)
	if err != nil {
		if err := os.Remove(dst); err != nil {
			return fmt.Errorf("error copying and removing file: %v", err)
		}
		return fmt.Errorf("error copying file: %v", err)
	}

	if err := fn(r, w); err != nil {
		return err
	}

	return nil
}

func Copy(fm *scan.FileMeta, move, safe bool) error {
	var rh, wh []byte

	dst := filepath.Join(fm.Root, fm.FileDirectory)
	if err := os.MkdirAll(dst, os.ModeDir); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	dst = filepath.Join(dst, fm.Info.Name())
	err := copy(fm.Path, dst, func(r io.ReadSeekCloser, w io.ReadWriteSeeker) error {
		if !safe {
			// Unsafe mode, skip checksum verification.
			return nil
		}

		// Seek to the beginning of the file.
		// This is needed because io.Copy() will copy the file from the current offset.
		// If we don't seek to the beginning of the file, the checksum will be empty.
		// This is because io.Copy() will copy the file from the end of the file.
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %v", err)
		}

		_, err = w.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %v", err)
		}

		rh, err = hashutil.Hash(r)
		if err != nil {
			return err
		}

		wh, err = hashutil.Hash(w)
		if err != nil {
			return err
		}

		if err := hashutil.Verify(rh, wh); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("#\tSource\t\t=>", util.TrimText(fm.Path))
	fmt.Println("#\tDestination\t=>", util.TrimText(dst))
	fmt.Println("#\tSafe\t\t=>", safe)
	if safe {
		fmt.Println("#\tR\t\t=> (CRC32)", hex.EncodeToString(rh))
		fmt.Println("#\tW\t\t=> (CRC32)", hex.EncodeToString(wh), "=> OK")
	}

	if move {
		if err := os.Remove(fm.Path); err != nil {
			panic(err)
		}
	}
	return nil
}
