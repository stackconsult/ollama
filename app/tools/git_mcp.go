//go:build windows || darwin || linux

package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GitMCP implements Model Context Protocol for Git operations
type GitMCP struct {
	workingDir string
}

// NewGitMCP creates a new GitMCP tool instance
func NewGitMCP(workingDir string) *GitMCP {
	if workingDir == "" {
		workingDir, _ = os.Getwd()
	}
	return &GitMCP{
		workingDir: workingDir,
	}
}

func (g *GitMCP) Name() string {
	return "git_mcp"
}

func (g *GitMCP) Description() string {
	return "Execute git operations following Model Context Protocol. Supports clone, status, log, diff, add, commit, push, pull, and branch operations."
}

func (g *GitMCP) Prompt() string {
	return ""
}

func (g *GitMCP) Schema() map[string]any {
	schemaBytes := []byte(`{
		"type": "object",
		"properties": {
			"operation": {
				"type": "string",
				"description": "Git operation to perform",
				"enum": ["clone", "status", "log", "diff", "add", "commit", "push", "pull", "branch", "checkout", "init"]
			},
			"repository": {
				"type": "string",
				"description": "Repository URL (for clone operation)"
			},
			"path": {
				"type": "string",
				"description": "Local path for the repository or files"
			},
			"message": {
				"type": "string",
				"description": "Commit message (for commit operation)"
			},
			"branch": {
				"type": "string",
				"description": "Branch name (for branch/checkout operations)"
			},
			"files": {
				"type": "array",
				"items": {
					"type": "string"
				},
				"description": "Files to add (for add operation)"
			},
			"options": {
				"type": "object",
				"description": "Additional options for the git command",
				"properties": {
					"limit": {
						"type": "integer",
						"description": "Limit number of results (for log operation)"
					},
					"oneline": {
						"type": "boolean",
						"description": "Show one line per entry (for log operation)"
					},
					"cached": {
						"type": "boolean",
						"description": "Show cached/staged changes (for diff operation)"
					}
				}
			}
		},
		"required": ["operation"]
	}`)
	var schema map[string]any
	if err := json.Unmarshal(schemaBytes, &schema); err != nil {
		return nil
	}
	return schema
}

type GitOperationResult struct {
	Operation string `json:"operation"`
	Output    string `json:"output"`
	Error     string `json:"error,omitempty"`
	Path      string `json:"path,omitempty"`
}

func (g *GitMCP) Execute(ctx context.Context, args map[string]any) (any, string, error) {
	operation, ok := args["operation"].(string)
	if !ok {
		return nil, "", fmt.Errorf("operation parameter is required")
	}

	var result *GitOperationResult
	var err error

	switch operation {
	case "clone":
		result, err = g.executeClone(ctx, args)
	case "status":
		result, err = g.executeStatus(ctx, args)
	case "log":
		result, err = g.executeLog(ctx, args)
	case "diff":
		result, err = g.executeDiff(ctx, args)
	case "add":
		result, err = g.executeAdd(ctx, args)
	case "commit":
		result, err = g.executeCommit(ctx, args)
	case "push":
		result, err = g.executePush(ctx, args)
	case "pull":
		result, err = g.executePull(ctx, args)
	case "branch":
		result, err = g.executeBranch(ctx, args)
	case "checkout":
		result, err = g.executeCheckout(ctx, args)
	case "init":
		result, err = g.executeInit(ctx, args)
	default:
		return nil, "", fmt.Errorf("unsupported operation: %s", operation)
	}

	if err != nil {
		return nil, "", err
	}

	// Format output for the model
	outputText := fmt.Sprintf("Git %s operation completed.\n", operation)
	if result.Output != "" {
		outputText += result.Output
	}
	if result.Error != "" {
		outputText += fmt.Sprintf("\nWarnings/Errors: %s", result.Error)
	}

	return result, outputText, nil
}

