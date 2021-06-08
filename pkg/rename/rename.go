// Package rename deals with renaming files and list images in directories
package rename

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jolsfd/imagenamer-go/pkg/metadata"
	"github.com/rwcarlsen/goexif/exif"
)

type FileInformation struct {
	Path          string
	SourceName    string
	TargetName    string
	FileName      string
	NewFileName   string
	FileExtension string
	CopyNumber    int
	Status        string
}

// BuildFileAttributes takes an source name as string. It assign values into the FileAttributes struct.
func (f *FileInformation) BuildFileAttributes(sourceName string) {
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
func (f *FileInformation) GetNewFileName(format string, imageExif *exif.Exif) error {
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
func (f *FileInformation) GetTargetName() error {
	f.TargetName = filepath.Join(f.Path, f.NewFileName+f.FileExtension)

	for CheckFileExists(f.Path, f.TargetName) {
		newFileName := f.NewFileName + "~" + strconv.Itoa(f.CopyNumber)
		f.TargetName = filepath.Join(f.Path, newFileName+f.FileExtension)
		f.CopyNumber++
	}

	return nil
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

// CheckFileExists checks if a file is in a given dir.
func CheckFileExists(path string, sourceName string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Join(path, file.Name()) == sourceName {
			return true
		}
	}
	return false
}

// ListImagesInDir search through a directory for files with extensions that match. If a directory is excluded it will skip this directory.
func ListImagesInDir(rootPath string, extensions []string, excludedDirs []string, safeRename bool, safePrefixes []string) (list []string, err error) {
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && find(excludedDirs, filepath.Base(path)) {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			if safeRename {
				if !checkSafePrefix(filepath.Base(path), safePrefixes) && find(extensions, filepath.Ext(path)) {
					list = append(list, path)
				}
			} else {
				if find(extensions, filepath.Ext(path)) {
					list = append(list, path)
				}
			}
		}
		return nil
	})
	return list, err
}

// GetFileInformation returns file informations from source names.
func GetFileInformation(sourceNames []string, format string, debug bool) (files []FileInformation, tableData [][]string, err error) {
	for _, sourceName := range sourceNames {
		// Init FileAttributes struct.
		var image FileInformation
		image.BuildFileAttributes(sourceName)

		// Get image exif.
		imageExif, err := metadata.GetExif(sourceName)
		if err != nil {
			image.Status = color.RedString("fail")
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Build new filename.
		err = image.GetNewFileName(format, imageExif)
		if err != nil {
			image.Status = color.RedString("fail")
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Build new target name.
		err = image.GetTargetName()
		if err != nil {
			image.Status = color.RedString("fail")
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Set image status.
		image.Status = color.GreenString("ok")

		// Append information to table.
		tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})

		// Append image to file sclice.
		files = append(files, image)
	}

	return files, tableData, nil
}

// RenameImages renames images.
func RenameImages(files []FileInformation) error {
	for _, file := range files {
		// Rename image.
		err := os.Rename(file.SourceName, file.TargetName)
		if err != nil {
			log.Print(file.SourceName, color.RedString(err.Error()))
		}
	}

	return nil
}
