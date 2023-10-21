package filefinder

import (
	"testing"
)

func TestFileFinder(t *testing.T) {
	// Test case 1: First path that exists
	paths1 := []string{
		"nonexistent_file1.txt",
		"nonexistent_file2.txt",
		"filefinder.go",
	}

	FileFinder1 := New(paths1)
	firstExistingPath1 := FileFinder1.FirstExists()
	if firstExistingPath1 != "filefinder.go" {
		t.Errorf("Expected 'filefinder.go', but got '%s'", firstExistingPath1)
	}

	// Test case 2: No existing paths
	paths2 := []string{
		"nonexistent_file1.txt",
		"nonexistent_file2.txt",
	}

	FileFinder2 := New(paths2)
	firstExistingPath2 := FileFinder2.FirstExists()
	if firstExistingPath2 != "" {
		t.Errorf("Expected an empty string, but got '%s'", firstExistingPath2)
	}

	// Test case 3: Any path exists
	paths3 := []string{
		"nonexistent_file1.txt",
		"filefinder.go",
	}

	FileFinder3 := New(paths3)
	anyPathExists3 := FileFinder3.AnyExists()
	if !anyPathExists3 {
		t.Errorf("Expected 'true', but got 'false'")
	}

	// Test case 4: No paths exist
	paths4 := []string{
		"nonexistent_file1.txt",
		"nonexistent_file2.txt",
	}

	FileFinder4 := New(paths4)
	anyPathExists4 := FileFinder4.AnyExists()
	if anyPathExists4 {
		t.Errorf("Expected 'false', but got 'true'")
	}
}
