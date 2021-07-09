package rename_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/metadata"
	"github.com/jolsfd/imagenamer-go/pkg/rename"
)

var sourceName = filepath.Join("..", "testdata", "test_image.jpg")
var path = filepath.Join("..", "testdata")

func TestCheckFileExists(t *testing.T) {
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
		Path:          path,
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
	templateString := "IMG_{{.DateTime}}_{{.CameraModel}}"
	exif, err := metadata.GetExif(sourceName)
	if err != nil {
		t.Error(err)
	}
	image.BuildFileInformation(sourceName)

	wantNewFilename := "IMG_20200409_220822_Pixel3a"
	err = image.GetNewFileName(templateString, exif)
	if err != nil {
		t.Error(err)
	}

	if wantNewFilename != image.NewFileName {
		t.Errorf("NewFilename = %s, want = %s", image.NewFileName, wantNewFilename)
	}

	wantTargetName := filepath.Join(path, "IMG_20200409_220822_Pixel3a~2.jpg")
	image.GetTargetName()

	if wantTargetName != image.TargetName {
		t.Errorf("NewTargetName = %s, want = %s", image.TargetName, wantTargetName)
	}
}

func TestListImagesInDir(t *testing.T) {
	exclude := []string{"exclude"}
	extensions := []string{".jpg"}
	safeRename := true
	safePrefixes := []string{"test"}

	// Test with SafeRename and exclude.
	want := []string{filepath.Join(path, "IMG_20200409_220822_Pixel3a.jpg")}

	files, err := rename.ListImagesInDir(path, extensions, exclude, safeRename, safePrefixes)
	if err != nil {
		t.Error(err)
	}
	for i, file := range files {
		if file != want[i] {
			t.Errorf("got = %s, want = %s", file, want[i])
		}
	}

	// Test without SafeRename and exclude.
	safeRename = false
	exclude = []string{}
	want = []string{filepath.Join(path, "IMG_20200409_220822_Pixel3a.jpg"), filepath.Join(path, "exclude", "excludeImage.jpg"), sourceName}
	files, err = rename.ListImagesInDir(path, extensions, exclude, safeRename, safePrefixes)
	if err != nil {
		t.Error(err)
	}
	for i, file := range files {
		if file != want[i] {
			t.Errorf("got = %s, want = %s", file, want[i])
		}
	}

}

func TestGetFileInformation(t *testing.T) {
	debug := true
	templateString := "IMG_{{.DateTime}}_{{.CameraModel}}"
	sourceNames := []string{
		filepath.Join(path, "IMG_20200409_220822_Pixel3a.jpg"),
		filepath.Join(path, "exclude", "excludeImage.jpg"),
		filepath.Join(path, "test_image.jpg")}

	wantTableData := [][]string{
		{"IMG_20200409_220822_Pixel3a.jpg", "IMG_20200409_220822_Pixel3a.jpg", "ok"},
		{"excludeImage.jpg", ".jpg", "fail"},
		{"test_image.jpg", "IMG_20200409_220822_Pixel3a~2.jpg", "ok"},
	}
	_, tableData, err := rename.GetFileInformation(sourceNames, templateString, debug)
	if err != nil {
		t.Error(err)
	}

	for i := range tableData {
		for j := range tableData[i] {
			if tableData[i][j] != wantTableData[i][j] {
				t.Errorf("got = %s, want = %s", tableData[i][j], wantTableData[i][j])
			}
		}
	}
}

func TestRenameImages(t *testing.T) {
	targetName := filepath.Join(path, "IMG_20200409_220822_Pixel3a~2.jpg")
	err := rename.RenameImages([]rename.FileInformation{{
		Path:          path,
		SourceName:    sourceName,
		TargetName:    targetName,
		FileName:      "test_image",
		NewFileName:   "IMG_20200409_220822_Pixel3a",
		FileExtension: ".jpg",
		CopyNumber:    2,
		Status:        rename.StatusOk,
	}})
	if err != nil {
		t.Error(err)
	}

	// Clean up.
	os.Rename(targetName, sourceName)
}
