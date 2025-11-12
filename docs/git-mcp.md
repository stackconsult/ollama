# Git MCP Tool

The Git MCP (Model Context Protocol) tool provides Git repository management capabilities for Ollama models that support function calling. This allows LLMs to interact with Git repositories to perform common version control operations.

## Features

The git_mcp tool supports the following operations:

- **init**: Initialize a new Git repository
- **clone**: Clone a repository from a URL
- **status**: Show the working tree status
- **log**: Show commit logs
- **diff**: Show changes between commits or working tree
- **add**: Add file contents to the index
- **commit**: Record changes to the repository
- **push**: Update remote refs along with associated objects
- **pull**: Fetch from and integrate with another repository or local branch
- **branch**: List, create, or delete branches
- **checkout**: Switch branches or restore working tree files

## Supported Models

The git_mcp tool is available for models with function calling capabilities:

- gpt-oss models
- qwen3 models
- deepseek-v3 and deepseek-r1 models
- llama3 models
- gemma3 models
- qwq models

## Usage

### Basic Operations

#### Initialize a Repository

```json
{
  "operation": "init",
  "path": "/path/to/new/repo"
}
```

#### Clone a Repository

```json
{
  "operation": "clone",
  "repository": "https://github.com/username/repo.git",
  "path": "/path/to/clone/destination"
}
```

#### Check Status

```json
{
  "operation": "status",
  "path": "/path/to/repo"
}
```

#### View Commit History

```json
{
  "operation": "log",
  "path": "/path/to/repo",
  "options": {
    "limit": 10,
    "oneline": true
  }
}
```

#### View Changes

```json
{
  "operation": "diff",
  "path": "/path/to/repo",
  "options": {
    "cached": false
  }
}
```

### Making Changes

#### Add Files

```json
{
  "operation": "add",
  "path": "/path/to/repo",
  "files": ["file1.txt", "file2.txt"]
}
```

Or add all changes:

```json
{
  "operation": "add",
  "path": "/path/to/repo"
}
```

#### Commit Changes

```json
{
  "operation": "commit",
  "path": "/path/to/repo",
  "message": "Add new features"
}
```

#### Push to Remote

```json
{
  "operation": "push",
  "path": "/path/to/repo"
}
```

#### Pull from Remote

```json
{
  "operation": "pull",
  "path": "/path/to/repo"
}
```

### Branch Management

#### List Branches

```json
{
  "operation": "branch",
  "path": "/path/to/repo"
}
```

#### Create a Branch

```json
{
  "operation": "branch",
  "path": "/path/to/repo",
  "branch": "feature-branch"
}
```

#### Switch Branches

```json
{
  "operation": "checkout",
  "path": "/path/to/repo",
  "branch": "feature-branch"
}
```

## Schema

The git_mcp tool accepts the following parameters:

- **operation** (required, string): The Git operation to perform. Must be one of: `clone`, `status`, `log`, `diff`, `add`, `commit`, `push`, `pull`, `branch`, `checkout`, `init`
- **repository** (string): Repository URL (required for clone operation)
- **path** (string): Local path for the repository or files
- **message** (string): Commit message (required for commit operation)
- **branch** (string): Branch name (for branch/checkout operations)
- **files** (array of strings): Files to add (for add operation)
- **options** (object): Additional options for the git command
  - **limit** (integer): Limit number of results (for log operation)
  - **oneline** (boolean): Show one line per entry (for log operation)
  - **cached** (boolean): Show cached/staged changes (for diff operation)

## Security Considerations

- The tool sanitizes Git URLs to remove authentication information when displaying output
- All operations are executed within the specified working directory or the server's configured working directory
- The tool validates that git is installed on the system before performing operations
- Path traversal is prevented by cleaning and validating all path inputs

## Examples

### Example 1: Initialize and Commit

```
User: "Initialize a new git repository in /tmp/myproject, add all files, and commit with message 'Initial commit'"

Model uses git_mcp:
1. operation: "init", path: "/tmp/myproject"
2. operation: "add", path: "/tmp/myproject"
3. operation: "commit", path: "/tmp/myproject", message: "Initial commit"
```

### Example 2: Clone and Check Status

```
User: "Clone the ollama repository and show me its status"

Model uses git_mcp:
1. operation: "clone", repository: "https://github.com/ollama/ollama.git", path: "/tmp/ollama"
2. operation: "status", path: "/tmp/ollama"
```

### Example 3: View Recent Changes

```
User: "Show me the last 5 commits and the current diff"

Model uses git_mcp:
1. operation: "log", path: "/current/repo", options: {"limit": 5, "oneline": true}
2. operation: "diff", path: "/current/repo"
```

## Error Handling

The tool provides detailed error messages for common issues:

- Missing required parameters (e.g., operation, repository for clone, message for commit)
- Git command failures (e.g., merge conflicts, authentication failures)
- Invalid paths or non-existent repositories
- Missing git installation

## Platform Support

The git_mcp tool is available on:

- Windows
- macOS (Darwin)
- Linux

The tool requires Git to be installed and available in the system PATH.

## Integration with Ollama

The git_mcp tool is automatically registered when using models that support function calling. It can be used in both:

- **Tools mode**: Single-turn tool execution
- **Agent mode**: Multi-turn conversational agent with tool access

The tool respects the server's configured working directory for relative paths.

## Helper Functions

The tool includes several helper functions for Git operations:

- `IsGitRepository(path)`: Checks if a directory is a Git repository
- `FindGitRoot(startPath)`: Finds the root of a Git repository
- `ValidateGitInstalled()`: Verifies Git is installed on the system
- `SanitizeGitURL(gitURL)`: Removes sensitive information from Git URLs

These helpers ensure safe and reliable Git operations across different environments.
