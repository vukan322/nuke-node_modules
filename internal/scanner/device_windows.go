//go:build windows

package scanner

import (
	"io/fs"
)

func getDevice(info fs.FileInfo) uint64 {
	return 0
}
