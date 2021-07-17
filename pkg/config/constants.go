package config

// Default config file.
const (
	DefaultFolderName     = ".imagenamer-go"
	DefaultConfigFileName = "config.yaml"
	DefaultConfigName     = "config"
	DefaultConfigType     = "yaml"
	DefaultSeparator      = "~"

	Template    = "template"
	Extensions  = "extensions"
	SafeStrings = "safe_strings"
	Separator   = "copy_separator"
)

// Default Values.
var (
	DefaultTemplateString = "IMG_{{.DateTime}}_{{.CameraModel}}"
	DefaultExtensions     = []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".dng"}
	DefaultSafePrefixes   = []string{"IMG"}
)
