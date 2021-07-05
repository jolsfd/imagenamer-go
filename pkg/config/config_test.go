package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func TestGetConfigDir(t *testing.T) {
	homeDir, err := homedir.Dir()
	checkError(err, t)

	want := filepath.Join(homeDir, ".imagenamer-go")
	got := config.GetConfigDir()

	if want != got {
		t.Errorf("Config directory = %s; want = %s;", got, want)
	}
}

func TestGetConfigFile(t *testing.T) {
	homeDir, err := homedir.Dir()
	checkError(err, t)

	want := filepath.Join(homeDir, ".imagenamer-go/config.yaml")
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

func TestLoadNoConfig(t *testing.T) {
	viper.Reset()
	configDir := "../testdata/NotExist"

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
	configDir := "../testdata/temp"
	configFile := "../testdata/temp/config.yaml"

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
