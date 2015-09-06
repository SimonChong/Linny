package packer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/simonchong/linny/common"
)

func Unpack(conf common.ConfigLinny, adPackFile string) bool {
	target := adPackFile
	if !strings.HasSuffix(target, ".adpack") {
		target += ".adpack"
	}
	target, err := filepath.Abs(target)
	if err != nil {
		if _, err2 := os.Stat(target); os.IsNotExist(err2) {
			fmt.Println("Invalid adpack file")
		}
		return false
	}
	dest := strings.TrimSuffix(target, ".adpack")
	fmt.Println(target, dest)

	err2 := unzip(target, dest)
	if err2 != nil {
		fmt.Println("Unzip Error:", err2)
		return false
	}
	confP := &conf
	confP.ContentRoot = dest
	conf.Save()
	return true
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extract := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extract(f)
		if err != nil {
			return err
		}
	}

	return nil
}
