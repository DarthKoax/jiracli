package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain_Help(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	out, err := exec.Command(binary, "help").CombinedOutput()
	if err != nil {
		t.Fatalf("help command failed: %v\n%s", err, out)
	}
	output := string(out)
	if !strings.Contains(output, "jiracli") {
		t.Errorf("help output missing 'jiracli': %s", output)
	}
	if !strings.Contains(output, "init") {
		t.Errorf("help output missing 'init': %s", output)
	}
	if !strings.Contains(output, "connect") {
		t.Errorf("help output missing 'connect': %s", output)
	}
}

func TestMain_Version(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	out, err := exec.Command(binary, "version").CombinedOutput()
	if err != nil {
		t.Fatalf("version command failed: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "jiracli") {
		t.Errorf("version output missing 'jiracli': %s", out)
	}
}

func TestMain_Init(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	tmpdir, err := os.MkdirTemp("", "jiracli-init-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	out, err := exec.Command(binary, "init", "--dir", tmpdir).CombinedOutput()
	if err != nil {
		t.Fatalf("init command failed: %v\n%s", err, out)
	}

	configPath := filepath.Join(tmpdir, "config.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("init did not create config at %s", configPath)
	}

	output := string(out)
	if !strings.Contains(output, configPath) {
		t.Errorf("init output missing config path: %s", output)
	}
}

func TestMain_InitAlreadyExists(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	tmpdir, err := os.MkdirTemp("", "jiracli-init-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	_, err = exec.Command(binary, "init", "--dir", tmpdir).CombinedOutput()
	if err != nil {
		t.Fatalf("first init failed: %v", err)
	}

	out, err := exec.Command(binary, "init", "--dir", tmpdir).CombinedOutput()
	if err == nil {
		t.Error("second init should have failed, but succeeded")
	}
	output := string(out)
	if !strings.Contains(output, "already exists") {
		t.Errorf("expected 'already exists' error, got: %s", output)
	}
}

func TestMain_UnknownCommand(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	out, err := exec.Command(binary, "nonexistent").CombinedOutput()
	if err == nil {
		t.Error("unknown command should have failed")
	}
	output := string(out)
	if !strings.Contains(output, "Unknown command") {
		t.Errorf("expected 'Unknown command' error, got: %s", output)
	}
}

func TestMain_NoArgs(t *testing.T) {
	binary := buildBinary(t)
	defer os.Remove(binary)

	out, err := exec.Command(binary).CombinedOutput()
	if err == nil {
		t.Error("no args should have failed")
	}
	output := string(out)
	if !strings.Contains(output, "USAGE") && !strings.Contains(output, "Usage") {
		t.Errorf("expected usage output, got: %s", output)
	}
}

func buildBinary(t *testing.T) string {
	t.Helper()
	tmpdir, err := os.MkdirTemp("", "jiracli-build-*")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(tmpdir, "jiracli")
	cmd := exec.Command("go", "build", "-o", binary, ".")
	cmd.Dir = filepath.Join(getProjectRoot(t), "cmd", "jiracli")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build binary: %v\n%s", err, out)
	}
	return binary
}

func getProjectRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find project root")
		}
		dir = parent
	}
}
