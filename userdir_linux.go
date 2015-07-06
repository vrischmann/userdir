// +build linux

// Package userdir provides functions to get user directories
package userdir

import (
	"os"
	"path/filepath"
)

// GetDataHome returns the user data directory.
func GetDataHome() string {
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome != "" {
		return xdgDataHome
	}

	return filepath.Join(getUserHome(), ".local", "share")
}

// GetConfigHome returns the user config directory.
func GetConfigHome() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		return xdgConfigHome
	}

	return filepath.Join(getUserHome(), ".config")
}

func getUserHome() string {
	return os.Getenv("HOME")
}
