package config_test

import (
	"path/filepath"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func TestGetConfigDir(t *testing.T) {
	homeDir, err := homedir.Dir()
	checkError(err, t)

	want := filepath.Join(homeDir, ".config/imagenamer-go")
	got := config.GetConfigDir()

	if want != got {
		t.Errorf("Config directory = %s; want = %s;", got, want)
	}
}

func TestGetConfigFile(t *testing.T) {
	homeDir, err := homedir.Dir()
	checkError(err, t)

	want := filepath.Join(homeDir, ".config/imagenamer-go/config.yaml")
	got := config.GetConfigFile()

	if want != got {
		t.Errorf("Config path = %s; want = %s;", got, want)
	}
}

func TestDefaultConfig(t *testing.T) {
	viper.Reset()
	config.DefaultConfig()

	// want
	wantFormat := "IMG_DATETIME_MODEL"
	wantExtensions := []string{".jpg", ".JPG", ".jpeg", ".JPEG", ".png", ".dng"}
	wantSafePrefixes := []string{"IMG"}

	// got
	format := viper.GetString("Format")
	extensions := viper.GetStringSlice("Extensions")
	safePrefixes := viper.GetStringSlice("SafePrefixes")

	if wantFormat != format {
		t.Errorf("Format = %s; want = %s;", format, wantFormat)
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
	configDir := "../testdata/imagenamer-go"

	viper.SetConfigName(config.DefaultConfigName)
	viper.SetConfigType(config.DefaultConfigType)
	viper.AddConfigPath(configDir)

	err := viper.ReadInConfig()
	err = config.CheckLoadError(err)
	checkError(err, t)
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
