// +build windows

// Package userdir provides functions to get user directories
package userdir

import (
	"strings"
	"syscall"
	"unsafe"
)

// GetDataHome returns the user data directory.
func GetDataHome() string {
	return getRoamingAppDataDir()
}

// GetConfigHome returns the user config directory.
func GetConfigHome() string {
	return getRoamingAppDataDir()
}

type guid struct {
	data1 uint32
	data2 uint16
	data3 uint16
	data4 [8]byte
}

var (
	modshell32               = syscall.NewLazyDLL("shell32.dll")
	modole32                 = syscall.NewLazyDLL("ole32.dll")
	procSHGetKnownFolderPath = modshell32.NewProc("SHGetKnownFolderPath")
	procCoTaskMemFree        = modole32.NewProc("CoTaskMemFree")

	roamingAppData = guid{
		0x3EB685DB,
		0x65F9,
		0x4CF6,
		[8]byte{0xA0, 0x3A, 0xE3, 0xEF, 0x65, 0x72, 0x9F, 0x3D},
	}
)

func coTaskMemFree(ptr uintptr) {
	procCoTaskMemFree.Call(ptr)
}

func getRoamingAppDataDir() string {
	dwFlags := uint32(0)
	var pwstr uintptr

	ret, _, _ := procSHGetKnownFolderPath.Call(
		uintptr(unsafe.Pointer(&roamingAppData)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&pwstr)),
	)
	defer coTaskMemFree(pwstr)

	if ret != 0 {
		return ""
	}

	return normalizeWindowsPath(utf16PtrToString(pwstr))
}

func normalizeWindowsPath(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}

func utf16PtrToString(str uintptr) string {
	// TODO(vincent): see if we can do anything about go vet complaining
	ptr := unsafe.Pointer(str)
	return syscall.UTF16ToString((*[1 << 16]uint16)(ptr)[:])
}
