package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry entry for the common directory path
type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	absDirPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{
		absPath: absDirPath,
	}
}

func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(dirEntry.absPath, className)
	data, err := ioutil.ReadFile(fileName)
	return data, dirEntry, err
}

func (dirEntry *DirEntry) String() string {
	return dirEntry.absPath
}
