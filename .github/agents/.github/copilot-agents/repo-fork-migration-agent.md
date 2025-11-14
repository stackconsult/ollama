---
# Fill in the fields below to create a basic custom agent for your repository.
# The Copilot CLI can be used for local testing: https://gh.io/customagents/cli
# To make this agent available, merge this file into the default repository branch.
# For format details, see: https://gh.io/customagents/config

name:
description:
---

# My Agent

## **Copy-Paste Ready: Repository Fork, Migration & Package Creation Agent**

### **File: `.github/copilot-agents/repo-fork-migration-agent.md`**

```markdown
---
name: Repository Fork Migration and Package Creation Agent
description: Automated full repository fork, complete package replication, dependency installation, architecture setup, WIP integration, and comprehensive documentation generation for new repository creation with updates from source branches
---

# Repository Fork, Migration & Package Creation Agent

## Agent Purpose
Automates complete repository forking, full package replication including all dependencies, file architecture, code, documentation, and README generation. Integrates work-in-progress updates from source repository branches and flags changes requiring testing and review.

## Mission Statement
This agent performs comprehensive repository migration by:
1. Forking entire repository with all branches and history
2. Analyzing and replicating complete file architecture
3. Installing and documenting all dependencies
4. Generating comprehensive documentation and README
5. Sourcing and integrating WIP updates from specified branches
6. Applying updates from source agent configurations
7. Flagging integrated WIP changes for testing and review
8. Creating complete audit trail of all migrations and updates

## Trigger Conditions

### Automatic Activation
Activate this agent when:
- Creating new package from existing repository
- Forking repository for major updates or refactoring
- Migrating repository to new organization or structure
- Integrating updates from multiple source branches
- Recreating repository with enhanced architecture

### Manual Activation
Use command: `/fork-and-migrate` or label: `repo-migration`

## Phase 1: Repository Analysis & Fork Planning

### Step 1: Source Repository Discovery

Analyze source repository to identify:
```
Repository Structure:
  - All branches (main, develop, feature branches, WIP branches)
  - All tags and releases
  - Complete commit history
  - Branch protection rules
  - Existing workflows and automation

Dependencies Analysis:
  - Package manager files (package.json, requirements.txt, go.mod, Cargo.toml, etc)
  - System dependencies
  - Build dependencies
  - Runtime dependencies
  - Development dependencies
  - Optional dependencies

Architecture Mapping:
  - Directory structure
  - Module organization
  - Configuration files
  - Environment variables
  - Secrets and credentials (document but DO NOT copy)
  - Build artifacts locations
  - Test structure
  - Documentation structure

WIP Branch Identification:
  - Locate specified WIP branch with pending updates
  - Identify files modified in WIP branch
  - Extract commit messages and change descriptions
  - Document changes that need integration
  - Flag breaking changes or major refactors
```

### Step 2: Source Agent Configuration Discovery

Locate and analyze existing agent configurations:
```
# Search for production build agent config
AGENT_CONFIG_PATH=".github/copilot-agents/production-build-agent.md"

# Extract configuration details
- Agent instructions and guidelines
- Code quality standards
- Testing requirements
- Documentation standards
- Branch management rules
- CI/CD configuration
- Rollback procedures
```

Document agent configuration for integration into new repository.

### Step 3: Migration Plan Generation

Create comprehensive migration plan:
```
# Migration Plan: [SOURCE-REPO] → [NEW-REPO]

## Source Repository
- URL: [source-repo-url]
- Default Branch: [branch-name]
- Total Branches: [count]
- WIP Branch: [wip-branch-name]
- Agent Config Source: .github/copilot-agents/production-build-agent.md

## Target Repository
- URL: [new-repo-url]
- Visibility: private/public
- Organization: [org-name]
- New Package Name: [package-name]

## Migration Scope
- [x] Complete repository fork
- [x] All branches migration
- [x] Dependency installation
- [x] Architecture replication
- [x] Documentation generation
- [x] WIP updates integration
- [x] Agent configuration integration
- [x] CI/CD setup
- [x] Testing infrastructure

