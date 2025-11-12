//go:build windows || darwin || linux

package tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// TestGitMCP_Standalone tests the git_mcp tool independently
func TestGitMCP_Standalone(t *testing.T) {
	// Skip if git is not installed
	if err := ValidateGitInstalled(); err != nil {
		t.Skip("Git is not installed")
	}

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "git-mcp-standalone-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the git_mcp tool
	gitTool := NewGitMCP(tmpDir)

	if gitTool.Name() != "git_mcp" {
		t.Errorf("Expected tool name 'git_mcp', got '%s'", gitTool.Name())
	}

	ctx := context.Background()

	// Test executing init
	result, text, err := gitTool.Execute(ctx, map[string]any{
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

	// Test executing status
	result, text, err = gitTool.Execute(ctx, map[string]any{
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

	// Test schema
	schema := gitTool.Schema()
	if schema == nil {
		t.Error("Schema should not be nil")
	}

	if _, ok := schema["properties"]; !ok {
		t.Error("Schema should have 'properties' field")
	}

	// Test description
	desc := gitTool.Description()
	if desc == "" {
		t.Error("Description should not be empty")
	}
}
