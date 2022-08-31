package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, errors.Wrapf(err, "failed to stat path: %s", path)
	}

	return fileInfo.IsDir(), nil
}

func GetFilePathsInDir(dirPath string) ([]string, error) {
	isDir, err := IsDirectory(dirPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to check whether given path is directory: %s", dirPath)
	}
	if !isDir {
		return nil, errors.Errorf("dirPath is not directory: %s", dirPath)
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to readdir: %s", dirPath)
	}

	filePaths := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePaths = append(filePaths, filepath.Join(dirPath, file.Name()))
	}

	return filePaths, nil
}
