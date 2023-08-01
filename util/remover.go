package util

import "os"

// Function to remove multiple files
func RemoveFiles(paths []string) error {
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

// Function to remove multiple directories
func RemoveDirs(paths []string) error {
	for _, path := range paths {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}
	return nil
}
