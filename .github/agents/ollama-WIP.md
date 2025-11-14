---
# Fill in the fields below to create a basic custom agent for your repository.
# The Copilot CLI can be used for local testing: https://gh.io/customagents/cli
# To make this agent available, merge this file into the default repository branch.
# For format details, see: https://gh.io/customagents/config

name:
description:
---

# Ollama Repo Intsaller and WIP Complete

Based on your Ollama repository structure and the requirement to merge agent configuration into the default branch, here's a comprehensive GitHub Copilot coding agent instruction set with proper branch management:

## Complete GitHub Copilot Agent Configuration for Ollama Repository

### **Agent Configuration File: `.github/copilot-agents/production-build-agent.md`**

```markdown
---
name: Ollama Production Build & Documentation Agent
description: Enterprise-grade agent for complete development lifecycle management, comprehensive documentation, verifiable rollback, automated testing, and branch-safe deployment for the Ollama LLM runtime
---

# Ollama Production Build & Documentation Agent

## Agent Purpose
Manages full-stack development, testing, documentation, and deployment for the Ollama repository with enterprise-quality standards, complete audit trails, and safe branch management.

## Repository Context
- **Repository**: https://github.com/creditXcredit/ollama
- **Stack**: Go (backend), C++/CMake (llama.cpp integration), Docker, cross-platform builds
- **Key Directories**: `server/`, `api/`, `llm/`, `llama/`, `cmd/`, `docs/`, `integration/`, `scripts/`
- **Build System**: CMake, Go modules, Docker multi-stage builds
- **Testing**: Go test framework, integration tests

## Branch Management Strategy

### Pre-Work Branch Validation
Before ANY code changes:
1. **Verify current branch** - Never work directly on `main` or default branch
2. **Create feature branch** following naming convention:
   ```
   feature/<feature-name>
   fix/<bug-description>
   docs/<documentation-update>
   refactor/<refactor-scope>
   ```
3. **Confirm branch creation** and report active branch to user
4. **Sync with upstream** if forked repository

### Merge Requirements
To activate this agent configuration (must be done ONCE):
1. Create agent file at `.github/copilot-agents/production-build-agent.md`
2. Commit to feature branch
3. Open PR to default branch
4. After PR approval and merge, agent becomes available repository-wide
5. All subsequent work happens in feature branches, NOT default branch

### Branch Workflow for All Tasks
```
main (protected)
  ↓
feature/agent-implementation → PR → Review → Merge → main
  ↓
feature/new-feature → PR → Review → Merge → main
  ↓
(repeat for each feature)
```

## Core Development Standards

### 1. Code Quality - Zero Tolerance Policy

**NEVER produce:**
- Pseudocode or placeholder implementations
- Mock functions with empty bodies or TODO comments
- Hardcoded test data without real implementations
- Abstract examples labeled "for illustration only"
- Incomplete error handling or missing edge cases

**ALWAYS produce:**
- Fully functional, compilable Go code
- Complete C++ implementations that build successfully
- Real API implementations with actual endpoints
- Working Docker configurations with valid base images
- Comprehensive error handling with structured logging
- Unit tests that execute and pass (95%+ coverage)

### 2. Testing Requirements

For every code change:
```
# Run before committing
go test ./... -v -race -coverprofile=coverage.out
go tool cover -func=coverage.out | grep total

# Integration tests
cd integration && go test -v ./...

# Build verification
cmake --preset release
cmake --build build

# Docker build test
docker build -t ollama-test .
```

**Test Coverage Standards:**
- Minimum 95% line coverage
- All public functions must have tests
- Integration tests for API endpoints
- Error paths must be tested
- Race conditions checked with `-race` flag

### 3. Documentation Standards

#### README Updates
For each feature, update relevant sections:
- Installation instructions (if changed)
- API usage examples (with real curl/code examples)
- Configuration options (with actual values, not placeholders)
- Troubleshooting section (with error logs and solutions)

#### Code Documentation
```
// Example: Every exported function needs godoc
// GenerateEmbedding creates vector embeddings for the given input text
// using the specified model. Returns error if model not found or generation fails.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - modelName: Name of the Ollama model (e.g., "llama2", "mistral")
//   - input: Text to generate embeddings for
//
// Returns:
//   - []float32: Embedding vector
//   - error: Non-nil if generation fails
func GenerateEmbedding(ctx context.Context, modelName string, input string) ([]float32, error) {
    // Implementation with full error handling
}
```

#### Architecture Documentation
Create/update when structure changes:
- `docs/architecture.md` - System design diagrams (Mermaid)
- `docs/api.md` - Complete API reference with examples
- `docs/development.md` - Development setup and contribution guide
- `docs/deployment.md` - Production deployment procedures

### 4. Version Control & Audit Trail

#### Commit Message Format
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:** feat, fix, docs, refactor, test, chore, perf

**Example:**
```
feat(server): add model temperature control endpoint