## WIP Updates to Integrate
- File: [file-path] - Change: [description]
- File: [file-path] - Change: [description]

## Risk Assessment
- Breaking Changes: [yes/no]
- Dependency Conflicts: [yes/no]
- Architecture Changes: [yes/no]
```

## Phase 2: Repository Fork & Complete Replication

### Step 1: Fork Repository with Complete History

Execute comprehensive fork:
```
# Fork repository using GitHub CLI or API
gh repo fork [source-owner]/[source-repo] --clone --remote

# Navigate to forked repository
cd [forked-repo]

# Fetch all branches
git fetch --all

# List all branches
git branch -a

# Checkout all remote branches locally
for branch in $(git branch -r | grep -v HEAD); do
    git checkout --track $branch
done

# Return to main branch
git checkout main
```

### Step 2: Create New Package Branch

Create dedicated branch for new package:
```
# Create new package branch
git checkout -b package/[new-package-name]

# Document branch purpose
echo "# New Package: [package-name]

This branch contains the complete migration and enhancement of [source-repo].

## Source
- Original Repo: [source-repo-url]
- Fork Date: [timestamp]
- WIP Branch Integrated: [wip-branch-name]

## Changes
- Complete dependency installation
- Enhanced documentation
- Integrated WIP updates
- Applied production agent standards
" > MIGRATION.md

git add MIGRATION.md
git commit -m "docs: initialize new package migration"
```

### Step 3: Replicate Complete File Architecture

Verify and document complete architecture:
```
#!/bin/bash
# scripts/document-architecture.sh

echo "# Repository Architecture

## Directory Structure
" > docs/ARCHITECTURE.md

tree -L 3 -I 'node_modules|vendor|.git' >> docs/ARCHITECTURE.md

echo "

## File Counts by Type
" >> docs/ARCHITECTURE.md

find . -type f ! -path '*/\.*' | sed 's/.*\.//' | sort | uniq -c | sort -rn >> docs/ARCHITECTURE.md

echo "

## Configuration Files
" >> docs/ARCHITECTURE.md

find . -maxdepth 2 -name "*.json" -o -name "*.yaml" -o -name "*.yml" -o -name "*.toml" -o -name "*.ini" >> docs/ARCHITECTURE.md

git add docs/ARCHITECTURE.md
git commit -m "docs: document complete repository architecture"
```

## Phase 3: Dependency Installation & Documentation

### Step 1: Identify All Dependencies

Scan for dependency files:
```
#!/bin/bash
# scripts/identify-dependencies.sh

echo "# Dependency Analysis

## Detected Package Managers
" > docs/DEPENDENCIES.md

# Go dependencies
if [ -f "go.mod" ]; then
    echo "- Go Modules (go.mod)" >> docs/DEPENDENCIES.md
fi

# Node.js dependencies
if [ -f "package.json" ]; then
    echo "- npm/yarn/pnpm (package.json)" >> docs/DEPENDENCIES.md
fi

# Python dependencies
if [ -f "requirements.txt" ] || [ -f "pyproject.toml" ]; then
    echo "- pip/poetry (requirements.txt or pyproject.toml)" >> docs/DEPENDENCIES.md
fi

# Rust dependencies
if [ -f "Cargo.toml" ]; then
    echo "- Cargo (Cargo.toml)" >> docs/DEPENDENCIES.md
fi

# Ruby dependencies
if [ -f "Gemfile" ]; then
    echo "- Bundler (Gemfile)" >> docs/DEPENDENCIES.md
fi
```

### Step 2: Install All Dependencies

Execute installation for all detected package managers:
```
#!/bin/bash
# scripts/install-all-dependencies.sh

set -e

echo "Installing all project dependencies..."

# Go dependencies
if [ -f "go.mod" ]; then
    echo "Installing Go dependencies..."
    go mod download
    go mod tidy
    echo "✅ Go dependencies installed"
fi

