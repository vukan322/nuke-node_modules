//go:build unix

package scanner

import (
	"io/fs"
	"syscall"
)

func getDevice(info fs.FileInfo) uint64 {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0
	}
	return uint64(stat.Dev)
}
