
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveConfigPathPrefersEnv(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "custom-config.yaml")
	if err := os.WriteFile(configPath, []byte("env: test\n"), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}

	t.Setenv("CONFIG_PATH", configPath)

	got, err := resolveConfigPath()
	if err != nil {
		t.Fatalf("resolveConfigPath() error = %v", err)
	}
	if got != configPath {
		t.Fatalf("resolveConfigPath() = %q, want %q", got, configPath)
	}
}

func TestResolveConfigPathFallsBackToDefaultCandidates(t *testing.T) {
	dir := t.TempDir()
	configDir := filepath.Join(dir, "config")
	configPath := filepath.Join(configDir, "config.yaml")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		t.Fatalf("mkdir config dir: %v", err)
	}
	if err := os.WriteFile(configPath, []byte("env: test\n"), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() {
		_ = os.Chdir(cwd)
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	t.Setenv("CONFIG_PATH", "")

	got, err := resolveConfigPath()
	if err != nil {
		t.Fatalf("resolveConfigPath() error = %v", err)
	}
	if got != configPath {
		t.Fatalf("resolveConfigPath() = %q, want %q", got, configPath)
	}
}
