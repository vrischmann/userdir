// +build darwin

// Package userdir provides functions to get user directories
package userdir

import (
	"os"
	"path/filepath"
)

// GetDataHome returns the user data directory.
func GetDataHome() string {
	return filepath.Join(getUserHome(), "Library")
}

// GetConfigHome returns the user config directory.
func GetConfigHome() string {
	return filepath.Join(getUserHome(), "Library", "Preferences")
}

func getUserHome() string {
	return os.Getenv("HOME")
}
