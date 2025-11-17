package scanner

import "time"

type FolderInfo struct {
	Path         string
	Size         int64
	ModifiedTime time.Time
}

type ScanResult struct {
	Folders    []FolderInfo
	TotalSize  int64
	TotalCount int
}
