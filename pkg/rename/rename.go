// Package rename deals with renaming files and list images in directories
package rename

type FileAttributes struct {
	Path          string
	SourceName    string
	TargetName    string
	FileName      string
	NewFileName   string
	FileExtension string
	CopyNumber    int
}
