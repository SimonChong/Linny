package common

import (
	"errors"
	"io/ioutil"
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

func GetWrappedContent(path string, root string) (string, error) {

	content, err0 := ioutil.ReadFile(path)
	if err0 != nil {
		return "", err0
	}
	header, err1 := ioutil.ReadFile(root + "/header.frag")
	if err1 != nil {
		return "", err1
	}
	footer, err2 := ioutil.ReadFile(root + "/footer.frag")
	if err2 != nil {
		return "", err2
	}

	rtn := string(header) + string(content) + string(footer)

	return rtn, nil
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
