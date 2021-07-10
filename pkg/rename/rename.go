// Package rename deals with renaming files and list images in directories
package rename

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/jolsfd/imagenamer-go/pkg/metadata"
	"github.com/jolsfd/imagenamer-go/pkg/validation"
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

type FileNameTemplate struct {
	CameraModel string
	DateTime    string
}

// BuildFileAttributes takes an source name as string. It assign values into the FileAttributes struct.
func (f *FileInformation) BuildFileInformation(sourceName string) {
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
func (f *FileInformation) GetNewFileName(templateString string, imageExif *exif.Exif) error {
	var output bytes.Buffer

	// Get exif values.
	dateTime, err := metadata.GetDateTime(imageExif)
	if err != nil {
		return err
	}

	cameraModel, err := metadata.GetCameraModel(imageExif)
	if err != nil {
		return err
	}

	// Create template.
	templ, err := template.New("").Parse(templateString)
	if err != nil {
		return err
	}

	// Execute template.
	err = templ.Execute(&output, FileNameTemplate{cameraModel, dateTime})
	if err != nil {
		return err
	}

	// Assign value.
	f.NewFileName = output.String()

	// Return error
	return nil
}

// GetTargetName assign the target name into a FileAttributes struct.
func (f *FileInformation) GetTargetName(targetNames []string) error {
	f.TargetName = filepath.Join(f.Path, f.NewFileName+f.FileExtension)

	// Check if source name equals target name.
	if f.SourceName == f.TargetName {
		return nil
	}

	// Check if file exists in dir or in list.
	for CheckFileExists(f.Path, f.TargetName) || find(targetNames, f.TargetName) {
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
func checkSafeStrings(baseName string, safeStrings []string) bool {
	for _, safeString := range safeStrings {
		match, _ := regexp.MatchString(safeString, baseName)
		if match {
			return match
		} else {
			continue
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
func ListImagesInDir(rootPath string, extensions []string, excludedDirs []string, safeRename bool, safeStrings []string) (list []string, err error) {
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && find(excludedDirs, filepath.Base(path)) {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			if safeRename {
				if !checkSafeStrings(filepath.Base(path), safeStrings) && find(extensions, filepath.Ext(path)) {
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
func GetFileInformation(sourceNames []string, templateString string, debug bool) (files []FileInformation, tableData [][]string, err error) {
	var targetNames []string

	for _, sourceName := range sourceNames {
		// Init FileAttributes struct.
		var image FileInformation
		image.BuildFileInformation(sourceName)

		// Get image exif.
		imageExif, err := metadata.GetExif(sourceName)
		if err != nil {
			image.Status = color.RedString(StatusFail)
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Build new filename.
		err = image.GetNewFileName(templateString, imageExif)
		if err != nil {
			image.Status = color.RedString(StatusFail)
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Build new target name.
		err = image.GetTargetName(targetNames)
		if err != nil {
			image.Status = color.RedString(StatusFail)
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Check forbidden characters.
		err = validation.CheckForbiddenCharacters(image.NewFileName)
		if err != nil {
			image.Status = color.RedString(StatusFail)
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Check length.
		err = validation.CheckTargetLength(image.TargetName)
		if err != nil {
			image.Status = color.RedString(StatusFail)
			tableData = append(tableData, []string{image.FileName + image.FileExtension, image.NewFileName + image.FileExtension, image.Status})
			if debug {
				log.Print(color.RedString(err.Error()))
			}
			continue
		}

		// Set image status.
		image.Status = color.GreenString(StatusOk)

		// Append information to table.
		tableData = append(tableData, []string{filepath.Base(image.SourceName), filepath.Base(image.TargetName), image.Status})

		// Append image to file sclice.
		files = append(files, image)

		// Append target name.
		targetNames = append(targetNames, image.TargetName)
	}

	return files, tableData, nil
}

// RenameImages renames images.
func RenameImages(files []FileInformation) error {
	for _, file := range files {
		// Check if file exists.
		if CheckFileExists(file.Path, file.TargetName) {
			log.Print(file.SourceName, color.RedString(FileExistsError))
		}

		// Rename image.
		err := os.Rename(file.SourceName, file.TargetName)
		if err != nil {
			log.Print(file.SourceName, color.RedString(err.Error()))
		}
	}

	return nil
}
