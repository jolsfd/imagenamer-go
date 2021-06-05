// Package rename deals with renaming files and list images in directories
package rename

import (
	"path/filepath"
	"strings"
)

type FileAttributes struct {
	Path          string
	SourceName    string
	TargetName    string
	FileName      string
	NewFileName   string
	FileExtension string
	CopyNumber    int
}

// BuildFileAttributes takes an source name as string. It assign values into the FileAttributes struct.
func (f *FileAttributes) BuildFileAttributes(sourceName string) {
	// Get Values.
	baseName := filepath.Base(sourceName)
	extension := filepath.Ext(baseName)
	fileName := strings.TrimSuffix(baseName, filepath.Ext(baseName))

	// Assign Values into struct.
	f.Path = filepath.Dir(sourceName)
	f.SourceName = sourceName
	f.FileName = fileName
	f.FileExtension = extension
	f.CopyNumber = DefaultCopyNumber
}
