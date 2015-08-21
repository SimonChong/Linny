package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetFile(baseDir string, name string) (string, error) {

	// fmt.Println(Conf.ContentRoot, baseDir, name)

	absBaseDir, e1 := filepath.Abs(baseDir)
	if e1 != nil {
		return "", e1
	}
	absBaseDir = filepath.Clean(absBaseDir)
	absFile, e2 := filepath.Abs(absBaseDir + "/" + name)
	if e2 != nil {
		return "", e2
	}
	absFile = filepath.Clean(absFile)

	if strings.HasPrefix(absFile, absBaseDir) {
		content, err := ioutil.ReadFile(absFile)
		return string(content), err
	}
	return "", errors.New("Invalid Path :" + absFile + " | " + absBaseDir)
}

func GetAdFile(root string, name string) (string, error) {
	adsDir := root + "/ad/"
	fmt.Println("Get Ad File: ", adsDir)
	return GetFile(adsDir, name)
}

func GetResource(root string, name string) (string, error) {
	return GetFile(root, name)
}

func GetWrappedContent(root string, name string) (string, error) {

	content, err0 := GetAdFile(root, name)
	if err0 != nil {
		return "", err0
	}
	header, err1 := GetResource(root, "header.frag")
	if err1 != nil {
		return "", err1
	}
	footer, err2 := GetResource(root, "footer.frag")
	if err2 != nil {
		return "", err2
	}

	rtn := header + content + footer

	return rtn, nil
}
