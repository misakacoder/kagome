package file

import (
	"os"
	"path/filepath"
)

func ExistFile(filepath string) bool {
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func ExistDir(directory string) bool {
	info, err := os.Stat(directory)
	return err == nil && info.IsDir()
}

func WriteFile(filename string, data []byte) {
	os.Mkdir(filepath.Dir(filename), os.ModePerm)
	os.WriteFile(filename, data, os.ModePerm)
}
