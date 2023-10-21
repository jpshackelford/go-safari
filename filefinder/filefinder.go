package filefinder

import "os"

type FileFinder struct {
	paths []string
}

// NewFileFinder creates a new FileFinder object with the provided array of strings.
func New(paths []string) *FileFinder {
	return &FileFinder{paths: paths}
}

// firstExists returns the first file path that exists on the file system.
func (ff *FileFinder) FirstExists() string {
	for _, path := range ff.paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path
		}
	}
	return ""
}

// anyExist returns true if any of the file paths exist, and false if none exist.
func (ff *FileFinder) AnyExists() bool {
	for _, path := range ff.paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return true
		}
	}
	return false
}
