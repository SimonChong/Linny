package common

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ResolveSecure(baseDir string, path string) (string, error) {
	absBaseDir, e1 := filepath.Abs(baseDir)
	if e1 != nil {
		return "", e1
	}
	absBaseDir = filepath.Clean(absBaseDir)
	absFile, e2 := filepath.Abs(absBaseDir + "/" + path)
	if e2 != nil {
		return "", e2
	}
	absFile = filepath.Clean(absFile)
	if strings.HasPrefix(absFile, absBaseDir) {
		return absFile, nil
	}
	return "", errors.New("Invalid Path :" + absFile + " | " + absBaseDir)
}

func FileExistsHTML(path string) (bool, string) {

	if strings.HasSuffix(path, ".html") {
		if _, err := os.Stat(path); err == nil {
			return true, path
		}
		return false, path
	} else if _, err := os.Stat(path + ".html"); err == nil {
		return true, path + ".html"
	}
	return false, path
}
