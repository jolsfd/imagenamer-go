package config

// Default config file.
const (
	DefaultFolderName     = ".imagenamer-go"
	DefaultConfigFileName = "config.yaml"
	DefaultConfigName     = "config"
	DefaultConfigType     = "yaml"

	Template    = "template"
	Extensions  = "extensions"
	SafeStrings = "safe_strings"
)

// Default Values.
var (
	DefaultTemplateString = "IMG_{{.DateTime}}_{{.CameraModel}}"
	DefaultExtensions     = []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".dng"}
	DefaultSafePrefixes   = []string{"IMG"}
)
