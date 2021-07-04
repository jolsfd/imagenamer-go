package validation_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/validation"
)

func TestCheckForbiddenChars(t *testing.T) {
	if runtime.GOOS == "windows" {
		// Forbidden chars.
		fileName := "w<>i:\"n|d?o*ws"

		err := validation.CheckForbiddenCharacters(fileName)
		checkCharError(err, t)

		// Not forbidden chars.
		fileName = "windows"

		err = validation.CheckForbiddenCharacters(fileName)
		if err != nil {
			t.Error(err)
		}

	} else if runtime.GOOS != "windows" {
		// Forbidden chars.
		fileName := "li/nux"

		err := validation.CheckForbiddenCharacters(fileName)
		checkCharError(err, t)

		// Not forbidden chars.
		fileName = "linux"

		err = validation.CheckForbiddenCharacters(fileName)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCheckMaxLenght(t *testing.T) {
	if runtime.GOOS == "windows" {
		targetName := strings.Repeat("w", 260)

		err := validation.CheckTargetLength(targetName)
		if err != nil {
			t.Error(err)
		}

		targetName = strings.Repeat("w", 261)
		err = validation.CheckTargetLength(targetName)
		if err == nil {
			t.Error("string should be too long")
		}

	} else if runtime.GOOS != "windows" {
		targetName := strings.Repeat("l", 255)

		err := validation.CheckTargetLength(targetName)
		if err != nil {
			t.Error(err)
		}

		targetName = strings.Repeat("l", 256)
		err = validation.CheckTargetLength(targetName)
		if err == nil {
			t.Error("string should be too long")
		}
	}
}

func checkCharError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("%v invalid chars\n", err)
	}
}
