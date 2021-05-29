package config_test

import "testing"

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
