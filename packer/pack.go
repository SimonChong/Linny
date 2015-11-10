package packer

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"path"
	// "log"
	"os"
	"strings"
)

func Pack(f string) {
	log.Println("Packing: ", f)

	zip := new(AdPackFile)
	zip.Create(path.Base(f))
	zip.AddAll(f, false)
	zip.Close()
}

type packWriterFunc func(info os.FileInfo, file io.Reader, entryName string) (err error)

type AdPackFile struct {
	Writer *zip.Writer
	Name   string
}

// Create new file adPackFile
func (z *AdPackFile) Create(name string) error {
	// check extension .zip
	if strings.HasSuffix(name, ".adpack") != true {
		name = name + ".adpack"
	}
	z.Name = name
	file, err := os.Create(z.Name)
	if err != nil {
		return err
	}
	z.Writer = zip.NewWriter(file)
	return nil
}

// AddAll adds all files from dir in archive, recursively.
// Directories receive a zero-size entry in the archive, with a trailing slash in the header name, and no compression
func (z *AdPackFile) AddAll(dir string, includeCurrentFolder bool) error {
	dir = path.Clean(dir)
	return addAll(dir, dir, includeCurrentFolder, func(info os.FileInfo, file io.Reader, entryName string) (err error) {

		// Create a header based off of the fileinfo
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// If it's a file, set the compression method to deflate (leave directories uncompressed)
		if !info.IsDir() {
			header.Method = zip.Deflate
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// Set the header's name to what we want--it may not include the top folder
		header.Name = entryName

		// Add a trailing slash if the entry is a directory
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		// Get a writer in the archive based on our header
		writer, err := z.Writer.CreateHeader(header)
		if err != nil {
			return err
		}

		// If we have a file to write (i.e., not a directory) then pipe the file into the archive writer
		if file != nil {
			if _, err := io.Copy(writer, file); err != nil {
				return err
			}
		}

		return nil
	})
}

func (z *AdPackFile) Close() error {
	err := z.Writer.Close()
	return err
}

func getSubDir(dir string, rootDir string, includeCurrentFolder bool) (subDir string) {

	subDir = strings.Replace(dir, rootDir, "", 1)

	if includeCurrentFolder {
		parts := strings.Split(rootDir, string(os.PathSeparator))
		subDir = path.Join(parts[len(parts)-1], subDir)
	}

	return
}

// addAll is used to recursively go down through directories and add each file and directory to an archive, based on an packWriterFunc given to it
func addAll(dir string, rootDir string, includeCurrentFolder bool, writerFunc packWriterFunc) error {

	// Get a list of all entries in the directory, as []os.FileInfo
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Loop through all entries
	for _, info := range fileInfos {

		full := path.Join(dir, info.Name())

		// If the entry is a file, get an io.Reader for it
		var file io.Reader
		if !info.IsDir() {
			file, err = os.Open(full)
			if err != nil {
				return err
			}
		}

		// Write the entry into the archive
		subDir := getSubDir(dir, rootDir, includeCurrentFolder)
		entryName := path.Join(subDir, info.Name())
		if err := writerFunc(info, file, entryName); err != nil {
			return err
		}

		// If the entry is a directory, recurse into it
		if info.IsDir() {
			addAll(full, rootDir, includeCurrentFolder, writerFunc)
		}
	}

	return nil
}
