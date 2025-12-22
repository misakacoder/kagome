package os

import (
	"os"
)

func ExistFile(filepath string) bool {
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func ExistDir(directory string) bool {
	info, err := os.Stat(directory)
	return err == nil && info.IsDir()
}