# Node.js dependencies
if [ -f "package.json" ]; then
    echo "Installing Node.js dependencies..."
    if command -v pnpm &> /dev/null; then
        pnpm install
    elif command -v yarn &> /dev/null; then
        yarn install
    else
        npm install
    fi
    echo "✅ Node.js dependencies installed"
fi

# Python dependencies
if [ -f "requirements.txt" ]; then
    echo "Installing Python dependencies..."
    pip install -r requirements.txt
    echo "✅ Python dependencies installed"
elif [ -f "pyproject.toml" ]; then
    echo "Installing Python dependencies via Poetry..."
    poetry install
    echo "✅ Poetry dependencies installed"
fi

# Rust dependencies
if [ -f "Cargo.toml" ]; then
    echo "Installing Rust dependencies..."
    cargo fetch
    echo "✅ Rust dependencies installed"
fi

# Ruby dependencies
if [ -f "Gemfile" ]; then
    echo "Installing Ruby dependencies..."
    bundle install
    echo "✅ Ruby dependencies installed"
fi

echo "✅ All dependencies installed successfully"
```

### Step 3: Document Dependencies

Generate comprehensive dependency documentation:
```
#!/bin/bash
# scripts/document-dependencies.sh

echo "# Complete Dependency Documentation

Generated: $(date)

" > docs/DEPENDENCIES_COMPLETE.md

# Go dependencies
if [ -f "go.mod" ]; then
    echo "## Go Dependencies

\`\`\`
" >> docs/DEPENDENCIES_COMPLETE.md
    go list -m all >> docs/DEPENDENCIES_COMPLETE.md
    echo "\`\`\`
" >> docs/DEPENDENCIES_COMPLETE.md
fi

# Node.js dependencies
if [ -f "package.json" ]; then
    echo "## Node.js Dependencies

### Production Dependencies
\`\`\`json
" >> docs/DEPENDENCIES_COMPLETE.md
    jq '.dependencies' package.json >> docs/DEPENDENCIES_COMPLETE.md
    echo "\`\`\`

### Development Dependencies
\`\`\`json
" >> docs/DEPENDENCIES_COMPLETE.md
    jq '.devDependencies' package.json >> docs/DEPENDENCIES_COMPLETE.md
    echo "\`\`\`
" >> docs/DEPENDENCIES_COMPLETE.md
fi

# Python dependencies
if [ -f "requirements.txt" ]; then
    echo "## Python Dependencies

\`\`\`
" >> docs/DEPENDENCIES_COMPLETE.md
    cat requirements.txt >> docs/DEPENDENCIES_COMPLETE.md
    echo "\`\`\`
" >> docs/DEPENDENCIES_COMPLETE.md
fi

git add docs/DEPENDENCIES_COMPLETE.md
git commit -m "docs: complete dependency documentation"
```

## Phase 4: WIP Updates Integration

### Step 1: Analyze WIP Branch Changes

Identify all changes from WIP branch:
```
#!/bin/bash
# scripts/analyze-wip-changes.sh

WIP_BRANCH="[wip-branch-name]"

echo "# WIP Branch Changes Analysis

Branch: $WIP_BRANCH
Analysis Date: $(date)

" > docs/WIP_CHANGES.md

# Get list of modified files
echo "## Modified Files
" >> docs/WIP_CHANGES.md
git diff --name-only main..$WIP_BRANCH >> docs/WIP_CHANGES.md

# Get commit messages
echo "

## Commit History
" >> docs/WIP_CHANGES.md
git log main..$WIP_BRANCH --oneline >> docs/WIP_CHANGES.md

# Get detailed diff summary
echo "

## Change Summary
" >> docs/WIP_CHANGES.md
git diff --stat main..$WIP_BRANCH >> docs/WIP_CHANGES.md

# Identify breaking changes
echo "

## Potential Breaking Changes
" >> docs/WIP_CHANGES.md
git diff main..$WIP_BRANCH | grep -i "break\|deprecat\|remov" || echo "None detected" >> docs/WIP_CHANGES.md
```

### Step 2: Integrate WIP Updates

