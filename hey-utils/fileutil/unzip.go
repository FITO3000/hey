package fileutil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(source, dest string) error {
	if Exists(source) {
		if err := PrepareDirectory(dest); err != nil {
			return err
		}
		archive, err := zip.OpenReader(source)
		if err != nil {
			return err
		}
		defer archive.Close()

		for _, f := range archive.File {
			filePath := filepath.Join(dest, f.Name)

			if !strings.HasPrefix(filePath, filepath.Clean(dest)+string(os.PathSeparator)) {
				return fmt.Errorf("invalid file path: %s", filePath)
			}
			if f.FileInfo().IsDir() {
				os.MkdirAll(filePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			fileInArchive, err := f.Open()
			if err != nil {
				return err
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				return err
			}

			dstFile.Close()
			fileInArchive.Close()
		}
		return nil
	} else {
		return fmt.Errorf("file: %s does not exist", source)
	}
}
