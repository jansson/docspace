package space

import (
	"os"
	"io/ioutil"
	"log"
)

// filesystem provides simple methods used for accessing the filesystem.
type filesystem interface {
	// Directories returns an array of all directories under the given path, but does not traverse sub folders.
	Directories(path string) []string

	// Files returns an array of all files under the given path, but does not traverse sub folders.
	Files(path string) []string

	// Content returns the actual byte content of the specified file, or nil if the file didn't exist.
	Content(path string) []byte
}

// filesystemFS implements the filesystem interface and provides real disk access.
type filesystemFS struct{}

func (fs filesystemFS) Directories(path string) []string { return []string{} }
func (fs filesystemFS) Files(path string) []string       {	
	dir, err := os.Open(path)
	checkErr(err)
	
	defer dir.Close()
	
	fi, err := dir.Stat()
	checkErr(err)
	
	files := make([]string, 0);
	
	if fi.IsDir() {
		fis, err := ioutil.ReadDir(path)
		checkErr(err)
		
		for _, fileinfo := range fis {
			if !fileinfo.IsDir() {
				files = append(files, fileinfo.Name())
			}
		}
	}
	
	return files
}
func (fs filesystemFS) Content(path string) []byte {
	file, err := ioutil.ReadFile(path)
	checkErr(err)
	return file
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err);
	}
}