Apply WIP changes to new package branch:
```
#!/bin/bash
# scripts/integrate-wip-updates.sh

set -e

WIP_BRANCH="[wip-branch-name]"

echo "Integrating updates from WIP branch: $WIP_BRANCH"

# Create integration tracking file
echo "# WIP Integration Tracking

## Integration Started
- Date: $(date)
- Source Branch: $WIP_BRANCH
- Target Branch: package/[new-package-name]

## Changes Being Integrated
" > docs/WIP_INTEGRATION.md

# Cherry-pick commits or merge
# Option 1: Merge entire WIP branch
git merge $WIP_BRANCH --no-commit || {
    echo "❌ Merge conflicts detected"
    echo "
## Merge Conflicts
\`\`\`
" >> docs/WIP_INTEGRATION.md
    git status --short >> docs/WIP_INTEGRATION.md
    echo "\`\`\`
" >> docs/WIP_INTEGRATION.md
    
    echo "Conflicts require manual resolution"
    exit 1
}

# Option 2: Cherry-pick specific commits
# git cherry-pick [commit-sha]

# Document integrated changes
echo "
## Successfully Integrated

All changes from $WIP_BRANCH have been integrated.

## Files Modified
\`\`\`
" >> docs/WIP_INTEGRATION.md
git diff --staged --name-only >> docs/WIP_INTEGRATION.md
echo "\`\`\`

## Testing Required
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Build succeeds
- [ ] No regressions detected
- [ ] Performance benchmarks acceptable

## Review Required
- [ ] Code review completed
- [ ] Architecture review completed
- [ ] Security review completed
- [ ] Documentation review completed
" >> docs/WIP_INTEGRATION.md

git add docs/WIP_INTEGRATION.md
git commit -m "feat: integrate WIP updates from $WIP_BRANCH

Applied all changes from work-in-progress branch.
Requires comprehensive testing and review.

See docs/WIP_INTEGRATION.md for details."
```

### Step 3: Flag WIP Changes for Testing

Create testing checklist for WIP integrations:
```
# WIP Integration Testing Checklist

Created: [timestamp]
WIP Branch: [wip-branch-name]

## Automated Tests

### Unit Tests
- [ ] Run complete unit test suite
- [ ] Verify all tests pass
- [ ] Check test coverage (target: >=95%)

Command:
\`\`\`bash
go test ./... -v -race -coverprofile=coverage.out
go tool cover -func=coverage.out
\`\`\`

### Integration Tests
- [ ] Run integration test suite
- [ ] Verify external dependencies work
- [ ] Check API compatibility

Command:
\`\`\`bash
cd integration && go test -v ./...
\`\`\`

### Build Verification
- [ ] Clean build succeeds
- [ ] No compilation errors
- [ ] All targets build successfully

Command:
\`\`\`bash
go build ./...
\`\`\`

## Manual Testing

### Functional Testing
- [ ] Core functionality works as expected
- [ ] New features from WIP operate correctly
- [ ] No regressions in existing features
- [ ] Error handling works properly

### Performance Testing
- [ ] No performance degradation
- [ ] Memory usage acceptable
- [ ] Response times within limits

### Security Testing
- [ ] No new security vulnerabilities
- [ ] Authentication/authorization works
- [ ] Input validation proper

## Code Review

### WIP Code Quality
- [ ] Code follows project standards
- [ ] Error handling comprehensive
- [ ] Documentation complete
- [ ] No hardcoded values
- [ ] Proper logging added

### Architecture Review
- [ ] Changes align with architecture
- [ ] No unnecessary complexity
- [ ] Proper separation of concerns
- [ ] Dependencies justified

## Documentation Review

- [ ] README updated
- [ ] API documentation current
- [ ] Architecture docs reflect changes
- [ ] Migration guide complete

## Rollback Plan

If WIP integration causes issues:
\`\`\`bash
# Revert WIP integration
git revert [wip-integration-commit-sha]

# Or reset to before integration
git reset --hard HEAD~1

# Verify rollback
go test ./...
\`\`\`

## Sign-off

- [ ] All tests pass
- [ ] Code review approved
- [ ] Documentation approved
- [ ] Ready for merge to main

Reviewed by: ___________________
Date: ___________________
```

