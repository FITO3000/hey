package fileutil

import (
	"fmt"
	"os"
	"path"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func IsDirectory(filename string) bool {
	stat, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return stat.IsDir()
}

func JoinPath(elem ...string) string {
	return path.Join(elem...)
}

func PrepareDirectory(path string) error {
	if Exists(path) {
		if !IsDirectory(path) {
			return fmt.Errorf("%s is not a directory", path)
		}
		return nil
	} else {
		return os.MkdirAll(path, os.ModePerm)
	}
}
