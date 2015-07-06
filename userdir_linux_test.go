// +build linux

package userdir_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vrischmann/userdir"
)

func TestGetDataHome(t *testing.T) {
	d := userdir.GetDataHome()
	require.True(t, strings.HasSuffix(d, ".local/share"))
}

func TestGetConfigHome(t *testing.T) {
	d := userdir.GetConfigHome()
	require.True(t, strings.HasSuffix(d, ".config"))
}
