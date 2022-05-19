package fileutil

import (
	"io"
	"net/http"
	"os"
	"path"
)

func Download(url, destFolder string) error {
	PrepareDirectory(destFolder)
	destFile := JoinPath(destFolder, path.Base(url))
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err

}
