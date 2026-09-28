package llvm

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestNewMemoryBufferFromFile(t *testing.T) {
	file, err := os.Create(filepath.Join(t.TempDir(), "memory-buffer.ll"))
	if err != nil {
		t.Fatal(err)
	}
	path := file.Name()

	want := []byte("define\x00void")
	if _, err := file.Write(want); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	buffer, err := NewMemoryBufferFromFile(path)
	if err != nil {
		t.Fatalf("NewMemoryBufferFromFile returned error: %v", err)
	}
	if buffer.IsNil() {
		t.Fatal("NewMemoryBufferFromFile returned a nil buffer")
	}
	defer buffer.Dispose()

	if got := buffer.Bytes(); !bytes.Equal(got, want) {
		t.Errorf("buffer bytes = %q, want %q", got, want)
	}
}

func TestNewMemoryBufferFromFileEmpty(t *testing.T) {
	file, err := os.Create(filepath.Join(t.TempDir(), "memory-buffer-empty.ll"))
	if err != nil {
		t.Fatal(err)
	}
	path := file.Name()
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	buffer, err := NewMemoryBufferFromFile(path)
	if err != nil {
		t.Fatalf("NewMemoryBufferFromFile returned error: %v", err)
	}
	if buffer.IsNil() {
		t.Fatal("NewMemoryBufferFromFile returned a nil buffer for an empty file")
	}
	defer buffer.Dispose()

	if got := buffer.Bytes(); len(got) != 0 {
		t.Errorf("empty buffer has %d bytes, want 0", len(got))
	}
}

func TestNewMemoryBufferFromFileMissingReturnsPathError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.ll")
	buffer, err := NewMemoryBufferFromFile(path)
	if !buffer.IsNil() {
		t.Fatal("NewMemoryBufferFromFile returned a buffer for a missing file")
	}
	if err == nil {
		t.Fatal("NewMemoryBufferFromFile returned nil error for a missing file")
	}

	var pathErr *os.PathError
	if !errors.As(err, &pathErr) {
		t.Fatalf("NewMemoryBufferFromFile error type = %T, want *os.PathError", err)
	}
	if pathErr.Op != "stat" {
		t.Errorf("PathError operation = %q, want %q", pathErr.Op, "stat")
	}
}

// NewMemoryBufferFromFile must record the file as an input of the running test,
// so that `go test` invalidates a cached result when the file changes. LLVM
// opens the file in C, which the os package cannot see, so the stat is what
// puts the path in the test log. Run in a child process because a test cannot
// read its own log until it exits.
func TestNewMemoryBufferFromFileRegistersTestInput(t *testing.T) {
	const childEnv = "GO_LLVM_TESTLOG_CHILD"
	if target := os.Getenv(childEnv); target != "" {
		buffer, err := NewMemoryBufferFromFile(target)
		if err == nil {
			buffer.Dispose()
		}
		return
	}

	dir := t.TempDir()
	target := filepath.Join(dir, "input.ll")
	if err := os.WriteFile(target, []byte("define\x00void"), 0o666); err != nil {
		t.Fatal(err)
	}
	testlog := filepath.Join(dir, "testlog.txt")

	cmd := exec.Command(os.Args[0], "-test.run="+t.Name(), "-test.testlogfile="+testlog)
	cmd.Env = append(os.Environ(), childEnv+"="+target)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("child process failed: %v\n%s", err, out)
	}

	log, err := os.ReadFile(testlog)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(log, []byte(target)) {
		t.Errorf("%s was not recorded as a test input; test log:\n%s", target, log)
	}
}
