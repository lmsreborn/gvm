package classpath

import "strings"

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(path, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}
