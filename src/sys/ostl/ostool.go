package ostl

import (
	"runtime"
)

// 系统工具 ·

// IsWindows ...
func IsWindows() bool {
	if "windows" == runtime.GOOS {
		return true
	}
	return false
}

// IsLinux ...
func IsLinux() bool {
	if "linux" == runtime.GOOS {
		return true
	}
	return false
}

// IsMac ...
func IsMac() bool {
	if "darwin" == runtime.GOOS {
		return true
	}
	return false
}