func (g *GitMCP) executeClone(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	repository, ok := args["repository"].(string)
	if !ok || repository == "" {
		return nil, fmt.Errorf("repository URL is required for clone operation")
	}

	targetPath := g.workingDir
	if path, ok := args["path"].(string); ok && path != "" {
		targetPath = path
	}

	// Ensure the target directory exists
	parentDir := filepath.Dir(targetPath)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	cmd := exec.CommandContext(ctx, "git", "clone", repository, targetPath)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "clone",
		Output:    string(output),
		Path:      targetPath,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("clone failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeStatus(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmd := exec.CommandContext(ctx, "git", "-C", path, "status")
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "status",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("status failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeLog(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmdArgs := []string{"-C", path, "log"}

	// Handle options
	if options, ok := args["options"].(map[string]any); ok {
		if limit, ok := options["limit"].(float64); ok && limit > 0 {
			cmdArgs = append(cmdArgs, fmt.Sprintf("-n%d", int(limit)))
		}
		if oneline, ok := options["oneline"].(bool); ok && oneline {
			cmdArgs = append(cmdArgs, "--oneline")
		}
	}

	cmd := exec.CommandContext(ctx, "git", cmdArgs...)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "log",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("log failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeDiff(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmdArgs := []string{"-C", path, "diff"}

	// Handle options
	if options, ok := args["options"].(map[string]any); ok {
		if cached, ok := options["cached"].(bool); ok && cached {
			cmdArgs = append(cmdArgs, "--cached")
		}
	}

	cmd := exec.CommandContext(ctx, "git", cmdArgs...)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "diff",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
	}

	return result, nil
}

func (g *GitMCP) executeAdd(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmdArgs := []string{"-C", path, "add"}

	// Handle files parameter
	if filesRaw, ok := args["files"]; ok {
		switch files := filesRaw.(type) {
		case []interface{}:
			for _, f := range files {
				if fileStr, ok := f.(string); ok {
					cmdArgs = append(cmdArgs, fileStr)
				}
			}
		case []string:
			cmdArgs = append(cmdArgs, files...)
		case string:
			cmdArgs = append(cmdArgs, files)
		default:
			return nil, fmt.Errorf("invalid files parameter type")
		}
	} else {
		// Default to add all
		cmdArgs = append(cmdArgs, ".")
	}

	cmd := exec.CommandContext(ctx, "git", cmdArgs...)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "add",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("add failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeCommit(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	message, ok := args["message"].(string)
	if !ok || message == "" {
		return nil, fmt.Errorf("commit message is required")
	}

	cmd := exec.CommandContext(ctx, "git", "-C", path, "commit", "-m", message)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "commit",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("commit failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executePush(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmd := exec.CommandContext(ctx, "git", "-C", path, "push")
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "push",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("push failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executePull(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmd := exec.CommandContext(ctx, "git", "-C", path, "pull")
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "pull",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("pull failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeBranch(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	cmdArgs := []string{"-C", path, "branch"}

	if branch, ok := args["branch"].(string); ok && branch != "" {
		cmdArgs = append(cmdArgs, branch)
	}

	cmd := exec.CommandContext(ctx, "git", cmdArgs...)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "branch",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("branch failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeCheckout(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	branch, ok := args["branch"].(string)
	if !ok || branch == "" {
		return nil, fmt.Errorf("branch name is required for checkout operation")
	}

	cmd := exec.CommandContext(ctx, "git", "-C", path, "checkout", branch)
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "checkout",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("checkout failed: %w", err)
	}

	return result, nil
}

func (g *GitMCP) executeInit(ctx context.Context, args map[string]any) (*GitOperationResult, error) {
	path := g.getRepoPath(args)

	// Ensure the directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	cmd := exec.CommandContext(ctx, "git", "-C", path, "init")
	output, err := cmd.CombinedOutput()

	result := &GitOperationResult{
		Operation: "init",
		Output:    string(output),
		Path:      path,
	}

	if err != nil {
		result.Error = err.Error()
		return result, fmt.Errorf("init failed: %w", err)
	}

	return result, nil
}

// getRepoPath returns the repository path from args or uses the working directory
func (g *GitMCP) getRepoPath(args map[string]any) string {
	if path, ok := args["path"].(string); ok && path != "" {
		// Clean the path and make it absolute if it's relative
		if !filepath.IsAbs(path) {
			absPath, err := filepath.Abs(path)
			if err == nil {
				return absPath
			}
		}
		return filepath.Clean(path)
	}
	return g.workingDir
}

// IsGitRepository checks if the given path is a git repository
func IsGitRepository(path string) bool {
	gitDir := filepath.Join(path, ".git")
	info, err := os.Stat(gitDir)
	return err == nil && info.IsDir()
}

// FindGitRoot finds the root of the git repository starting from the given path
func FindGitRoot(startPath string) (string, error) {
	currentPath, err := filepath.Abs(startPath)
	if err != nil {
		return "", err
	}

	for {
		if IsGitRepository(currentPath) {
			return currentPath, nil
		}

		parent := filepath.Dir(currentPath)
		if parent == currentPath {
			// Reached the root of the filesystem
			return "", fmt.Errorf("not a git repository (or any of the parent directories)")
		}
		currentPath = parent
	}
}

// ValidateGitInstalled checks if git is installed and available
func ValidateGitInstalled() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not in PATH: %w", err)
	}
	return nil
}

// SanitizeGitURL sanitizes a git URL to remove sensitive information
func SanitizeGitURL(gitURL string) string {
	// Remove authentication information from URLs
	if strings.Contains(gitURL, "@") {
		parts := strings.Split(gitURL, "@")
		if len(parts) >= 2 {
			// Keep only the part after the @
			return parts[len(parts)-1]
		}
	}
	return gitURL
}