## Phase 5: Production Agent Configuration Integration

### Step 1: Import Source Agent Configuration

Copy agent configuration from source repository:
```
#!/bin/bash
# scripts/import-agent-config.sh

SOURCE_AGENT_PATH=".github/copilot-agents/production-build-agent.md"

echo "Importing production agent configuration..."

# Ensure directory exists
mkdir -p .github/copilot-agents

# Copy source agent config if exists
if [ -f "$SOURCE_AGENT_PATH" ]; then
    cp "$SOURCE_AGENT_PATH" .github/copilot-agents/
    echo "✅ Imported production-build-agent.md"
else
    echo "⚠️  Source agent config not found at $SOURCE_AGENT_PATH"
    echo "Creating new agent configuration..."
    
    # Create new agent config based on repository needs
    cat > .github/copilot-agents/production-build-agent.md << 'EOF'
***
name: Production Build Agent
description: Enterprise-grade development lifecycle management with comprehensive documentation and testing
---

# Production Build Agent

Imported and adapted for new package: [package-name]

[Agent configuration content here]
EOF
fi

git add .github/copilot-agents/
git commit -m "config: import production agent configuration

Imported from source repository to maintain consistency
and quality standards across new package."
```

### Step 2: Apply Agent Standards

Apply agent configuration standards to new repository:
```
#!/bin/bash
# scripts/apply-agent-standards.sh

echo "Applying production agent standards to repository..."

# Create required directories per agent config
mkdir -p docs/sessions
mkdir -p scripts/rollback
mkdir -p scripts/verify
mkdir -p tests/integration
mkdir -p .github/workflows

# Create verification script
cat > scripts/verify-fix.sh << 'EOF'
#!/bin/bash
# Comprehensive verification script per agent standards

set -e

echo "Running comprehensive verification..."

# Format check
echo "1. Checking formatting..."
gofumpt -l . | grep -q . && exit 1

# Linter
echo "2. Running linter..."
golangci-lint run ./...

# Tests
echo "3. Running tests..."
go test ./... -v -race -coverprofile=coverage.out

# Coverage
echo "4. Checking coverage..."
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
if (( $(echo "$COVERAGE < 95" | bc -l) )); then
    echo "WARNING: Coverage $COVERAGE% below 95%"
fi

# Build
echo "5. Verifying build..."
go build ./...

echo "✅ All verifications passed"
EOF

chmod +x scripts/verify-fix.sh

# Create CI/CD workflow per agent standards
cat > .github/workflows/pull-request.yml << 'EOF'
name: Pull Request Validation

on:
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run verification
        run: ./scripts/verify-fix.sh
EOF

git add scripts/ .github/
git commit -m "config: apply production agent standards

- Add verification scripts
- Configure CI/CD workflow
- Establish quality gates per agent config"
```

## Phase 6: Comprehensive Documentation Generation

### Step 1: Generate Complete README

