package scanner

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New("/tmp", 14, false, false)

	if s.rootPath != "/tmp" {
		t.Errorf("expected rootPath /tmp, got %s", s.rootPath)
	}

	if s.days != 14 {
		t.Errorf("expected days 14, got %d", s.days)
	}

	if s.verbose {
		t.Error("expected verbose false")
	}

	if s.includeHidden {
		t.Error("expected includeHidden false")
	}
}

func TestScan_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	s := New(tmpDir, 0, false, false)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 0 {
		t.Errorf("expected 0 folders, got %d", result.TotalCount)
	}
}

func TestScan_FindsNodeModules(t *testing.T) {
	tmpDir := t.TempDir()

	nmPath := filepath.Join(tmpDir, "node_modules")
	if err := os.Mkdir(nmPath, 0755); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().AddDate(0, 0, -30)
	if err := os.Chtimes(nmPath, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 14, false, false)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 1 {
		t.Errorf("expected 1 folder, got %d", result.TotalCount)
	}

	if len(result.Folders) != 1 {
		t.Fatalf("expected 1 folder in results, got %d", len(result.Folders))
	}

	if result.Folders[0].Path != nmPath {
		t.Errorf("expected path %s, got %s", nmPath, result.Folders[0].Path)
	}
}

func TestScan_SkipsRecentNodeModules(t *testing.T) {
	tmpDir := t.TempDir()

	nmPath := filepath.Join(tmpDir, "node_modules")
	if err := os.Mkdir(nmPath, 0755); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 14, false, false)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 0 {
		t.Errorf("expected 0 folders (recent), got %d", result.TotalCount)
	}
}

func TestScan_SkipsHiddenDirectories(t *testing.T) {
	tmpDir := t.TempDir()

	hiddenDir := filepath.Join(tmpDir, ".hidden")
	if err := os.Mkdir(hiddenDir, 0755); err != nil {
		t.Fatal(err)
	}

	nmPath := filepath.Join(hiddenDir, "node_modules")
	if err := os.Mkdir(nmPath, 0755); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().AddDate(0, 0, -30)
	if err := os.Chtimes(nmPath, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 0, false, false)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 0 {
		t.Errorf("expected 0 folders (hidden skipped), got %d", result.TotalCount)
	}
}

func TestScan_IncludesHiddenWithFlag(t *testing.T) {
	tmpDir := t.TempDir()

	hiddenDir := filepath.Join(tmpDir, ".hidden")
	if err := os.Mkdir(hiddenDir, 0755); err != nil {
		t.Fatal(err)
	}

	nmPath := filepath.Join(hiddenDir, "node_modules")
	if err := os.Mkdir(nmPath, 0755); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().AddDate(0, 0, -30)
	if err := os.Chtimes(nmPath, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 0, false, true)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 1 {
		t.Errorf("expected 1 folder (hidden included), got %d", result.TotalCount)
	}
}

func TestScan_MultipleNodeModules(t *testing.T) {
	tmpDir := t.TempDir()

	project1 := filepath.Join(tmpDir, "project1")
	project2 := filepath.Join(tmpDir, "project2")

	if err := os.Mkdir(project1, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(project2, 0755); err != nil {
		t.Fatal(err)
	}

	nm1 := filepath.Join(project1, "node_modules")
	nm2 := filepath.Join(project2, "node_modules")

	if err := os.Mkdir(nm1, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(nm2, 0755); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().AddDate(0, 0, -30)
	if err := os.Chtimes(nm1, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(nm2, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 14, false, false)
	result, err := s.Scan()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalCount != 2 {
		t.Errorf("expected 2 folders, got %d", result.TotalCount)
	}
}

func TestDelete_RemovesNodeModules(t *testing.T) {
	tmpDir := t.TempDir()

	nmPath := filepath.Join(tmpDir, "node_modules")
	if err := os.Mkdir(nmPath, 0755); err != nil {
		t.Fatal(err)
	}

	testFile := filepath.Join(nmPath, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	oldTime := time.Now().AddDate(0, 0, -30)
	if err := os.Chtimes(nmPath, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}

	s := New(tmpDir, 0, false, false)
	scanResult, _ := s.Scan()

	deleteResult, err := s.Delete(scanResult)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if deleteResult.TotalCount != 1 {
		t.Errorf("expected 1 deleted folder, got %d", deleteResult.TotalCount)
	}

	if _, err := os.Stat(nmPath); !os.IsNotExist(err) {
		t.Error("node_modules should be deleted")
	}
}
