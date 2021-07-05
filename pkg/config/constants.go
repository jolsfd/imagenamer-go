package config

// Default config file.
const (
	DefaultFolderName     = ".imagenamer-go"
	DefaultConfigFileName = "/config.yaml"
	DefaultConfigName     = "config"
	DefaultConfigType     = "yaml"
	DefaultNoConfigError  = "no config file"
)

// Default Values.
var (
	DefaultFormat       = "IMG_DATETIME_MODEL"
	DefaultExtensions   = []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".dng"}
	DefaultSafePrefixes = []string{"IMG"}
)