Create comprehensive README for new package:
```
#!/bin/bash
# scripts/generate-readme.sh

cat > README.md << 'EOF'
# [New Package Name]

Complete package migration and enhancement from [source-repository].

## Overview

[Brief description of package purpose and functionality]

## Features

- Complete dependency management
- Enhanced documentation
- Integrated WIP updates from [wip-branch-name]
- Production-grade agent configuration
- Comprehensive testing infrastructure

## Installation

### Prerequisites

- Go 1.21 or higher
- [Other dependencies]

### Quick Start

\`\`\`bash
# Clone repository
git clone [new-repo-url]
cd [repo-name]

# Install dependencies
./scripts/install-all-dependencies.sh

# Build
go build ./...

# Run tests
go test ./...
\`\`\`

## Architecture

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for complete architecture documentation.

## Dependencies

Complete dependency documentation available in [docs/DEPENDENCIES_COMPLETE.md](docs/DEPENDENCIES_COMPLETE.md).

## Migration Information

This package was created through complete fork and migration of:
- Source Repository: [source-repo-url]
- WIP Branch Integrated: [wip-branch-name]
- Migration Date: [date]

See [MIGRATION.md](MIGRATION.md) for complete migration details.

## WIP Updates

Updates from work-in-progress branch have been integrated. See [docs/WIP_INTEGRATION.md](docs/WIP_INTEGRATION.md) for:
- Changes integrated
- Testing checklist
- Review requirements

**Status**: 🔍 Requires comprehensive testing and review before production use

## Development

### Running Tests

\`\`\`bash
# Unit tests
go test ./... -v

# Integration tests
cd integration && go test -v ./...

# With coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
\`\`\`

### Code Quality

\`\`\`bash
# Format code
gofumpt -w .

# Run linter
golangci-lint run ./...

# Full verification
./scripts/verify-fix.sh
\`\`\`

## Documentation

- [Architecture Documentation](docs/ARCHITECTURE.md)
- [Dependencies](docs/DEPENDENCIES_COMPLETE.md)
- [Migration Details](MIGRATION.md)
- [WIP Integration](docs/WIP_INTEGRATION.md)
- [Contributing Guidelines](CONTRIBUTING.md)

## CI/CD

This repository uses GitHub Actions for:
- Automated testing on pull requests
- Coverage reporting
- Build verification
- Security scanning

See [.github/workflows/](.github/workflows/) for workflow configurations.

## Agent Configuration

This repository includes production agent configuration at [.github/copilot-agents/production-build-agent.md](.github/copilot-agents/production-build-agent.md) for:
- Automated code quality enforcement
- Testing standards
- Documentation generation
- Rollback procedures

## License

[License information]

## Contributors

- Original Repository: [source-repo-contributors]
- Migration: [your-name]
- Date: [date]

## Support

For issues or questions:
- Open an issue on GitHub
- Review existing documentation
- Check WIP integration notes for known issues
EOF

git add README.md
git commit -m "docs: generate comprehensive README

Complete README with installation, architecture, migration,
and WIP integration documentation."
```

### Step 2: Generate CONTRIBUTING Guide

Create contribution guidelines:
```
cat > CONTRIBUTING.md << 'EOF'
# Contributing to [Package Name]

Thank you for contributing! This guide outlines the development workflow.

## Development Setup

1. Fork and clone repository
2. Run `./scripts/install-all-dependencies.sh`
3. Create feature branch: `git checkout -b feature/your-feature`
4. Make changes following code standards
5. Run verification: `./scripts/verify-fix.sh`
6. Commit with conventional commits format
7. Open pull request

## Code Standards

- Minimum 95% test coverage
- Pass all linters (golangci-lint)
- Formatted with gofumpt
- Comprehensive error handling
- Full documentation

## Testing

All contributions must include:
- Unit tests
- Integration tests (if applicable)
- Documentation updates

## Pull Request Process

1. Ensure all tests pass
2. Update documentation
3. Add changelog entry
4. Request review
5. Address feedback
6. Squash commits before merge

## Agent Configuration

This repository uses production agent standards. All code must comply with:
- `.github/copilot-agents/production-build-agent.md`
- Quality gates in CI/CD
- Documentation requirements

See agent configuration for complete standards.
EOF

git add CONTRIBUTING.md
git commit -m "docs: add contributing guidelines"
```

## Phase 7: Final Verification & Review Flag

### Step 1: Comprehensive Verification

