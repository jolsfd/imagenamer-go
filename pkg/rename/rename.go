// Package rename deals with renaming files and list images in directories
package rename

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jolsfd/imagenamer-go/pkg/metadata"
	"github.com/rwcarlsen/goexif/exif"
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

// GetNewFileName assign the new filename into a FileAttributes struct.
func (f *FileAttributes) GetNewFileName(format string, imageExif *exif.Exif) error {
	dateTime, err := metadata.GetDateTime(imageExif)
	if err != nil {
		return err
	}
	cameraModel, err := metadata.GetCameraModel(imageExif)
	if err != nil {
		return err
	}
	format = strings.ReplaceAll(format, "DATETIME", dateTime)
	format = strings.ReplaceAll(format, "MODEL", cameraModel)

	f.NewFileName = format
	return nil
}

// GetTargetName assign the target name into a FileAttributes struct.
func (f *FileAttributes) GetTargetName() {
	f.TargetName = filepath.Join(f.Path, f.NewFileName+f.FileExtension)

	for CheckFileExists(f.Path, f.TargetName) {
		newFileName := f.NewFileName + "~" + strconv.Itoa(f.CopyNumber)
		f.TargetName = filepath.Join(f.Path, newFileName+f.FileExtension)
		f.CopyNumber++
	}
}

// find checks if a value is in a sclice and return true or false.
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// checkSafePrefix checks if a file name contains a safe prefix.
func checkSafePrefix(baseName string, list []string) bool {
	for _, safePrefix := range list {
		// Check lenght
		if len(baseName) < len(safePrefix) {
			continue
		} else if baseName[:len(safePrefix)] == safePrefix {
			return true
		}
	}
	return false
}
