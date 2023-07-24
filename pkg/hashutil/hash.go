package hashutil

import (
	"aliffatulmf/flus/pkg/scan"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type FileHash struct {
	scan.FileMeta
	Checksum []byte
}

func NewFileHash(file scan.FileMeta, checksum []byte) *FileHash {
	return &FileHash{
		FileMeta: file,
		Checksum: checksum,
	}
}

func Hash(file *os.File) ([]byte, error) {
	return hashFile(file)
}

func HashFile(file scan.FileMeta) (*FileHash, error) {
	fo, err := os.Open(file.Path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", file.Info.Name(), err)
	}
	defer fo.Close()

	cs, err := hashFile(fo)
	if err != nil {
		return nil, err
	}

	return &FileHash{
		FileMeta: file,
		Checksum: cs,
	}, nil
}

func Verify(file *os.File, checksum []byte) (bool, error) {
	cs, err := hashFile(file)
	if err != nil {
		return false, err
	}

	return isEqual(cs, checksum), nil
}

func VerifyFile(file *FileHash) (bool, error) {
	fo, err := os.Open(file.Path)
	if err != nil {
		return false, fmt.Errorf("error opening file %s: %v", file.Info.Name(), err)
	}
	defer fo.Close()

	cs, err := hashFile(fo)
	if err != nil {
		return false, err
	}

	return isEqual(cs, file.Checksum), nil
}

func hashFile(f *os.File) ([]byte, error) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, fmt.Errorf("error hashing file %s: %v", f.Name(), err)
	}

	hstr := fmt.Sprintf("%x", h.Sum(nil))
	return []byte(hstr), nil
}

func isEqual(a, b []byte) bool {
	return fmt.Sprintf("%x", a) == fmt.Sprintf("%x", b)
}

func ConvertToHexString(sum []byte) []byte {
	s := fmt.Sprintf("%x", sum)
	return []byte(s)
}
