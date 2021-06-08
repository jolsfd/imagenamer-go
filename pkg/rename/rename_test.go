package rename_test

import (
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/rename"
)

func TestCheckFileExists(t *testing.T) {
	path := "../testdata"
	fileExist := "../testdata/test_image.jpg"
	fileNotExist := "../testdata/NO_image.jpg"

	if !rename.CheckFileExists(path, fileExist) {
		t.Error("FileExist should exist")
	}

	if rename.CheckFileExists(path, fileNotExist) {
		t.Errorf("FileExist should not exist")
	}
}
