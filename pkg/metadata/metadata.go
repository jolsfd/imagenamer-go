// Package metadata handles image exif and gets values from the exif.
package metadata

import (
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

// GetExif requires a source name from an image as an argument and returns an exif type with the metadata of the image and an error.
func GetExif(sourceName string) (imageExif *exif.Exif, err error) {
	// Open image.
	file, err := os.Open(sourceName)
	if err != nil {
		return imageExif, err
	}
	defer file.Close()

	// Decode exif.
	imageExif, err = exif.Decode(file)
	if err != nil {
		return imageExif, err
	}
	return imageExif, err
}

// GetCameraModel returns a string with the camera model of the exif and an error.
func GetCameraModel(imageExif *exif.Exif) (cameraModel string, err error) {
	rawModel, err := imageExif.Get(exif.Model)
	if err != nil {
		return "", err
	}
	cameraModel, err = rawModel.StringVal()
	if err != nil {
		return "", err
	}
	cameraModel = strings.ReplaceAll(cameraModel, " ", "")
	return cameraModel, nil
}
