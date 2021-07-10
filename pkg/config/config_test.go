package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/spf13/viper"
)

func TestGetConfigDir(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	checkError(err, t)

	want := filepath.Join(homeDir, config.DefaultFolderName)
	got := config.GetConfigDir()

	if want != got {
		t.Errorf("Config directory = %s; want = %s;", got, want)
	}
}

func TestGetConfigFile(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	checkError(err, t)

	want := filepath.Join(homeDir, config.DefaultFolderName, config.DefaultConfigFileName)
	got := config.GetConfigFile()

	if want != got {
		t.Errorf("Config path = %s; want = %s;", got, want)
	}
}

func TestDefaultConfig(t *testing.T) {
	viper.Reset()
	config.DefaultConfig()

	// want
	wantTemplateString := "IMG_{{.DateTime}}_{{.CameraModel}}"
	wantExtensions := []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".dng"}
	wantSafePrefixes := []string{"IMG"}

	// got
	templateString := viper.GetString(config.Template)
	extensions := viper.GetStringSlice(config.Extensions)
	safePrefixes := viper.GetStringSlice(config.SafeStrings)

	if wantTemplateString != templateString {
		t.Errorf("Template = %s; want = %s;", templateString, wantTemplateString)
	}

	for index, extension := range extensions {
		if extension != wantExtensions[index] {
			t.Errorf("Extension = %v; want = %v;", extension, wantExtensions[index])
		}
	}

	for index, safePrefix := range safePrefixes {
		if safePrefix != wantSafePrefixes[index] {
			t.Errorf("Extension = %v; want = %v;", safePrefix, wantSafePrefixes[index])
		}
	}
}

func TestLoadConfig(t *testing.T) {
	viper.Reset()
	configDir := filepath.Join("..", "testdata", "imagenamer-go")

	viper.SetConfigName(config.DefaultConfigName)
	viper.SetConfigType(config.DefaultConfigType)
	viper.AddConfigPath(configDir)

	err := viper.ReadInConfig()
	err = config.CheckLoadError(err)
	checkError(err, t)
}

func TestLoadNoConfig(t *testing.T) {
	viper.Reset()
	configDir := filepath.Join("..", "testdata", "NotExist")

	viper.SetConfigName(config.DefaultConfigName)
	viper.SetConfigType(config.DefaultConfigType)
	viper.AddConfigPath(configDir)

	err := viper.ReadInConfig()
	err = config.CheckLoadError(err)

	if err.Error() != config.DefaultNoConfigError {
		t.Errorf("want = %s; got = %v", config.DefaultNoConfigError, err)
	}
}

func TestWriteConfigFile(t *testing.T) {
	viper.Reset()
	var err error
	configDir := filepath.Join("..", "testdata", "temp")
	configFile := filepath.Join("..", "testdata", "temp", "config.yaml")

	viper.SetConfigName(config.DefaultConfigName)
	viper.SetConfigType(config.DefaultConfigType)

	config.DefaultConfig()

	err = config.WriteConfigFile(configDir, configFile)
	checkError(err, t)

	err = os.Remove(configFile)
	checkError(err, t)

	err = os.Remove(configDir)
	checkError(err, t)
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
