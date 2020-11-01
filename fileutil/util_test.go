package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMkdirAll_SilentFailure(t *testing.T) {
	dirPath := "myapplication/secrets"
	t.Cleanup(func() {
		if err := os.RemoveAll(dirPath); err != nil {
			t.Error("Could not remove directory")
		}
	})
	// Evil attacker creates the directory ahead of time
	// with full 777 permissions.
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		t.Fatalf("Could not write directory: %v", err)
	}
	// Now our application attempts to write to the directory
	// with 700 permissions to only allow current user read/write/exec.
	err := MkdirAll(dirPath)
	if err == nil {
		t.Error("Expected error, received nil")
	}
	if !strings.Contains(err.Error(), "already exists with wrong permissions") {
		t.Errorf("Received wrong error %v", err)
	}
}

func TestWriteFile_SilentFailure(t *testing.T) {
	dirPath := "myapplication/secrets"
	t.Cleanup(func() {
		if err := os.RemoveAll(dirPath); err != nil {
			t.Error("Could not remove directory")
		}
	})
	// We create a directory with 777 permissions.
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		t.Fatalf("Could not write directory: %v", err)
	}
	secretFile := filepath.Join(dirPath, "credentials.txt")
	if err := ioutil.WriteFile(secretFile, []byte("password"), 0777); err != nil {
		t.Fatalf("Could not write file: %v", err)
	}
	// Now our application attempts to write to the file
	// to only allow current user read/write/exec.
	err := WriteFile(secretFile, []byte("password"))
	if err == nil {
		t.Error("Expected error, received nil")
	}
	if !strings.Contains(err.Error(), "already exists with wrong permissions") {
		t.Errorf("Received wrong error %v", err)
	}
}
