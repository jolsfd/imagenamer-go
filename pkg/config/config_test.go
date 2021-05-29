package config_test

import (
	"path/filepath"
	"testing"

	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/mitchellh/go-homedir"
)

func TestGetHomeDir(t *testing.T) {
	homeDir, err := homedir.Dir()
	checkError(err, t)

	want := filepath.Join(homeDir, ".config/imagenamer-go")
	got := config.GetConfigDir()

	if want != got {
		t.Errorf("Config directory = %s; want = %s;", got, want)
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
