# Ollama Tools

This package provides tool implementations for Ollama models that support function calling. Tools enable LLMs to interact with external systems and perform actions beyond text generation.

## Available Tools

### Browser Tools

Browser tools simulate a full browser environment, allowing models to search, navigate, and interact with web pages.

- **browser.search**: Search the web and get structured results
- **browser.open**: Open web pages and extract content
- **browser.find**: Find specific text within web pages

**Supported Models**: gpt-oss models

**Files**:
- `browser.go`: Core browser functionality and page management
- `browser_websearch.go`: Web search implementation
- `browser_crawl.go`: Web page crawling
- `browser_test.go`: Browser tool tests

### Web Search Tools

Simpler web search tools that provide basic search and fetch capabilities without full browser simulation.

- **web_search**: Search the web using Ollama's web search API
- **web_fetch**: Fetch and extract content from web pages

**Supported Models**: qwen3, deepseek-v3

**Files**:
- `web_search.go`: Web search implementation
- `web_fetch.go`: Web page fetching

### Git MCP Tool

Git MCP (Model Context Protocol) tool provides Git repository management capabilities.

- **git_mcp**: Perform Git operations (clone, status, log, diff, add, commit, push, pull, branch, checkout, init)

**Supported Models**: gpt-oss, qwen3, deepseek-v3, deepseek-r1, llama3, gemma3, qwq

**Files**:
- `git_mcp.go`: Git operations implementation
- `git_mcp_test.go`: Git tool tests

**Features**:
- Repository initialization and cloning
- Status and log viewing
- File staging and committing
- Branch management
- Push and pull operations
- Secure URL sanitization

## Tool Architecture

### Tool Interface

All tools implement the `Tool` interface defined in `tools.go`:

```go
type Tool interface {
    Name() string
    Description() string
    Schema() map[string]any
    Execute(ctx context.Context, args map[string]any) (any, string, error)
    Prompt() string
}
```

### Registry

The `Registry` manages available tools and their execution:

- **NewRegistry()**: Creates a new tool registry
- **Register(tool Tool)**: Adds a tool to the registry
- **Get(name string)**: Retrieves a tool by name
- **Execute(ctx, name, args)**: Executes a tool with given arguments
- **AvailableTools()**: Returns all tools as schema maps for API calls

### Tool Execution Flow

1. Model requests tool execution with name and arguments
2. Registry looks up the tool by name
3. Tool validates arguments against schema
4. Tool executes the operation
5. Tool returns results (both structured data and text for the model)

## Adding New Tools

To add a new tool:

1. Create a new file (e.g., `my_tool.go`)
2. Implement the `Tool` interface
3. Define the JSON schema for parameters
4. Implement the `Execute` method
5. Add tests in a corresponding `_test.go` file
6. Register the tool in `app/ui/ui.go`
7. Add model support check if needed
8. Document the tool in `docs/`

### Example Tool Structure

```go
package tools

import (
    "context"
    "encoding/json"
    "fmt"
)

type MyTool struct{}

func (t *MyTool) Name() string {
    return "my_tool"
}

func (t *MyTool) Description() string {
    return "Description of what the tool does"
}

func (t *MyTool) Schema() map[string]any {
    schemaBytes := []byte(`{
        "type": "object",
        "properties": {
            "param1": {
                "type": "string",
                "description": "Description of param1"
            }
        },
        "required": ["param1"]
    }`)
    var schema map[string]any
    json.Unmarshal(schemaBytes, &schema)
    return schema
}

func (t *MyTool) Prompt() string {
    return ""
}

func (t *MyTool) Execute(ctx context.Context, args map[string]any) (any, string, error) {
    param1, ok := args["param1"].(string)
    if !ok {
        return nil, "", fmt.Errorf("param1 is required")
    }
    
    // Perform the tool operation
    result := doSomething(param1)
    
    // Return both structured data and text for the model
    return result, "Operation completed successfully", nil
}
```

## Testing

All tools should have comprehensive tests covering:

- Schema validation
- Parameter handling
- Error cases
- Successful operations
- Edge cases

Run tests:

```bash
cd app/tools
go test -v .
```

## Build Tags

Tools may use build tags to restrict compilation to specific platforms:

- `//go:build windows || darwin`: macOS and Windows only
- `//go:build windows || darwin || linux`: All major platforms

The git_mcp tool is available on all major platforms.

## Security Considerations

When implementing tools:

- Validate all inputs
- Sanitize sensitive information (passwords, tokens) from outputs
- Use context for cancellation support
- Limit resource usage (timeouts, rate limits)
- Validate paths to prevent directory traversal
- Handle errors gracefully

## Documentation

Each tool should have:

1. Inline documentation (godoc comments)
2. Schema documentation (parameter descriptions)
3. Usage examples in `docs/` directory
4. Test coverage demonstrating usage

See `docs/git-mcp.md` for a comprehensive documentation example.

## Related Files

- `app/ui/ui.go`: Tool registration and model support checks
- `docs/`: Tool documentation
- `integration/tools_test.go`: Integration tests

## References

- [Ollama API Documentation](../docs/api.md)
- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Git MCP Tool Documentation](../docs/git-mcp.md)
