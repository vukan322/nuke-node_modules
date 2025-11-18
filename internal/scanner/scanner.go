package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/vukan322/nuke-node_modules/internal/util"
)

type Scanner struct {
	rootPath      string
	days          int
	verbose       bool
	includeHidden bool
	rootDev       uint64
	cutoffTime    time.Time
}

func New(rootPath string, days int, verbose, includeHidden bool) *Scanner {
	return &Scanner{
		rootPath:      rootPath,
		days:          days,
		verbose:       verbose,
		includeHidden: includeHidden,
		cutoffTime:    time.Now().AddDate(0, 0, -days),
	}
}

func (s *Scanner) Scan() (*ScanResult, error) {
	stat, err := os.Stat(s.rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat root path: %w", err)
	}

	s.rootDev = getDevice(stat)

	result := &ScanResult{
		Folders: make([]FolderInfo, 0),
	}

	err = filepath.WalkDir(s.rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if s.verbose {
				fmt.Fprintf(os.Stderr, "Warning: cannot access %s: %v\n", path, err)
			}
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		if d.Type()&os.ModeSymlink != 0 {
			if s.verbose {
				fmt.Fprintf(os.Stderr, "Skipping symlink: %s\n", path)
			}
			return fs.SkipDir
		}

		if !s.includeHidden {
			base := filepath.Base(path)
			if len(base) > 0 && base[0] == '.' && path != s.rootPath {
				return fs.SkipDir
			}
		}

		info, err := d.Info()
		if err != nil {
			if s.verbose {
				fmt.Fprintf(os.Stderr, "Warning: cannot get info for %s: %v\n", path, err)
			}
			return nil
		}

		if getDevice(info) != s.rootDev {
			if s.verbose {
				fmt.Fprintf(os.Stderr, "Skipping different filesystem: %s\n", path)
			}
			return fs.SkipDir
		}

		if d.Name() == "node_modules" {
			if info.ModTime().Before(s.cutoffTime) {
				size := calculateSize(path)
				folder := FolderInfo{
					Path:         path,
					Size:         size,
					ModifiedTime: info.ModTime(),
				}
				result.Folders = append(result.Folders, folder)
				result.TotalSize += size
				result.TotalCount++

				if s.verbose {
					fmt.Printf("Found: %s (%s, modified %s)\n",
						path,
						util.FormatSize(size),
						info.ModTime().Format("2006-01-02"),
					)
				}
			} else if s.verbose {
				fmt.Printf("Skipping recent: %s (modified %s)\n",
					path,
					info.ModTime().Format("2006-01-02"),
				)
			}
			return fs.SkipDir
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walk error: %w", err)
	}

	return result, nil
}

func (s *Scanner) Delete(result *ScanResult) (*ScanResult, error) {
	deleted := &ScanResult{
		Folders: make([]FolderInfo, 0),
	}

	for _, folder := range result.Folders {
		err := os.RemoveAll(folder.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete %s: %v\n", folder.Path, err)
			continue
		}

		deleted.Folders = append(deleted.Folders, folder)
		deleted.TotalSize += folder.Size
		deleted.TotalCount++

		if s.verbose {
			fmt.Printf("Deleted: %s (%s)\n", folder.Path, util.FormatSize(folder.Size))
		}
	}

	return deleted, nil
}

func calculateSize(path string) int64 {
	var size int64
	_ = filepath.WalkDir(path, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err == nil {
				size += info.Size()
			}
		}
		return nil
	})
	return size
}
