//go:build windows || darwin || linux

package tools

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGitMCP_Name(t *testing.T) {
	git := NewGitMCP("")
	if git.Name() != "git_mcp" {
		t.Errorf("Expected name 'git_mcp', got '%s'", git.Name())
	}
}

func TestGitMCP_Description(t *testing.T) {
	git := NewGitMCP("")
	desc := git.Description()
	if desc == "" {
		t.Error("Description should not be empty")
	}
}

func TestGitMCP_Schema(t *testing.T) {
	git := NewGitMCP("")
	schema := git.Schema()
	if schema == nil {
		t.Error("Schema should not be nil")
	}

	// Verify schema has required fields
	if _, ok := schema["properties"]; !ok {
		t.Error("Schema should have 'properties' field")
	}
	if _, ok := schema["required"]; !ok {
		t.Error("Schema should have 'required' field")
	}
}

func TestGitMCP_Execute_InvalidOperation(t *testing.T) {
	git := NewGitMCP("")
	ctx := context.Background()

	_, _, err := git.Execute(ctx, map[string]any{
		"operation": "invalid_operation",
	})

	if err == nil {
		t.Error("Expected error for invalid operation")
	}
}

func TestGitMCP_Execute_MissingOperation(t *testing.T) {
	git := NewGitMCP("")
	ctx := context.Background()

	_, _, err := git.Execute(ctx, map[string]any{})

	if err == nil {
		t.Error("Expected error when operation is missing")
	}
}

func TestGitMCP_Init(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	result, text, err := git.Execute(ctx, map[string]any{
		"operation": "init",
		"path":      tmpDir,
	})

	if err != nil {
		t.Errorf("Init operation failed: %v", err)
	}

	if result == nil {
		t.Error("Result should not be nil")
	}

	if text == "" {
		t.Error("Output text should not be empty")
	}

	// Verify .git directory was created
	gitDir := filepath.Join(tmpDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Error(".git directory was not created")
	}
}

func TestGitMCP_Status(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory and initialize a git repo
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	result, text, err := git.Execute(ctx, map[string]any{
		"operation": "status",
		"path":      tmpDir,
	})

	if err != nil {
		t.Errorf("Status operation failed: %v", err)
	}

	if result == nil {
		t.Error("Result should not be nil")
	}

	if text == "" {
		t.Error("Output text should not be empty")
	}
}

func TestGitMCP_AddAndCommit(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory and initialize a git repo
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Configure git user for the test
	exec.Command("git", "-C", tmpDir, "config", "user.email", "test@example.com").Run()
	exec.Command("git", "-C", tmpDir, "config", "user.name", "Test User").Run()

	// Create a test file
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	// Test add operation
	_, _, err = git.Execute(ctx, map[string]any{
		"operation": "add",
		"path":      tmpDir,
		"files":     []string{"test.txt"},
	})

	if err != nil {
		t.Errorf("Add operation failed: %v", err)
	}

	// Test commit operation
	_, _, err = git.Execute(ctx, map[string]any{
		"operation": "commit",
		"path":      tmpDir,
		"message":   "Test commit",
	})

	if err != nil {
		t.Errorf("Commit operation failed: %v", err)
	}
}

func TestGitMCP_Branch(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory and initialize a git repo
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Configure git user and create initial commit
	exec.Command("git", "-C", tmpDir, "config", "user.email", "test@example.com").Run()
	exec.Command("git", "-C", tmpDir, "config", "user.name", "Test User").Run()
	os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test"), 0644)
	exec.Command("git", "-C", tmpDir, "add", ".").Run()
	exec.Command("git", "-C", tmpDir, "commit", "-m", "Initial commit").Run()

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	// Test branch listing
	result, text, err := git.Execute(ctx, map[string]any{
		"operation": "branch",
		"path":      tmpDir,
	})

	if err != nil {
		t.Errorf("Branch operation failed: %v", err)
	}

	if result == nil {
		t.Error("Result should not be nil")
	}

	if text == "" {
		t.Error("Output text should not be empty")
	}
}