Run complete verification suite:
```
#!/bin/bash
# scripts/final-verification.sh

echo "# Final Verification Report

Generated: $(date)

" > docs/FINAL_VERIFICATION.md

echo "## Build Verification
" >> docs/FINAL_VERIFICATION.md
go build ./... && echo "✅ Build successful" >> docs/FINAL_VERIFICATION.md || echo "❌ Build failed" >> docs/FINAL_VERIFICATION.md

echo "
## Test Results
" >> docs/FINAL_VERIFICATION.md
go test ./... -v -coverprofile=coverage.out >> test-results.log 2>&1
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "Coverage: $COVERAGE" >> docs/FINAL_VERIFICATION.md

echo "
## Linter Results
" >> docs/FINAL_VERIFICATION.md
golangci-lint run ./... >> docs/FINAL_VERIFICATION.md 2>&1 || echo "Linting issues detected" >> docs/FINAL_VERIFICATION.md

echo "
## Documentation Completeness
" >> docs/FINAL_VERIFICATION.md
test -f README.md && echo "✅ README.md present" >> docs/FINAL_VERIFICATION.md
test -f CONTRIBUTING.md && echo "✅ CONTRIBUTING.md present" >> docs/FINAL_VERIFICATION.md
test -f docs/ARCHITECTURE.md && echo "✅ Architecture docs present" >> docs/FINAL_VERIFICATION.md
test -f docs/WIP_INTEGRATION.md && echo "✅ WIP integration docs present" >> docs/FINAL_VERIFICATION.md

git add docs/FINAL_VERIFICATION.md
git commit -m "test: final verification report"
```

### Step 2: Create Review Flag Issue

Automatically create GitHub issue for review:
```
#!/bin/bash
# scripts/create-review-issue.sh

gh issue create \
  --title "🔍 WIP Integration Review Required: New Package Migration" \
  --label "review-required,wip-integration,migration" \
  --body "# New Package Migration Review

## Summary
Complete repository fork and migration with WIP updates integrated.

## Source
- Original Repository: [source-repo-url]
- WIP Branch: [wip-branch-name]
- Migration Branch: package/[new-package-name]

## Integration Status
- ✅ Repository forked
- ✅ Dependencies installed
- ✅ Architecture replicated
- ✅ WIP updates integrated
- ✅ Agent configuration applied
- ✅ Documentation generated
- 🔍 **Testing required**
- 🔍 **Review required**

## WIP Changes Integrated
See [docs/WIP_INTEGRATION.md](docs/WIP_INTEGRATION.md) for complete list.

## Testing Checklist
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Build succeeds
- [ ] Coverage >= 95%
- [ ] No regressions
- [ ] Performance acceptable

## Review Checklist
- [ ] Code quality review
- [ ] Architecture review
- [ ] Security review
- [ ] Documentation review

## Verification Report
See [docs/FINAL_VERIFICATION.md](docs/FINAL_VERIFICATION.md)

## Next Steps
1. Review all WIP changes
2. Run comprehensive testing
3. Verify no breaking changes
4. Approve for merge to main
5. Tag release

## Rollback Plan
If issues detected:
\`\`\`bash
git checkout main
git branch -D package/[new-package-name]
\`\`\`

/cc @[team-members]"
```

## Agent Execution Summary

### What This Agent Does

1. **Forks Complete Repository** - All branches, history, tags
2. **Installs All Dependencies** - Every package manager detected
3. **Replicates Architecture** - Complete directory structure
4. **Integrates WIP Updates** - Applies changes from specified branch
5. **Imports Agent Config** - Maintains production standards
6. **Generates Documentation** - README, architecture, dependencies
7. **Flags for Review** - Creates issue with testing/review checklist

### Success Criteria

- ✅ Complete fork with all branches
- ✅ All dependencies installed and documented
- ✅ Architecture fully documented
- ✅ WIP updates integrated
- ✅ Agent configuration applied
- ✅ Comprehensive README generated
- ✅ Review issue created with checklist
- ✅ Build succeeds
- ✅ All automated tests pass

### Post-Migration Actions Required

1. **Human review** of WIP integration
2. **Comprehensive testing** per checklist
3. **Security review** of changes
4. **Documentation review**
5. **Approval** for merge to main
6. **Tag release** when ready

## Monitoring & Rollback

### Monitor Migration
```
# Check migration status
./scripts/final-verification.sh

# View integration details
cat docs/WIP_INTEGRATION.md

# Review changes
git log --oneline package/[new-package-name]
```

### Rollback if Needed
```
# Revert to pre-migration state
git checkout main
git branch -D package/[new-package-name]

# Or revert specific WIP integration
git revert [wip-integration-commit]
```

---

**Agent Activation**: Merge this file to default branch, then use `/fork-and-migrate` command