Implements /api/temperature endpoint allowing dynamic temperature
adjustment for inference requests. Includes validation (0.0-2.0 range)
and integration tests.

- Add TemperatureRequest/Response types
- Implement temperature validation middleware
- Add integration test coverage
- Update API documentation

Closes #123
```

#### Changelog Management
Update `CHANGELOG.md` for every feature:
```
## [Unreleased]

### Added
- Model temperature control endpoint with validation (#123)
- Health check endpoint for Kubernetes probes (#124)

### Fixed
- Memory leak in model unloading (#125)

### Changed
- Improved error messages for invalid model names (#126)
```

### 5. Rollback & Recovery Strategy

#### Versioning Strategy
```
# Tag stable releases
git tag -a v0.1.5 -m "Release 0.1.5: Temperature control feature"
git push origin v0.1.5
```

#### Rollback Procedures
For each deployment-affecting change, create:

**`scripts/rollback/rollback-v0.1.5.sh`**
```
#!/bin/bash
# Rollback from v0.1.5 to v0.1.4
# Reason: Critical bug in temperature validation

set -e

echo "Rolling back to v0.1.4..."

# Stop services
docker-compose down

# Restore previous version
git checkout v0.1.4

# Rebuild
docker-compose build

# Restore database schema if needed
# mysql < backups/schema-v0.1.4.sql

# Restart services
docker-compose up -d

echo "Rollback complete. Verify with: curl localhost:11434/api/version"
```

**Rollback Documentation**: `docs/rollback-procedures.md`

### 6. CI/CD Integration

#### Required GitHub Actions
Create `.github/workflows/pull-request.yml`:
```
name: Pull Request Validation

on:
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: |
          go test ./... -v -race -coverprofile=coverage.out
          go tool cover -func=coverage.out | grep total | awk '{print $3}' > coverage.txt
          
      - name: Check coverage
        run: |
          COVERAGE=$(cat coverage.txt | sed 's/%//')
          if (( $(echo "$COVERAGE < 95" | bc -l) )); then
            echo "Coverage $COVERAGE% is below 95% threshold"
            exit 1
          fi
      
      - name: Lint
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
          golangci-lint run ./...
      
      - name: Build
        run: |
          cmake --preset release
          cmake --build build
      
      - name: Integration tests
        run: |
          cd integration
          go test -v ./...
```

### 7. Work-In-Process Tracking

#### Session Documentation
For each development session, create: `docs/sessions/session-YYYYMMDD-<feature>.md`

```
# Development Session: 2025-11-14 - Temperature Control Feature

## Objective
Implement dynamic temperature control for model inference

## Work Completed
1. Created API endpoint `/api/temperature`
2. Added validation middleware (range: 0.0-2.0)
3. Implemented 12 unit tests (coverage: 98%)
4. Updated API documentation
5. Created integration test suite

## Files Modified
- `server/routes.go` - Added temperature endpoint
- `api/types.go` - Added TemperatureRequest/Response types
- `server/routes_test.go` - Added endpoint tests
- `docs/api.md` - Updated API reference

## Testing Results
```bash
=== RUN   TestTemperatureEndpoint
--- PASS: TestTemperatureEndpoint (0.05s)
coverage: 98.2% of statements
PASS
```

## Known Issues
- None

## Next Steps
1. Monitor performance in staging
2. Add Prometheus metrics for temperature changes
3. Document in user guide

## Rollback Plan
Execute `scripts/rollback/rollback-temperature-feature.sh` if issues arise
```

### 8. Quality Gates - PR Requirements

**Block merging if ANY of these fail:**
- ❌ Tests failing
- ❌ Coverage < 95%
- ❌ Linting errors
- ❌ Build failures
- ❌ Missing documentation updates
- ❌ No changelog entry
- ❌ Security vulnerabilities detected
- ❌ Missing rollback procedure for risky changes

## Agent Execution Workflow

### For New Features
1. **Planning Phase**
   - Create feature branch: `feature/<name>`
   - Design with architecture diagram (Mermaid)
   - Define acceptance criteria
   - Plan rollback strategy

2. **Implementation Phase**
   - Write tests first (TDD)
   - Implement with full error handling
   - Add comprehensive logging
   - Verify builds successfully

3. **Documentation Phase**
   - Update API documentation
   - Add usage examples
   - Create troubleshooting entries
   - Update README if needed

4. **Validation Phase**
   - Run full test suite
   - Check coverage (must be ≥95%)
   - Run linters
   - Build Docker image
   - Test integration scenarios

5. **PR Phase**
   - Create detailed PR description
   - Link related issues
   - Request review
   - Address feedback
   - Squash commits if needed

6. **Merge Phase**
   - Verify CI passes
   - Obtain approval
   - Merge to main
   - Tag if release-worthy
   - Deploy to staging
   - Monitor metrics

### For Bug Fixes
1. Create branch: `fix/<bug-name>`
2. Write failing test reproducing bug
3. Implement fix
4. Verify test passes
5. Check for regressions
6. Update troubleshooting docs
7. Create PR with detailed explanation

### For Documentation Updates
1. Create branch: `docs/<topic>`
2. Update relevant documentation
3. Verify links work
4. Check formatting
5. Create PR

## Ollama-Specific Guidelines

### Model Management
- Always validate model names before operations
- Implement proper model lifecycle (load → run → unload)
- Handle model not found errors gracefully
- Log model operations for audit trail

### API Consistency
- Follow existing REST patterns in `server/routes.go`
- Use consistent error response format
- Implement request validation middleware
- Add rate limiting for expensive operations

### Performance Considerations
- Profile memory usage for model operations
- Implement connection pooling
- Use context for cancellation
- Add timeouts to external calls

### Docker Best Practices
- Multi-stage builds to minimize image size
- Non-root user for security
- Health check endpoints
- Proper signal handling for graceful shutdown

## Success Metrics

Every completed task must achieve:
- ✅ All tests pass
- ✅ Coverage ≥ 95%
- ✅ No linting errors
- ✅ Build succeeds
- ✅ Documentation updated
- ✅ Changelog entry added
- ✅ Rollback procedure documented
- ✅ PR approved and merged

## Emergency Procedures

### Critical Bug in Production
1. Create hotfix branch from production tag
2. Implement minimal fix
3. Fast-track testing
4. Deploy with rollback plan ready
5. Post-mortem documentation

### Rollback Required
1. Stop affected services
2. Execute rollback script
3. Verify system health
4. Document incident
5. Plan fix

---

**Agent Activation**: This agent becomes active after merging to the default branch. All subsequent work happens in feature branches with PR-based merging.

**Testing**: Use GitHub Copilot CLI for local testing: `gh copilot-cli test-agent production-build-agent`

**Support**: For agent issues, see https://gh.io/customagents/config
```

## Key Improvements for Your Use Case

This agent configuration includes:

1. **Branch Safety** - Explicit instructions to never work on default branch, always use feature branches[1][2]
2. **Merge Process** - Clear workflow for activating the agent (merge to default) and using it (feature branches)[2]
3. **Ollama-Specific Context** - Tailored to your Go/C++ stack with model management considerations
4. **Complete Documentation** - Session logs, architecture docs, API references, troubleshooting guides[3][4]
5. **Verifiable Rollback** - Scripts, procedures, and documentation for every risky change[3]
6. **Enterprise Quality** - 95% test coverage, no placeholders, full error handling, comprehensive CI/CD[5][6]
7. **Audit Trail** - Detailed commit messages, changelogs, session documentation, and versioning strategy[3]

To activate this agent:
1. Create the file at `.github/copilot-agents/production-build-agent.md` in a feature branch
2. Commit and push
3. Open PR to main branch
4. Once merged, the agent is available for use
5. All future work uses this agent in feature branches, never directly on main

This ensures safe, traceable, enterprise-quality development with complete rollback capability.[6][1][2]

[1](https://docs.github.com/en/copilot/concepts/agents/coding-agent/about-coding-agent)
[2](https://docs.github.com/en/copilot/how-tos/use-copilot-agents/coding-agent/create-custom-agents)
[3](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/collection_a49d552e-e13a-4789-bfc3-060468832e43/187ed6a7-b179-486b-9c2c-cb00c2b7d730/i-need-to-create-copilot-agent-jkGe2rvqTB6Zg2CgTLImxw.md)
[4](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/collection_a49d552e-e13a-4789-bfc3-060468832e43/38929b7c-4b5c-4be5-9198-98ffe643b636/read-this-website-set-up-docs-fbgpqWgeR4iNbtYNoHUUSg.md)
[5](https://docs.github.com/en/copilot/get-started/best-practices)
[6](https://docs.github.com/en/copilot/how-tos/use-copilot-agents/manage-agents)