func TestIsGitRepository(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Should not be a git repository yet
	if IsGitRepository(tmpDir) {
		t.Error("Directory should not be identified as a git repository")
	}

	// Create .git directory
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.MkdirAll(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git directory: %v", err)
	}

	// Should now be identified as a git repository
	if !IsGitRepository(tmpDir) {
		t.Error("Directory should be identified as a git repository")
	}
}

func TestFindGitRoot(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Create a subdirectory
	subDir := filepath.Join(tmpDir, "subdir", "nested")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Find git root from subdirectory
	root, err := FindGitRoot(subDir)
	if err != nil {
		t.Errorf("Failed to find git root: %v", err)
	}

	// Compare cleaned paths
	expectedRoot := filepath.Clean(tmpDir)
	foundRoot := filepath.Clean(root)

	if foundRoot != expectedRoot {
		t.Errorf("Expected git root '%s', got '%s'", expectedRoot, foundRoot)
	}
}

func TestValidateGitInstalled(t *testing.T) {
	err := ValidateGitInstalled()
	if err != nil {
		t.Logf("Git is not installed: %v", err)
		t.Skip("Git is not installed on this system")
	}
}

func TestSanitizeGitURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "https://github.com/user/repo.git",
			expected: "https://github.com/user/repo.git",
		},
		{
			input:    "https://username:password@github.com/user/repo.git",
			expected: "github.com/user/repo.git",
		},
		{
			input:    "git@github.com:user/repo.git",
			expected: "github.com:user/repo.git",
		},
	}

	for _, tt := range tests {
		result := SanitizeGitURL(tt.input)
		if result != tt.expected {
			t.Errorf("SanitizeGitURL(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestGitMCP_Log(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory and initialize a git repo
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Configure git user and create initial commit
	exec.Command("git", "-C", tmpDir, "config", "user.email", "test@example.com").Run()
	exec.Command("git", "-C", tmpDir, "config", "user.name", "Test User").Run()
	os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test"), 0644)
	exec.Command("git", "-C", tmpDir, "add", ".").Run()
	exec.Command("git", "-C", tmpDir, "commit", "-m", "Initial commit").Run()

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	// Test log operation with options
	result, text, err := git.Execute(ctx, map[string]any{
		"operation": "log",
		"path":      tmpDir,
		"options": map[string]any{
			"limit":   float64(1),
			"oneline": true,
		},
	})

	if err != nil {
		t.Errorf("Log operation failed: %v", err)
	}

	if result == nil {
		t.Error("Result should not be nil")
	}

	if text == "" {
		t.Error("Output text should not be empty")
	}
}

func TestGitMCP_Diff(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory and initialize a git repo
	tmpDir, err := os.MkdirTemp("", "git-mcp-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Configure git user and create initial commit
	exec.Command("git", "-C", tmpDir, "config", "user.email", "test@example.com").Run()
	exec.Command("git", "-C", tmpDir, "config", "user.name", "Test User").Run()
	os.WriteFile(filepath.Join(tmpDir, "test.txt"), []byte("original content"), 0644)
	exec.Command("git", "-C", tmpDir, "add", ".").Run()
	exec.Command("git", "-C", tmpDir, "commit", "-m", "Initial commit").Run()

	// Modify the file
	os.WriteFile(filepath.Join(tmpDir, "test.txt"), []byte("modified content"), 0644)

	git := NewGitMCP(tmpDir)
	ctx := context.Background()

	// Test diff operation
	result, text, err := git.Execute(ctx, map[string]any{
		"operation": "diff",
		"path":      tmpDir,
	})

	if err != nil {
		t.Errorf("Diff operation failed: %v", err)
	}

	if result == nil {
		t.Error("Result should not be nil")
	}

	// Note: text might be empty if there are no differences
	_ = text
}
