package rename_test

import (
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/metadata"
	"github.com/jolsfd/imagenamer-go/pkg/rename"
)

const sourceName = "../testdata/test_image.jpg"

func TestCheckFileExists(t *testing.T) {
	path := "../testdata"
	fileNotExist := "../testdata/NO_image.jpg"

	if !rename.CheckFileExists(path, sourceName) {
		t.Error("FileExist should exist")
	}

	if rename.CheckFileExists(path, fileNotExist) {
		t.Errorf("FileExist should not exist")
	}
}

func TestBuildFileInformation(t *testing.T) {
	want := rename.FileInformation{
		Path:          "../testdata",
		SourceName:    sourceName,
		TargetName:    "",
		FileName:      "test_image",
		NewFileName:   "",
		FileExtension: ".jpg",
		CopyNumber:    2,
	}
	got := rename.FileInformation{}
	got.BuildFileInformation(sourceName)

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

func TestGetValues(t *testing.T) {
	var image rename.FileInformation
	format := "IMG_DATETIME_MODEL"
	exif, err := metadata.GetExif(sourceName)
	if err != nil {
		t.Error(err)
	}
	image.BuildFileInformation(sourceName)

	wantNewFilename := "IMG_20200409_220822_Pixel3a"
	err = image.GetNewFileName(format, exif)
	if err != nil {
		t.Error(err)
	}

	if wantNewFilename != image.NewFileName {
		t.Errorf("NewFilename = %s, want = %s", image.NewFileName, wantNewFilename)
	}

	wantTargetName := "../testdata/IMG_20200409_220822_Pixel3a.jpg"
	image.GetTargetName()

	if wantTargetName != image.TargetName {
		t.Errorf("NewTargetName = %s, want = %s", image.TargetName, wantTargetName)
	}
}
