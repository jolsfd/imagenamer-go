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

func TestBuildFileInformation(t *testing.T) {
	want := rename.FileInformation{
		Path:          "../testdata",
		SourceName:    "../testdata/test_image.jpg",
		TargetName:    "",
		FileName:      "test_image",
		NewFileName:   "",
		FileExtension: ".jpg",
		CopyNumber:    2,
	}
	got := rename.FileInformation{}
	got.BuildFileInformation("../testdata/test_image.jpg")

	if want.Path != got.Path {
		t.Errorf("Path = %s, want = %s", want.Path, got.Path)
	}

	if want.SourceName != got.SourceName {
		t.Errorf("SourceName = %s, want = %s", want.SourceName, got.SourceName)
	}

	if want.TargetName != got.TargetName {
		t.Errorf("TargetName = %s, want = %s", want.TargetName, got.TargetName)
	}

	if want.FileName != got.FileName {
		t.Errorf("FileName = %s, want = %s", want.FileName, got.FileName)
	}

	if want.NewFileName != got.NewFileName {
		t.Errorf("NewFileName = %s, want = %s", want.NewFileName, got.NewFileName)
	}

	if want.FileExtension != got.FileExtension {
		t.Errorf("FileExtension = %s, want = %s", want.FileExtension, got.FileExtension)
	}

	if want.CopyNumber != got.CopyNumber {
		t.Errorf("CopyNumber = %d, want = %d", want.CopyNumber, got.CopyNumber)
	}
}
