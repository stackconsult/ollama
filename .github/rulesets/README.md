# GitHub Repository Rulesets

This directory contains GitHub repository rulesets that define branch protection rules for the Ollama repository.

## Available Rulesets

### 1. Main Branch Protection (`main-branch-protection.json`)

Protects the default branch (main) with the following rules:

- **Branch Deletion**: Prevents deletion of the main branch
- **Force Push Prevention**: Disallows force pushes and other non-fast-forward updates
- **Pull Request Requirements**:
  - At least 1 approving review required
  - Latest push must be approved (prevents last-minute changes without review)
  - Stale reviews are dismissed when new commits are pushed
  - All review threads must be resolved before merging
- **Required Status Checks**:
  - `test / go_mod_tidy` - Ensures Go modules are properly tidied
  - `test / test (ubuntu-latest)` - Tests on Ubuntu
  - `test / test (macos-latest)` - Tests on macOS
  - `test / test (windows-latest)` - Tests on Windows
  - `test / patches` - Verifies patches apply cleanly
  - Strict mode enabled (branch must be up-to-date with base branch)
- **Linear History**: Requires a linear commit history (no merge commits)
- **Bypass Actors**: Repository administrators can bypass rules when creating pull requests

### 2. Release Branch Protection (`release-branch-protection.json`)

Protects release branches (matching `release/*` and `v*` patterns) with stricter requirements:

- **Branch Deletion**: Prevents deletion of release branches
- **Force Push Prevention**: Disallows force pushes and other non-fast-forward updates
- **Pull Request Requirements**:
  - At least 2 approving reviews required (stricter than main)
  - Latest push must be approved
  - Stale reviews are dismissed when new commits are pushed
  - All review threads must be resolved before merging
- **Linear History**: Requires a linear commit history
- **Bypass Actors**: Repository administrators can bypass rules when creating pull requests

## How to Import Rulesets

### Via GitHub Web UI

1. Navigate to your repository on GitHub
2. Go to **Settings** → **Rules** → **Rulesets**
3. Click **New ruleset** → **Import a ruleset**
4. Upload the JSON file from this directory
5. Review the configuration and click **Create** or **Update**

### Via GitHub CLI

```bash
# Import main branch protection
gh api repos/{owner}/{repo}/rulesets \
  --method POST \
  --input .github/rulesets/main-branch-protection.json

# Import release branch protection
gh api repos/{owner}/{repo}/rulesets \
  --method POST \
  --input .github/rulesets/release-branch-protection.json
```

### Via GitHub API

```bash
# Using curl
curl -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  https://api.github.com/repos/{owner}/{repo}/rulesets \
  -d @.github/rulesets/main-branch-protection.json
```

## Customization

You can customize these rulesets by:

1. Modifying the JSON files directly
2. Adjusting the number of required reviewers
3. Adding or removing required status checks
4. Changing the target branch patterns
5. Modifying bypass actors

After making changes, re-import the ruleset through the GitHub UI or API.

## References

- [GitHub Rulesets Documentation](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets)
- [Available Rules for Rulesets](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/available-rules-for-rulesets)
- [GitHub Ruleset Recipes](https://github.com/github/ruleset-recipes)

## Notes

- Rulesets are enforced at the repository or organization level, not through files in the repository
- These JSON files serve as documentation and templates that can be imported into GitHub
- Changes to these files do not automatically update the active rulesets - they must be re-imported
- Administrators with bypass permissions can still push directly to protected branches if needed
