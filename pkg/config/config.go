// Package config handles writing config files and opening config files.
package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// GetConfigDir returns a string with the path to the config folder.
func GetConfigDir() string {
	configPath, err := os.UserHomeDir()
	if err != nil {
		return err.Error()
	}
	return filepath.Join(configPath, DefaultFolderName)
}

// GetConfigFile returns a string with the absolute path to the settings.
func GetConfigFile() string {
	return filepath.Join(GetConfigDir(), DefaultConfigFileName)
}

// WriteConfigFile write the config file to the config dir and return error.
func WriteConfigFile(configDir string, configFile string) (err error) {
	// Check directory.
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		if err = os.Mkdir(configDir, 0755); err != nil {
			return err
		}
	}

	// Check file.
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		file, err := os.Create(configFile)
		if err != nil {
			return err
		}

		// Close file.
		defer file.Close()
	}

	// Write config.
	err = viper.WriteConfigAs(configFile)
	return err
}

// DefaultConfig set default values to config.
func DefaultConfig() {
	viper.SetDefault(Template, DefaultTemplateString)
	viper.SetDefault(Extensions, DefaultExtensions)
	viper.SetDefault(SafeStrings, DefaultSafePrefixes)
}