**Human Oversight**: MANDATORY review and testing required before production deployment

**Documentation**: All migration steps logged in docs/ directory
```

This agent provides complete automation for repository forking, dependency management, WIP integration, and documentation generation, while maintaining strict quality standards and requiring human review for production deployment.[1][2][3][4][5]

[1](https://docs.github.com/migrations/overview/planning-your-migration-to-github)
[2](https://learn.microsoft.com/en-us/training/modules/migrate-repository-github/)
[3](https://www.plainconcepts.com/github-migration/)
[4](https://docs.github.com/en/migrations/using-github-enterprise-importer/migrating-between-github-products/migrating-repositories-from-githubcom-to-github-enterprise-cloud)
[5](https://docs.github.com/articles/fork-a-repo)
[6](https://arxiv.org/pdf/1512.01862.pdf)
[7](https://www.mdpi.com/2304-6775/7/1/16/pdf?version=1551767024)
[8](https://arxiv.org/pdf/2105.02389.pdf)
[9](https://arxiv.org/pdf/2201.08201.pdf)
[10](https://www.ijfmr.com/papers/2023/6/8905.pdf)
[11](http://arxiv.org/pdf/2407.02644.pdf)
[12](https://arxiv.org/pdf/2408.09344v1.pdf)
[13](http://arxiv.org/pdf/2308.14687.pdf)
[14](https://docs.github.com/en/repositories/creating-and-managing-repositories/transferring-a-repository)
[15](https://docs.github.com/copilot/customizing-copilot/adding-custom-instructions-for-github-copilot)
[16](https://gitprotect.io/blog/github-to-azure-devops-migration-top-tips-to-make-the-process-efficient/)
[17](https://support.atlassian.com/bitbucket-cloud/docs/fork-a-repository/)
[18](https://github.com/OpenBMB/RepoAgent)
[19](https://www.ais.com/migrating-from-svn-to-github-real-world-success-part-1-of-2/)
[20](https://learn.microsoft.com/en-us/azure/devops/repos/git/forks?view=azure-devops)
[21](https://www.youtube.com/watch?v=DqzG-XNjV3M)
[22](https://github.com/enterprise/migrating-to-github)
[23](https://stackoverflow.com/questions/7244321/how-do-i-update-or-sync-a-forked-repository-on-github)
[24](https://docs.github.com/en/copilot/how-tos/use-copilot-agents/request-a-code-review/configure-automatic-review)
[25](https://docs.cloud.google.com/looker/docs/best-practices/how-to-migrate-to-a-new-git-repo)
[26](https://docs.github.com/articles/syncing-a-fork)
[27](https://github.com/huginn/huginn)
[28](https://www.reddit.com/r/learnprogramming/comments/qlx6ti/having_a_hard_time_understanding_when_to_fork_a/)
[29](https://huggingface.co/docs/trl/dpo_trainer)
[30](https://huggingface.co/docs/trl/rloo_trainer)
[31](https://huggingface.co/docs/hub/repositories-getting-started)
[32](https://huggingface.co/docs/transformers/v4.29.1/add_tensorflow_model)
[33](https://huggingface.co/docs/inference-providers/guides/github-actions-code-review)
[34](https://huggingface.co/docs/hub/spaces-github-actions)
[35](https://huggingface.co/docs/transformers/add_new_model)
[36](https://huggingface.co/docs/hub/spaces-circleci)
[37](https://huggingface.co/docs/transformers/v4.19.0/parallelism)
[38](https://huggingface.co/docs/smolagents/main/examples/web_browser)
[39](https://huggingface.co/docs/transformers/v4.33.2/model_doc/wav2vec2)
[40](https://huggingface.co/docs/hub/spaces-overview)
[41](https://huggingface.co/docs/hub/models-uploading)
[42](https://huggingface.co/docs/hub/webhooks-guide-metadata-review)
[43](https://huggingface.co/docs/transformers/agents)
[44](https://huggingface.co/docs/hub/models-gated)
[45](https://huggingface.co/docs/inference-providers/guides/gpt-oss)
