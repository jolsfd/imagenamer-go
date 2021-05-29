// config handles writing config files and opening config files.
package config

import (
	"os"
	"path/filepath"
)

// GetConfigDir returns a string with the path to the config folder.
func GetConfigDir() string {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return err.Error()
	}
	return filepath.Join(configPath, "imagenamer-go")
}
