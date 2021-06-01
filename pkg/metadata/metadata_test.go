package metadata_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/metadata"
)

const pathToImage = "../testdata/test_image.jpg"

func TestLoadImage(t *testing.T) {
	_, err := metadata.GetExif(pathToImage)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCameraModel(t *testing.T) {
	want := "Pixel3a"
	exif, err := metadata.GetExif(pathToImage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*exif)
	got, err := metadata.GetCameraModel(exif)
	if err != nil {
		log.Fatal(err)
	}

	if got != want {
		t.Errorf("Camera model = %s; want = %s;", got, want)
	}
}

func TestGetDateTime(t *testing.T) {
	want := "20200409_220822"
	exif, err := metadata.GetExif(pathToImage)
	if err != nil {
		log.Fatal(err)
	}
	got, err := metadata.GetDateTime(exif)
	if err != nil {
		log.Fatal(err)
	}

	if got != want {
		t.Errorf("Datetime = %s; want = %s;", got, want)
	}
}
