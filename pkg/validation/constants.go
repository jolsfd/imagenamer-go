package validation

import "regexp"

var (
	// windowsForbiddenCharRegex is used to match the strings that contain forbidden
	// characters in Windows' file names. This does not include also forbidden
	// forward and back slash characters because their presence will cause a new
	// directory to be created.
	windowsForbiddenCharRegex = regexp.MustCompile(`<|>|:|"|\||\?|\*`)
	linuxForbiddenCharRegex   = regexp.MustCompile("/")
)

const (
	windowsMaxLength = 260
	unixMaxBytes     = 255
)
