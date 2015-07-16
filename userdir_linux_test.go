// +build linux

package userdir_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vrischmann/userdir"
)

func TestGetDataHome(t *testing.T) {
	d := userdir.GetDataHome()
	require.True(t, strings.HasSuffix(d, ".local/share"))
}

func TestGetDataHomeFromVariable(t *testing.T) {
	tmp := os.Getenv("XDG_DATA_HOME")
	os.Setenv("XDG_DATA_HOME", "/tmp/foo")

	d := userdir.GetDataHome()
	require.Equal(t, "/tmp/foo", d)

	os.Setenv("XDG_DATA_HOME", tmp)
}

func TestGetConfigHome(t *testing.T) {
	d := userdir.GetConfigHome()
	require.True(t, strings.HasSuffix(d, ".config"))
}

func TestGetConfigHomeFromVariable(t *testing.T) {
	tmp := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", "/tmp/foo")

	d := userdir.GetConfigHome()
	require.Equal(t, "/tmp/foo", d)

	os.Setenv("XDG_CONFIG_HOME", tmp)
}
