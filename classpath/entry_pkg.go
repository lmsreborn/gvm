package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// PkgEntry entry for package like *.zip or *.jar
type PkgEntry struct {
	absPath string
}

func newPkgEntry(path string) *PkgEntry {
	absPkgPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &PkgEntry{
		absPath: absPkgPath,
	}
}

func (pkgEntry *PkgEntry) readClass(className string) ([]byte, Entry, error) {
	pkg, err := zip.OpenReader(pkgEntry.absPath)
	if err != nil {
		panic(err)
	}

	defer pkg.Close()

	for _, file := range pkg.File {
		if file.Name == className {
			fileRC, err := file.Open()
			if err != nil {
				return nil, nil, err
			}

			defer fileRC.Close()

			data, err := ioutil.ReadAll(fileRC)

			if err != nil {
				return nil, nil, err
			}

			return data, pkgEntry, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (pkgEntry *PkgEntry) String() string {
	return pkgEntry.absPath
}
