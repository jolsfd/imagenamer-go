package metadata_test

import (
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
