package validation

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// CheckForbiddenCharacters is responsible for ensuring that target file names
// do not contain forbidden characters for the current OS.
func CheckForbiddenCharacters(fileName string) error {
	if runtime.GOOS == "windows" {
		if windowsForbiddenCharRegex.MatchString(fileName) {
			return errors.New(
				strings.Join(
					windowsForbiddenCharRegex.FindAllString(fileName, -1),
					",",
				),
			)
		}
	} else if runtime.GOOS != "windows" {
		if linuxForbiddenCharRegex.MatchString(fileName) {
			return errors.New(
				strings.Join(
					linuxForbiddenCharRegex.FindAllString(fileName, -1),
					",",
				),
			)
		}
	}

	return nil
}

// CheckTargetLength is responsible for ensuring that the target name length
// does not exceed the maximum value on each supported operating system.
func CheckTargetLength(targetName string) error {
	// Get the standalone filename
	filename := filepath.Base(targetName)

	if runtime.GOOS == "windows" && len([]rune(filename)) > windowsMaxLength {
		// max length of 260 characters in windows
		return fmt.Errorf("%d characters", windowsMaxLength)

	} else if runtime.GOOS != "windows" && len([]byte(filename)) > unixMaxBytes {
		// max length of 255 bytes on Linux and other unix-based OSes
		return fmt.Errorf("%d bytes", unixMaxBytes)
	}

	return nil
}
